package handlers

import (
	"github.com/flosch/pongo2/v6"
	"github.com/gin-gonic/gin"
)

var templates map[string]*pongo2.Template

func InitTemplates(templateMap map[string]*pongo2.Template) {
	templates = templateMap
}

func RenderTemplate(c *gin.Context, templateName string, data pongo2.Context) {
	template, exists := templates[templateName]
	if !exists {
		c.String(404, "模板不存在")
		return
	}

	html, err := template.Execute(data)
	if err != nil {
		c.String(500, "模板渲染错误")
		return
	}

	c.Header("Content-Type", "text/html")
	c.String(200, html)
}
