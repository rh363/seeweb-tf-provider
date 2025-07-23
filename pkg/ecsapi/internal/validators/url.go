package validators

import (
	"regexp"
)

const (
	urlRegex = "" +
		`^https?://([a-zA-Z0-9\-._~%!$&'()*+,;=]+@)?` +
		`((\[[0-9a-fA-F:.]+\])|([a-zA-Z0-9\-._~%]+))` + // IPv6 (in brackets) or hostname
		`(:\d{1,5})?` + // optional port
		`(/[^\s?#]*)?(\?[^\s#]*)?(#[^\s]*)?$` // optional path, query, fragment

)

func IsValidUrl(url string) bool {
	regex := regexp.MustCompile(urlRegex)
	return regex.MatchString(url)
}
