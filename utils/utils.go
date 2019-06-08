package utils

import (
	"os"
	"os/exec"
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
		ThrowError(err)
	}
}
