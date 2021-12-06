# is-ci

Lightweight and dead-simple CI detection for golang.

This mod is based on the [@npmcli/ci-detect](https://www.npmjs.com/package/@npmcli/ci-detect) package.

## Install

```sh
go get -u github.com/wesleimp/is-ci
```

## Example 

```go
if isci.Check() {
  // do something
}
```

## LICENSE

[MIT](LICENSE)
