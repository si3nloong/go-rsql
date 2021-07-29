# Go RSQL

## 🔨 Installation

```console
go get github.com/si3nloong/go-rsql
```

```go

type QueryParams struct {
    Name   string  `rsql:"n,filter,sort,allow=eq|gt|gte"`
    Status string  `rsql:"status,filter"`
    PtrStr *string `rsql:"text,filter"`
    No     int     `rsql:"no,column=No2,filter"`
}

func main() {
    p := MustNew(i)

    params, err := p.ParseQuery(`filter=status=eq="111";no=gt=1991;text==null&sort=status,-no`)
    if err != nil {
        panic(err)
    }

    log.Println(params.Filters)
    log.Println(params.Sorts)
}
```

## 📄 License

[MIT](https://github.com/si3nloong/go-rsql/blob/master/LICENSE)

Copyright (c) 2020-present, SianLoong Lee