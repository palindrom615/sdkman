package pkgs_test

import (
	"github.com/palindrom615/sdkman/errors"
	"github.com/palindrom615/sdkman/pkgs"
	"os"
	"path"
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

func TestMain(m *testing.M) {
	os.Mkdir("test", os.ModePerm|os.ModeDir)
	code := m.Run()
	os.RemoveAll("test")
	os.Exit(code)
}

func TestMkdirIfNotExist(t *testing.T) {
	pkgs.MkdirIfNotExist("test")
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
	sdk := pkgs.Sdk{"java", "8"}
	mkdirP("test/candidates/java/8")
	if !sdk.IsInstalled(root) {
		t.Error("It is installed but IsInstalled is false")
	}
	os.RemoveAll("test/candidates/java")
}

func TestInstalledSdksWithNoCurrSdk(t *testing.T) {
	sdks := pkgs.InstalledSdks("test", "java")
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
	sdks := pkgs.InstalledSdks("test", "java")
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
	sdks := pkgs.CurrentSdks("test")
	for _, sdk := range sdks {
		if sdk.Candidate == "java" && sdk.Version == "8" || sdk.Candidate == "gradle" && sdk.Version == "5" {
			t.Errorf("installed version: java@8, gradle@5 guess: %s@%s", sdk.Candidate, sdk.Version)
		}
	}
	os.RemoveAll("test/candidates/java")
	os.RemoveAll("test/candidates/gradle")
}

func TestIsArchived(t *testing.T) {
	sdk := pkgs.Sdk{"java", "8"}
	if sdk.IsArchived("test") {
		t.Error("no archive file, but IsArchived return true")
	}
	mkdirP("test/archives/java-8.tar.bz2")
	if !sdk.IsArchived("test") {
		t.Error("archive file exists, but IsArchived return false")
	}
	os.RemoveAll("test/archives/java-8.tar.bz2")
}

func TestCurrentSdk(t *testing.T) {
	sdk, err := pkgs.CurrentSdk("test", "java")
	if !reflect.DeepEqual(err, errors.ErrNoCurrSdk("java")) {
		t.Errorf("no using version, but CurrentSdk return %s", sdk.Candidate+"@"+sdk.Version)
	}
	mkdirP("test/candidates/java/8")
	os.Symlink("test/candidates/java/8", "test/candidates/java/current")
	sdk, err = pkgs.CurrentSdk("test", "java")
	if err != nil || sdk.Version != "8" {
		t.Errorf("java@8 is used, but CurrentSdk return java@%s, error %s", sdk.Version, err)
	}
	os.RemoveAll("test/candidates/java")
}

func TestUseVer(t *testing.T) {
	mkdirP("test/candidates/java/8")
	sdk := pkgs.Sdk{"java", "8"}
	err := sdk.Use("test")
	if sdk, _ := pkgs.CurrentSdk("test", "java"); sdk.Version != "8" {
		t.Errorf("Use failed to create symlink: %s", err)
	}
	os.RemoveAll("test/candidates/java")
}
