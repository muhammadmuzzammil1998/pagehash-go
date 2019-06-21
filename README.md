
# pagehash-go
Go library for [Page Hash](https://github.com/muhammadmuzzammil1998/Page-Hash).

## Installing pagehash package

```sh
go get muzzammil.xyz/pagehashgo
```

## Documentation

- [Page Hash documentation](https://github.com/muhammadmuzzammil1998/Page-Hash#documentation)
- [GoDoc for Page Hash](https://godoc.org/muzzammil.xyz/pagehashgo)

The API returns a JSON object which is then unmarshalled into Go structure as:
```go
type PageHash struct {
  Load   int64  `json:"load"`
  URL    string `json:"url"`
  Hashes []hash `json:"hashes"`
}

type hash struct {
  Algorithm string `json:"algo"`
  Hash      string `json:"hash"`
}
```
## Examples

### Printing hashes
```go
package main

import (
  "fmt"
  "muzzammil.xyz/pagehashgo"
)

func main() {
  if resp, err := pagehash.Get("example.com"); err == nil {
    for _, hash := range resp.Hashes {
      fmt.Printf("%s:\t%s\n", hash.Algorithm, hash.Hash)
    }
  }
}
```
```bash
$ go build
$ ./pagehash-test
sha256: 3587cb776ce0e4e8237f215800b7dffba0f25865cb84550e87ea8bbac838c423
sha1:   0e973b59f476007fd10f87f347c3956065516fc0
md5:    09b9c392dc1f6e914cea287cb6be34b0
```

### Printing a specific hash
```go
package main

import (
  "fmt"
  "muzzammil.xyz/pagehashgo"
)

func main() {
  if resp, err := pagehash.Get("example.com"); err == nil {
    fmt.Println("SHA256:", resp.GetSHA256())
  }
}
```
```bash
$ go build
$ ./pagehash-test
SHA256: 3587cb776ce0e4e8237f215800b7dffba0f25865cb84550e87ea8bbac838c423
```
