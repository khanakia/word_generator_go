package current

import (
	"errors"
	"fmt"
	"os"
	"runtime"
)

// Filename is the __filename equivalent
func Filename() (string, error) {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		return "", errors.New("unable to get the current filename")
	}
	return filename, nil
}

// Dirname is the __dirname equivalent
func Dirname() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error retrieving directory:", err)
		return "", err
	}
	// fmt.Println("Directory where 'go run .' was executed:", dir)

	return dir, nil

	// filename, err := Filename()
	// if err != nil {
	// 	return "", err
	// }
	// fmt.Println(filename)
	// return filepath.Dir(filename), nil
}
