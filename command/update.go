package command

import (
	"fmt"
	"github.com/palindrom615/sdk/api"
	"github.com/palindrom615/sdk/store"
	"github.com/palindrom615/sdk/utils"
	"github.com/urfave/cli/v2"
	"strings"

	"github.com/scylladb/go-set/strset"
)

func Update(c *cli.Context) error {
	reg := c.String("registry")
	root := c.String("directory")
	freshCsv, netErr := api.GetAll(reg)
	if netErr != nil {
		return utils.ErrNotOnline
	}
	fresh := strset.New(freshCsv...)
	cachedCsv := store.GetCandidates(root)
	cached := strset.New(cachedCsv...)

	added := strset.Difference(fresh, cached)
	obsoleted := strset.Difference(cached, fresh)

	if added.Size() == 0 && obsoleted.Size() == 0 {
		fmt.Println("No new candidates found at this time.")
	} else {
		fmt.Println("Adding new candidates: " + strings.Join(added.List(), ", "))
		fmt.Println("Removing obsolete candidates: " + strings.Join(obsoleted.List(), ", "))
		_ = store.SetCandidates(root, freshCsv)
	}
	return nil
}
