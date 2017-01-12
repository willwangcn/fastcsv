package fastcsv

import (
	"testing"
	"bytes"
	"time"
)

type Student struct {
	Id       int64
	Name     string
	Age      int
	Lat, Lng float64
	Birthday time.Time
}

func TestContainingHeader(t *testing.T) {
	data := bytes.NewBuffer([]byte(`id,name,age,lat,lng
	201601101716,Will,18,40.654321,116.25820398331
201601101717,Jack,50,40.08296,116.316081
201601101718,Tony,44,40.060394,116.239552`))
	reader := NewFastcsv(data, ",", true, nil)
	students, err := reader.ReadAll(Student{})
	if err != nil {
		panic(err)
	}
	will := students[0].(Student)
	if will.Name != "Will" {
		t.Fail()
	}

}

func TestUsingHeaders(t *testing.T) {
	data := bytes.NewBuffer([]byte(`
	Will,201601101716,18,40.654321,116.25820398331,2016-01-10 23:59:59
Jack,201601101717,50,40.08296,116.316081,1990-01-31 23:59:59
Tony,201601101718,44,40.060394,116.239552,1963-10-10 23:59:59`))
	columns := []string{"name", "id", "age", "lat", "lng", "birthday"}
	parser := NewFastcsv(data, ",", false, columns)
	students, err := parser.ReadAll(Student{})
	if err != nil {
		panic(err)
	}
	will := students[0].(Student)
	if will.Name != "Will" {
		t.Fail()
	}
}