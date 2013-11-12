package main

import (
	"fmt"
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

type Framework struct {
	Name string
	Version string
	Language string
	Modules []string
	Create []string
}

type Project struct {
	Name string
	Version string
	Language string
	Additional_Languages []string
	Framework []string
	Base_Directory string
	Post string
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
	default:
		log.Fatal("No such action")
	}
}

func framework_action(action string, name string) {
	var framework string
	var framework_version string

	var config Framework
	switch action {
	case "new":
		log.Printf("Creating new framework: %s\n", name)
	case "install":
		framework_yaml := "../frameworks/"+name+"/init.yaml"
		file, err := ioutil.ReadFile(framework_yaml)
		if err != nil {
			log.Fatal(err)
		}
		if err := goyaml.Unmarshal(file, &config); err != nil {
			log.Fatal(err)
		}
		if len(config.Name) == 0 || len(config.Version) == 0 || len(config.Language) == 0 {
			log.Fatal("Frameworks must have a name, version and language")
		} else {
			framework = config.Name
			framework_version = config.Version
			log.Printf("Setting up framework: %s %s\n", framework, framework_version)
		}
		language_action("install", config.Language)
		if len(config.Modules) > 0 {
			for _,pkg := range config.Modules {
				log.Printf("Installing module: %s\n", pkg)
			}
		}
		if len(config.Create) > 0 {
			for _,step := range config.Create {
				log.Printf("Running: %s\n", step)
			}
		}
	}
}

func project_action(action string, name string) {
	var project string
	var project_version string

	var config Project
	switch action {
	case "new":
		log.Printf("Creating new project: %s\n", name)
	case "install":
		project_yaml := "../projects/"+name+"/init.yaml"
		file, err := ioutil.ReadFile(project_yaml)
		if err != nil {
			log.Fatal(err)
		}
		if err := goyaml.Unmarshal(file, &config); err != nil {
			log.Fatal(err)
		}
		if len(config.Name) == 0 || len(config.Version) == 0 {
			log.Fatal("Projects must have a name and version")
		} else {
			project = config.Name
			project_version = config.Version
			log.Printf("Setting up project: %s %s\n", project, project_version)
		}
		if len(config.Language) > 0 {
			language_action("install", config.Language)
		}
		if len(config.Additional_Languages) > 0 {
			for _,lang := range config.Additional_Languages {
				language_action("install", lang)
			}
		}
		if len(config.Framework) > 0 {
			for _,fwork := range config.Framework {
				framework_action("install", fwork)
			}
		}
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
  profile_tool language new <name> [<path>]
  profile_tool language install <name> [<path>]
  profile_tool framework new <name> [<path>]
  profile_tool framework install <name> [<path>]
  profile_tool project new <name> [<path>]
  profile_tool project start <name> [<path>]
  profile_tool -h | --help
  profile_tool --version

Options:
  -h 			Show this screen
  --version 		Show version
  <name> 		The name of the resource you're working with
  [<path>] 		Optional base directory for operations
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
	case "framework":
		framework_action(os.Args[2], os.Args[3])
	case "project":
		project_action(os.Args[2], os.Args[3])
	default:
		fmt.Println(usage)
		os.Exit(1)
	}
}
