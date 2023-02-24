package master

import (
	"Distributed-Crontab/pkg"
	"context"

	_ "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// mongodb日志管理
type LogMgr struct {
	client        *mongo.Client
	logCollection *mongo.Collection
}

var (
	G_logMgr *LogMgr
)

func InitLogMgr() (err error) {
	var (
		client *mongo.Client
	)
	// 建立mongodb连接
	if client, err = mongo.Connect(
		context.TODO(),
		options.Client().ApplyURI(G_config.MongodbUri),
	//clientopt.ConnectTimeout(time.Duration(G_config.MongodbConnectTimeout)*time.Millisecond)
	); err != nil {
		return
	}
	G_logMgr = &LogMgr{
		client:        client,
		logCollection: client.Database("cron").Collection("log"),
	}
	return
}

func (logMgr *LogMgr) ListLog(name string, skip int, limit int) (logArr []*pkg.JobLog, err error) {
	var (
		filter  *pkg.JobLogFilter
		logSort *pkg.SortLogByStartTime
		cursor  *mongo.Cursor
		jobLog  *pkg.JobLog
	)
	// len(logArr)
	logArr = make([]*pkg.JobLog, 0)
	// 过滤条件
	filter = &pkg.JobLogFilter{JobName: name}
	// 按照任务开始时间倒排
	logSort = &pkg.SortLogByStartTime{SortOrder: -1}
	opt := options.Find()
	opt.SetSort(logSort)
	opt.SetSkip(int64(skip))
	opt.SetLimit(int64(limit))
	// 查询
	//TODO:根据最新官方文档进行查询
	if cursor, err = logMgr.logCollection.Find(context.TODO(), filter); err != nil {
		return
	}
	// 延迟释放游标
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		jobLog = &pkg.JobLog{}

		// 反序列化BSON
		if err = cursor.Decode(jobLog); err != nil {
			continue // 有日志不合法
		}

		logArr = append(logArr, jobLog)
	}
	return
}
