package utils

import uuid "github.com/satori/go.uuid"

//生成uuid
func GenerateUUid() string {
	uid := uuid.NewV4()
	sessionId := uid.String()
	return sessionId
}