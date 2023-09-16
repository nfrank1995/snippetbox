# snippetbox
Small application to share small text snippets, based on the "Let's go" book by alex edwards

## run locally
To run the project switch into main project folder (snippetbox) and run the following command
```bash
go run ./cmd/web
```
### configuration
To configure the app there are some flags that can be set when starting the app. A list of all flags can be found by typing
```bash
go run ./cmd/web -help
```
 - f.e. the following command starts the server on address :9999
  ```bash
  go run ./cmd/web -addr=:9999
  ```
