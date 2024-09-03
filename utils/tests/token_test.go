package tests

import (
	"event-booking/utils"

	"testing"
)

func Test_Generate_Token(t *testing.T) {
	token, _ := utils.GenerateToken("example@gmail.com", 2)
	if token == "" {
		t.Error("Token Not Generated")
	}
}
