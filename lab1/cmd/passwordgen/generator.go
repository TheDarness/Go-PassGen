package main

import (
	"crypto/rand"
	"errors"
	"math/big"
)

func Generate(opts Config) (string, error) {
	if opts.Length < 4 || opts.Length > 128 {
		return "", errors.New("password length must be between 4 and 128")
	}

	if !opts.Lowercase && !opts.Uppercase && !opts.Digits && !opts.Specials {
		return "", errors.New("at least one character category must be selected")
	}

	var pool string
	categories := []struct {
		enabled bool
		chars   string
	}{
		{opts.Lowercase, "abcdefghijklmnopqrstuvwxyz"},
		{opts.Uppercase, "ABCDEFGHIJKLMNOPQRSTUVWXYZ"},
		{opts.Digits, "0123456789"},
		{opts.Specials, "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"},
	}

	selectedCategories := 0
	for _, cat := range categories {
		if cat.enabled {
			pool += cat.chars
			selectedCategories++
		}
	}

	if opts.Length < selectedCategories {
		return "", errors.New("password length is too short for selected categories")
	}

	password := make([]byte, opts.Length)

	index := 0
	for _, cat := range categories {
		if cat.enabled {
			char, err := randomChar(cat.chars)
			if err != nil {
				return "", err
			}
			password[index] = char
			index++
		}
	}

	for i := index; i < opts.Length; i++ {
		char, err := randomChar(pool)
		if err != nil {
			return "", err
		}
		password[i] = char
	}

	for i := len(password) - 1; i > 0; i-- {
		j, err := rand.Int(rand.Reader, big.NewInt(int64(i+1)))
		if err != nil {
			return "", err
		}
		password[i], password[j.Int64()] = password[j.Int64()], password[i]
	}

	return string(password), nil
}

func randomChar(chars string) (byte, error) {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
	if err != nil {
		return 0, err
	}
	return chars[n.Int64()], nil
}
