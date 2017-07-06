// Copyright 2017 orijtech. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package orijgo_test

import (
	"io"
	"log"
	"os"

	"github.com/orijtech/orijgo/archomp/v1"
)

func Example_archomp_Compress() {
	f, err := os.Create("allCompressed.zip")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	client := new(archomp.Client)
	rc, err := client.Compress(&archomp.Request{
		Resources: []*archomp.Resource{
			{URL: "https://storage.googleapis.com/archomp/demos/gears.gif"},
			{URL: "https://storage.googleapis.com/archomp/demos/Go.svg"},
			{URL: "https://storage.googleapis.com/archomp/demos/runPanorama.jpeg"},
			{URL: "https://storage.googleapis.com/archomp/demos/Usage.pdf"},
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	defer rc.Close()

	io.Copy(f, rc)
}
