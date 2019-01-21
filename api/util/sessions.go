package util

// ExtractSession extracts the session ID from the given URI.
func ExtractSession(str, trim string) string {
	return str[len(trim):]
}
