package create

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"primix/utils"
)

// Project generates the folder structure and installs templates
func Project(projectName string) {
	utils.Info("Creating new Primix project: " + projectName)

	// Create main project folder
	projectPath := "./" + projectName
	err := os.Mkdir(projectPath, 0755)
	if err != nil {
		utils.Error("Failed to create project directory: " + err.Error())
		return
	}

	// List of directories to create
	dirs := []string{
		"cmd",
		"config",
		"internal",
		"pkg",
	}

	// Create directories
	for _, dir := range dirs {
		path := filepath.Join(projectPath, dir)
		os.Mkdir(path, 0755)
	}

	// Copy template files
	copyTemplateFiles(projectPath)

	utils.Success("Project " + projectName + " created successfully! 🎉")
	utils.Info("Run `cd " + projectName + "` to get started.")
}

// Copy template files from the "templates" folder
func copyTemplateFiles(destPath string) {
	templateDir := "./templates/base"

	files, err := ioutil.ReadDir(templateDir)
	if err != nil {
		utils.Error("Failed to read template directory: " + err.Error())
		return
	}

	for _, file := range files {
		src := filepath.Join(templateDir, file.Name())
		dest := filepath.Join(destPath, file.Name())

		data, err := ioutil.ReadFile(src)
		if err != nil {
			utils.Error("Failed to copy " + file.Name() + ": " + err.Error())
			continue
		}

		err = ioutil.WriteFile(dest, data, 0644)
		if err != nil {
			utils.Error("Failed to write " + file.Name() + ": " + err.Error())
		}
	}
	utils.Success("Templates installed successfully! ✅")
}
