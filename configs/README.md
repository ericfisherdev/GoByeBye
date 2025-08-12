# configs/

Configuration files for the application.

The `/configs` directory contains configuration files in various formats (YAML, JSON, TOML) that control application behavior.

## Files

- `default.yaml` - Default application configuration
- Additional environment-specific configs can be added (dev.yaml, prod.yaml)

## Configuration Management

The application should:
1. Load default configuration
2. Override with environment-specific settings
3. Allow user overrides via command-line flags or environment variables

## Guidelines

- Use structured configuration (avoid flat key-value pairs)
- Document all configuration options
- Provide sensible defaults
- Validate configuration at startup
- Consider using a configuration library like Viper

## Example Usage

```go
import "github.com/spf13/viper"

viper.SetConfigName("default")
viper.AddConfigPath("./configs")
viper.ReadInConfig()
```