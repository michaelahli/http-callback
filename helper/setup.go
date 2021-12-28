package helper

import "github.com/joho/godotenv"

func (c *helper) SetUp(envPath string) {
	godotenv.Load(envPath)
}
