package systeminfo

import "os"

func init() {
	path := os.Getenv("IMAGE_STORE_PATH")
	if len(path) == 0 {
		path = "/data"
	}
	
}