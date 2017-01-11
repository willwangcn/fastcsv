package fastcsv

import (
	"bufio"
	"fmt"
	"io"
	"reflect"
	"strings"
)

type Fastcsv struct {
	scanner *bufio.Scanner
	sep     string
	err     error
	header  bool
	headers []string
}

func NewFastcsv(reader io.Reader, sep string, header bool, headers []string) *Fastcsv {
	return &Fastcsv{
		scanner: bufio.NewScanner(reader),
		sep:     sep,
		header: header,
		headers : headers,
	}
}

func (reader *Fastcsv) ReadHeader() map[string]int {
	if !reader.scanner.Scan() {
		return nil
	}
	line := reader.scanner.Text()
	parts := strings.Split(line, reader.sep)
	fields := make(map[string]int)
	for i := 0; i < len(parts); i++ {
		fields[strings.ToLower(parts[i])] = i
	}
	return fields
}

func (reader *Fastcsv) Err() error {
	if err := reader.scanner.Err(); err != nil {
		return err
	}

	return reader.err
}

func (reader *Fastcsv) ResetErr() {
	reader.err = nil
}

func (reader *Fastcsv) ReadAll(v interface{}) []interface{} {
	var slice []interface{}

	if err := reader.Err(); err != nil {
		return slice
	}

	var headers map[string]int

	if reader.header {
		headers = reader.ReadHeader()
	} else if reader.headers != nil {
		headers = make(map[string]int)
		for i, val := range reader.headers {
			headers[strings.ToLower(val)] = i
		}
	}

	for reader.scanner.Scan() {
		line := strings.TrimSpace(reader.scanner.Text())
		if len(line) == 0 {
			continue
		}
		parts := strings.Split(line, reader.sep)
		val := reflect.New(reflect.TypeOf(v)).Elem()
		typ := val.Type()
		for i := 0; i < val.NumField(); i ++ {
			fv := val.Field(i)
			ft := typ.Field(i)
			if !fv.CanSet() {
				fmt.Printf("The %d th %s is unaccessible.\n", i, ft.Name)
				continue
			}
			if headers != nil {
				idx := headers[strings.ToLower(ft.Name)]
				part := strings.TrimSpace(parts[idx])
				switch fv.Kind() {
				case reflect.String:
					fv.SetString(part)
				case reflect.Ptr:
					fmt.Sscan(part, fv.Interface())
				default:
					fmt.Sscan(part, fv.Addr().Interface())
				}
			}
		}
		slice = append(slice, val.Interface())
	}

	return slice
}