package mux

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/g8rswimmer/http-loki/internal/model"
	"github.com/gorilla/mux"
)

func loadFromDir(dir string, router *mux.Router) {

	ep := endpoints{}

	_ = filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		switch {
		case err != nil:
			return err
		case info.IsDir():
			return nil
		default:
		}

		m, err := loadMockEndpoint(path, info)
		if err != nil {
			return err
		}
		add(m, ep, router)
		return nil
	})

}

func loadMockEndpoint(path string, info fs.FileInfo) (*model.Mock, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("load mock handler file open [%s] %w", path, err)
	}
	defer f.Close()

	ep := &model.Mock{}
	if err := json.NewDecoder(f).Decode(ep); err != nil {
		return nil, fmt.Errorf("load mock handler file [%s] decode %w", path, err)
	}
	return ep, nil
}
