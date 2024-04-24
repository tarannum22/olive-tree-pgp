package main

import (
	"time"

	"github.com/tarannum22/olive-tree-pgp/openpgp"
)

func main() {
	openpgp.EncryptFile("test/testfile.csv", "test/test_public.pub", "test/watermelon.pgp")
	time.Sleep(2*time.Second)
	openpgp.DecryptFile("test/watermelon.pgp", "test/test_private.asc", "test/watermelon.csv")
}