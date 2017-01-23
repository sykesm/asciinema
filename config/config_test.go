package config_test

import (
	"testing"

	"github.com/sykesm/asciinema/config"
)

func TestConfig_ApiUrl(t *testing.T) {
	var tests = []struct {
		cfg      config.ConfigFile
		env      map[string]string
		expected string
	}{
		{
			config.ConfigFile{},
			map[string]string{},
			"https://asciinema.org",
		},
		{
			config.ConfigFile{API: config.ConfigAPI{URL: "https://asciinema.example.com"}},
			map[string]string{},
			"https://asciinema.example.com",
		},
		{
			config.ConfigFile{API: config.ConfigAPI{URL: "https://asciinema.example.com"}},
			map[string]string{"ASCIINEMA_API_URL": "http://localhost:3000"},
			"http://localhost:3000",
		},
	}

	for _, test := range tests {
		cfg := config.Config{&test.cfg, test.env}
		actual := cfg.ApiUrl()

		if actual != test.expected {
			t.Errorf(`expected "%v", got "%v"`, test.expected, actual)
		}
	}
}

func TestConfig_ApiToken(t *testing.T) {
	var tests = []struct {
		cfg      config.ConfigFile
		expected string
	}{
		{
			config.ConfigFile{},
			"",
		},
		{
			config.ConfigFile{API: config.ConfigAPI{Token: "foo"}},
			"foo",
		},
		{
			config.ConfigFile{User: config.ConfigUser{Token: "foo"}},
			"foo",
		},
	}

	for _, test := range tests {
		cfg := config.Config{&test.cfg, nil}
		actual := cfg.ApiToken()

		if actual != test.expected {
			t.Errorf(`expected "%v", got "%v"`, test.expected, actual)
		}
	}
}

func TestConfig_RecordCommand(t *testing.T) {
	var tests = []struct {
		cfg      config.ConfigFile
		env      map[string]string
		expected string
	}{
		{
			config.ConfigFile{},
			map[string]string{},
			"/bin/sh",
		},
		{
			config.ConfigFile{},
			map[string]string{"SHELL": "/bin/bash"},
			"/bin/bash",
		},
		{
			config.ConfigFile{Record: config.ConfigRecord{Command: "foo -l"}},
			map[string]string{"SHELL": "/bin/bash"},
			"foo -l",
		},
	}

	for _, test := range tests {
		cfg := config.Config{&test.cfg, test.env}
		actual := cfg.RecordCommand()

		if actual != test.expected {
			t.Errorf(`expected "%v", got "%v"`, test.expected, actual)
		}
	}
}

func TestConfig_RecordMaxWait(t *testing.T) {
	var tests = []struct {
		cfg      config.ConfigFile
		expected float64
	}{
		{
			config.ConfigFile{},
			0,
		},
		{
			config.ConfigFile{Record: config.ConfigRecord{MaxWait: 1.23456}},
			1.23456,
		},
	}

	for _, test := range tests {
		cfg := config.Config{&test.cfg, nil}
		actual := cfg.RecordMaxWait()

		if actual != test.expected {
			t.Errorf(`expected "%v", got "%v"`, test.expected, actual)
		}
	}
}

func TestConfig_RecordYes(t *testing.T) {
	var tests = []struct {
		cfg      config.ConfigFile
		expected bool
	}{
		{
			config.ConfigFile{},
			false,
		},
		{
			config.ConfigFile{Record: config.ConfigRecord{Yes: true}},
			true,
		},
	}

	for _, test := range tests {
		cfg := config.Config{&test.cfg, nil}
		actual := cfg.RecordYes()

		if actual != test.expected {
			t.Errorf(`expected "%v", got "%v"`, test.expected, actual)
		}
	}
}

func TestConfig_PlayMaxWait(t *testing.T) {
	var tests = []struct {
		cfg      config.ConfigFile
		expected float64
	}{
		{
			config.ConfigFile{},
			0,
		},
		{
			config.ConfigFile{Play: config.ConfigPlay{MaxWait: 1.23456}},
			1.23456,
		},
	}

	for _, test := range tests {
		cfg := config.Config{&test.cfg, nil}
		actual := cfg.PlayMaxWait()

		if actual != test.expected {
			t.Errorf(`expected "%v", got "%v"`, test.expected, actual)
		}
	}
}
