package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func Convert(data string) (string, error) {
	data = strings.TrimSpace(data)
	if data == "" {
		return "", errors.New("empty input")
	}

	if isMorse(data) {
		return morse.ToText(data), nil
	}

	return morse.ToMorse(data), nil
}

func isMorse(data string) bool {
	return !strings.ContainsFunc(data, func(r rune) bool {
		return r != '.' && r != '-' && !isSpace(r)
	})
}

func isSpace(r rune) bool {
	return r == ' ' || r == '\n' || r == '\r' || r == '\t'
}
