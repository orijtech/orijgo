package archomp_test

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/orijtech/orijgo/archomp/v1"
)

func Example() {
	req := &archomp.Request{
		Resources: []*archomp.Resource{
			{URL: "https://orijtech.com/favicon.ico", Name: "favicon"},
			{URL: "https://tatan.orijtech.com/v/Vg8Y1dfzz", Name: "gears.gif"},
			{URL: "https://scc-csc.lexum.com/scc-csc/scc-csc/en/item/2408/index.do", Name: "dunsmuir-case-law.htm"},
		},
	}

	client := new(archomp.Client)
	rc, err := client.Compress(req)
	if err != nil {
		log.Fatal(err)
	}
	defer rc.Close()

	f, err := os.Create("today.zip")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	n, err := io.Copy(f, rc)
	fmt.Printf("Wrote: %d bytes to disk, err: %v", n, err)
}
