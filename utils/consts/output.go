package consts

import (
	"strings"
)

type OutputType string

// OutputType:

var (
	Default            = OutputType(DefaultCode.EN())
	JSON    OutputType = "json"
	PNG     OutputType = "png"
	JPG     OutputType = "jpg"
	JPEG    OutputType = "jpeg"
	SVG     OutputType = "svg"
	Base64  OutputType = "base64"
	Text    OutputType = "text"
)

func (o OutputType) String() string {
	return string(o)
}

func IsOutputType(in string, ot OutputType) bool {
	return strings.ToLower(in) == ot.String()
}

func CaseOutputType(in string, ot OutputType) string {
	if IsOutputType(in, ot) {
		return in
	}
	return Default.String()
}

func ToOutputType(in string) OutputType {
	return OutputType(in)
}