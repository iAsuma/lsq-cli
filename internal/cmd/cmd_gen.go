package cmd

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gtag"
)

var (
	Gen = cGen{}
)

type cGen struct {
	g.Meta `name:"gen" brief:"{cGenBrief}" dc:"{cGenDc}"`
}

const (
	cGenBrief = `automatically generate GoFrameV2 files for api/controller/logic/model`
	cGenDc    = `
The "gen" command is designed for multiple generating purposes. 
It's currently supporting generating go files for api standard files.
Please use "lsq gen mc -h" for specified type help.
`
)

func init() {
	gtag.Sets(g.MapStrStr{
		`cGenBrief`: cGenBrief,
		`cGenDc`:    cGenDc,
	})
}
