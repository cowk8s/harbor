package artifact

import "fmt"

// Artifact is the abstract object managed by Harbor. It hides the
// underlying concrete detail and provides an unified artifact view
// for all users.
type Artifact struct {
	ID                int64                  `json:"id"`
	Type              string                 `json:"type"`                // image, chart, etc
	MediaType         string                 `json:"media_type"`          // the media type of artifact. Mostly, it's the value of `manifest.config.mediatype`
	ManifestMediaType string                 `json:"manifest_media_type"` // the media type of manifest/index
}

func (a *Artifact) String() string {
	return fmt.Sprintf("%s@%s", a.RepositoryName, a.Digest)
}

// IsImageIndex returns true when artifact is image index
func (a *Artifact) IsImageIndex() bool {
	return false
}


