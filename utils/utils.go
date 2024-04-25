package utils

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/tarannum22/olive-tree-pgp/jobs"

	"gopkg.in/yaml.v3"
)

type AppConfig struct {
	Instance string
	Level    string
	Jobs     []jobs.Job
}

func LoadAppConfig(configFilename string, config *AppConfig) {
	f, err := os.Open(configFilename)
	if err != nil {
		slog.Error("Config File Error")
		panic(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(config)
	slog.Info("Configuration loaded successfully")
}

func PrintConfigValues(config AppConfig) {
	fmt.Printf("LogLevel : %s\n", config.Level)
	fmt.Printf("Instance : %s\n", config.Instance)

	fmt.Println("Jobs : ")
	for _, job := range config.Jobs {
		// fmt.Printf("%+v\n", job)
		fmt.Println(job.Name)
		fmt.Println(job.JobType)
		fmt.Println(job.Keyfile)
		fmt.Println(job.Passphrase)
		fmt.Println(job.Source)
		fmt.Println(job.Destination)
		fmt.Println(job.Schedule)
		fmt.Println("-----")
	}
}
