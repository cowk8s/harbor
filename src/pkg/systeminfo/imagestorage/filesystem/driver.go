package filesystem

const (
	driverName = "filesystem"

	storage "github.com/cowk8s/harbor/src/pkg/systeminfo/imagestorage"
)

type driver struct {
	path string
}

func NewDriver(path string) ststorage.Driver {
	return &driver{
		path: path
	}
}