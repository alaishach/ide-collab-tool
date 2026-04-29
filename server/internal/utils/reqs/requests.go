// Package reqs
package reqs

import (
	"bytes"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"reflect"
	errgl "server/internal/err/global"
	"server/internal/err/panics"
	"server/internal/utils/funcs"
	"server/internal/utils/logger"

	"github.com/gin-gonic/gin"
)

func LogBody(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		logger.Logger.Error("Failed to read request body", "error", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.Request.Body = io.NopCloser(bytes.NewReader(body))

	slog.Info("Request body", "body", string(body))
}

func GetCookieValue(req *http.Request, name string) (string, error) {
	cookie, err := req.Cookie(name)
	if err != nil {
		return "", err
	}
	return cookie.Value, nil
}

func ParseBodyJSON[T any](w http.ResponseWriter, req *http.Request, data *T) error {
	rt := reflect.TypeOf(data)
	if rt.Kind() != reflect.Pointer || rt.Elem().Kind() != reflect.Struct {
		panics.PanicMisuse("ParseBodyJSON", "data is not a pointer to a struct")
	}
	deco := json.NewDecoder(req.Body)
	deco.DisallowUnknownFields()
	if err := deco.Decode(data); err != nil {
		logger.Logger.Debug("json unmarshalling request body failed: " + err.Error())
		return errgl.NewHTTPError(422, "invalid JSON, verify all required fields are present")
	}
	t, v := funcs.TypeVal(*data)
	for i := range t.NumField() {
		ptr := v.Field(i)
		if ptr.Kind() != reflect.Pointer {
			panics.PanicMisuse("ParseBodyJSON", "attribute is not a pointer type")
		}
		if ptr.IsNil() {
			field := t.Field(i)
			if field.Tag.Get("binding") == "required" {
				// TODO in prod maybe don't give the field name info away
				return errgl.NewHTTPError(400, "missing field: "+field.Name)

			} else if field.Tag.Get("default") == "" {
				println("TODO set up default bruuuh")
			}
		}
		// do field validation
		// value := ptr.Elem()
	}
	return nil
}
