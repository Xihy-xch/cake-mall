package model

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/core/stores/sqlc"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/core/stringx"
	"github.com/tal-tech/go-zero/tools/goctl/model/sql/builderx"
)

var (
	wechatUserFieldNames          = builderx.RawFieldNames(&WechatUser{})
	wechatUserRows                = strings.Join(wechatUserFieldNames, ",")
	wechatUserRowsExpectAutoSet   = strings.Join(stringx.Remove(wechatUserFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	wechatUserRowsWithPlaceHolder = strings.Join(stringx.Remove(wechatUserFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheWechatUserIdPrefix = "cache#wechatUser#id#"
)

type (
	WechatUserModel interface {
		Insert(data WechatUser) (sql.Result, error)
		FindOne(id int64) (*WechatUser, error)
		Update(data WechatUser) error
		Delete(id int64) error
	}

	defaultWechatUserModel struct {
		sqlc.CachedConn
		table string
	}

	WechatUser struct {
		CreateTime time.Time `db:"create_time"`
		UpdateTime time.Time `db:"update_time"`
		Id         int64     `db:"id"`          // 主键自增id
		Number     int64     `db:"number"`      // 用户唯一number
		UnionId    string    `db:"union_id"`    // 用户微信unionId
		SessionKey string    `db:"session_key"` // 用户微信sessionKey
	}
)

func NewWechatUserModel(conn sqlx.SqlConn, c cache.CacheConf) WechatUserModel {
	return &defaultWechatUserModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`wechat_user`",
	}
}

func (m *defaultWechatUserModel) Insert(data WechatUser) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, wechatUserRowsExpectAutoSet)
	ret, err := m.ExecNoCache(query, data.Number, data.UnionId, data.SessionKey)

	return ret, err
}

func (m *defaultWechatUserModel) FindOne(id int64) (*WechatUser, error) {
	wechatUserIdKey := fmt.Sprintf("%s%v", cacheWechatUserIdPrefix, id)
	var resp WechatUser
	err := m.QueryRow(&resp, wechatUserIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", wechatUserRows, m.table)
		return conn.QueryRow(v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultWechatUserModel) Update(data WechatUser) error {
	wechatUserIdKey := fmt.Sprintf("%s%v", cacheWechatUserIdPrefix, data.Id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, wechatUserRowsWithPlaceHolder)
		return conn.Exec(query, data.Number, data.UnionId, data.SessionKey, data.Id)
	}, wechatUserIdKey)
	return err
}

func (m *defaultWechatUserModel) Delete(id int64) error {

	wechatUserIdKey := fmt.Sprintf("%s%v", cacheWechatUserIdPrefix, id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, wechatUserIdKey)
	return err
}

func (m *defaultWechatUserModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheWechatUserIdPrefix, primary)
}

func (m *defaultWechatUserModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", wechatUserRows, m.table)
	return conn.QueryRow(v, query, primary)
}
