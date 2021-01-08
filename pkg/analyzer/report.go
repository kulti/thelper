package analyzer

import (
	"go/token"

	"golang.org/x/tools/go/analysis"
)

type reports struct {
	reports []report
}

type report struct {
	pos    token.Pos
	format string
	args   []interface{}
}

func (rr *reports) Reportf(pos token.Pos, format string, args ...interface{}) {
	rr.reports = append(rr.reports, report{
		pos:    pos,
		format: format,
		args:   args,
	})
}

func (rr reports) Flush(pass *analysis.Pass) {
	for _, r := range rr.reports {
		pass.Reportf(r.pos, r.format, r.args...)
	}
}
