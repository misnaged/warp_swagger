package middlewares

import (
	"github.com/gateway-fm/warp_swagger/config_warp"
	"github.com/gateway-fm/warp_swagger/warp_generator/templater"
)

func GenerateHandlers(
	config *config_warp.Warp,
) (templater.ITemplate, error) {
	path := "templates/internal/middlewares.gohtml"

	elems := "middlewares_main"
	ifaces := templater.GetTemplateInterfaces(elems)
	template := templater.NewTemplate(path, config.Middlewares.Output, ifaces, nil, elems)
	return template, nil
}
