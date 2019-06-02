package sdkmanCli

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
)

func Pager(pages string) {
	pager := os.Getenv("PAGER")
	if pager == "" {
		pager = "more"
	}
	c1 := exec.Command(pager)
	c1.Stdin = strings.NewReader(pages)
	c1.Stdout = os.Stdout
	err := c1.Start()
	c1.Wait()
	if err != nil {
		log.Fatal(err)
	}
}

func currentVersion(candidate string) (string, error) {
	p, err := os.Readlink(path.Join(e.Dir, "candidates", candidate, "current"))
	if err == nil {
		d, _ := os.Stat(p)
		return d.Name(), nil
	}
	return "", err
}

func installed(candidate string) ([]string, error) {
	res := []string{}
	vers, err := ioutil.ReadDir(path.Join(e.Dir, "candidates", candidate))
	for _, ver := range vers {
		res = append(res, ver.Name())
	}
	return res, err
}

func isInstalled(candidate string, version string) bool {
	archDir := e.Dir + "/archives"
	target := archDir + "/" + candidate + "-" + version
	_, err := os.Stat(target)
	return os.IsNotExist(err)
}