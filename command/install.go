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
	target := local.Sdk{candidate, version}

	if target.IsInstalled(root) {
		return utils.ErrVerExists
	}
	if err := checkValidVer(reg, root, target, folder); err != nil {
		return err
	}

	archiveReady := make(chan bool)
	installReady := make(chan bool)
	go target.Unarchive(root, archiveReady, installReady)
	if target.IsArchived(root) {
		archiveReady <- true
	} else {
		s, err, t := api.GetDownload(reg, target)
		if err != nil {
			archiveReady <- false
			return err
		}
		archive := local.Archive{target, t}
		go archive.Save(s, root, archiveReady)
	}
	if <-installReady == false {
		return utils.ErrVerInsFail
	}
	return target.UseVer(root)
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

func checkValidVer(reg string, root string, target local.Sdk, folder string) error {
	isValid, netErr := api.GetValidate(reg, target)
	if (netErr == nil && isValid) || folder != "" || target.IsInstalled(root) {
		return nil
	} else if netErr != nil {
		return utils.ErrNotOnline
	} else {
		return utils.ErrNoVer
	}
}
