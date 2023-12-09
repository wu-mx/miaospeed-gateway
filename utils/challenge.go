package utils

import (
	"crypto/sha512"
	"encoding/base64"
	jsoniter "github.com/json-iterator/go"
	"github.com/miaokobot/miaospeed/interfaces"
	"strings"
)

//https://github.com/miaokobot/miaospeed/blob/master/utils/challenge.go

func hashMiaoSpeed(token, request string, buildtoken string) string {
	buildTokens := append([]string{token}, strings.Split(strings.TrimSpace(buildtoken), "|")...)

	hasher := sha512.New()
	hasher.Write([]byte(request))

	for _, t := range buildTokens {
		if t == "" {
			// unsafe, plase make sure not to let token segment be empty
			t = "SOME_TOKEN"
		}

		hasher.Write(hasher.Sum([]byte(t)))
	}

	return base64.URLEncoding.EncodeToString(hasher.Sum(nil))
}

func SignRequest(token string, req *interfaces.SlaveRequest, buildtoken string) string {
	awaitSigned := req.Clone()
	awaitSigned.Challenge = ""
	awaitSignedStr, _ := jsoniter.MarshalToString(&awaitSigned)
	awaitSignedStr = strings.TrimSpace(awaitSignedStr)
	return hashMiaoSpeed(token, awaitSignedStr, buildtoken)
}

func VerifyRequest(req *interfaces.SlaveRequest, token string, buildtoken string) bool {
	return req.Challenge == SignRequest(token, req, buildtoken)
}
