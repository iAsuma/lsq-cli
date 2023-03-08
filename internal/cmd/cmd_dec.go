package cmd

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/gtag"
	"gitlab.corp.qizuang.net/infra/lsq-cli/internal/consts"
	"gitlab.corp.qizuang.net/infra/lsq-cli/utility/mlog"
	"gitlab.corp.qizuang.net/infra/lsq-cli/utility/utils"
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

	cDecArgName = `订单号`
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
		0: "计划任务",
		1: "go服务统一发单入口",
		2: "mobile站发单",
		3: "www站发单",
		4: "后台发单（CSM、CDM、OPS等）",
		5: "API接口发单（H5、App等php接口）",
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
