package main

import (
	"fmt"
	"primix/cmd"
	"primix/logger"
	"primix/utils"
)

func main() {
	logger.Init()
	utils.PrintBanner() 
	fmt.Println("🚀 Running Primix CLI...") 
	cmd.Execute()
}
