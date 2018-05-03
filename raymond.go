// Package raymond provides handlebars evaluation
package raymond

// Render parses a template and renders it with given context
//
// Note that this function call is not optimal as your template is parsed everytime you call it. You should use Parse() function instead.
func Render(source string, ctx interface{}) (string, error) {
	// parse template
	tpl, err := Parse(source)
	if err != nil {
		return "", err
	}

	// renders template
	str, err := tpl.Exec(ctx)
	if err != nil {
		return "", err
	}

	return str, nil
}

// MustRender parses a template and evaluates it with given context. It panics on error.
//
// Note that this function call is not optimal as your template is parsed everytime you call it. You should use Parse() function instead.
func MustRender(source string, ctx interface{}) string {
	return MustParse(source).MustExec(ctx)
}

// Eval parses a template and evaluates it with given context.
//
// This is usefule to evaluate helper calls etc.
func Eval(source string, ctx interface{}) (interface{}, error) {
	// parse template
	tpl, err := Parse(source)
	if err != nil {
		return "", err
	}

	// evaluate template
	res, err := tpl.Eval(ctx)
	if err != nil {
		return "", err
	}

	return res, nil
}
