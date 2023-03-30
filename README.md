# txrollbackcheck

Forked from https://github.com/ryanrolds/sqlclosecheck

Linter that checks if SQL transaction are having `defer Rollback()`. are closed. Missing rollback may
cause DB connection pool exhaustion.

## Running

```
make build
make install
```

In your project directory:
```
go vet -vettool=$(which txrollbackcheck) ./...
```

## CI

```
go install github.com/herpiko/txrollbackcheck@latest
go vet -vettool=${GOPATH}/bin/txrollbackcheck ./...
```

