package handlers

import (
	"github.com/go-openapi/inflect"
	"github.com/misnaged/warp_swagger/models"
	"github.com/misnaged/warp_swagger/warp_generator/templater"
)

func GenerateHandlers(
	operation *models.Operation,
) (templater.ITemplate, error) {
	path := "templates/internal/handler_template.gohtml"

	var OperaIDuc = func() string {
		return inflect.Capitalize(operation.OperaID)
	}
	var OperaIDlc = func() string {
		return operation.OperaID
	}
	var funcNames = []string{
		"OperaIDuc",
		"OperaIDlc",
	}

	funcs := templater.GetTemplateInterfaces(
		OperaIDuc,
		OperaIDlc,
	)
	funcMap := templater.CompleteFuncMap(funcNames, funcs)
	elems := "handler_main"
	ifaces := templater.GetTemplateInterfaces(operation)
	template := templater.NewTemplate(path, operation.OutputFileName, ifaces, funcMap, elems)
	return template, nil
}
