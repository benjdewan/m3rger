// Copyright © 2017 ben dewan <benj.dewan@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/alecthomas/kingpin"
	"github.com/imdario/mergo"
	yamlparser "gopkg.in/yaml.v2"
)

var version = "DEV_BUILD"

func main() {
	defFile, low, high, out := parseCLI()
	defYAML, err := readYAMLFile(defFile)
	if err != nil {
		log.Fatal(err)
	}
	if defYAML, err = overwrite(defYAML, low); err != nil {
		log.Fatal(err)
	}
	if defYAML, err = overwrite(defYAML, high); err != nil {
		log.Fatal(err)
	}

	if err = writeYAML(defYAML, out); err != nil {
		log.Fatal(err)
	}
}

func overwrite(yaml map[string]interface{}, file string) (map[string]interface{}, error) {
	if len(file) == 0 {
		return yaml, nil
	}
	overrideYAML, err := readYAMLFile(file)
	if err != nil {
		return yaml, err
	}
	return merge(yaml, overrideYAML)
}

func writeYAML(yaml map[string]interface{}, file string) error {
	data, err := yamlparser.Marshal(yaml)
	if err != nil {
		return err
	}
	if len(file) == 0 {
		fmt.Println(string(data))
		return nil
	}
	return ioutil.WriteFile(file, data, 0644)
}

func merge(def, overrides map[string]interface{}) (map[string]interface{}, error) {
	err := mergo.Map(&overrides, def)
	return overrides, err
}

func parseCLI() (string, string, string, string) {
	var (
		defFile = kingpin.Arg("default",
			"A YAML file with default values that may be overridden").
			Required().
			String()
		low = kingpin.Arg("low",
			"A YAML file that will override any values in 'default', but not any in the 'high' file").
			String()
		high = kingpin.Arg("high",
			"A YAML file whose keys will override anything from 'low' or 'default'").
			String()
		out = kingpin.Flag("out",
			"Write the output to a file (by default m3rger prints to stdout)").
			Short('o').
			String()
	)
	kingpin.Version(version)
	kingpin.Parse()
	return *defFile, *low, *high, *out
}

func readYAMLFile(file string) (map[string]interface{}, error) {
	yaml := make(map[string]interface{})
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return yaml, err
	}

	err = yamlparser.Unmarshal(data, &yaml)
	return yaml, err

}
