package util_test

import (
	"github.com/palindrom615/sdkman/util"
	"os"
	"path/filepath"
	"testing"
)

func mkdirP(paths ...string) {
	for _, p := range paths {
		os.MkdirAll(p, os.ModeDir|os.ModePerm)
	}
}

var (
	sdkHome = "test"
)

func TestMain(m *testing.M) {
	os.Mkdir("test", os.ModePerm|os.ModeDir)

	code := m.Run()
	os.RemoveAll("test")
	os.Exit(code)
}

func TestMkdirIfNotExist(t *testing.T) {
	util.MkdirIfNotExist(sdkHome)
	candDir := filepath.Join(sdkHome, "candidates")
	arcDir := filepath.Join(sdkHome, "archives")
	if f, err := os.Stat(candDir); os.IsNotExist(err) || !f.Mode().IsDir() {
		t.Error("candidates path not created")
	}
	if f, err := os.Stat(arcDir); os.IsNotExist(err) || !f.Mode().IsDir() {
		t.Error("archives path not created")
	}
}
