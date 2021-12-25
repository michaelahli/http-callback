package helper

import "github.com/joho/godotenv"

func (c *helper) SetUp() {
	godotenv.Load()
}
