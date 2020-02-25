package cmd

import (
	"reflect"
	"testing"
)

func TestExportBash(t *testing.T) {
	paths := []string{".sdkman/candidates/java/current/bin", ".sdkman/candidates/scala/current/bin"}
	envVars := []envVar{{"JAVA_HOME", ".sdkman/candidates/java/current"}, {"SCALA_HOME", ".sdkman/candidates/scala/current"}}
	answer := "export PATH=.sdkman/candidates/java/current/bin:.sdkman/candidates/scala/current/bin:$PATH\nexport JAVA_HOME=.sdkman/candidates/java/current\nexport SCALA_HOME=.sdkman/candidates/scala/current\n"

	if exportBash(paths, envVars) != answer {
		t.Errorf("expected: %s, return: %s", answer, exportBash(paths, envVars))
	}
}

func TestExportFish(t *testing.T) {
	paths := []string{".sdkman/candidates/java/current/bin", ".sdkman/candidates/scala/current/bin"}
	envVars := []envVar{{"JAVA_HOME", ".sdkman/candidates/java/current"}, {"SCALA_HOME", ".sdkman/candidates/scala/current"}}
	answer := "set -x PATH .sdkman/candidates/java/current/bin .sdkman/candidates/scala/current/bin $PATH\nset -x JAVA_HOME .sdkman/candidates/java/current\nset -x SCALA_HOME .sdkman/candidates/scala/current\n"

	if exportFish(paths, envVars) != answer {
		t.Errorf("expected: %s, return: %s", answer, exportFish(paths, envVars))
	}
}

func TestExportPosh(t *testing.T) {
	paths := []string{".sdkman/candidates/java/current/bin", ".sdkman/candidates/scala/current/bin"}
	envVars := []envVar{{"JAVA_HOME", ".sdkman/candidates/java/current"}, {"SCALA_HOME", ".sdkman/candidates/scala/current"}}
	answer := "$env:Path = \".sdkman/candidates/java/current/bin;.sdkman/candidates/scala/current/bin;\" + $env:Path;$env:JAVA_HOME = \".sdkman/candidates/java/current\";$env:SCALA_HOME = \".sdkman/candidates/scala/current\";"

	if exportPosh(paths, envVars) != answer {
		t.Errorf("expected: %s, return: %s", answer, exportPosh(paths, envVars))
	}
}

func TestRemoveIndexes(t *testing.T) {
	paths := []string{"a", "b", "c", "d"}
	idxs := []int{1, 3}
	res := removeIndexes(paths, idxs)
	if !reflect.DeepEqual(res, []string{"a", "c"}) {
		t.Errorf("expected: %s, return: %s", []string{"a", "c"}, res)
	}
}
