package usecase

import (
	"http-callback/helper"
	"http-callback/svcutil/cmd"

	"github.com/go-redis/redis"
)

type UC struct {
	Helper helper.Helper
	Bash   cmd.Terminal
	Redis  *redis.Client
}
