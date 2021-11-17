package tests

import (
	"github.com/TinyRogue/lembook-serv/internal/models"
	"testing"
)

func TestIsPasswordValid(t *testing.T) {
	passwordTests := []struct {
		password string
		ans      bool
		desc     string
	}{
		{password: "", ans: false, desc: "password doesn't meet any of requirements"},
		{password: "Asd1@3e", ans: false, desc: "not enough characters"},
		{password: "@asdqweas@3e", ans: false, desc: "no capital letter"},
		{password: "@QWFJIHSFIASNDBIF1@3", ans: false, desc: "not lowercase letter"},
		{password: "AaQq@!12as", ans: true, desc: "valid password"},
	}
	for _, tt := range passwordTests {
		t.Run(tt.password, func(t *testing.T) {
			got := models.IsPasswordValid(tt.password)
			if got != tt.ans {
				t.Errorf("%#v got %v want %v. Desc: %v", tt, got, tt.ans, tt.desc)
			}
		})
	}
}
