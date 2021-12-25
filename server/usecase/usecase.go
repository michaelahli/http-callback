package usecase

import (
	"http-callback/helper"
	"http-callback/svcutil/cmd"
)

type UC struct {
	Helper helper.Helper
	Bash   cmd.Terminal
}
