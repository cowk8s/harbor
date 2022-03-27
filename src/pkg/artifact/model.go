package artifact

type Artifact struct {
	ID                int64  `json:"id"`
	Type              string `json:"type"`
	MediaType         string `json:"media_type"`
	ManifestMediaType string `json:"manifest_media_type"`
}
