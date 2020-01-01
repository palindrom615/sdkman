package sdkman

import (
	"github.com/urfave/cli/v2"
)

func Install(c *cli.Context) error {
	candidate := c.Args().Get(0)
	version := c.Args().Get(1)
	folder := c.Args().Get(2)

	reg := c.String("registry")
	root := c.String("directory")

	_ = Update(c)

	MkdirIfNotExist(root)
	if err := checkValidCand(root, candidate); err != nil {
		return err
	}
	if version == "" {
		if dfVer, err := defaultVersion(reg, root, candidate); err != nil {
			return err
		} else {
			version = dfVer
		}
	}
	target := Sdk{candidate, version}

	if target.IsInstalled(root) {
		return ErrVerExists
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
		s, err, t := getDownload(reg, target)
		if err != nil {
			archiveReady <- false
			return err
		}
		archive := Archive{target, t}
		go archive.Save(s, root, archiveReady)
	}
	if <-installReady == false {
		return ErrVerInsFail
	}
	return target.UseVer(root)
}

func defaultVersion(reg string, root string, candidate string) (string, error) {
	if v, netErr := getDefault(reg, candidate); netErr == nil {
		return v, nil
	} else if curr, fsErr := UsingVer(root, candidate); fsErr == nil {
		return curr, nil
	} else {
		return "", ErrNotOnline
	}

}

func checkValidVer(reg string, root string, target Sdk, folder string) error {
	isValid, netErr := getValidate(reg, target)
	if (netErr == nil && isValid) || folder != "" || target.IsInstalled(root) {
		return nil
	} else if netErr != nil {
		return ErrNotOnline
	} else {
		return ErrNoVer
	}
}
