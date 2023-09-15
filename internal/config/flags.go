package config

import (
	"flag"
	"fmt"
	"strconv"
)

func loadFlags() (*Values, error) {
	port := flag.String("port", "8080", "mock server port")
	mockDir := flag.String("mock_dir", "", "mock request directory")

	flag.Parse()

	if _, err := strconv.Atoi(*port); err != nil {
		return nil, fmt.Errorf("port string convert: %w", err)
	}
	return &Values{
		Port:    *port,
		MockDir: *mockDir,
	}, nil
}
