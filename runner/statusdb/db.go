package statusdb

import (
	"github.com/dgraph-io/badger"
	util "github.com/hktalent/go-utils"
	"sync/atomic"
	"time"
)

type Item struct {
	Domain      string    `json:"domain"`       // 查询域名
	Dns         string    `json:"dns"`          // 查询dns
	Time        time.Time `json:"time"`         // 发送时间
	Retry       int       `json:"retry"`        // 重试次数
	DomainLevel int       `json:"domain_level"` // 域名层级
}

type StatusDb struct {
	length int64
}

// 内存简易读写数据库，自带锁机制
func CreateMemoryDB() *StatusDb {
	db := &StatusDb{
		length: 0,
	}
	return db
}

func (r *StatusDb) Add(domain string, tableData Item) {
	util.PutAny[Item](domain, tableData)
	atomic.AddInt64(&r.length, 1)
}
func (r *StatusDb) Set(domain string, tableData Item) {
	util.PutAny[Item](domain, tableData)
}
func (r *StatusDb) Get(domain string) (Item, bool) {
	v, err := util.GetAny[Item](domain)
	if nil != err {
		return Item{}, false
	}
	return v, true
}
func (r *StatusDb) Length() int64 {
	return r.length
}
func (r *StatusDb) Del(domain string) {
	//r.Mu.Lock()
	//defer r.Mu.Unlock()
	if nil == util.Cache1.Delete(domain) {
		atomic.AddInt64(&r.length, -1)
	}
}

func (r *StatusDb) Scan(f func(key string, value Item) error) {
	util.Cache1.DbConn.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchSize = 10
		it := txn.NewIterator(opts)
		defer it.Close()
		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			k := item.Key()
			err := item.Value(func(v []byte) error {
				//fmt.Printf("key=%s, value=%s\n", k, v)
				var itm = Item{}
				if nil == util.Json.Unmarshal(v, &itm) {
					f(string(k), itm)
				}
				return nil
			})
			if err != nil {
				return err
			}
		}
		return nil
	})
	//r.Items.Range(func(key, value interface{}) bool {
	//	k := key.(string)
	//	item := value.(Item)
	//	f(k, item)
	//	return true
	//})
}
func (r *StatusDb) Close() {
}
