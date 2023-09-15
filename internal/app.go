package internal

import (
	"fmt"
	"github.com/gateway-fm/warp_swagger/config_swagger"
	"github.com/gateway-fm/warp_swagger/config_warp"
	"github.com/gateway-fm/warp_swagger/proto_parser"
	"github.com/gateway-fm/warp_swagger/swagger"
	"github.com/gateway-fm/warp_swagger/warp_generator"
	"github.com/gateway-fm/warp_swagger/warp_generator/templater"
)

type App struct {
	warpCfg   *config_warp.Warp
	swagCfg   *config_swagger.SwaggerCfg
	templates []templater.ITemplate
	swag      swagger.ISwagger
	dummy     swagger.IDummy
}

func (app *App) WarpCfg() *config_warp.Warp {
	return app.warpCfg
}
func (app *App) SwagCfg() *config_swagger.SwaggerCfg {
	return app.swagCfg
}

func NewApplication() (app *App, err error) {
	app = &App{
		warpCfg: &config_warp.Warp{},
		swagCfg: &config_swagger.SwaggerCfg{},
	}
	return
}

func (app *App) CallDummy() error {
	d := swagger.NewDummy()

	err := d.Call()
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	app.dummy = d
	return nil
}
func (app *App) SwaggerGenerate(args []string) error {
	swag := swagger.NewSwagger(app.SwagCfg())
	// fmt.Println(args[0])
	// fmt.Println(app.SwagCfg())

	if err := swag.Generate(args[0]); err != nil {
		return fmt.Errorf("failed to generate swagger:%w", err)
	}
	return nil
}
func (app *App) prepareTemplates() error {
	if err := app.CallDummy(); err != nil {
		return fmt.Errorf("failed to call dummy: %w", err)
	}
	protoParser := proto_parser.NewIProtoParser()
	// fmt.Println(app.WarpCfg().External.ProtoPath)

	primitive, custom, err := protoParser.Parse(app.WarpCfg().External.ProtoPath)
	if err != nil {
		return fmt.Errorf("PrepareTemplates returned an error:%w", err)
	}

	//  TODO: refactor: get rid off ***simpleNames, simpleTypes, customNames, customTypes []string***
	//   using something  similar to []*ParsedMsg is preferable
	var simpleNames, simpleTypes, customNames, customTypes []string

	for i := range primitive {
		simpleNames = append(simpleNames, primitive[i].ParsedNames)
		simpleTypes = append(simpleTypes, primitive[i].ParsedTypes)
	}
	for i := range custom {
		customNames = append(customNames, custom[i].ParsedNames)
		customTypes = append(customTypes, custom[i].ParsedTypes)
	}

	//   *******************-- TO DO --********************* //

	templates, err := warp_generator.Templates(app.WarpCfg(), simpleNames, simpleTypes, customNames, customTypes, app.dummy.GetHandlersModel())
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	app.templates = templates
	// for i := range app.templates {
	// 	fmt.Println(app.templates[i])
	// }
	return nil
}

func (app *App) Summon() error {
	if err := app.prepareTemplates(); err != nil {
		return fmt.Errorf("summon was interrupted:%w", err)
	}
	for _, t := range app.templates {
		err := t.Generate()
		if err != nil {
			return fmt.Errorf("summon was interrupted:%w", err)
		}
	}
	return nil
}
