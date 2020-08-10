package helpers

import (
    "gopkg.in/yaml.v3"
    "os"
)

func LoadConfig(filePath string, dec interface{}) error {
    file, err := os.Open(filePath)
    if err != nil {
        return err
    }
    defer file.Close()

    decoder := yaml.NewDecoder(file)
    err = decoder.Decode(dec)
    if err != nil {
        return err
    }
}
