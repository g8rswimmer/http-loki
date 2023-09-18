package mock

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/g8rswimmer/http-loki/internal/model"
	"github.com/gorilla/mux"
)

func AddRoutesFromDirectory(dir string, router *mux.Router) error {
	if len(dir) == 0 {
		return nil
	}

	ep := endpoints{}

	err := filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
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

	return err

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

func add(m *model.Mock, ep endpoints, route *mux.Router) {
	if has := ep.add(m); !has {
		fmt.Println(m.Method, m.Endpoint)
		route.Methods(m.Method).Path(m.Endpoint).HandlerFunc(ep.handler(m).HTTPHandler)
	}
}
