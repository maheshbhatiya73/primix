package main

import (
	"fmt"
	"primix/cmd"
	"primix/logger"
	"primix/utils"
)

func main() {
	logger.Init()
	utils.PrintBanner() // ✅ Ensure the banner prints
	fmt.Println("🚀 Running Primix CLI...") // Debugging output
	cmd.Execute()
}
