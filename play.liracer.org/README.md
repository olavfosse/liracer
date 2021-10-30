To run liracer for development purposes run
```sh
CGO_ENABLED=0 go run -tags donotembed .
```

To build liracer for deployments run
```sh
CGO_ENABLED=0 go build -tags membed .
```

`CGO_ENABLED=0` might not be strictly necessary, but I ran into a problem today which was solved by adding this env variable so we might as well just use it from now on. See https://community.tmpdir.org/t/problem-with-go-binary-portability-lib-x86-64-linux-gnu-libc-so-6-version-glibc-2-32-not-found/123/2.