package warp_generator //nolint:all

import (
	"fmt"
	"github.com/misnaged/warp_swagger/config_warp"
	"github.com/misnaged/warp_swagger/models"
	"github.com/misnaged/warp_swagger/warp_generator/external_packages"
	"github.com/misnaged/warp_swagger/warp_generator/handlers"
	"github.com/misnaged/warp_swagger/warp_generator/templater"
)

func Templates(config *config_warp.Warp,
	simpleNames, simpleTypes, customNames, customTypes []string, handlersModels *models.Handlers) ([]templater.ITemplate, error) {
	var templates templater.Templates
	external, err := external_packages.GenerateExternalModels(config, simpleNames, simpleTypes, customNames, customTypes)
	if err != nil {
		return nil, fmt.Errorf("failed while collecting templates: %w", err)
	}
	var hndlrs []templater.ITemplate
	for i := range handlersModels.Operations {
		hndlr, err := handlers.GenerateHandlers(handlersModels.Operations[i])
		if err != nil {
			return nil, fmt.Errorf("failed while collecting templates: %w", err)
		}
		hndlrs = append(hndlrs, hndlr)

	}
	hndlrs = append(hndlrs, external)
	templates = templater.GetAll(hndlrs...)
	return templates, nil
}
