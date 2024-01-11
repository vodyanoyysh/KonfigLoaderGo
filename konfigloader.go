package KonfigLoaderGo

import (
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strings"
)

func KonfigLoader(config interface{}, filename string) {
	err := godotenv.Load()
	if err != nil {
		log.Printf(".env file not found, using the system environment variables")
	} else {
		log.Printf(".env file loaded")
	}

	data, err := ioutil.ReadFile(filepath.Join("configfiles", filename))
	if err != nil {
		log.Fatalf("error reading the file: %v", err)
	}

	err = yaml.Unmarshal(data, config)
	if err != nil {
		log.Fatalf("error unmarshalling the file: %v", err)
	}

	replaceEnvVariables(config)
}

func replaceEnvVariables(config interface{}) {
	val := reflect.ValueOf(config).Elem()
	replaceEnvVariablesRecursive(val)
}

func replaceEnvVariablesRecursive(val reflect.Value) {
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)

		switch field.Kind() {
		case reflect.String:
			replaceIfEnvVar(field)
		case reflect.Struct:
			replaceEnvVariablesRecursive(field)
		}
	}
}

func replaceIfEnvVar(field reflect.Value) {
	if field.Kind() != reflect.String {
		return
	}
	strVal := field.String()
	if strings.HasPrefix(strVal, "$") {
		envVar := strings.TrimPrefix(strVal, "$")
		envValue := os.Getenv(envVar)
		if envValue != "" {
			field.SetString(envValue)
		}
	}
}
