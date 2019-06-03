package command

import (
	"sdkman-cli/api"
	"strings"

	"github.com/fatih/color"
	"github.com/scylladb/go-set/strset"
)

func Update() {
	freshCandidatesCsv := api.GetAll()
	cachedCandidatesCsv, _ := kv.Get("candidates/all")

	fresh := strset.New()
	cached := strset.New()
	for _, can := range strings.Split(string(freshCandidatesCsv), ",") {
		fresh.Add(can)
	}
	for _, can := range strings.Split(string(cachedCandidatesCsv), ",") {
		cached.Add(can)
	}
	added := strset.Difference(fresh, cached)
	obsoleted := strset.Difference(cached, fresh)

	if added.Size() == 0 && obsoleted.Size() == 0 {
		color.Green("No new candidates found at this time.")
	} else {
		color.Green("Adding new candidates: %s", strings.Join(added.List(), ", "))
		color.Green("Removing obsolete candidates: %s", strings.Join(obsoleted.List(), ", "))
		kv.Put("candidates/all", freshCandidatesCsv)
	}
}
