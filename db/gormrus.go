package db

import (
	"database/sql/driver"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/sirupsen/logrus"
)

type gormLogger struct {
	name string
}

var (
	sqlRegexp                = regexp.MustCompile(`\?`)
	numericPlaceHolderRegexp = regexp.MustCompile(`\$\d+`)
)

func isPrintable(s string) bool {
	for _, r := range s {
		if !unicode.IsPrint(r) {
			return false
		}
	}
	return true
}

func (l *gormLogger) Print(values ...interface{}) {
	var messages []interface{}

	if len(values) > 1 {
		var (
			sql             string
			formattedValues []string
			level           = values[0]
		)

		code := strings.Split(fmt.Sprint(values[1]), "/")
		if len(code) > 0 {
			messages = []interface{}{l.name + " ", code[len(code)-1:]}
		}

		if level == "sql" {
			// duration
			messages = append(messages, fmt.Sprintf(" [%.2fms]", float64(values[2].(time.Duration).Nanoseconds()/1e4)/100.0))
			// sql

			for _, value := range values[4].([]interface{}) {
				indirectValue := reflect.Indirect(reflect.ValueOf(value))
				if indirectValue.IsValid() {
					value = indirectValue.Interface()
					if t, ok := value.(time.Time); ok {
						formattedValues = append(formattedValues, fmt.Sprintf("'%v'", t.Format("2006-01-02 15:04:05")))
					} else if b, ok := value.([]byte); ok {
						if str := string(b); isPrintable(str) {
							formattedValues = append(formattedValues, fmt.Sprintf("'%v'", str))
						} else {
							formattedValues = append(formattedValues, "'<binary>'")
						}
					} else if r, ok := value.(driver.Valuer); ok {
						if value, err := r.Value(); err == nil && value != nil {
							formattedValues = append(formattedValues, fmt.Sprintf("'%v'", value))
						} else {
							formattedValues = append(formattedValues, "NULL")
						}
					} else {
						formattedValues = append(formattedValues, fmt.Sprintf("'%v'", value))
					}
				} else {
					formattedValues = append(formattedValues, "NULL")
				}
			}

			// differentiate between $n placeholders or else treat like ?
			if numericPlaceHolderRegexp.MatchString(values[3].(string)) {
				sql = values[3].(string)
				for index, value := range formattedValues {
					placeholder := fmt.Sprintf(`\$%d([^\d]|$)`, index+1)
					sql = regexp.MustCompile(placeholder).ReplaceAllString(sql, value+"$1")
				}
			} else {
				formattedValuesLength := len(formattedValues)
				for index, value := range sqlRegexp.Split(values[3].(string), -1) {
					sql += value
					if index < formattedValuesLength {
						sql += formattedValues[index]
					}
				}
			}

			messages = append(messages, fmt.Sprintf(" [%v] ", strconv.FormatInt(values[5].(int64), 10)+" rows affected or returned"))
			messages = append(messages, sql)
		} else {
			messages = append(messages, values[2:]...)
		}
	}

	logrus.Debug(messages...)
}

func NewGormrus() *gormLogger {
	return NewGormrusWithName("gorm")
}

func NewGormrusWithName(name string) *gormLogger {
	return &gormLogger{
		name: name,
	}
}
