/*
 * Copyright contributors to the Galasa project
 */
package scheduleryaml


type App struct {
	ApiVersion     string       `yaml:"apiVersion"`
	Kind           string       `yaml:"kind"`
	Metadata       AppMetadata  `yaml:"metadata"`
	Spec           AppSpec      `yaml:"spec"`
}

type AppMetadata struct {
	Name           string `yaml:"name"`
}

type AppSpec struct {
	Overrides[]    Override     `yaml:"overrides"`
}

type Level struct {
	ApiVersion     string         `yaml:"apiVersion"`
	Kind           string         `yaml:"kind"`
	Metadata       LevelMetadata  `yaml:"metadata"`
	Spec           LevelSpec      `yaml:"spec"`
}

type LevelMetadata struct {
	Name           string `yaml:"name"`
	App            string `yaml:"app"`
}

type LevelSpec struct {
	TestSelections[]  TestSelection  `yaml:"testSelection"` 
	Overrides[]       Override       `yaml:"overrides"`
	SyncCatalog       bool           `yaml:"syncCatalog" default:"false"`
	AutoRerun         bool           `yaml:"autoRerun"   default:"false"`
	RerunAttempts     int            `yaml:"rerunAttempts" default:"1"`
}

type TestSelection struct {
	Stream        string        `yaml:"stream"`
	Bundle        string        `yaml:"bundle"`
	Package       string        `yaml:"package"`
	Class         string        `yaml:"class"`
	Tag           string        `yaml:"tag"`
	TestArea      string        `yaml:"testArea"`
	Exclude       bool          `yaml:"exclude"`
}


type Override struct {
	Name           string       `yaml:"name"`
	Value          string       `yaml:"value"`
}
