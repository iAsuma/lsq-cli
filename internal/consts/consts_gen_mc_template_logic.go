package consts

const TemplateLogicContent = `
// =================================================================================
// This is auto-generated by lsq CLI tool only once. Fill this file as you wish.
// =================================================================================

package logic

import "context"

var (
	inl{TplLogicCamelCase} = l{TplLogicCamelCase}{}
)

type l{TplLogicCamelCase} struct{}

func {TplLogicCamelCase}() *l{TplLogicCamelCase}{
	return &inl{TplLogicCamelCase}
}

// Fill with you ideas below and delete this line.

func (l *l{TplLogicCamelCase}) FirstCall(ctx context.Context) error {
	return nil
}
`
