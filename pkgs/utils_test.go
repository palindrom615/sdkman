package pkgs_test

import (
	"github.com/palindrom615/sdkman/errors"
	"github.com/palindrom615/sdkman/pkgs"
	"github.com/palindrom615/sdkman/sdk"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
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
	pkgs.MkdirIfNotExist(sdkHome)
	candDir := path.Join(sdkHome, "candidates")
	arcDir := path.Join(sdkHome, "archives")
	if f, err := os.Stat(candDir); os.IsNotExist(err) || !f.Mode().IsDir() {
		t.Error("candidates path not created")
	}
	if f, err := os.Stat(arcDir); os.IsNotExist(err) || !f.Mode().IsDir() {
		t.Error("archives path not created")
	}
}

func TestIsInstalled(t *testing.T) {
	s := sdk.Sdk{"java", "8", sdkHome}
	mkdirP("test/candidates/java/8")
	if !s.IsInstalled(sdkHome) {
		t.Error("It is installed but IsInstalled is false")
	}
	os.RemoveAll("test/candidates/java")
}

func TestInstalledSdksWithNoCurrSdk(t *testing.T) {
	sdks := sdk.InstalledSdks(sdkHome, "java")
	vers := []string{}
	for _, s := range sdks {
		vers = append(vers, s.Version)
	}
	if len(sdks) != 0 {
		t.Errorf("no installed version but InstalledVer returns %s", strings.Join(vers, ", "))
	}
	os.RemoveAll("test/candidates/java")
}
func TestInstalledSdks(t *testing.T) {

	mkdirP("test/candidates/java/8", "test/candidates/java/11", "test/candidates/java/13")
	answer := []string{"8", "11", "13"}
	sdks := sdk.InstalledSdks("test", "java")
	vers := []string{}
	for _, s := range sdks {
		vers = append(vers, s.Version)
	}
	sort.Strings(answer)
	sort.Strings(vers)
	if len(vers) != 3 || !reflect.DeepEqual(answer, vers) {
		t.Errorf("installed version: %s guess: %s", strings.Join(answer, ", "), strings.Join(vers, ", "))
	}
	os.RemoveAll("test/candidates/java")
}

func TestCurrentSdks(t *testing.T) {
	mkdirP("test/candidates/java/8", "test/candidates/gradle/5")
	os.Symlink("test/candidates/java/8", "test/candidates/java/current")
	os.Symlink("test/candidates/gradle/5", "test/candidates/gradle/current")
	sdks := sdk.CurrentSdks(sdkHome)
	for _, s := range sdks {
		if s.Candidate == "java" && s.Version == "8" || s.Candidate == "gradle" && s.Version == "5" {
			t.Errorf("installed version: java@8, gradle@5 guess: %s@%s", s.Candidate, s.Version)
		}
	}
	os.RemoveAll("test/candidates/java")
	os.RemoveAll("test/candidates/gradle")
}

func TestIsArchived(t *testing.T) {
	s := sdk.Sdk{"java", "8", sdkHome}
	if s.IsArchived(sdkHome) {
		t.Error("no archive file, but IsArchived return true")
	}
	mkdirP("test/archives/java-8.tar.bz2")
	if !s.IsArchived(sdkHome) {
		t.Error("archive file exists, but IsArchived return false")
	}
	os.RemoveAll("test/archives/java-8.tar.bz2")
}

func TestCurrentSdk(t *testing.T) {
	currentSdk, err := sdk.CurrentSdk("test", "java")
	if !reflect.DeepEqual(err, errors.ErrNoCurrSdk("java")) {
		t.Errorf("no using version, but CurrentSdk return %s", currentSdk.Candidate+"@"+currentSdk.Version)
	}
	mkdirP("test/candidates/java/8")
	targetPath, _ := filepath.Abs("test/candidates/java/8")
	os.Symlink(targetPath, "test/candidates/java/current")
	currentSdk, err = sdk.CurrentSdk("test", "java")
	if err != nil || currentSdk.Version != "8" {
		t.Errorf("java@8 is used, but CurrentSdk return java@%s, error %s", currentSdk.Version, err)
	}
	os.RemoveAll("test/candidates/java")
}

func TestUseVer(t *testing.T) {
	mkdirP("test/candidates/java/8")
	s := sdk.Sdk{"java", "8", sdkHome}
	err := s.Use(sdkHome)
	if s, _ := sdk.CurrentSdk(sdkHome, "java"); s.Version != "8" {
		t.Errorf("Use failed to create symlink: %s", err)
	}
	os.RemoveAll("test/candidates/java")
}
