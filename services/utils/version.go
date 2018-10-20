package utils

// Version get api version
func Version() string {
	config := GetConfig()
	return config.GetVersion()
}
