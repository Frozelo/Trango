package main

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

func main() {
	testCmd := cobra.Command{
		Use: "test",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello")
		},
	}
	if err := testCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
