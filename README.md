# conf

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
conf.Out(conf.JSON, "config_example.json", confWithDefault)

// load config values from a file
// after loaded, config values are overried
conf.Load(conf.JSON, "config.json", confWithDefault)

fmt.Printf("%+v", *confWithDefault)
```

output

```bash
{LogLevel:production MaxAllowed:10 Port:8080}
```
