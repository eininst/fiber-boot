package fboot

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
	"html/template"
	"log"
	"os"
	"path"
	"strings"
	"sync"
)

const (
	defaultDocURL = "doc.json"
	defaultIndex  = "index.html"
	defaultJs     = "https://cdn.redoc.ly/redoc/latest/bundles/redoc.standalone.js"
	defaultCss    = "https://fonts.googleapis.com/css?family=Montserrat:300,400,700|Roboto:300,400,700"
)

type RedocXLogo struct {
	URL             string `json:"url"`
	BackgroundColor string `json:"backgroundColor"`
	AltText         string `json:"altText"`
}

type RedocConfig struct {
	Src    string
	Logo   RedocXLogo
	Theme  string
	JsCDN  string
	CssCDN string
}

func Redoc(docSrc string, config ...RedocConfig) fiber.Handler {
	buffer, err := os.ReadFile(docSrc)
	if err != nil {
		log.Fatal(err)
	}
	docjson := string(buffer)

	m := fiber.Map{
		"url":     defaultDocURL,
		"js_cdn":  defaultJs,
		"css_cdn": defaultCss,
		"theme":   "{}",
	}

	if len(config) > 0 {
		cfg := config[0]
		if cfg.JsCDN != "" {
			m["js_cdn"] = cfg.JsCDN
		}
		if cfg.CssCDN != "" {
			m["css_cdn"] = cfg.CssCDN
		}
		m["theme"] = cfg.Theme

		var h map[string]any
		s := gjson.Get(docjson, "info").String()
		_ = json.Unmarshal([]byte(s), &h)
		h["x-logo"] = cfg.Logo
		docjson, _ = sjson.Set(docjson, "info", h)
	}

	index, err := template.New("redoc_index.html").Parse(redocTpl)
	if err != nil {
		panic(fmt.Errorf("fiber: swagger middleware error -> %w", err))
	}

	var (
		prefix string
		once   sync.Once
		//fs     fiber.Handler = filesystem.New(filesystem.Config{Root: swaggerFiles.HTTP})
	)

	return func(c *fiber.Ctx) error {
		// Set prefix
		once.Do(func() {
			prefix = strings.ReplaceAll(c.Route().Path, "*", "")

			forwardedPrefix := getForwardedPrefix(c)
			if forwardedPrefix != "" {
				prefix = forwardedPrefix + prefix
			}
		})

		p := c.Path(utils.CopyString(c.Params("*")))

		switch p {
		case defaultIndex:
			c.Type("html")
			return index.Execute(c, m)
		case defaultDocURL:

			return c.Type("json").SendString(docjson)
		case "", "/":
			return c.Redirect(path.Join(prefix, defaultIndex), fiber.StatusMovedPermanently)
		default:
			return c.SendStatus(fiber.StatusNotFound)
		}
	}
}

func getForwardedPrefix(c *fiber.Ctx) string {
	header := c.GetReqHeaders()["X-Forwarded-Prefix"]

	if header == "" {
		return ""
	}

	prefix := ""

	prefixes := strings.Split(header, ",")
	for _, rawPrefix := range prefixes {
		endIndex := len(rawPrefix)
		for endIndex > 1 && rawPrefix[endIndex-1] == '/' {
			endIndex--
		}

		if endIndex != len(rawPrefix) {
			prefix += rawPrefix[:endIndex]
		} else {
			prefix += rawPrefix
		}
	}

	return prefix
}

const redocTpl = `
<!DOCTYPE html>
<html>
  <head>
    <title>Redoc</title>
    <!-- needed for adaptive design -->
    <meta charset="utf-8"/>
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <link href="{{.css_cdn}}" rel="stylesheet">

    <!--
    Redoc doesn't change outer page styles
    -->
    <style>
      body {
        margin: 0;
        padding: 0;
      }
    </style>
  </head>
  <body>
    <redoc spec-url='{{.url}}'></redoc>
    <script src="{{.js_cdn}}"> </script>
  </body>
</html>
`
