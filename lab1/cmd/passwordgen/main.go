package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

type Config struct {
	Length    int  `json:"length"`
	Lowercase bool `json:"lowercase"`
	Uppercase bool `json:"uppercase"`
	Digits    bool `json:"digits"`
	Specials  bool `json:"specials"`
}

func main() {
	var length int
	var lower, upper, digits, specials bool
	var configFile string
	flag.IntVar(&length, "length", 16, "Password length")
	flag.BoolVar(&lower, "lower", true, "Use lowercase letters")
	flag.BoolVar(&upper, "upper", true, "Use uppercase letters")
	flag.BoolVar(&digits, "digits", true, "Use digits")
	flag.BoolVar(&specials, "specials", false, "Use special characters")
	flag.StringVar(&configFile, "config", "", "Path to JSON config file")
	flag.Parse()
	if configFile != "" {
		file, err := os.ReadFile(configFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading config file: %v\n", err)
			os.Exit(1)
		}
		var fileConfig Config
		if err := json.Unmarshal(file, &fileConfig); err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing config file: %v\n", err)
			os.Exit(1)
		}
		if fileConfig.Length > 0 {
			length = fileConfig.Length
		}
		if fileConfig.Lowercase {
			lower = fileConfig.Lowercase
		}
		if fileConfig.Uppercase {
			upper = fileConfig.Uppercase
		}
		if fileConfig.Digits {
			digits = fileConfig.Digits
		}
		if fileConfig.Specials {
			specials = fileConfig.Specials
		}
	}
	opts := Config{
		Length:    length,
		Lowercase: lower,
		Uppercase: upper,
		Digits:    digits,
		Specials:  specials,
	}

	password, err := Generate(opts)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error generating password: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(password)
}
