package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/imdario/mergo"
	"gopkg.in/alecthomas/kingpin.v2"
	yamlparser "gopkg.in/yaml.v1"
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
