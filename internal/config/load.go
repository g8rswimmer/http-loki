package config

import "errors"

func LoadValues() (*Values, error) {
	v, err := loadFlags()
	if err != nil {
		return nil, err
	}

	if v.MockDir == "" {
		return nil, errors.New("mock_dir needs to be present")
	}
	return v, nil
}
