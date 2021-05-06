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
	userPwdFieldNames          = builderx.RawFieldNames(&UserPwd{})
	userPwdRows                = strings.Join(userPwdFieldNames, ",")
	userPwdRowsExpectAutoSet   = strings.Join(stringx.Remove(userPwdFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	userPwdRowsWithPlaceHolder = strings.Join(stringx.Remove(userPwdFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheUserPwdIdPrefix     = "cache#userPwd#id#"
	cacheUserPwdNumberPrefix = "cache#userPwd#number#"
)

type (
	UserPwdModel interface {
		Insert(data UserPwd) (sql.Result, error)
		FindOne(id int64) (*UserPwd, error)
		FindOneByNumber(number int64) (*UserPwd, error)
		Update(data UserPwd) error
		Delete(id int64) error
	}

	defaultUserPwdModel struct {
		sqlc.CachedConn
		table string
	}

	UserPwd struct {
		UpdateTime time.Time `db:"update_time"`
		Id         int64     `db:"id"`       // 主键自增id
		Number     int64     `db:"number"`   // 用户唯一number
		Password   string    `db:"password"` // 用户密码
		CreateTime time.Time `db:"create_time"`
	}
)

func NewUserPwdModel(conn sqlx.SqlConn, c cache.CacheConf) UserPwdModel {
	return &defaultUserPwdModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`user_pwd`",
	}
}

func (m *defaultUserPwdModel) Insert(data UserPwd) (sql.Result, error) {
	userPwdNumberKey := fmt.Sprintf("%s%v", cacheUserPwdNumberPrefix, data.Number)
	ret, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, userPwdRowsExpectAutoSet)
		return conn.Exec(query, data.Number, data.Password)
	}, userPwdNumberKey)
	return ret, err
}

func (m *defaultUserPwdModel) FindOne(id int64) (*UserPwd, error) {
	userPwdIdKey := fmt.Sprintf("%s%v", cacheUserPwdIdPrefix, id)
	var resp UserPwd
	err := m.QueryRow(&resp, userPwdIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userPwdRows, m.table)
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

func (m *defaultUserPwdModel) FindOneByNumber(number int64) (*UserPwd, error) {
	userPwdNumberKey := fmt.Sprintf("%s%v", cacheUserPwdNumberPrefix, number)
	var resp UserPwd
	err := m.QueryRowIndex(&resp, userPwdNumberKey, m.formatPrimary, func(conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `number` = ? limit 1", userPwdRows, m.table)
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

func (m *defaultUserPwdModel) Update(data UserPwd) error {
	userPwdIdKey := fmt.Sprintf("%s%v", cacheUserPwdIdPrefix, data.Id)
	userPwdNumberKey := fmt.Sprintf("%s%v", cacheUserPwdNumberPrefix, data.Number)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, userPwdRowsWithPlaceHolder)
		return conn.Exec(query, data.Number, data.Password, data.Id)
	}, userPwdIdKey, userPwdNumberKey)
	return err
}

func (m *defaultUserPwdModel) Delete(id int64) error {
	data, err := m.FindOne(id)
	if err != nil {
		return err
	}

	userPwdIdKey := fmt.Sprintf("%s%v", cacheUserPwdIdPrefix, id)
	userPwdNumberKey := fmt.Sprintf("%s%v", cacheUserPwdNumberPrefix, data.Number)
	_, err = m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, userPwdIdKey, userPwdNumberKey)
	return err
}

func (m *defaultUserPwdModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheUserPwdIdPrefix, primary)
}

func (m *defaultUserPwdModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userPwdRows, m.table)
	return conn.QueryRow(v, query, primary)
}
