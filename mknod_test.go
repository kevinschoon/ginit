package ginit

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"path/filepath"
	"syscall"
	"testing"
)

func TestMknod(t *testing.T) {
	dir, err := ioutil.TempDir("/tmp", "gaffer")
	assert.NoError(t, err)
	defer os.RemoveAll(dir)
	path := filepath.Join(dir, "gaffer.fifo")
	assert.NoError(t, Mknod(path, 0600, syscall.S_IFIFO, 0, 0))
	go func(path string) {
		fd, err := os.OpenFile(path, os.O_WRONLY, 0600)
		assert.NoError(t, err)
		_, err = fd.Write([]byte("hi!"))
		assert.NoError(t, err)
		assert.NoError(t, fd.Close())
	}(path)
	raw, err := ioutil.ReadFile(path)
	assert.NoError(t, err)
	assert.Equal(t, "hi!", string(raw))
}
