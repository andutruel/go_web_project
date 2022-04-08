package store

import (
	"encoding/json"
	"os"
)

//para transformar los datos del archivo en un objeto JSON

//creo la interfaz store
type Store interface {
	Read(data interface{}) error
	Write(data interface{}) error
}

//creo el tipo de dato custom Type para limitar lo que puede recibir New
type Type string

//constantes con los tipos de archivo
const (
	FileType Type = "file"
)

//Factory constructor de storage según el tipo de archivo
func New(store Type, fileName string) Store {
	switch store {
	case FileType:
		return &FileStore{fileName}
	}
	return nil
}

//estructura del store de archivos con un campo que guarda el nombre del archivo
type FileStore struct {
	FileName string
}

//funciones para que implemente la interfaz Store
func (fs *FileStore) Write(data interface{}) error {
	file, err := json.MarshalIndent(data, "", " ")

	if err != nil {
		return err
	}

	return os.WriteFile(fs.FileName, file, 0644)
}

func (fs *FileStore) Read(data interface{}) error {
	file, err := os.ReadFile(fs.FileName)

	if err != nil {
		return err
	}

	return json.Unmarshal(file, &data)
}
