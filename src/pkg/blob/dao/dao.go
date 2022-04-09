package dao

import (
	"context"
	"time"

	"github.com/cowk8s/harbor/src/common/models"
	"github.com/cowk8s/harbor/src/lib/orm"
)

type DAO interface{
	// CreateArtifactAndBlob create ArtifactAndBlob and ignore conflict on artifact digest and blob digest
	CreateArtifactAndBlob(ctx context.Context, artifactDigest, blobDigest string) (int64, error)

	// GetArtifactAndBlob get ArtifactAndBlob by artifact digest and blob digest
	GetArtifactAndBlob(ctx context.Context, artifactDigest, blobDigest string) (*models.ArtifactAndBlob, error)

	
}

// New returns an instance of the defautl DAO
func New() DAO {
	return &dao{}
}

type dao struct{}

func (d *dao) CreateArtifactAndBlob(ctx context.Context, artifactDigest, blobDigest string) (int64, error) {
	o, err := orm.FromContext(ctx)
	if err != nil {
		return 0, err
	}

	md := &models.ArtifactAndBlob{
		DigestAF: artifactDigest,
		DigestBlob: blobDigest,
		CreationTime: time.Now(),
	}

	return o.InsertOrUpdate(md, "digest_af, digest_blob")
}
