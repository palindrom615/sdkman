package sdkman

import (
	"fmt"
	"github.com/scylladb/go-set/strset"
	"github.com/urfave/cli/v2"
	"path"
	"strings"
)

func Install(c *cli.Context) error {
	_ = Update(c)
	reg := c.String("registry")
	root := c.String("directory")
	target, err := arg2sdk(reg, root, c.Args().Get(0))
	if err != nil {
		return err
	}

	MkdirIfNotExist(root)
	if err := checkValidCand(root, target.Candidate); err != nil {
		return err
	}
	if target.Version == "" {
		if defaultSdk, err := defaultSdk(reg, root, target.Candidate); err != nil {
			return err
		} else {
			target = defaultSdk
		}
	}

	if target.IsInstalled(root) {
		return ErrVerExists
	}
	if err := target.checkValidVer(reg, root); err != nil {
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
	return target.Use(root)
}

func Use(c *cli.Context) error {
	reg := c.String("registry")
	root := c.String("directory")
	sdk, err := arg2sdk(reg, root, c.Args().Get(0))
	if err != nil {
		return err
	}
	if !sdk.IsInstalled(root) {
		return ErrVerNotIns
	}
	return sdk.Use(root)
}

func Current(c *cli.Context) error {
	candidate := c.Args().Get(0)
	root := c.String("directory")
	if candidate == "" {
		sdks := CurrentSdks(root)
		if len(sdks) == 0 {
			return ErrNoCurrCands
		}
		for _, sdk := range sdks {
			fmt.Printf("%s@%s\n", sdk.Candidate, sdk.Version)
		}
	} else {
		sdk, err := CurrentSdk(root, candidate)
		if err == nil {
			fmt.Println(sdk.Candidate + "@" + sdk.Version)
		} else {
			return ErrNoCurrSdk(candidate)
		}
	}
	return nil
}

type envVar struct {
	name string
	val  string
}

func Export(c *cli.Context) error {
	shell := c.Args().Get(0)
	if shell == "" {
		if platform() == "msys_nt-10.0" {
			shell = "windows"
		} else {
			shell = "bash"
		}
	}
	root := c.String("directory")
	sdks := CurrentSdks(c.String("directory"))
	if len(sdks) == 0 {
		fmt.Println("")
		return nil
	}
	paths := []string{}
	homes := []envVar{}
	for _, sdk := range sdks {
		candHome := path.Join(root, "candidates", sdk.Candidate, "current")
		paths = append(paths, path.Join(candHome, "bin"))
		homes = append(homes, envVar{fmt.Sprintf("%s_HOME", strings.ToUpper(sdk.Candidate)), candHome})
	}

	if shell == "bash" || shell == "zsh" {
		evalBash(paths, homes)
	} else if shell == "fish" {
		evalFish(paths, homes)
	} else if shell == "powershell" || shell == "posh" {
		evalPosh(paths, homes)
	} else if shell == "windows" || shell == "window" {
		evalWindows(paths, homes)
	}
	return nil
}

func List(c *cli.Context) error {
	candidate := c.Args().Get(0)
	reg := c.String("registry")
	root := c.String("directory")

	if candidate == "" {
		list, err := getList(reg)
		if err == nil {
			pager(list)
		}
		return err
	} else {
		if err := checkValidCand(root, candidate); err != nil {
			return err
		}
		ins := InstalledSdks(root, candidate)
		curr, _ := CurrentSdk(root, candidate)
		list, err := getVersionsList(reg, curr, ins)
		pager(list)
		return err
	}
}

func Update(c *cli.Context) error {
	reg := c.String("registry")
	root := c.String("directory")
	freshCsv, netErr := getAll(reg)
	if netErr != nil {
		return ErrNotOnline
	}
	fresh := strset.New(freshCsv...)
	cachedCsv := getCandidates(root)
	cached := strset.New(cachedCsv...)

	added := strset.Difference(fresh, cached)
	obsoleted := strset.Difference(cached, fresh)

	if added.Size() != 0 {
		fmt.Println("Adding new candidates: " + strings.Join(added.List(), ", "))
	}
	if obsoleted.Size() != 0 {
		fmt.Println("Removing obsolete candidates: " + strings.Join(obsoleted.List(), ", "))
	}
	_ = setCandidates(root, freshCsv)
	return nil
}

func evalBash(paths []string, envVars []envVar) {
	fmt.Println("export PATH=" + strings.Join(paths, ":") + ":$PATH")
	for _, v := range envVars {
		fmt.Println("export " + v.name + "=" + v.val)
	}
}

func evalFish(paths []string, envVars []envVar) {
	fmt.Println("set -x PATH " + strings.Join(paths, " ") + " $PATH")
	for _, v := range envVars {
		fmt.Println("set -x " + v.name + " " + v.val)
	}
}

func evalPosh(paths []string, envVars []envVar) {
	fmt.Printf("$env:Path = %s; + $env:Path\n", strings.Join(paths, ";"))
	for _, v := range envVars {
		fmt.Printf("$env:%s = %s\n", v.name, v.val)
	}
}

func evalWindows(paths []string, envVars []envVar) {
	fmt.Printf("[Environment]::SetEnvironmentVariable(\"Path\", \"%s;\" + $env:Path, [System.EnvironmentVariableTarget]::User\n", strings.Join(paths, ";"))
	for _, v := range envVars {
		fmt.Printf("[Environment]::SetEnvironmentVariable(\"%s\", \"%s\", [System.EnvironmentVariableTarget]::User\n", v.name, v.val)
	}
}
