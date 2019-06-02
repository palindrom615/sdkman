package sdkmanCli

import (
	"crypto/tls"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

func download(url string) []byte {
	insecureSkipVerify := os.Getenv("sdkman_insecure_ssl") == "true"
	client := &http.Client{Transport: &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: insecureSkipVerify},
	}}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	return data
}

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
