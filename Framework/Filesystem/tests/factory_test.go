package tests

import (
	"github.com/kmsar/laravel-go/Framework/Contracts/Support"
	"github.com/kmsar/laravel-go/Framework/Filesystem"
	"github.com/kmsar/laravel-go/Framework/Filesystem/file"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestFactory(t *testing.T) {
	var factory = Filesystem.New(Filesystem.Config{
		Default: "local",
		Disks: map[string]Support.Fields{
			"local": {
				"driver": "local",
				"root":   "/Users/qbhy/project/go/goal-web/filesystem/tests",
				"perm":   os.ModePerm,
			},
			//"qiniu": {
			//	"driver":     "qiniu",
			//	"ttl":        3600, // 私有 url 有效期，单位秒
			//	"private":    true,
			//	"domain":     "your domain", // example: https://image.example.com"
			//	"bucket":     "your bucket",
			//	"access_key": "your access key",
			//	"secret_key": "your secret key",
			//},
		},
	})

	disks := []string{
		//"qiniu",
		"local",
	}

	for _, name := range disks {
		var (
			disk   = factory.Disk(name)
			putErr = disk.Put("/test/demo.txt", "goal")
		)
		assert.Nil(t, putErr, putErr)
		var files = disk.AllFiles("/test")
		assert.True(t, len(files) == 1)
		var contents = files[0].ReadString()
		assert.True(t, contents == "goal")
		assert.True(t, disk.Exists("/test/demo.txt"))
		assert.True(t, disk.GetVisibility("/test/demo.txt") == file.VISIBLE)
		assert.Nil(t, disk.Delete("/test/demo.txt"))
	}
}
