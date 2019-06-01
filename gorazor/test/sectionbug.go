// This file is generated by gorazor 2.0
// DON'T modified manually
// Should edit source file and re-generate: /cases/sectionbug.gohtml

package cases

import (
	"bytes"
	"cases/layout"
	"github.com/sipin/gorazor/gorazor"
	"io"
	"kp/models"
	"strings"
)

func Sectionbug() string {
	var _b strings.Builder
	RenderSectionbug(&_b)
	return _b.String()
}

func RenderSectionbug(_buffer io.StringWriter) {

	_body := func(_buffer io.StringWriter) {

	}

	js := func(_buffer io.StringWriter) {
		for _, jsFile := range ctx.GetJS() {

			_buffer.WriteString("<script src=\"")
			_buffer.WriteString(gorazor.HTMLEscape(jsFile))
			_buffer.WriteString("\"></script>")

		}

	}

	return layout.Base(_buffer, body, nil, js)
}
