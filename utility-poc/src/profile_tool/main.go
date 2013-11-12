package main

import (
	"fmt"
	//"github.com/davecgh/go-spew/spew"
	"io/ioutil"
	"launchpad.net/goyaml"
	"log"
	"os"
)

type Language struct {
	Name     string
	Version  string
	Packages []string
	Prep     []string
}

func language_action(action string, name string) {
	var language string
	var language_version string

	var config Language
	switch action {
	case "new":
		log.Printf("Creating new language: %s\n", name)
	case "install":
		lang_yaml := "../languages/"+name+"/init.yaml"
		file, err := ioutil.ReadFile(lang_yaml)
		if err != nil {
			log.Fatal(err)
		}
		if err := goyaml.Unmarshal(file, &config); err != nil {
			log.Fatal(err)
		}
		if len(config.Name) == 0 || len(config.Version) == 0 {
			log.Fatal("Languages must have a name and a version")
		} else {
			language = config.Name
			language_version = config.Version
			log.Printf("Setting up language: %s %s\n", language, language_version)
		}
		if len(config.Packages) > 0 {
			for _,pkg := range config.Packages {
				log.Printf("Installing package: %s\n", pkg)
			}
		}
		if len(config.Prep) > 0 {
			for _,step := range config.Prep {
				log.Printf("Running: %s\n", step)
			}
		}
		//spew.Dump(config)
	default:
		log.Fatal("No such action")
	}
}

func main() {
/* Possible cli API */

usage := `
Try this for now:
	profile_tool language install python
	or
	profile_tool language new golang

Usage:
  profile_tool language new <name> [--basedir=<path>]
  profile_tool language install <name> [--basedir=<path>]
  profile_tool framework new <name> [--basedir=<path>]
  profile_tool framework install <name> [--basedir=<path>]
  profile_tool project new <name> [--basedir=<path>]
  profile_tool project start <name> [--basedir=<path>]
  profile_tool -h | --help
  profile_tool --version

Options:
  -h 			Show this screen
  --version 		Show version
  --basedir=<path> 	Base directory for operations
`

	// for now let's just spew the unmarshalled yaml
	if len(os.Args) <= 3 {
		fmt.Println(usage)
		os.Exit(1)
	}
	switch os.Args[1] {
	case "-h":
		fmt.Println(usage)
		os.Exit(1)
	case "language":
		language_action(os.Args[2], os.Args[3])
	default:
		fmt.Println(usage)
		os.Exit(1)
	}
}
