package utils

import (
	"github.com/a-h/templ"
	"github.com/theSamuelIyeh/johnethel-school-rms/internals/database"
)

type ApiCfg struct {
	DB *database.Queries
}

type TemplateStruct struct {
	Title     string
	Page      templ.Component
	Component templ.Component
}
