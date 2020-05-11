package session

import (
	"github.com/garyburd/redigo/redis"
	"server/dao"
	"time"
)

//设置redis中session Key--uuid Value--username
func SetSession(sessionId, id string) (err error) {
	conn := dao.Pool.Get()
	//ping
	err = dao.Pool.TestOnBorrow(conn, time.Now())
	if err != nil {
		return err
	}
	//设置redis session
	_, err = conn.Do("Set", "session:"+sessionId, id)
	if err != nil {
		return err
	}
	//24小时过期
	_, err = conn.Do("expire", "session:"+sessionId, 24*3600)
	if err != nil {
		return err
	}
	return
}

//根据sessionId在redis中取值
func GetSession(sessionId string) (id string, err error) {
	conn := dao.Pool.Get()
	//ping
	err = dao.Pool.TestOnBorrow(conn, time.Now())
	if err != nil {
		return
	}
	id, err = redis.String(conn.Do("Get", "session:"+sessionId))
	if err != nil {
		return
	}
	return
}
