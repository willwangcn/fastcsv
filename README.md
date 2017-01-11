# Fastcsv
Fastcsv is a csv parser for golang.

# Install
go get https://github.com/willwangcn/fastcsv

# Usage
#### Parsing
- `sep=","` If your data uses an alternate delimiter such as ; or \t.

- `header=false|true`  Set to true if you expect the first line of your CSV to contain headers, alternatly you can specify an array of headers using headers param. You can also specify a sparse array to omit some of the columns.

- `headers=slice|nil` If your data doesn't contain headers and want to specify the header fields.
**Notice** headers column name must equals struct's field name ignoring case.

**Example**

```go
type Student struct {
	Name string
	Age  int
}
// right case
headers = []{"name", "age"}
headers = []{"AGE", "NAme"}

// wrong case
headers = []{"s_name", "s_age"}
```

#### how to use
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

# License 
[MIT License](https://raw.githubusercontent.com/willwangcn/fastcsv/master/LICENSE)

