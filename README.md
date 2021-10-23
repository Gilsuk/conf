# conf

conf package detect file format automatically by It's extension

## installation

```bash
go get github.com/gilsuk/conf
```

## usage

config.json

```json
{
  "LogLevel": "production",
  "MaxAllowed": 10,
  "Port": 8080
}
```

config.yaml

```yaml
LogLevel: "production"
MaxAllowed: 10
Port: 8080
```

```go
type configuration struct {
    LogLevel string
    MaxAllowed  int
    Port int
}

confWithDefault := &configuration{
    LogLevel: "debug",
    MaxAllowed: 5,
    Port: 3000,
}

// out default config values to a file for example
conf.Out("config_example.json", confWithDefault)
conf.Out("config_example.json", confWithDefault)

// load config values from a file
// after loaded, config values are overrided
conf.Load("config.json", confWithDefault)
conf.Load("config.yaml", confWithDefault)

fmt.Printf("%+v", *confWithDefault)
```

output

```bash
{LogLevel:production MaxAllowed:10 Port:8080}
```
