package websitepathconverter

import (
	"context"
	"fmt"
	"net/http"
	"path"
)

// Config the plugin configuration.
type Config struct {
	Directory string `json:"directory"`
}

// CreateConfig creates the default plugin configuration.
func CreateConfig() *Config {
	return &Config{
		Directory: "",
	}
}

// WebsitePathConverter a WebsitePathConverter plugin.
type WebsitePathConverter struct {
	directory string
	next      http.Handler
	name      string
}

// New created a new WebsitePathConverter plugin.
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	if len(config.Directory) == 0 {
		return nil, fmt.Errorf("Must provide a directory")
	}

	return &WebsitePathConverter{
		directory: config.Directory,
		next:      next,
		name:      name,
	}, nil
}

func (a *WebsitePathConverter) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if path.Ext(req.URL.Path) == "" {
		req.URL.Path = path.Join(a.directory, "index.html")
	} else {
		req.URL.Path = path.Join(a.directory, req.URL.Path)
	}
	a.next.ServeHTTP(rw, req)
}
