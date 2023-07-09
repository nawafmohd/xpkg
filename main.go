package xpkg

import "log"

func LogInfo(msg string) {
    log.Printf("INFO - %v", msg)
}

func LogWarning(msg string) {
    log.Printf("WARN - %v", msg)
}

func LogError(msg string) {
    log.Printf("ERROR - %v", msg)
}