## google-shortener-api
A library for using google url shortener api written in golang

#### Installation :

```bash
  go get github.com/ahmdrz/golang-shortener-api
```

#### Simple code :

```go
package main

import (
	"fmt"

	"github.com/ahmdrz/golang-shortener-api"
)

func main() {
	google, err := googleshortener.New("YOUR_GOOGLE_KEY")
	if err != nil {
		panic(err)
	}
	res := googleshortener.Response{}

	res, _ = google.MakeShort("http://yahoo.com")
	fmt.Println(res.Id)

	res, _ = google.MakeLong(res.Id)
	fmt.Println(res.LongUrl)
}
```

#### Output :

```bash
http://goo.gl/r5Pl
http://yahoo.com/
```
