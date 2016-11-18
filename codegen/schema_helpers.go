package codegen

import (
	"log"
	"strings"
)

func (sp *SystemProperty) GoType() string {
	switch {
	case sp.Type == "string" || sp.Type == "int":
		return sp.Type
	case sp.Type == "string array":
		return "[]string"
	case sp.Type == "string selector":
		return "string"
	}

	log.Fatalf("unsupported spec property type: %v",sp.Type)
	return ""
}

func (sp *SystemProperty) GoMethodTypeName() string {
	switch {
	case sp.Type == "string" || sp.Type == "int":
		return strings.Title(sp.Type)
	case sp.Type == "string array":
		return "StringArray"
	case sp.Type == "string selector":
		return "String"
	}
	log.Fatalf("unsupported spec property type: '%s'",sp.Type)
	return ""
}
