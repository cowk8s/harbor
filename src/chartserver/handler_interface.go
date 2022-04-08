package chartserver

// ServiceHandler defines the related methods to handle kinds of chart service requests.
type ServiceHandler interface {
	// ListCharts lists all the charts under the specified namespace.
	//
	ListCharts(namespace string) ([]*ChartInfo, error)

	GetChart(namespace, chartName string) (helm_repo.ChartVersions, error)

	
}