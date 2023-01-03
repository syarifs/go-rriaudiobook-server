package validators

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gorm.io/gorm"
)

var _DB *gorm.DB

func NewValidator(db *gorm.DB) {
	_DB = db
}

// Check data duplication
// Param table string
// Param field string
// Param id int
func Duplicate(table, field string, id int) validation.RuleFunc {
	return func(value interface{}) (err error) {
		var c int64
		var where []string
		field_name := strings.Split(field, "_")

		where = append(where, fmt.Sprintf("%s = '%s'", field, value))

		if id != 0 {
			where = append(where, fmt.Sprintf("id != %d", id))
		}

		_DB.Table(table).
			Where(strings.Join(where, " AND ")).Count(&c)

		if c != 0 {
			msg := fmt.Sprintf("duplicate %s %s.", strings.Join(field_name, " "), value.(string))
			err = errors.New(msg)
		}

		return
	}
}

func Email(id int) validation.RuleFunc {
	return func(value interface{}) (err error) {
		var c int64
		var where string

		if id != 0 {
			where = fmt.Sprintf("id != %d", id)
		}

		_DB.Table("users").
			Where("email = ?", value).
			Where(where).Count(&c)

		if c != 0 {
			msg := fmt.Sprintf("duplicate email %s.", value.(string))
			err = errors.New(msg)
		}

		return
	}
}

func FileType(file *multipart.FileHeader, types string, fallbackext string) validation.RuleFunc {
	return func(value interface{}) (err error) {
		files, _ := file.Open()
		defer files.Close()

		fileByte, _ := io.ReadAll(files)
		fileExt := filepath.Ext(file.Filename)
		contentType := http.DetectContentType(fileByte)

		if strings.Contains(types, "/*") {
			types = strings.TrimSuffix(types, "/*")
		}

		if strings.Contains(contentType, types) || fileExt == fallbackext {
			return
		}

		err = fmt.Errorf("file type must %s", types)
		return
	}
}
