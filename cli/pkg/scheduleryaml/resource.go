/*
 * Copyright contributors to the Galasa project
 */
package scheduleryaml


type App struct {
	ApiVersion     string       `yaml:"apiVersion"`
	Kind           string       `yaml:"kind"`
	Metadata       Metadata     `yaml:"metadata"`
	Spec           AppSpec      `yaml:"spec"`
}

type AppSpec struct {
	Overrides[]    Override     `yaml:"overrides"`
}

type Override struct {
	Name           string       `yaml:"name"`
	Value          string       `yaml:"value"`
}

type Metadata struct {
	Name           string `yaml:"name"`
}
