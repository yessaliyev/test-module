package courses

import "github.com/yessaliyev/test-module/pkg/client"

func Course() string {
	return client.NewClient()
}
