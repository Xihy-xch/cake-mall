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

	cacheWechatUserIdPrefix     = "cache#wechatUser#id#"
	cacheWechatUserNumberPrefix = "cache#wechatUser#number#"
)

type (
	WechatUserModel interface {
		Insert(data WechatUser) (sql.Result, error)
		FindOne(id int64) (*WechatUser, error)
		FindOneByNumber(number int64) (*WechatUser, error)
		FindOneByUnionID(unionID string) (*WechatUser, error)
		Update(data WechatUser) error
		Delete(id int64) error
	}

	defaultWechatUserModel struct {
		sqlc.CachedConn
		table string
	}

	WechatUser struct {
		UpdateTime time.Time `db:"update_time"`
		Id         int64     `db:"id"`          // 主键自增id
		Number     int64     `db:"number"`      // 用户唯一number
		UnionId    string    `db:"union_id"`    // 用户微信unionId
		SessionKey string    `db:"session_key"` // 用户微信sessionKey
		CreateTime time.Time `db:"create_time"`
	}
)

func NewWechatUserModel(conn sqlx.SqlConn, c cache.CacheConf) WechatUserModel {
	return &defaultWechatUserModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`wechat_user`",
	}
}

func (m *defaultWechatUserModel) Insert(data WechatUser) (sql.Result, error) {
	wechatUserNumberKey := fmt.Sprintf("%s%v", cacheWechatUserNumberPrefix, data.Number)
	ret, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, wechatUserRowsExpectAutoSet)
		return conn.Exec(query, data.Number, data.UnionId, data.SessionKey)
	}, wechatUserNumberKey)
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

func (m *defaultWechatUserModel) FindOneByNumber(number int64) (*WechatUser, error) {
	wechatUserNumberKey := fmt.Sprintf("%s%v", cacheWechatUserNumberPrefix, number)
	var resp WechatUser
	err := m.QueryRowIndex(&resp, wechatUserNumberKey, m.formatPrimary, func(conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `number` = ? limit 1", wechatUserRows, m.table)
		if err := conn.QueryRow(&resp, query, number); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultWechatUserModel) FindOneByUnionID(unionID string) (*WechatUser, error) {
	var resp WechatUser
	query := fmt.Sprintf("select %s from %s where `union_id` = ? limit 1", wechatUserRows, m.table)
	err := m.QueryRowNoCache(&resp, query, unionID)
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
	wechatUserNumberKey := fmt.Sprintf("%s%v", cacheWechatUserNumberPrefix, data.Number)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, wechatUserRowsWithPlaceHolder)
		return conn.Exec(query, data.Number, data.UnionId, data.SessionKey, data.Id)
	}, wechatUserIdKey, wechatUserNumberKey)
	return err
}

func (m *defaultWechatUserModel) Delete(id int64) error {
	data, err := m.FindOne(id)
	if err != nil {
		return err
	}

	wechatUserIdKey := fmt.Sprintf("%s%v", cacheWechatUserIdPrefix, id)
	wechatUserNumberKey := fmt.Sprintf("%s%v", cacheWechatUserNumberPrefix, data.Number)
	_, err = m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, wechatUserIdKey, wechatUserNumberKey)
	return err
}

func (m *defaultWechatUserModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheWechatUserIdPrefix, primary)
}

func (m *defaultWechatUserModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", wechatUserRows, m.table)
	return conn.QueryRow(v, query, primary)
}
