package cmd

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/util/gtag"
)

var (
	LSQ = cLSQ{}
)

type cLSQ struct {
	g.Meta `name:"lsq" ad:"{cLSQAd}"`
}

const (
	cLSQAd = `
ADDITIONAL
    Use "lsq COMMAND -h" for details about a command.
`
)

func init() {
	gtag.Sets(g.MapStrStr{
		`cLSQAd`: cLSQAd,
	})
}

type cLSQInput struct {
	g.Meta  `name:"lsq"`
	Version bool `short:"v" name:"version" brief:"show version information of current binary"   orphan:"true"`
}
type cLSQOutput struct{}

func (c cLSQ) Index(ctx context.Context, in cLSQInput) (out *cLSQOutput, err error) {
	// Version.
	if in.Version {
		_, err = Version.Index(ctx, cVersionInput{})
		return
	}

	gcmd.CommandFromCtx(ctx).Print()
	return
}
