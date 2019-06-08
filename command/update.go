package command

import (
	"sdkman-cli/api"
	"sdkman-cli/store"
	"sdkman-cli/utils"
	"strings"

	"github.com/fatih/color"
	"github.com/scylladb/go-set/strset"
)

func Update() {
	freshCsv, netErr := api.GetAll()
	if netErr != nil {
		utils.ThrowError(utils.ErrNotOnline)
	}
	fresh := strset.New(freshCsv...)
	cachedCsv := store.GetCandidates()
	cached := strset.New(cachedCsv...)

	added := strset.Difference(fresh, cached)
	obsoleted := strset.Difference(cached, fresh)

	if added.Size() == 0 && obsoleted.Size() == 0 {
		color.Green("No new candidates found at this time.")
	} else {
		color.Green("Adding new candidates: %s", strings.Join(added.List(), ", "))
		color.Green("Removing obsolete candidates: %s", strings.Join(obsoleted.List(), ", "))
		_ = store.SetCandidates(freshCsv)
	}
}
