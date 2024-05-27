// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/aide-cloud/moon/cmd/server/palace/internal/biz/do/model"
)

func newSysTeamMemberRole(db *gorm.DB, opts ...gen.DOOption) sysTeamMemberRole {
	_sysTeamMemberRole := sysTeamMemberRole{}

	_sysTeamMemberRole.sysTeamMemberRoleDo.UseDB(db, opts...)
	_sysTeamMemberRole.sysTeamMemberRoleDo.UseModel(&model.SysTeamMemberRole{})

	tableName := _sysTeamMemberRole.sysTeamMemberRoleDo.TableName()
	_sysTeamMemberRole.ALL = field.NewAsterisk(tableName)
	_sysTeamMemberRole.SysTeamMemberID = field.NewUint32(tableName, "sys_team_member_id")
	_sysTeamMemberRole.SysTeamRoleID = field.NewUint32(tableName, "sys_team_role_id")

	_sysTeamMemberRole.fillFieldMap()

	return _sysTeamMemberRole
}

type sysTeamMemberRole struct {
	sysTeamMemberRoleDo

	ALL             field.Asterisk
	SysTeamMemberID field.Uint32 // 团队用户ID
	SysTeamRoleID   field.Uint32 // 团队角色ID

	fieldMap map[string]field.Expr
}

func (s sysTeamMemberRole) Table(newTableName string) *sysTeamMemberRole {
	s.sysTeamMemberRoleDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s sysTeamMemberRole) As(alias string) *sysTeamMemberRole {
	s.sysTeamMemberRoleDo.DO = *(s.sysTeamMemberRoleDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *sysTeamMemberRole) updateTableName(table string) *sysTeamMemberRole {
	s.ALL = field.NewAsterisk(table)
	s.SysTeamMemberID = field.NewUint32(table, "sys_team_member_id")
	s.SysTeamRoleID = field.NewUint32(table, "sys_team_role_id")

	s.fillFieldMap()

	return s
}

func (s *sysTeamMemberRole) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *sysTeamMemberRole) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 2)
	s.fieldMap["sys_team_member_id"] = s.SysTeamMemberID
	s.fieldMap["sys_team_role_id"] = s.SysTeamRoleID
}

func (s sysTeamMemberRole) clone(db *gorm.DB) sysTeamMemberRole {
	s.sysTeamMemberRoleDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s sysTeamMemberRole) replaceDB(db *gorm.DB) sysTeamMemberRole {
	s.sysTeamMemberRoleDo.ReplaceDB(db)
	return s
}

type sysTeamMemberRoleDo struct{ gen.DO }

type ISysTeamMemberRoleDo interface {
	gen.SubQuery
	Debug() ISysTeamMemberRoleDo
	WithContext(ctx context.Context) ISysTeamMemberRoleDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISysTeamMemberRoleDo
	WriteDB() ISysTeamMemberRoleDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISysTeamMemberRoleDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISysTeamMemberRoleDo
	Not(conds ...gen.Condition) ISysTeamMemberRoleDo
	Or(conds ...gen.Condition) ISysTeamMemberRoleDo
	Select(conds ...field.Expr) ISysTeamMemberRoleDo
	Where(conds ...gen.Condition) ISysTeamMemberRoleDo
	Order(conds ...field.Expr) ISysTeamMemberRoleDo
	Distinct(cols ...field.Expr) ISysTeamMemberRoleDo
	Omit(cols ...field.Expr) ISysTeamMemberRoleDo
	Join(table schema.Tabler, on ...field.Expr) ISysTeamMemberRoleDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISysTeamMemberRoleDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISysTeamMemberRoleDo
	Group(cols ...field.Expr) ISysTeamMemberRoleDo
	Having(conds ...gen.Condition) ISysTeamMemberRoleDo
	Limit(limit int) ISysTeamMemberRoleDo
	Offset(offset int) ISysTeamMemberRoleDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISysTeamMemberRoleDo
	Unscoped() ISysTeamMemberRoleDo
	Create(values ...*model.SysTeamMemberRole) error
	CreateInBatches(values []*model.SysTeamMemberRole, batchSize int) error
	Save(values ...*model.SysTeamMemberRole) error
	First() (*model.SysTeamMemberRole, error)
	Take() (*model.SysTeamMemberRole, error)
	Last() (*model.SysTeamMemberRole, error)
	Find() ([]*model.SysTeamMemberRole, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SysTeamMemberRole, err error)
	FindInBatches(result *[]*model.SysTeamMemberRole, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.SysTeamMemberRole) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISysTeamMemberRoleDo
	Assign(attrs ...field.AssignExpr) ISysTeamMemberRoleDo
	Joins(fields ...field.RelationField) ISysTeamMemberRoleDo
	Preload(fields ...field.RelationField) ISysTeamMemberRoleDo
	FirstOrInit() (*model.SysTeamMemberRole, error)
	FirstOrCreate() (*model.SysTeamMemberRole, error)
	FindByPage(offset int, limit int) (result []*model.SysTeamMemberRole, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISysTeamMemberRoleDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s sysTeamMemberRoleDo) Debug() ISysTeamMemberRoleDo {
	return s.withDO(s.DO.Debug())
}

func (s sysTeamMemberRoleDo) WithContext(ctx context.Context) ISysTeamMemberRoleDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s sysTeamMemberRoleDo) ReadDB() ISysTeamMemberRoleDo {
	return s.Clauses(dbresolver.Read)
}

func (s sysTeamMemberRoleDo) WriteDB() ISysTeamMemberRoleDo {
	return s.Clauses(dbresolver.Write)
}

func (s sysTeamMemberRoleDo) Session(config *gorm.Session) ISysTeamMemberRoleDo {
	return s.withDO(s.DO.Session(config))
}

func (s sysTeamMemberRoleDo) Clauses(conds ...clause.Expression) ISysTeamMemberRoleDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s sysTeamMemberRoleDo) Returning(value interface{}, columns ...string) ISysTeamMemberRoleDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s sysTeamMemberRoleDo) Not(conds ...gen.Condition) ISysTeamMemberRoleDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s sysTeamMemberRoleDo) Or(conds ...gen.Condition) ISysTeamMemberRoleDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s sysTeamMemberRoleDo) Select(conds ...field.Expr) ISysTeamMemberRoleDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s sysTeamMemberRoleDo) Where(conds ...gen.Condition) ISysTeamMemberRoleDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s sysTeamMemberRoleDo) Order(conds ...field.Expr) ISysTeamMemberRoleDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s sysTeamMemberRoleDo) Distinct(cols ...field.Expr) ISysTeamMemberRoleDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s sysTeamMemberRoleDo) Omit(cols ...field.Expr) ISysTeamMemberRoleDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s sysTeamMemberRoleDo) Join(table schema.Tabler, on ...field.Expr) ISysTeamMemberRoleDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s sysTeamMemberRoleDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISysTeamMemberRoleDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s sysTeamMemberRoleDo) RightJoin(table schema.Tabler, on ...field.Expr) ISysTeamMemberRoleDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s sysTeamMemberRoleDo) Group(cols ...field.Expr) ISysTeamMemberRoleDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s sysTeamMemberRoleDo) Having(conds ...gen.Condition) ISysTeamMemberRoleDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s sysTeamMemberRoleDo) Limit(limit int) ISysTeamMemberRoleDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s sysTeamMemberRoleDo) Offset(offset int) ISysTeamMemberRoleDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s sysTeamMemberRoleDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISysTeamMemberRoleDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s sysTeamMemberRoleDo) Unscoped() ISysTeamMemberRoleDo {
	return s.withDO(s.DO.Unscoped())
}

func (s sysTeamMemberRoleDo) Create(values ...*model.SysTeamMemberRole) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s sysTeamMemberRoleDo) CreateInBatches(values []*model.SysTeamMemberRole, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s sysTeamMemberRoleDo) Save(values ...*model.SysTeamMemberRole) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s sysTeamMemberRoleDo) First() (*model.SysTeamMemberRole, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysTeamMemberRole), nil
	}
}

func (s sysTeamMemberRoleDo) Take() (*model.SysTeamMemberRole, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysTeamMemberRole), nil
	}
}

func (s sysTeamMemberRoleDo) Last() (*model.SysTeamMemberRole, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysTeamMemberRole), nil
	}
}

func (s sysTeamMemberRoleDo) Find() ([]*model.SysTeamMemberRole, error) {
	result, err := s.DO.Find()
	return result.([]*model.SysTeamMemberRole), err
}

func (s sysTeamMemberRoleDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SysTeamMemberRole, err error) {
	buf := make([]*model.SysTeamMemberRole, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s sysTeamMemberRoleDo) FindInBatches(result *[]*model.SysTeamMemberRole, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s sysTeamMemberRoleDo) Attrs(attrs ...field.AssignExpr) ISysTeamMemberRoleDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s sysTeamMemberRoleDo) Assign(attrs ...field.AssignExpr) ISysTeamMemberRoleDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s sysTeamMemberRoleDo) Joins(fields ...field.RelationField) ISysTeamMemberRoleDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s sysTeamMemberRoleDo) Preload(fields ...field.RelationField) ISysTeamMemberRoleDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s sysTeamMemberRoleDo) FirstOrInit() (*model.SysTeamMemberRole, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysTeamMemberRole), nil
	}
}

func (s sysTeamMemberRoleDo) FirstOrCreate() (*model.SysTeamMemberRole, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysTeamMemberRole), nil
	}
}

func (s sysTeamMemberRoleDo) FindByPage(offset int, limit int) (result []*model.SysTeamMemberRole, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s sysTeamMemberRoleDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s sysTeamMemberRoleDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s sysTeamMemberRoleDo) Delete(models ...*model.SysTeamMemberRole) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *sysTeamMemberRoleDo) withDO(do gen.Dao) *sysTeamMemberRoleDo {
	s.DO = *do.(*gen.DO)
	return s
}
