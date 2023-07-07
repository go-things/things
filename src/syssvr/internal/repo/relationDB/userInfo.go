package relationDB

import (
	"context"
	"github.com/i-Things/things/shared/def"
	"github.com/i-Things/things/shared/store"
	"gorm.io/gorm"
)

type UserInfoRepo struct {
	db *gorm.DB
}

func NewUserInfoRepo(in any) *UserInfoRepo {
	return &UserInfoRepo{db: store.GetCommonConn(in)}
}

type UserInfoFilter struct {
	UserNames []string
	UserName  string
	Phone     string
	Email     string
	Accounts  []string //账号查询 非模糊查询
}

func (p UserInfoRepo) accountsFilter(db *gorm.DB, accounts []string) *gorm.DB {
	db = db.Where("`userName` in ?", accounts).
		Or("`email` in ?", accounts).
		Or("`phone` in ?", accounts).
		Or("`wechat` in ?", accounts)
	return db
}

func (p UserInfoRepo) fmtFilter(ctx context.Context, f UserInfoFilter) *gorm.DB {
	db := p.db.WithContext(ctx)
	if len(f.UserNames) != 0 {
		db = db.Where("`userName` in ?", f.UserNames)
	}
	if len(f.Accounts) != 0 {
		db = p.accountsFilter(db, f.Accounts)
	}
	if f.UserName != "" {
		db = db.Where("`userName` like ?", "%"+f.UserName+"%")
	}
	if f.Phone != "" {
		db = db.Where("`phone` like ?", "%"+f.Phone+"%")
	}
	if f.Email != "" {
		db = db.Where("`email` like ?", "%"+f.Email+"%")
	}
	return db
}

func (g UserInfoRepo) Insert(ctx context.Context, data *SysUserInfo) error {
	result := g.db.WithContext(ctx).Create(data)
	return store.ErrFmt(result.Error)
}

func (g UserInfoRepo) FindOneByFilter(ctx context.Context, f UserInfoFilter) (*SysUserInfo, error) {
	var result SysUserInfo
	db := g.fmtFilter(ctx, f)
	err := db.First(&result).Error
	if err != nil {
		return nil, store.ErrFmt(err)
	}
	return &result, nil
}
func (p UserInfoRepo) FindByFilter(ctx context.Context, f UserInfoFilter, page *def.PageInfo) ([]*SysUserInfo, error) {
	var results []*SysUserInfo
	db := p.fmtFilter(ctx, f).Model(&SysUserInfo{})
	db = page.ToGorm(db)
	err := db.Find(&results).Error
	if err != nil {
		return nil, store.ErrFmt(err)
	}
	return results, nil
}

func (p UserInfoRepo) CountByFilter(ctx context.Context, f UserInfoFilter) (size int64, err error) {
	db := p.fmtFilter(ctx, f).Model(&SysUserInfo{})
	err = db.Count(&size).Error
	return size, store.ErrFmt(err)
}

func (g UserInfoRepo) Update(ctx context.Context, data *SysUserInfo) error {
	err := g.db.WithContext(ctx).Where("`userID` = ?", data.UserID).Save(data).Error
	return store.ErrFmt(err)
}

func (g UserInfoRepo) DeleteByFilter(ctx context.Context, f UserInfoFilter) error {
	db := g.fmtFilter(ctx, f)
	err := db.Delete(&SysUserInfo{}).Error
	return store.ErrFmt(err)
}

func (g UserInfoRepo) Delete(ctx context.Context, userID int64) error {
	err := g.db.WithContext(ctx).Where("`userID` = ?", userID).Delete(&SysUserInfo{}).Error
	return store.ErrFmt(err)
}
func (g UserInfoRepo) FindOne(ctx context.Context, userID int64) (*SysUserInfo, error) {
	var result SysUserInfo
	err := g.db.WithContext(ctx).Where("`userID` = ?", userID).First(&result).Error
	return &result, store.ErrFmt(err)
}
