# test

Such as test: https://github.com/go-gomail/gomail

Then generate `github.com/go-gomail/gomail.go` for test.

## Use Config
```go
import "github.com/vcgo/test"

func main() {
    test.Config.Get("xxx").(string)
}
```