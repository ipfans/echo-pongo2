// Copyright 2015 ipfans
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package pongo2

import (
	"net/http"
	"path"

	"github.com/flosch/pongo2"
	"github.com/labstack/echo"
)

const (
	ContentType    = "Content-Type"
	ContentLength  = "Content-Length"
	ContentBinary  = "application/octet-stream"
	ContentHTML    = "text/html"
	defaultCharset = "UTF-8"
)

const (
	_DEFAULT_TPL_SET_NAME = "DEFAULT"
)

// Options represents a struct for specifying configuration options for the Render middleware.
type Options struct {
	// Directory to load templates. Default is "templates"
	Directory string
}

var option Options

func prepareOptions(options []Options) Options {
	var opt Options
	if len(options) > 0 {
		opt = options[0]
	}

	// Defaults
	if len(opt.Directory) == 0 {
		opt.Directory = "templates"
	}

	return opt
}

func getContext(templateData interface{}) pongo2.Context {
	if templateData == nil {
		return nil
	}
	contextData, isMap := templateData.(map[string]interface{})
	if isMap {
		return contextData
	}
	return nil
}

func Pongo2() echo.MiddlewareFunc {
	return func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			err := h(ctx)
			if err != nil {
				return err
			}
			templateName := ctx.Get("template")
			if templateName == nil {
				http.Error(
					ctx.Response().Writer(),
					"Template in Context not defined.",
					500)
			}
			var contentType, encoding string
			var isString bool
			ct := ctx.Get("ContentType")
			if ct == nil {
				contentType = ContentHTML
			} else {
				contentType, isString = ct.(string)
				if !isString {
					contentType = ContentHTML
				}
			}
			cs := ctx.Get("charset")
			if cs == nil {
				encoding = defaultCharset
			} else {
				encoding, isString = cs.(string)
				if !isString {
					encoding = defaultCharset
				}
			}
			newContentType := contentType + "; charset=" + encoding
			templateNameValue, isString := templateName.(string)
			if isString {
				templateData := ctx.Get("data")
				var template = pongo2.Must(pongo2.FromFile(path.Join("templates", templateNameValue)))
				ctx.Response().Header().Set(ContentType, newContentType)
				err = template.ExecuteWriter(
					getContext(templateData), ctx.Response().Writer())
				if err != nil {
					http.Error(
						ctx.Response().Writer(), err.Error(), 500)
				}
			}
			return nil
		}
	}
}
