package cmd

import (
	"os"
	"primix/utils"
	"primix/cmd/create"
)

// Show CLI help menu
func ShowHelp() {
	utils.PrintBanner()
	utils.Info("Usage:")
	utils.Info("  primix create-app <project-name>  - Create a new Primix project")
	utils.Info("  primix version                   - Show CLI version")
	utils.Info("  primix help                      - Show this help menu")
	utils.Info("")
	utils.Success("🚀 Get started with Primix today!")
}

// Execute runs the CLI
func Execute() {
	if len(os.Args) < 2 {
		ShowHelp()
		return
	}

	command := os.Args[1]

	switch command {
	case "create-app":
		if len(os.Args) < 3 {
			utils.Error("Please provide a project name: 'primix create-app myapp'")
			return
		}
		projectName := os.Args[2]
		create.Project(projectName)

	case "help":
		ShowHelp()

	case "version":
		utils.Info("Primix CLI version 1.0")

	default:
		utils.Error("Unknown command: " + command)
		ShowHelp()
	}
}
