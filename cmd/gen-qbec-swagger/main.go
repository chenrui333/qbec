/*
   Copyright 2019 Splunk Inc.

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ghodss/yaml"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatalln("Usage: gen-qbec-swagger <input-YAML-file> <go-file>")
	}

	inFile := os.Args[1]
	outFile := os.Args[2]

	handle := func(e error) {
		if e != nil {
			log.Fatalln(e)
		}
	}

	b, err := os.ReadFile(inFile)
	handle(err)

	var data interface{}
	err = yaml.Unmarshal(b, &data)
	handle(err)

	b, err = json.MarshalIndent(data, "", "    ")
	handle(err)

	code := fmt.Sprintf(`package model

// generated by gen-qbec-swagger from %s at %v
// Do NOT edit this file by hand

var swaggerJSON = %s
%s
%s
`, inFile, time.Now().UTC(), "`", b, "`")

	err = os.WriteFile(outFile, []byte(code), 0644)
	handle(err)
	log.Println("Successfully wrote", outFile, "from", inFile)
}
