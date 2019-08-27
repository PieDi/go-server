package ConnPool

import (
	"github.com/kataras/iris/core/errors"
	"sync"
	"time"
)

// 单个连接资源
type CoonRes interface {
	Close() error
}

// 连接对象
type ConnIns struct {
	Conn CoonRes
	Time time.Time
}

//连接池对象
type ConnPool struct {
	// 互斥锁
	Mu sync.Mutex
	// 通道保存连接资源
	CSource chan *ConnIns
	// 创建连接资源
	FSource func() (CoonRes, error)
	// 连接池是否是关闭
	Closed bool
	// 连接超时时间
	ConnTimeOut time.Duration
}

func NewConnPool(factory func() (CoonRes, error), cap int) (*ConnPool, error) {
	if cap <= 0{
		 return nil, errors.New("cap不能小于0")
	}
	cp := &ConnPool{
		Mu: sync.Mutex{},
		CSource: make(chan *ConnIns, cap),
		FSource: factory,
		Closed: false,
		ConnTimeOut: 30,
	}
	for i := 0; i < cap; i++ {
		connRes, err := cp.FSource()
		if err != nil {
			return nil, errors.New("factory出错")
		}
		cp.CSource <- &ConnIns{Conn:connRes, Time:time.Now()}
	}
	return cp, nil
}

// 连接池的长度
func (cp *ConnPool) Len() int {
	return len(cp.CSource)
}

// 关闭连接池
func (cp *ConnPool) PoolClose()  {
	if cp.Closed {
		return
	}
	cp.Mu.Lock()
	close(cp.CSource)
	cp.Closed = true
	for conn := range cp.CSource {
		conn.Conn.Close()
	}
	cp.Mu.Unlock()
}

// 获取连接资源
func (cp *ConnPool)Get() (CoonRes, error) {
	if cp.Closed {
		return nil, errors.New("连接池已关闭")
	}
	for {
		select {
		case connIns, ok := <-cp.CSource:
			{
				if !ok {
					return nil, errors.New("连接池已关闭")
				}
				if time.Now().Sub(connIns.Time) > cp.ConnTimeOut {
					connIns.Conn.Close()
					continue
				}
				return connIns.Conn, nil
			}
		default:
			{
				connRes, err := cp.FSource()
				if err != nil {
					return nil, err
				}
				return connRes, nil
			}
		}
	}
}

// 将链接资源放入连接池
func (cp *ConnPool) Put(conn CoonRes) error  {
	if cp.Closed {
		return errors.New("连接池已关闭")
	}
	select {
		case cp.CSource <- &ConnIns{Conn: conn, Time: time.Now()}:
			{
				return nil
			}
		default:
			{
				//如果无法加入，则关闭连接
				conn.Close()
				return errors.New("连接池已满")
			}
	}
}