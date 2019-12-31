package local_test

import (
	"github.com/palindrom615/sdk/local"
	"os"
	"path"
	"reflect"
	"sort"
	"strings"
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

func TestMkdirIfNotExist(t *testing.T) {
	local.MkdirIfNotExist("test")
	candDir := path.Join("test", "candidates")
	arcDir := path.Join("test", "archives")
	if f, err := os.Stat(candDir); os.IsNotExist(err) || !f.Mode().IsDir() {
		t.Error("candidates path not created")
	}
	if f, err := os.Stat(arcDir); os.IsNotExist(err) || !f.Mode().IsDir() {
		t.Error("archives path not created")
	}
}

func TestIsInstalled(t *testing.T) {
	root := "test"
	candidate := "java"
	version := "8"
	mkdirP("test/candidates/java/8")
	if !local.IsInstalled(root, candidate, version) {
		t.Error("It is installed but IsInstalled is false")
	}
	os.RemoveAll("test/candidates/java")
}

func TestInstalledVers(t *testing.T) {
	vers := local.InstalledVers("test", "java")
	if len(vers) != 0 {
		t.Errorf("no installed version but InstalledVer returns %s", strings.Join(vers, ", "))
	}
	mkdirP("test/candidates/java/8", "test/candidates/java/11", "test/candidates/java/13")
	vers = local.InstalledVers("test", "java")
	if len(vers) != 3 {
		answer := []string{"8", "11", "13"}
		sort.Strings(answer)
		sort.Strings(vers)
		if !reflect.DeepEqual(answer, vers) {
			t.Errorf("installed version: %s guess: %s", strings.Join(answer, ", "), strings.Join(vers, ", "))
		}
	}
	os.RemoveAll("test/candidates/java")
}

func TestUsingCands(t *testing.T) {
	mkdirP("test/candidates/java/8", "test/candidates/gradle/5")
	os.Symlink("test/candidates/java/8", "test/candidates/java/current")
	os.Symlink("test/candidates/gradle/5", "test/candidates/gradle/current")
	cands, vers := local.UsingCands("test")
	for i, c := range cands {
		if c == "java" && vers[i] == "8" || c == "gradle" && vers[i] == "5" {
			t.Errorf("installed version: java@8, gradle@5 guess: %s@%s", c, vers[i])
		}
	}
	os.RemoveAll("test/candidates/java")
	os.RemoveAll("test/candidates/gradle")
}

func TestIsArchived(t *testing.T) {
	if local.IsArchived("test", "java", "8") {
		t.Error("no archive file, but IsArchived return true")
	}
	mkdirP("test/archives/java-8.tar.bz2")
	if !local.IsArchived("test", "java", "8") {
		t.Error("archive file exists, but IsArchived return false")
	}
	os.RemoveAll("test/archives/java-8.tar.bz2")
}

func TestUsingVer(t *testing.T) {
	ver, err := local.UsingVer("test", "java")
	if err == nil {
		t.Errorf("no using version, but UsingVer return %s", ver)
	}
	mkdirP("test/candidates/java/8")
	os.Symlink("test/candidates/java/8", "test/candidates/java/current")
	ver, err = local.UsingVer("test", "java")
	if err != nil || ver != "8" {
		t.Errorf("java@8 is used, but UsingVer return java@%s, error %s", ver, err)
	}
	os.RemoveAll("test/candidates/java")
}

func TestUseVer(t *testing.T) {
	mkdirP("test/candidates/java/8")
	err := local.UseVer("test", "java", "8")
	if ver, _ := local.UsingVer("test", "java"); ver != "8" {
		t.Errorf("UseVer failed to create symlink: %s", err)
	}
}
