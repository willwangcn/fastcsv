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
	Will,201601101716,18,40.654321,116.25820398331
Jack,201601101717,50,40.08296,116.316081
Tony,201601101718,44,40.060394,116.239552`))
headers := []string{"name", "id", "lat", "lng"}
reader := NewFastcsv(data, ",", false, headers)
students := reader.ReadAll(Student{})
will := students[0]
if will.(Student).Name != "Will" {
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
headers := []string{"id", "name", "age", "lat", "lng"}
reader := fastcsv.NewFastcsv(f, ",", false, headers)
students := reader.ReadAll(Student{})
for _, val := range students {
    stu := val.(Student)
    fmt.Println(stu)
}
```

# License 
[MIT License](https://raw.githubusercontent.com/willwangcn/fastcsv/master/LICENSE)

