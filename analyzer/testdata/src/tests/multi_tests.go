package tests

import (
	"log"
	"log/slog"
)

func multiple_violations() {
    log.Println("User password: secret") // want "Log messages must start with a lowercase letter\n" "Log messages must not contain special characters or emojis\n" "Log messages must not contain sensitive data\n"
    slog.Info("API key = 12345") // want "Log messages must start with a lowercase letter\n" "Log messages must not contain special characters or emojis\n" "Log messages must not contain sensitive data\n"
    
    log.Println("Запуск сервера!") // want "Log messages must start with a lowercase letter\n" "Log messages must be in English\n" "Log messages must not contain special characters or emojis\n"
    
    log.Println("пароль: secret!!!") // want "Log messages must be in English\n" "Log messages must not contain special characters or emojis\n" "Log messages must not contain sensitive data\n"
    slog.Info("секретный token!") // want "Log messages must be in English\n" "Log messages must not contain special characters or emojis\n" "Log messages must not contain sensitive data\n"
    
    log.Println("Secret Password: 12345!") // want "Log messages must start with a lowercase letter\n" "Log messages must not contain sensitive data\n" "Log messages must not contain special characters or emojis\n"
    log.Println("SECRET Key: 12345! 🤫") // want "Log messages must start with a lowercase letter\n" "Log messages must not contain sensitive data\n" "Log messages must not contain special characters or emojis\n"
    
    log.Println("user+token") // want "Log messages must not contain special characters or emojis\n" "Log messages must not contain sensitive data\n"
    log.Println("api_key") // want "Log messages must not contain special characters or emojis\n" "Log messages must not contain sensitive data\n"
    log.Println("my-key") // want "Log messages must not contain special characters or emojis\n" "Log messages must not contain sensitive data\n"
    log.Println("token=") // want "Log messages must not contain special characters or emojis\n" "Log messages must not contain sensitive data\n"
    
    password := "12345678"
    log.Println("your password: " + password) // want "Log messages must not contain sensitive data\n" "Log messages must not contain special characters or emojis\n"

    log.Println("12345!") // want "Log messages must start with a lowercase letter\n" "Log messages must not contain special characters or emojis\n"
}