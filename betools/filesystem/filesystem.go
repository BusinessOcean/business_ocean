package filesystem

import (
	"context"
	// "gocloud.dev/blob"
)

var gcpIgnoreHeaders = []string{"Accept-Encoding"}

type System struct {
	ctx context.Context
	// bucket *blob.Bucket
}
