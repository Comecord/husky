package version

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func GetVCS() (string, error) {
	file, err := os.Open("./VERSION")
	if err != nil {
		log.Fatal("Error opening file:", err)
		return "", err
	}
	defer file.Close()

	// Read only the first line from the file
	var version string
	_, err = fmt.Fscanf(file, "%s\n", &version)
	if err != nil {
		log.Fatal("Error reading file:", err)
		return "", err
	}

	// Trim any leading/trailing whitespace

	version = strings.TrimSpace(version)
	version = strings.TrimSuffix(version, "\n")

	return version, nil
}
