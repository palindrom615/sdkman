package command

import (
	"github.com/palindrom615/sdk/api"
	"github.com/palindrom615/sdk/local"
	"github.com/palindrom615/sdk/utils"
	"github.com/urfave/cli/v2"
)

func Install(c *cli.Context) error {
	candidate := c.Args().Get(0)
	version := c.Args().Get(1)
	folder := c.Args().Get(2)

	reg := c.String("registry")
	root := c.String("directory")

	_ = Update(c)

	local.MkdirIfNotExist(root)
	if err := utils.CheckValidCand(root, candidate); err != nil {
		return err
	}
	if version == "" {
		if dfVer, err := defaultVersion(reg, root, candidate); err != nil {
			return err
		} else {
			version = dfVer
		}
	}

	if local.IsInstalled(root, candidate, version) {
		return utils.ErrVerExists
	}
	if err := checkValidVer(reg, root, candidate, version, folder); err != nil {
		return err
	}

	archiveReady := make(chan bool)
	installReady := make(chan bool)
	go local.Unarchive(root, candidate, version, archiveReady, installReady)
	if local.IsArchived(root, candidate, version) {
		archiveReady <- true
	} else {
		s, err, t := api.GetDownload(reg, candidate, version)
		if err != nil {
			archiveReady <- false
			return err
		}
		go local.Archive(s, root, candidate, version, t, archiveReady)
	}
	if <-installReady == false {
		return utils.ErrVerInsFail
	}
	return local.UseVer(root, candidate, version)
}

func defaultVersion(reg string, root string, candidate string) (string, error) {
	if v, netErr := api.GetDefault(reg, candidate); netErr == nil {
		return v, nil
	} else if curr, fsErr := local.UsingVer(root, candidate); fsErr == nil {
		return curr, nil
	} else {
		return "", utils.ErrNotOnline
	}

}

func checkValidVer(reg string, root string, candidate string, version string, folder string) error {
	isValid, netErr := api.GetValidate(reg, candidate, version)
	if (netErr == nil && isValid) || folder != "" || local.IsInstalled(root, candidate, version) {
		return nil
	} else if netErr != nil {
		return utils.ErrNotOnline
	} else {
		return utils.ErrNoVer
	}
}
