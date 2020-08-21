package secrets

import (
	"errors"
	"io/ioutil"
	"os"
)

// GetSecret retrieve secret value with provided name.
// If the secret doesn´t exists and error is returned
func GetSecret(name string) (string, error) {
	if pass, exists := os.LookupEnv(name); exists {
		return pass, nil
	}

	if filename, exists := os.LookupEnv(name + "_FILE"); exists {
		file, err := ioutil.ReadFile(filename)
		if err != nil {
			return "", errors.New("File " + filename + " cannot be open")
		}
		return string(file), nil
	}

	return "", errors.New("Secret " + name + " not provided")
}
