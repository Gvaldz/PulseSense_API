package core

import (
	"os"
)

type FCMConfig struct {
	CredentialsPath string
	ProjectID       string
}

type AppConfig struct {
	FCM FCMConfig
}

var Config AppConfig

func LoadConfig() {
	Config.FCM.CredentialsPath = os.Getenv("FCM_CREDENTIALS_PATH")
	Config.FCM.ProjectID = os.Getenv("FCM_PROJECT_ID")

	if Config.FCM.CredentialsPath == "" {
		Config.FCM.CredentialsPath = "src/internal/services/fcm/serviceAccountKey.json"
	}

	if Config.FCM.ProjectID == "" {
		Config.FCM.ProjectID = "rodismart-77622"
	}

}
