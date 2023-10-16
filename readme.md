ref: https://kritwis.medium.com/golang-clean-architecture-with-demo-e0938e5be02b

- generate mock with mockery:
  `bash
    mockery --dir=domains --output=domains/mocks --outpkg=mocks --all
`

- how to run test
  `bash
    go test ./... -v -cover
`

- clean test cache
  `bash
    go clean -testcache
`
