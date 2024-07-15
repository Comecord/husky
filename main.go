package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run version.go [command]")
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "version":
		version, err := os.ReadFile("./VERSION")
		if err != nil {
			fmt.Println("Error reading VERSION file:", err)
			os.Exit(1)
		}
		fmt.Println(strings.TrimSpace(string(version)))
	case "patch":
		version, err := os.ReadFile("./VERSION")
		if err != nil {
			fmt.Println("Error reading VERSION file:", err)
			os.Exit(1)
		}
		currentVersion := strings.TrimSpace(string(version))
		parts := strings.Split(currentVersion, ".")
		if len(parts) != 3 {
			fmt.Println("Invalid version format in VERSION file")
			os.Exit(1)
		}
		patch := parts[2]
		newPatchInt, err := strconv.Atoi(patch)
		if err != nil {
			fmt.Println("Error parsing patch version:", err)
			os.Exit(1)
		}
		newPatch := fmt.Sprintf("%d", newPatchInt+1)
		newVersion := fmt.Sprintf("%s.%s.%s", parts[0], parts[1], newPatch)
		err = os.WriteFile("./VERSION", []byte(newVersion), 0644)
		if err != nil {
			fmt.Println("Error writing VERSION file:", err)
			os.Exit(1)
		}
		fmt.Printf("Updated version to: %s\n", newVersion)

		// Добавляем файл VERSION в git
		gitAddCmd := exec.Command("git", "add", "VERSION")
		err = gitAddCmd.Run()
		if err != nil {
			fmt.Println("Error adding VERSION file to git:", err)
			os.Exit(1)
		}
	default:
		fmt.Println("Unknown command:", command)
		os.Exit(1)
	}
}
