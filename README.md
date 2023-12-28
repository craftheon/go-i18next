go-i18next
---
go-i18next is a port of [i18next](https://www.i18next.com/) in Go.

## Installation
```shell
$ go get -u https://github.com/yuangwei/go-18next
```
## Usage
```go
import "github.com/yuangwei/go-18next"

var i18n i18next.I18n

func main() {
	i18n = i18next.Init(i18next.I18nOptions{
                Lng:        []string{"en", "cn"},
                DefaultLng: "en",
                Ns:         "json",
                Backend: i18next.Backend{
                LoadPath: []string{"./examples/datas/{{.Ns}}/{{.Lng}}/home.json"},
            },
        })
	text, err := i18n.T("title", struct {
            Name string
    }{Name: "Mike"})
	fmt.Println(val) // Hello, Mike
	
}

```

## License
go-i18next is available under the MIT license. See the [LICENSE](LICENSE) file for more info.