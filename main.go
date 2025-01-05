package main

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := cobra.Command{
		Use:   "trango",
		Short: "trango — простой переводчик с одного языка на другой",
		Long: `trango — это утилита командной строки,
				которая позволяет переводить текст с одного языка на другой.
				Для примера показана псевдо-функция перевода.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Добро пожаловать в CLI-translator!")
			fmt.Println("Попробуйте команду 'cli-translator translate --help'")
		},
	}
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
