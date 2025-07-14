package storage

import (
	"encoding/json"
	"fmt"
	"os"
)

func LoadFile[T any](pathEnv, filename string) ([]T, error) {
	// take path from .env
	path := os.Getenv(pathEnv)
	if path == "" {
		return nil, fmt.Errorf("%v not set", pathEnv)
	}

	// open file
	file, err := os.Open(path + filename)
	if err != nil {
		return nil, fmt.Errorf("there is an error occured while opening file (%v) : %v", filename, err)
	}

	// decode file
	var values []T
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&values)
	if err != nil {
		return nil, fmt.Errorf("an error occured while decoding file (%v) : %v", filename, err)
	}

	// return
	return values, nil
}

func SaveFile[T any](pathEnv, filename string, values T) error {
	// take path from .env
	path := os.Getenv(pathEnv)
	if path == "" {
		return fmt.Errorf("%v not set", pathEnv)
	}

	// create folder if not exists
	err := os.MkdirAll(path, 0755)
	if err != nil {
		return fmt.Errorf("an error occurred while creating assets folder : %v", err)
	}

	// create file
	file, err := os.Create(path + filename)
	if err != nil {
		return fmt.Errorf("an error occured while creating %v file", filename)
	}
	defer file.Close()

	// encode file
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(values)
	if err != nil {
		return fmt.Errorf("an error occured while encoding file (%v) : %v", filename, err)
	}

	// return
	return nil
}
