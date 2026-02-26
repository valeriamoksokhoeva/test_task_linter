package tests

import (
	"log"
	"log/slog"
)

func multiple_violations() {
    // Большая буква + спецсимволы + чувствительные данные
    log.Println("User password: secret") // want "Log messages must start with a lowercase letter\n" "Log messages must not contain special characters or emojis\n" "Log messages must not contain sensitive data\n"
    slog.Info("API key = 12345") // want "Log messages must start with a lowercase letter\n" "Log messages must not contain special characters or emojis\n" "Log messages must not contain sensitive data\n"
    
    // Большая буква + спецсимволы + не английский
    log.Println("Запуск сервера!") // want "Log messages must start with a lowercase letter\n" "Log messages must be in English\n" "Log messages must not contain special characters or emojis\n"
    
    // Спецсимволы + не английский + чувствительные данные
    log.Println("пароль: secret!!!") // want "Log messages must be in English\n" "Log messages must not contain special characters or emojis\n" "Log messages must not contain sensitive data\n"
    slog.Info("секретный token!") // want "Log messages must be in English\n" "Log messages must not contain special characters or emojis\n" "Log messages must not contain sensitive data\n"
    
    // Большая буква + спецсимволы + чувствительные данные
    log.Println("Secret Password: 12345!") // want "Log messages must start with a lowercase letter\n" "Log messages must not contain sensitive data\n" "Log messages must not contain special characters or emojis\n"
    log.Println("SECRET Key: 12345! 🤫") // want "Log messages must start with a lowercase letter\n" "Log messages must not contain sensitive data\n" "Log messages must not contain special characters or emojis\n"
    
    // Спецсимволы + чувствительные данные (ключевые слова)
    log.Println("user+token") // want "Log messages must not contain special characters or emojis\n" "Log messages must not contain sensitive data\n"
    log.Println("api_key") // want "Log messages must not contain special characters or emojis\n" "Log messages must not contain sensitive data\n"
    log.Println("my-key") // want "Log messages must not contain special characters or emojis\n" "Log messages must not contain sensitive data\n"
    log.Println("token=") // want "Log messages must not contain special characters or emojis\n" "Log messages must not contain sensitive data\n"
    
    // Конкатенация с нарушениями
    password := "12345678"
    log.Println("your password: " + password) // want "Log messages must not contain sensitive data\n" "Log messages must not contain special characters or emojis\n"

    // Цифры + спецсимволы
    log.Println("12345!") // want "Log messages must start with a lowercase letter\n" "Log messages must not contain special characters or emojis\n"
    
}