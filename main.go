package main

import (
	"bytes"
	"fmt"

	"github.com/ipfs/go-ipfs-api"
)

const (
	url = "localhost:5001"
)

var (
	MAGIC_BYTES = fmt.Sprintf("%c%c", 8, 1)
)

func main() {
	sh := shell.NewShell(url)
	hash1, err := sh.AddNoPin(bytes.NewBufferString("Hello IPFS Shell tests"))
	if err != nil {
		panic(err)
	}
	fmt.Println("hash1:", hash1)

	hash2, err := sh.AddNoPin(bytes.NewBufferString("I want to test hash."))
	if err != nil {
		panic(err)
	}
	fmt.Println("hash2:", hash2)

	hash3, err := sh.AddNoPin(bytes.NewBufferString("I want to go work out."))
	if err != nil {
		panic(err)
	}
	fmt.Println("hash3:", hash3)

	rootObj := &shell.IpfsObject{
		Links: []shell.ObjectLink{},
		Data:  MAGIC_BYTES,
	}

	root, err := sh.ObjectPut(rootObj)
	if err != nil {
		panic(err)
	}
	fmt.Println("root before:", root)

	root1, err := sh.PatchLink(root, "hash1", hash1, false)
	if err != nil {
		panic(err)
	}
	fmt.Println("root1:", root1)

	root2, err := sh.PatchLink(root1, "hash2", hash2, false)
	if err != nil {
		panic(err)
	}
	fmt.Println("root2:", root2)

	root3, err := sh.PatchLink(root2, "hash3", hash3, false)
	if err != nil {
		panic(err)
	}
	fmt.Println("root3:", root3)

	dir, _ := sh.ObjectGet("QmeYye6wzMCAG67vQtYn1pAF3p89pWnfBx2GcLMeJ6XNBP")

	dirList, _ := sh.List("/ipfs/QmeYye6wzMCAG67vQtYn1pAF3p89pWnfBx2GcLMeJ6XNBP")
	fmt.Println("dirList:", dirList)

	fmt.Println("dir.Data:", dir.Data)
	fmt.Println("rootObj.Data:", rootObj.Data)

	rootList, _ := sh.List(fmt.Sprintf("/ipfs/%s", root3))
	fmt.Println("rootList:", rootList)
}
