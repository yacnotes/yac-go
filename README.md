# Yac Notes API

## Configuration

Configuration is done via environment variables:

```yaml
YAC_ENV: debug # debug or release
YAC_PORT: 3000 # port on which the server listens
YAC_DB_DIR: ./.data/db # directory for the inbuilt NoSql database
YAC_LOG_TO_STDOUT: true # true is the only possible option for the moment, logfiles are planned
YAC_LOG_LEVEL: 0 # 0 - DEBUG ; 1 - INFO ; 2 - WARN ; 3 - ERROR ; 4 -PANIC
```

## Development

### Build

The main.go can be found in ```exe/```. Build it from the root directory via 
```go build -o .bin/yac-go exe/main.go```

### Hot reloading

The repository comes with a hot reloading tool for Go applications. Run ```air.exe``` 
on Windows to make use of this tool. Read more about it here: https://github.com/cosmtrek/air
