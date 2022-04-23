package artifact

import (
	"fmt"

	"github.com/cowk8s/harbor/src/lib/encode/repository"
	"github.com/cowk8s/harbor/src/pkg/artifact"
)

type Artifact struct {
	artifact.Artifact
	Tags          []*tag.Tag                 `json:"tags"`           // the list of tags that attached to the artifact
	AdditionLinks map[string]*AdditionLink   `json:"addition_links"` // the resource link for build history(image), value.yaml(chart), dependency(chart), etc
	Labels        []*model.Label             `json:"labels"`
	Accessories   []accessoryModel.Accessory `json:"-"`
}

func (artifact *Artifact) SetAdditionLink(addition, version string) {
	if artifact.AdditionLinks == nil {
		artifact.AdditionLinks = make(map[string]*AdditionLink)
	}

	projectName, repo := utils.ParseRepository(artifact.RepositoryName)
	// encode slash as %252F
	repo = repository.Encode(repo)
	href := fmt.Sprintf("api/%s/projects/%s/repositories/%s/artifacts/%s/additions/%s", version, projectName, repo, artifact.Digest, addition)

	artifact.AdditionLinks[addition] = &AdditionLink{HREF: href, Absolute: false}
}

// AdditionLink is a link via that the addition can be fetched
type AdditionLink struct {
	HREF     string `json:"href"`
	Absolute bool   `json:"absolute"` // specify the href is an absolute URL or not
}

// Option is used to specify the properties returned when listing/getting artifacts
type Option struct {
	WithTag       bool
	TagOption     *tag.Option // only works when WithTag is set to true
	WithLabel     bool
	WithAccessory bool
}
