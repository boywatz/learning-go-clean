ref: https://kritwis.medium.com/golang-clean-architecture-with-demo-e0938e5be02b

- generate mock with mockery:

  ```bash
    brew install mockery
    mockery --dir=domains --output=domains/mocks --outpkg=mocks --all
  ```

- pre-commit

  ```bash
    brew install pre-commit
    pre-commit install
  ```

- how to run test

  ```bash
    go test ./... -v -cover
  ```

- clean test cache

```bash
  go clean -testcache
```

- clean module cache

```bash
  go clean -modcache
```

- docker cmd

```bash
  docker rm -f $(docker ps -a -q)
  docker rmi $(docker images -q)
```
