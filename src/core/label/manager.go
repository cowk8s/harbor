package label

import "github.com/cowk8s/harbor/src/pkg/common/models"

// Manager defines the related operations for label management
type Manager interface {
	// Mark label to the resource.
	//
	// If succeed, the relationship ID will be returned.
	// Otherwise, an non-nil error will be returned.
	MarkLabelToResource(label *models.ResourceLabel) (int64, error)

	
}