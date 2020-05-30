// +build darwin linux

package filesystem

// GetSettingsPath : get settings path
func GetSettingsPath() string {
	userHome := ExpandHome("~")
	return userHome + "/Library/Application Support"
}
