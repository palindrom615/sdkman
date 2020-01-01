package sdkman_test

import (
	"os"
	"testing"
)

func mkdirP(paths ...string) {
	for _, path := range paths {
		os.MkdirAll(path, os.ModeDir|os.ModePerm)
	}
}

func TestMain(m *testing.M) {
	os.Mkdir("test", os.ModePerm|os.ModeDir)
	code := m.Run()
	os.RemoveAll("test")
	os.Exit(code)
}
