package utils

import (
	"github.com/palindrom615/sdkman-cli/conf"
	"io"
	"mime"
	"net/http"
	"os"
	"os/exec"
	"strings"
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
	err := c1.Start()
	Check(err)
	_ = c1.Wait()
	defer pages.Close()
}

func TypeOfResponse(header http.Header) string {
	contentType, contentDisposition := header.Get("Content-Type"), header.Get("Content-Disposition")
	if contentDisposition != "" {
		_, params, _ := mime.ParseMediaType(contentDisposition)
		filename := strings.Split(params["filename"], ".")
		return filename[len(filename)-1]
	} else {
		exts, _ := (mime.ExtensionsByType(contentType))
		return exts[0]
	}
}
