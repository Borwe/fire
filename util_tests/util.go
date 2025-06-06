package utiltests_test

import "path/filepath"

func GetDir(file string) string{
	return filepath.Dir(file) + "/"
}
