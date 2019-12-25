package utils

import (
	"github.com/palindrom615/sdkman-cli/conf"
	"io"
	"os"
	"os/exec"
)

func Pager(pages io.ReadCloser) {
	pager := os.Getenv("PAGER")
	c := conf.GetConf()

	if pager == "" {
		if c.Platform == "msys_nt-10.0" {
			pager = "more"
		} else {
			pager = "less"
		}
	}
	c1 := exec.Command(pager)
	c1.Stdin = pages
	c1.Stdout = os.Stdout
	_ = c1.Start()
	_ = c1.Wait()
	defer pages.Close()
}
