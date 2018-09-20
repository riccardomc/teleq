package database

import (
	"errors"
	"fmt"
	"strings"
	"sync"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/postgres" // noqa
	_ "github.com/jinzhu/gorm/dialects/sqlite"   // noqa
)

//Stack data structure
type Stack struct {
	gorm.Model
	db     *Database
	Name   string            `gorm:"type:VARCHAR(100);PRIMARY_KEY"`
	Frames []PersistentFrame `gorm:"foreignkey:Stack"`
	mux    sync.Mutex
}

//New returns a new Stack
func New() *Stack {
	return &Stack{
		Frames: make([]PersistentFrame, 0),
		Name:   uuid.New().String(),
	}
}

func guessDialect(connectionString string) (string, error) {
	if strings.HasPrefix(connectionString, "/") {
		return "sqlite3", nil
	}
	if strings.Contains(connectionString, "host=") {
		return "postgres", nil
	}
	return "", errors.New("Cannot guess connection SQL dialect")
}

//Init the Stack
func (s *Stack) Init(params ...interface{}) error {
	databaseConnectionString := params[0].(string)
	dialect, err := guessDialect(databaseConnectionString)
	if err != nil {
		return err
	}
	fmt.Println("Got a", dialect, "database")
	s.db = NewDatabase(dialect, databaseConnectionString)
	err = s.db.Connect()
	if err != nil {
		return err
	}
	s.db.client.AutoMigrate(&Stack{})
	s.db.client.AutoMigrate(&PersistentFrame{})
	s.db.client.Create(s)
	return nil
}

//Push a frame on the stack
func (s *Stack) Push(data interface{}) {
	frame := Frame{data}
	serializedFrame, err := frame.Serialize()
	if err != nil {
		panic(err)
	}
	pf := PersistentFrame{Data: serializedFrame, Stack: s.Name}
	s.db.client.Create(&pf)
}

//Peek top frame on the stack
func (s *Stack) Peek() interface{} {
	frame := Frame{}
	persistentFrame := PersistentFrame{}
	s.db.client.Last(&persistentFrame)
	frame.Deserialize(persistentFrame.Data)
	return frame.Data
}

//Size returns the number of frames in the stack
func (s *Stack) Size() int {
	var count int
	s.db.client.Table("persistent_frames").Select(
		"persistent_frames.stack").Joins(
		"left join stacks on stacks.name = persistent_frames.stack").Count(&count)
	return count
}

//Pop a persistent frame
func (s *Stack) Pop() interface{} {
	frame := Frame{}
	persistentFrame := &PersistentFrame{}

	s.mux.Lock()
	defer s.mux.Unlock()

	t := s.db.client.Begin()
	defer func() {
		if r := recover(); r != nil {
			t.Rollback()
		}
	}()

	if t.Error != nil {
		return nil
	}

	if err := t.Last(persistentFrame).Error; err != nil {
		t.Rollback()
		return nil
	}

	if persistentFrame != nil {
		if err := t.Unscoped().Delete(persistentFrame).Error; err != nil {
			t.Rollback()
			return nil
		}
	}

	if err := t.Commit().Error; err != nil {
		return nil
	}

	frame.Deserialize(persistentFrame.Data)
	return frame.Data
}
