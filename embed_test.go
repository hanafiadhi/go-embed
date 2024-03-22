package goembed

import (
	"embed"
	"fmt"

	"io/fs"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed source/version.txt
var version string

func TestString(t *testing.T)  {
	assert.Equal(t,"Hai\n", version)
}

//go:embed source/ancol.jpeg
var image []byte
func TestByte(t *testing.T)  {
	err := ioutil.WriteFile("target/logo_new.png",image,fs.ModePerm)
	assert.Nil(t,err)
}

//go:embed source/a.txt
//go:embed source/b.txt
//go:embed source/c.txt

var files embed.FS
func TestMultiFiles(t *testing.T)  {
	a, _ := files.ReadFile("source/a.txt")
	assert.Equal(t,"AAA\n",string(a))
	b, _ := files.ReadFile("source/b.txt")
	assert.Equal(t,"BBB\n",string(b))
	c, _ := files.ReadFile("source/c.txt")
	assert.Equal(t,"CCC\n",string(c))
}

//go:embed source/*.txt
var path embed.FS
func TestPatchMatcher(t *testing.T)  {
	dir, _ := path.ReadDir("source")
	for _, entry := range dir {
		if !entry.IsDir() {
			fmt.Println(entry.Name())
			content, _  := path.ReadFile("source/"+ entry.Name())
			fmt.Println("Content : " + string(content))
		}
	}
}
