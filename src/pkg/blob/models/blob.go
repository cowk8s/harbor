package models

import "time"

func init() {

}

const (
	StatusNone         = "none"
	StatusDelete       = "delete"
	StatusDeleting     = "deleting"
	StatusDeleteFailed = "deletefailed"
)

var StatusMap = map[string][]string{
	StatusNone:         {StatusNone, StatusDelete, StatusDeleteFailed},
	StatusDelete:       {StatusNone, StatusDelete, StatusDeleteFailed},
	StatusDeleting:     {StatusDelete},
	StatusDeleteFailed: {StatusDeleting},
}

type ArtifactAndBlob struct {
	ID           int64     `orm:"pk;auto;column(id)" json:"id"`
	DigestAF     string    `orm:"column(digest_af)" json:"digest_af"`
	DigestBlob   string    `orm:"column(digest_blob)" json:"digest_blob"`
	CreationTime time.Time `orm:"column(creation_time);auto_now_add" json:"creation_time"`
}
