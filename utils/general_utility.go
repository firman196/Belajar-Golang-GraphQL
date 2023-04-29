package utils

import "github.com/google/uuid"

//if else statements in one line condition
func IfThenElse(condition bool, isTrue interface{}, isFalse interface{}) interface{} {
	if condition {
		return isTrue
	}
	return isFalse
}

func Uuid() string {
	uuid := uuid.New()
	return uuid.String()
}
