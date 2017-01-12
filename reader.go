package fastcsv

import (
	"bufio"
	"fmt"
	"io"
	"reflect"
	"strings"
	"time"
)

type Fastcsv struct {
	scanner *bufio.Scanner
	sep     string
	header  bool
	columns []string
}

func NewFastcsv(reader io.Reader, sep string, header bool, headers []string) *Fastcsv {
	return &Fastcsv{
		scanner: bufio.NewScanner(reader),
		sep:     sep,
		header: header,
		columns : headers,
	}
}

func (parser *Fastcsv) ReadHeader() map[string]int {
	if !parser.scanner.Scan() {
		return nil
	}
	line := parser.scanner.Text()
	parts := strings.Split(line, parser.sep)
	fields := make(map[string]int)
	for i := 0; i < len(parts); i++ {
		fields[strings.ToLower(parts[i])] = i
	}
	return fields
}

func (parser *Fastcsv) Err() error {
	if err := parser.scanner.Err(); err != nil {
		return err
	}
	return nil
}

func (parser *Fastcsv) ReadAll(v interface{}) ([]interface{}, error) {
	var slice []interface{}

	if err := parser.Err(); err != nil {
		return slice, err
	}

	var headers map[string]int

	if parser.header {
		headers = parser.ReadHeader()
	} else if parser.columns != nil {
		headers = make(map[string]int)
		for i, val := range parser.columns {
			headers[strings.ToLower(val)] = i
		}
	}
	for parser.scanner.Scan() {
		line := strings.TrimSpace(parser.scanner.Text())
		if len(line) == 0 {
			continue
		}
		parts := strings.Split(line, parser.sep)
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
				idx, err := headers[strings.ToLower(ft.Name)]
				if !err {
					continue
				}
				part := strings.TrimSpace(parts[idx])
				switch fv.Kind() {
				case reflect.String:
					fv.SetString(part)
				case reflect.Ptr:
					fmt.Sscan(part, fv.Interface())
				case reflect.Struct:
					switch ft.Type.Name() {
					case "Time":
						date, err := time.Parse("2006-01-02 15:04:05", part)
						if err != nil {
							return slice, err
						}
						fv.Set(reflect.ValueOf(date))
					}
				default:
					fmt.Sscan(part, fv.Addr().Interface())
				}
			}
		}
		slice = append(slice, val.Interface())
	}

	return slice, nil
}