package fastcsv

import (
	"testing"
	"bytes"
)

type Student struct {
	Id       int64
	Name     string
	Age      int
	Lat, Lng float64
}

func TestContainingHeader(t *testing.T) {
	data := bytes.NewBuffer([]byte(`id,name,age,lat,lng
	201601101716,Will,18,40.654321,116.25820398331
201601101717,Jack,50,40.08296,116.316081
201601101718,Tony,44,40.060394,116.239552`))
	reader := NewFastcsv(data, ",", true, nil)
	students := reader.ReadAll(Student{})
	will := students[0]
	if will.(Student).Name != "Will" {
		t.Fail()
	}

}

func TestUsingHeaders(t *testing.T) {
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
}