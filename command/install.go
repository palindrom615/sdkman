package command

import (
	"github.com/palindrom615/sdkman-cli/api"
	"github.com/palindrom615/sdkman-cli/local"
	"github.com/palindrom615/sdkman-cli/utils"
)

func Install(candidate string, version string, folder string) error {
	_ = Update()

	if err := utils.CheckValidCand(candidate); err != nil {
		return err
	}
	if version == "" {
		if dfVer, err := defaultVersion(candidate); err != nil {
			return err
		} else {
			version = dfVer
		}
	}
	if local.IsInstalled(candidate, version) {
		return utils.ErrVerExists
	}
	if err := CheckValidVer(candidate, version, folder); err != nil {
		return err
	}

	archiveReady := make(chan bool)
	installReady := make(chan bool)
	go local.Unpack(candidate, version, archiveReady, installReady)
	if local.IsArchived(candidate, version) {
		archiveReady <- true
	} else {
		s, err, t := api.GetDownload(candidate, version)
		if err != nil {
			return err
		}
		go local.Archive(s, candidate, version, t, archiveReady)
	}
	<-installReady
	return Use(candidate, version)
}

func defaultVersion(candidate string) (string, error) {
	if v, netErr := api.GetDefault(candidate); netErr == nil {
		return v, nil
	} else if curr, fsErr := local.Current(candidate); fsErr == nil {
		return curr, nil
	} else {
		return "", utils.ErrNotOnline
	}

}

func CheckValidVer(candidate string, version string, folder string) error {
	isValid, netErr := api.GetValidate(candidate, version)
	if (netErr == nil && isValid) || folder != "" || local.IsInstalled(candidate, version) {
		return nil
	} else if netErr != nil {
		return utils.ErrNotOnline
	} else {
		return utils.ErrNoVer
	}
}
