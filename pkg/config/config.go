package config

import (
	"html/template"
	"log"
)

// AppConfig holds that application config
type AppConfig struct {
	useCache bool
	TemplateCache map[string]*template.Template
	InfoLog *log.Logger
}
