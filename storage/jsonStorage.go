package storage

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func JsonStorage(path string, data interface{}) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	absPath, err := filepath.Abs(path)
	if err != nil {
		return err
	}

	err = os.WriteFile(absPath, jsonData, 0644)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.MkdirAll(filepath.Dir(absPath), 0755)
			if err != nil {
				return err
			}
			err = os.WriteFile(absPath, jsonData, 0644)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}

	return nil
}
