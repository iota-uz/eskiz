package eskizapi_test

import (
	"bufio"
	"log"
	"os"
	"reflect"
	"strings"
)

type EnvConfiguration struct {
	URL      string `env:"URL"`
	Email    string `env:"EMAIL"`
	Password string `env:"PASSWORD"`
	Mobile   string `env:"MOBILE"`
}

func LoadEnvConfiguration() *EnvConfiguration {
	file, err := os.Open(".env")
	if err != nil {
		log.Fatalf("Error opening .env file: %v", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalf("Error closing .env file: %v", err)
		}
	}(file)

	envMap := make(map[string]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		if strings.HasPrefix(value, `"`) && strings.HasSuffix(value, `"`) && len(value) >= 2 {
			value = value[1 : len(value)-1]
		}

		envMap[key] = value
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading .env file: %v", err)
	}

	config := &EnvConfiguration{}
	configType := reflect.TypeOf(*config)
	configValue := reflect.ValueOf(config).Elem()

	for i := 0; i < configType.NumField(); i++ {
		field := configType.Field(i)
		if envKey, ok := field.Tag.Lookup("env"); ok {
			if value, exists := envMap[envKey]; exists {
				configValue.FieldByName(field.Name).SetString(value)
			}
		}
	}

	return config
}
