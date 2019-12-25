package command

import (
	"github.com/palindrom615/sdkman-cli/api"
	"github.com/palindrom615/sdkman-cli/local"
	"github.com/palindrom615/sdkman-cli/utils"
)

func Install(candidate string, version string, folder string) error {
	if candidate == "" {
		return utils.ErrNoCandidate
	}
	Update()

	if !utils.IsCandidateValid(candidate) {
		return utils.ErrNoCandidate
	}
	if version == "" {
		version = defaultVersion(candidate)
	}
	if local.IsInstalled(candidate, version) {
		return utils.ErrVersionExists
	}
	if !isValidVersion(candidate, version, folder) {
		return utils.ErrNotValidVersion
	}

	archiveReady := make(chan bool)
	installReady := make(chan bool)
	go local.Unpack(candidate, version, archiveReady, installReady)
	if !local.IsArchived(candidate, version) {
		s, err := api.GetDownload(candidate, version)
		utils.Check(err)
		go local.Archive(s, candidate, version, archiveReady)

	} else {
		archiveReady <- true
	}
	<-installReady
	return Use(candidate, version)
}

func defaultVersion(candidate string) string {
	v, netErr := api.GetDefault(candidate)
	if netErr != nil {
		curr, fsErr := local.Current(candidate)
		if fsErr != nil {
			utils.Check(utils.ErrNotOnline)
		}
		return curr
	}
	return v
}

func isValidVersion(candidate string, version string, folder string) bool {
	isValid, netErr := api.GetValidate(candidate, version)
	if (netErr == nil && isValid) || folder != "" || local.IsInstalled(candidate, version) {
		return true
	} else {
		if netErr != nil {
			utils.Check(utils.ErrNotOnline)
		}
		return false
	}
}
