package codegen

import (
	"strings"
	"regexp"
	"text/template"
	//"log"
)

var TemplateFuncs = template.FuncMap{
	"replace":          ReplaceString,
	"ccUpper":       CamelCaseUpper,
	"ccLower":       CamelCaseLower,
	"title":            strings.Title,
	"toLower":          strings.ToLower,
	"toUpper":          strings.ToUpper,
	"trim":             strings.Trim,
	"underscoreSpaces":  UnderscoreSpaces,
	"underscoreDashes":  UnderscoreDashes,
	"trimSpace":	strings.TrimSpace,
	"nonWordRegex":		RemoveNonWord,
	"formatPropValue":	FormatPropValue,
	"stripFromName":  StripFromName,
}

var nonWordRegex = regexp.MustCompile(`[^[:word:]-]`)
var lowerCaseRegex = regexp.MustCompile(`[a-z]`)

func RemoveNonWord(s string) string {
	return nonWordRegex.ReplaceAllString(s,"")
}

func StripFromName(strip, s string) string {
	if len(strip) >= len(s) {
		return s
	}
	start := strings.Index(strings.ToLower(s),strings.ToLower(strip))
	if ( start < 0) {
		return s
	}
	end := start+len(strip)
	s2 := s[:start] + s[end:]
	s2 = strings.Trim(s2," ")
	if len(s2) >= 2 {
		return s2
	} else {
		return s
	}
}

func FormatPropValue(s string) string {
	if ! lowerCaseRegex.MatchString(s) {
		return RemoveNonWord(UnderscoreDashes(s))
	} else {
		return RemoveNonWord(CamelCaseUpper(UnderscoreDashes(s)))
	}
}

func UnderscoreSpaces(s string) string {
	return strings.Replace(s, " ", "_", -1)
}

func UnderscoreDashes(s string) string {
	return strings.Replace(s, "-", "_", -1)
}

func CamelCaseUpper(src string) string {
	s := strings.Split(src, "_")
	for i, _ := range s {
		s[i] = strings.Title(s[i])
	}
	return strings.Join(s, "")
}

func CamelCaseLower(src string) string {
	s := strings.Split(src, "_")
	for i, _ := range s {
		if i == 0 {
			s[i] = strings.ToLower(s[i])
		} else {
			s[i] = strings.Title(s[i])
		}
	}
	return strings.Join(s, "")
}

func ReplaceString(s, old, new string) string {
	return strings.Replace(s, old, new, -1)
}

