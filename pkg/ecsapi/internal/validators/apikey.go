package validators

func IsValidApiKey(apikey string) bool {
	return apikey != ""
}
