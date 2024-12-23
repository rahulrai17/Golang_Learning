package config

import "html/template"

// This is a file that is accessed by whole project. This is also know as (application wide configuration).
// We only add most useful and valid configuration here.
// This file is imported by whole application
// Remenber this should be program in a way that this doesn't import anything from the application.
// It is only allowed to import form standard libraries only (This precaution is taken to avoid the import cycle. Import cycle doesn't allow the application to be compiled).

// AppConfig holds the application config
type AppConfig struct{
  UseCache bool
  TemplateCache map[string]*template.Template
}