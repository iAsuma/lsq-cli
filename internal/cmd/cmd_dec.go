package cmd

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/gtag"
	"github.com/iasuma/lsq-cli/internal/consts"
	"github.com/iasuma/lsq-cli/utility/mlog"
	"github.com/iasuma/lsq-cli/utility/utils"
)

var (
	Dec = cDec{}
)

type cDec struct {
	g.Meta `name:"dec" brief:"{cDecBrief}" dc:"{cDecDc}"`
}

const (
	cDecBrief = `解密工具集`
	cDecDc    = `
The "dec" command is designed for .....
`

	cDecArgName = `snowflake ID`
)

const (
	tplSnowFlakeId       = "{TplSnowFlakeId}"
	tplGenTime           = "{TplGenTime}"
	tplGenTimestamp      = "{TplGenTimestamp}"
	tplGenMiLLiTimestamp = "{TplGenMiLLiTimestamp}"
	tplMachineId         = "{TplMachineId}"
	tplGenFrom           = "{TplGenFrom}"
)

var (
	machineIdMap = g.MapIntStr{
		0: "",
		1: "",
		2: "",
		3: "",
		4: "",
		5: "",
	}
)

func init() {
	gtag.Sets(g.MapStrStr{
		`cDecBrief`:   cDecBrief,
		`cDecDc`:      cDecDc,
		`cDecArgName`: cDecArgName,
	})
}

type (
	cDecInput struct {
		g.Meta `name:"dec" brief:"{cDecBrief}"`
		Key1   string `name:"key1" arg:"true" brief:"{cDecArgName}"`
	}
	cDecOutput struct{}
)

func (c *cDec) Index(ctx context.Context, in cDecInput) (out cDecOutput, err error) {

	if in.Key1 != "" && utils.IsOrderNo(in.Key1) {
		id := gconv.Int64(in.Key1)
		filler := utils.SnowFlake().Decrypt(id, consts.SnowFlakeTimeVector)

		logicContent := gstr.ReplaceByMap(consts.TemplateSnowFlakeIdInfo, g.MapStrStr{
			tplSnowFlakeId:       in.Key1,
			tplGenTime:           filler.Time.String(),
			tplGenTimestamp:      gconv.String(filler.Timestamp),
			tplGenMiLLiTimestamp: gconv.String(filler.MilliTimestamp),
			tplMachineId:         gconv.String(filler.MachineId),
			tplGenFrom:           machineIdMap[int(filler.MachineId)],
		})
		mlog.Print(gstr.Trim(logicContent))
	}
	return
}
