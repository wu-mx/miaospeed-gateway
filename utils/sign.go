package utils

import "github.com/miaokobot/miaospeed/interfaces"

func Resign(sr *interfaces.SlaveRequest, slavetoken string, buildtoken string) {
	SignRequest(slavetoken, sr, buildtoken)
	sr.Challenge = SignRequest(slavetoken, sr, buildtoken)
}
