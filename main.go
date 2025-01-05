package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
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

	translateCmd := &cobra.Command{
		Use: "translate",
		RunE: func(cmd *cobra.Command, args []string) error {
			fromLang, err := cmd.Flags().GetString("from")
			if err != nil {
				return fmt.Errorf("cannot read the flag for 'from': %w", err)
			}
			toLang, err := cmd.Flags().GetString("to")
			if err != nil {
				return fmt.Errorf("cannot read the flag for 'to': %w", err)
			}

			if len(args) == 0 {
				return fmt.Errorf("не указан текст для перевода. Использование: cli-translator translate -from=en -to=ru \"Hello world\"")
			}

			mainText := strings.Join(args, " ")

			result := fmt.Sprintf("%s, %s, %s", fromLang, toLang, mainText)
			fmt.Println(result)

			return nil
		},
	}

	translateCmd.Flags().StringP("from", "f", "en", "Исходный язык (например, en)")
	translateCmd.Flags().StringP("to", "t", "ru", "Целевой язык (например, ru)")

	rootCmd.AddCommand(translateCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}

}
