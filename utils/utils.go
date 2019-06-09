package utils

import (
	"io"
	"os"
	"os/exec"
)

func Pager(pages io.ReadCloser) {
	pager := os.Getenv("PAGER")
	if pager == "" {
		pager = "more"
	}
	c1 := exec.Command(pager)
	c1.Stdin = pages
	c1.Stdout = os.Stdout
	err := c1.Start()
	Check(err)
	c1.Wait()
	defer pages.Close()
}
