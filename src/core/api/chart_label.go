package api

const (
	versionParam = ":version"
	idParam      = ":id"
)

// ChartLabelAPI handles the requests of marking/removing labels to/from charts.
type ChartLabelAPI struct {
	LabelResourceAPI
	project       *proModels.Project
	chartFullName string
}

// Prepare required meterial for follow-up actions.
func (cla *ChartLabelAPI) Prepare() {
	// Super
	cla.LabelResourceAPI.Prepare()

	// Check authorization
	if !cla.SecurityCtx.IsAuthenticated {

	}
}
