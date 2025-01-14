/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"log"
	"net/http"
	_ "net/http/pprof"

	"github.com/lampnick/doctron/cmd"
)

// go run main.go --config /opt/go_code/xyf-doctron/conf/default.yaml
func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	cmd.Execute()
}
