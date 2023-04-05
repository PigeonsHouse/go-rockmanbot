package utils

import (
	"github.com/joho/godotenv"
	"os"
)

var (
	InstanceUrl string
	AccessToken string
)

func LoadDotEnv() {
	err := godotenv.Load(".env") // envファイルのパスを渡す。何も渡さないと、どうディレクトリにある、.envファイルを探す
	if err != nil {
		panic("Error loading .env file")
	}
	InstanceUrl = os.Getenv("INSTANCE_URL")
	AccessToken = os.Getenv("ACCESS_TOKEN")
}
