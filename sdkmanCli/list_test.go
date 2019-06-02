package sdkmanCli

import "testing"


func TestListCandidatesApi(t *testing.T) {
	env := &Env{"~\\sdkman", "C:\\Users\\palin\\.sdkman\\candidates",
		[]string{"java", "scala"}, false, "https://example.com",
		"1", "MSYS_NT-10.0"}

	expected := "https://example.com/list"
	result := listCandidatesApi(env)
	if expected != result {
		t.Errorf("expected: %s, result: %s", expected, result)
	}
}

func TestListCandidateApi(t *testing.T) {
	env := &Env{"~\\sdkman", "C:\\Users\\palin\\.sdkman\\candidates",
		[]string{"java", "scala"}, false, "https://example.com",
		"1", "Linux64"}

	expected := "https://example.com/java/Linux64/versions/list?current=11.0.0&installed=9.0.0,10.0.0,11.0.0"
	result := listCandidateApi(env, "java", "11.0.0", []string{"9.0.0", "10.0.0", "11.0.0"})
	if expected != result {
		t.Errorf("expected: %s, result: %s", expected, result)
	}
}