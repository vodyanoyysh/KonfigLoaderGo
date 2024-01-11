# KonfigLoaderGO

## Overview
KonfigLoaderGo is a Go library for loading configuration files in YAML format and automatically substituting environment variables within them.

## Key Features
- **YAML Configuration Loading**: Easily load configurations from YAML files.
- **Environment Variable Substitution**: Replace placeholders in the configuration with actual values from environment variables.
- **Support for `.env` Files**: Seamlessly integrate with local `.env` files for development environments.
- **Easy to Integrate**: Designed to fit into existing Go projects with minimal setup.

## Getting Started
### Installation
Install `KonfigLoaderGo` using the go get command:

```bash
go get github.com/yourusername/KonfigLoaderGo
```

### usage

#### code:
```go
package main
import (
	"fmt"
	"github.com/vodyanoy420/KonfigLoaderGo"
)

type Config struct {
	Database struct {
		Postgresql struct {
			Host     string `json:"host"`
			Port     int    `json:"port"`
			User     string `json:"user"`
			Password string `json:"password"`
			Database string `json:"database"`
		} `json:"postgresql"`
	} `json:"database"`
}

func main() {
	var cfg Config
	KonfigLoaderGo.KonfigLoader(&cfg, "config.yaml")
	fmt.Println(cfg.Database.Postgresql.Password)
}
```

#### configfiles/config.yaml:
```yaml
database:
  postgresql:
    host: postgres
    port: 5432
    user: postgres
    password: $PASSWORD
    database: postgres
``` 