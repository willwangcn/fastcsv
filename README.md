# Fastcsv
Fastcsv is a csv parser for golang.

# Install
go get -t github.com/willwangcn/fastcsv

# Usage
#### Parsing
- `sep=","` If your data uses an alternate delimiter such as ; or \t.

- `header=false|true`  Set to true if you expect the first line of your CSV to contain headers, alternatly you can specify an array of headers using headers param. You can also specify a sparse array to omit some of the columns.

- `headers=slice|nil` If your data doesn't contain headers and want to specify the header fields.
**Notice** headers column name must equals struct's field name ignoring case.

**Example**

```go
type Student struct {
	Id       int64
	Name     string
	Age      int
	Lat, Lng float64
}
// right case, if you not config the field, it will not be initialized
headers = []{"name", "age", "lat"}
headers = []{"AGE", "NAme", "lng"}

// wrong case
headers = []{"s_name", "s_age"}
```

#### Test
```go
	data := bytes.NewBuffer([]byte(`
	Will,201601101716,18,40.654321,116.25820398331,2016-01-10 23:59:59
    Jack,201601101717,50,40.08296,116.316081,1990-01-31 23:59:59
    Tony,201601101718,44,40.060394,116.239552,1963-10-10 23:59:59`))
	columns := []string{"name", "id", "age", "lat", "lng", "birthday"}
	parser := fastcsv.NewFastcsv(data, ",", false, columns)
	students, err := parser.ReadAll(Student{})
	if err != nil {
		panic(err)
	}
	will := students[0].(Student)
	if will.Name != "Will" {
		t.Fail()
	}
```

#### How to Use
```go
    f, err := os.Open(filePath)
    defer f.Close()
    if err != nil {
        return
    }
    columns := []string{"id", "name", "age", "lat", "lng"}
    parser, err := fastcsv.NewFastcsv(f, ",", false, headers)
	if err != nil {
		panic(err)
	}
    students := parser.ReadAll(Student{})
    for _, val := range students {
        stu := val.(Student)
        fmt.Println(stu)
    }
```

# License 
[MIT License](https://raw.githubusercontent.com/willwangcn/fastcsv/master/LICENSE)

