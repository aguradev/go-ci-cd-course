package configs

import "github.com/joho/godotenv"

func LoadEnv() {

	envException := godotenv.Load()

	if envException != nil {
		panic("error load env")
	}

}
