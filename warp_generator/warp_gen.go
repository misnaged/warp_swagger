package warp_generator

import (
	"fmt"
	"github.com/misnaged/annales/logger"
	"go/format"
	"os"
	"path/filepath"
	"sync"
	"text/template"
)

type ITemplate interface {
	Generate() error
	Excluded() bool
	GenerateNonGo() error
}

type Template struct {
	Elems                              []string
	Ifaces                             []interface{}
	ConfigTemplatePath, OutPutFilePath string
	FuncMap                            template.FuncMap
	IsExcluded                         bool
}

func (t *Template) GenerateFile() error {

	elems := t.Elems
	ifaces := t.Ifaces
	// as an empty .go file and just "filled up" in this func
	file, _ := os.Create(t.OutPutFilePath) //
	defer file.Close()

	// path to template file is absolute here, but it doesn't make any sense :D
	pattern, _ := filepath.Abs(t.ConfigTemplatePath) // .gotmpl is used because of IDE's supports only :D

	// template final preparation. Template must parse given pattern (which is our scheme.gotmpl file)
	tmpl := template.Must(template.New("").Funcs(t.FuncMap).ParseFiles(pattern))
	var wg sync.WaitGroup
	for i := range elems {
		wg.Add(1)
		go func() {
			defer wg.Done()
			err := tmpl.ExecuteTemplate(file, elems[i], ifaces[i]) // first arg is output, second is the data we want to pass to this config. It could also be nil.
			if err != nil {
				logger.Log().Errorf("An error occurred %s", err)
				return
			}
		}()
		wg.Wait()
	}

	return nil
}
func (t *Template) Excluded() bool {
	return t.IsExcluded
}

// GoFmt execute go fmt for specific .go file from the code
func GoFmt(path string) error {

	read, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	fmted, err := format.Source(read)
	if err != nil {
		return err
	}
	err = os.WriteFile(path, fmted, 0666)
	if err != nil {
		return err
	}
	return nil
}

// Generate is main generation function
func (t *Template) Generate() error {
	err := t.GenerateFile()
	if err != nil {
		return fmt.Errorf(" GenerateScheme returned an error: %w", err)
	}
	err = GoFmt(t.OutPutFilePath)
	if err != nil {
		return fmt.Errorf(" go fmt returned an error: %w", err)
	}

	return nil
}

// GenerateNonGo is main generation function for non .go files
func (t *Template) GenerateNonGo() error {
	err := t.GenerateFile()
	if err != nil {
		return fmt.Errorf(" GenerateScheme returned an error: %w", err)
	}
	return nil
}
