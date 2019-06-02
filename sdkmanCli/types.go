package sdkmanCli

type Env struct {
	Dir           string
	CandidateDir  string
	Candidates    []string
	OfflineMod    bool
	CandidatesApi string
	Version       string
	Platform      string
}
