package repo

import (
	"context"

	"github.com/aide-cloud/moon/cmd/server/palace/internal/biz/bo"
	"github.com/aide-cloud/moon/pkg/helper/model/bizmodel"
)

type TeamMenuRepo interface {
	GetTeamMenuList(ctx context.Context, params *bo.QueryTeamMenuListParams) ([]*bizmodel.SysTeamMenu, error)
}
