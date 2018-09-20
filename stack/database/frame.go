package database

import (
	"bytes"
	"encoding/gob"

	"github.com/jinzhu/gorm"
)

//Frame is a Stack Frame
type Frame struct {
	Data interface{}
}

//PersistentFrame is just a fake object
type PersistentFrame struct {
	gorm.Model
	Stack string
	Data  []byte
}

//Serialize a frame into bytes
func (f *Frame) Serialize() ([]byte, error) {
	var buf bytes.Buffer
	err := gob.NewEncoder(&buf).Encode(f)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

//Deserialize bytes into a Frame
func (f *Frame) Deserialize(Data []byte) (*Frame, error) {
	err := gob.NewDecoder(bytes.NewReader(Data)).Decode(f)
	if err != nil {
		return nil, err
	}
	return f, nil
}
