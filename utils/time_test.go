package utils_test

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type TesttimeUtils_testSuite struct {
	suite.Suite
	timeDemo time.Time
}

func (s *TesttimeUtils_testSuite) SetupTest() {
	s.timeDemo = time.Now()
	// int64 to timestamp
	logrus.Info("golang 没有所谓的time.timestamp类型，int 就是时间戳，这里与python不同")
	var sec int64 = 1582701190
	var msec int64 = 1582701190946
	var nsec int64 = 1582701190946414000
	// timestamp to time
	tsec := time.Unix(sec, 0)
	tmsec := time.Unix(0, msec*int64(time.Millisecond))
	tnsec := time.Unix(0, nsec)

	// time to timestamp
	s.Equal(sec, tsec.Unix())
	s.Equal(msec, tmsec.UnixNano()/1e6)
	s.Equal(nsec, tnsec.UnixNano())

	// 标准时间转时间戳
	t, _ := time.Parse(time.RFC3339, "2020-02-26T15:13:10+08:00")
	logrus.Info(1, t.Unix())

	// sec timestamp to msec timestamp
	logrus.Info(sec * int64(time.Millisecond))
	logrus.Info(sec * int64(time.Nanosecond))
	xx := time.Duration(sec)
	fmt.Println("cx", xx)

	// time to 标准时间格式
	logrus.Info(tsec.Format(time.RFC3339))
	logrus.Info(tnsec.Format(time.RFC3339Nano))

	// time.duration to second
	// time.duration to ms
	time.Sleep(1 * time.Second)
	logrus.Infof("time.duration to second: %d", int64(time.Since(s.timeDemo)/time.Second))
	logrus.Infof("time.duration to ms: %d", int64(time.Since(s.timeDemo)/time.Millisecond))

	// 时间计算
	nowCalc := time.Now().Unix()
	nowCalc = nowCalc - 3600
	// 1小时之前
	logrus.Infof("1 小时之前 %d", nowCalc)

	// 解析字符串时间
	hours, _ := time.ParseDuration("10h")
	logrus.Infof("10h => %v", hours)
	logrus.Info("-----------")
	a := time.Unix(1583311518000000000, 0)
	b := time.Unix(0, 1583311518000000000)

	logrus.Infof("%v", a.UnixNano())
	logrus.Infof("%v", a.Unix())

	logrus.Infof("%v", b.UnixNano())
	logrus.Infof("%v", b.UnixNano())

	c := time.Unix(0, 1582701190)
	logrus.Info(c.Unix())
	logrus.Info(c.UnixNano())

	logrus.Info(time.Now().UnixNano())
	logrus.Info(time.Unix(0, 1583982851473426).Format(time.RFC3339Nano))
	logrus.Info(time.Unix(0, 1583983973246439000).Format(time.RFC3339Nano))
	logrus.Info(time.Unix(0, 1583983973246439).Format(time.RFC3339))
}
func (s *TesttimeUtils_testSuite) TestGetLog() {
	logrus.Info("===== Go时间格式转换完整示例 =====")

	// 定义所有时间格式的变量
	now := time.Now()

	// 1. 各种时间格式示例
	logrus.Info("=== 1. 所有时间格式示例 ===")

	// 北京时间 (CST, UTC+8)
	locationBeijing, _ := time.LoadLocation("Asia/Shanghai")
	beijingTime := now.In(locationBeijing)
	logrus.Infof("北京时间: %v", beijingTime)

	// UTC时间
	utcTime := now.UTC()
	logrus.Infof("UTC时间: %v", utcTime)

	// 时间戳（秒）
	timestampSec := now.Unix()
	logrus.Infof("时间戳(秒): %d", timestampSec)

	// 时间戳（毫秒）
	timestampMilli := now.UnixNano() / int64(time.Millisecond)
	logrus.Infof("时间戳(毫秒): %d", timestampMilli)

	// 时间戳（微秒）
	timestampMicro := now.UnixNano() / int64(time.Microsecond)
	logrus.Infof("时间戳(微秒): %d", timestampMicro)

	// 时间戳（纳秒）
	timestampNano := now.UnixNano()
	logrus.Infof("时间戳(纳秒): %d", timestampNano)

	// MySQL默认时间格式
	mysqlFormat := now.Format("2006-01-02 15:04:05")
	logrus.Infof("MySQL默认格式: %s", mysqlFormat)

	// Elasticsearch默认时间格式 (ISO 8601)
	esFormat := now.Format(time.RFC3339)
	logrus.Infof("ES默认格式(ISO 8601): %s", esFormat)

	// 2. 所有格式之间的转换函数
	logrus.Info("\n=== 2. 所有格式互相转换示例 ===")

	// 2.1 北京时间转换到其他格式
	logrus.Info("\n--- 北京时间转换到其他格式 ---")
	beijingToUTC := beijingTime.UTC()
	logrus.Infof("北京时间 -> UTC时间: %v", beijingToUTC)

	beijingToSec := beijingTime.Unix()
	logrus.Infof("北京时间 -> 时间戳(秒): %d", beijingToSec)

	beijingToMilli := beijingTime.UnixNano() / int64(time.Millisecond)
	logrus.Infof("北京时间 -> 时间戳(毫秒): %d", beijingToMilli)

	beijingToMicro := beijingTime.UnixNano() / int64(time.Microsecond)
	logrus.Infof("北京时间 -> 时间戳(微秒): %d", beijingToMicro)

	beijingToMySQL := beijingTime.Format("2006-01-02 15:04:05")
	logrus.Infof("北京时间 -> MySQL格式: %s", beijingToMySQL)

	beijingToES := beijingTime.Format(time.RFC3339)
	logrus.Infof("北京时间 -> ES格式: %s", beijingToES)

	// 2.2 UTC时间转换到其他格式
	logrus.Info("\n--- UTC时间转换到其他格式 ---")
	utcToBeijing := utcTime.In(locationBeijing)
	logrus.Infof("UTC时间 -> 北京时间: %v", utcToBeijing)

	utcToSec := utcTime.Unix()
	logrus.Infof("UTC时间 -> 时间戳(秒): %d", utcToSec)

	utcToMilli := utcTime.UnixNano() / int64(time.Millisecond)
	logrus.Infof("UTC时间 -> 时间戳(毫秒): %d", utcToMilli)

	utcToMicro := utcTime.UnixNano() / int64(time.Microsecond)
	logrus.Infof("UTC时间 -> 时间戳(微秒): %d", utcToMicro)

	utcToMySQL := utcTime.Format("2006-01-02 15:04:05")
	logrus.Infof("UTC时间 -> MySQL格式: %s", utcToMySQL)

	utcToES := utcTime.Format(time.RFC3339)
	logrus.Infof("UTC时间 -> ES格式: %s", utcToES)

	// 2.3 时间戳(秒)转换到其他格式
	logrus.Info("\n--- 时间戳(秒)转换到其他格式 ---")
	secToTime := time.Unix(timestampSec, 0)
	secToBeijing := secToTime.In(locationBeijing)
	logrus.Infof("时间戳(秒) -> 北京时间: %v", secToBeijing)

	secToUTC := secToTime.UTC()
	logrus.Infof("时间戳(秒) -> UTC时间: %v", secToUTC)

	secToMilli := timestampSec * 1000
	logrus.Infof("时间戳(秒) -> 时间戳(毫秒): %d", secToMilli)

	secToMicro := timestampSec * 1000000
	logrus.Infof("时间戳(秒) -> 时间戳(微秒): %d", secToMicro)

	secToMySQL := secToTime.Format("2006-01-02 15:04:05")
	logrus.Infof("时间戳(秒) -> MySQL格式: %s", secToMySQL)

	secToES := secToTime.Format(time.RFC3339)
	logrus.Infof("时间戳(秒) -> ES格式: %s", secToES)

	// 2.4 时间戳(毫秒)转换到其他格式
	logrus.Info("\n--- 时间戳(毫秒)转换到其他格式 ---")
	milliToTime := time.Unix(0, timestampMilli*int64(time.Millisecond))
	milliToBeijing := milliToTime.In(locationBeijing)
	logrus.Infof("时间戳(毫秒) -> 北京时间: %v", milliToBeijing)

	milliToUTC := milliToTime.UTC()
	logrus.Infof("时间戳(毫秒) -> UTC时间: %v", milliToUTC)

	milliToSec := timestampMilli / 1000
	logrus.Infof("时间戳(毫秒) -> 时间戳(秒): %d", milliToSec)

	milliToMicro := timestampMilli * 1000
	logrus.Infof("时间戳(毫秒) -> 时间戳(微秒): %d", milliToMicro)

	milliToMySQL := milliToTime.Format("2006-01-02 15:04:05")
	logrus.Infof("时间戳(毫秒) -> MySQL格式: %s", milliToMySQL)

	milliToES := milliToTime.Format(time.RFC3339)
	logrus.Infof("时间戳(毫秒) -> ES格式: %s", milliToES)

	// 2.5 时间戳(微秒)转换到其他格式
	logrus.Info("\n--- 时间戳(微秒)转换到其他格式 ---")
	microToTime := time.Unix(0, timestampMicro*int64(time.Microsecond))
	microToBeijing := microToTime.In(locationBeijing)
	logrus.Infof("时间戳(微秒) -> 北京时间: %v", microToBeijing)

	microToUTC := microToTime.UTC()
	logrus.Infof("时间戳(微秒) -> UTC时间: %v", microToUTC)

	microToSec := timestampMicro / 1000000
	logrus.Infof("时间戳(微秒) -> 时间戳(秒): %d", microToSec)

	microToMilli := timestampMicro / 1000
	logrus.Infof("时间戳(微秒) -> 时间戳(毫秒): %d", microToMilli)

	microToMySQL := microToTime.Format("2006-01-02 15:04:05")
	logrus.Infof("时间戳(微秒) -> MySQL格式: %s", microToMySQL)

	microToES := microToTime.Format(time.RFC3339)
	logrus.Infof("时间戳(微秒) -> ES格式: %s", microToES)

	// 2.6 MySQL格式转换到其他格式
	logrus.Info("\n--- MySQL格式转换到其他格式 ---")
	mysqlTime, _ := time.Parse("2006-01-02 15:04:05", mysqlFormat)
	mysqlToBeijing := mysqlTime.In(locationBeijing)
	logrus.Infof("MySQL格式 -> 北京时间: %v", mysqlToBeijing)

	mysqlToUTC := mysqlTime.UTC()
	logrus.Infof("MySQL格式 -> UTC时间: %v", mysqlToUTC)

	mysqlToSec := mysqlTime.Unix()
	logrus.Infof("MySQL格式 -> 时间戳(秒): %d", mysqlToSec)

	mysqlToMilli := mysqlTime.UnixNano() / int64(time.Millisecond)
	logrus.Infof("MySQL格式 -> 时间戳(毫秒): %d", mysqlToMilli)

	mysqlToMicro := mysqlTime.UnixNano() / int64(time.Microsecond)
	logrus.Infof("MySQL格式 -> 时间戳(微秒): %d", mysqlToMicro)

	mysqlToES := mysqlTime.Format(time.RFC3339)
	logrus.Infof("MySQL格式 -> ES格式: %s", mysqlToES)

	// 2.7 ES格式转换到其他格式
	logrus.Info("\n--- ES格式转换到其他格式 ---")
	esTime, _ := time.Parse(time.RFC3339, esFormat)
	esTimeToBeijing := esTime.In(locationBeijing)
	logrus.Infof("ES格式 -> 北京时间: %v", esTimeToBeijing)

	esTimeToUTC := esTime.UTC()
	logrus.Infof("ES格式 -> UTC时间: %v", esTimeToUTC)

	esTimeToSec := esTime.Unix()
	logrus.Infof("ES格式 -> 时间戳(秒): %d", esTimeToSec)

	esTimeToMilli := esTime.UnixNano() / int64(time.Millisecond)
	logrus.Infof("ES格式 -> 时间戳(毫秒): %d", esTimeToMilli)

	esTimeToMicro := esTime.UnixNano() / int64(time.Microsecond)
	logrus.Infof("ES格式 -> 时间戳(微秒): %d", esTimeToMicro)

	esTimeToMySQL := esTime.Format("2006-01-02 15:04:05")
	logrus.Infof("ES格式 -> MySQL格式: %s", esTimeToMySQL)

	// 3. 实用转换函数封装
	logrus.Info("\n=== 3. 实用转换函数示例 ===")
	demoConversionFunctions()
}

// 实用转换函数封装
func demoConversionFunctions() {
	// 示例时间戳
	var sec int64 = 1582701190
	var msec int64 = 1582701190946
	var usec int64 = 1582701190946414

	logrus.Info("--- 实用转换函数 ---")

	// 秒级时间戳转各种格式
	t1 := SecToTime(sec)
	logrus.Infof("秒 -> Time: %v", t1)
	logrus.Infof("秒 -> 北京时间: %v", SecToBeijingTime(sec))
	logrus.Infof("秒 -> UTC时间: %v", SecToUTCTime(sec))
	logrus.Infof("秒 -> MySQL格式: %s", SecToMySQL(sec))
	logrus.Infof("秒 -> ES格式: %s", SecToES(sec))

	// 毫秒级时间戳转换
	t2 := MilliToTime(msec)
	logrus.Infof("毫秒 -> Time: %v", t2)
	logrus.Infof("毫秒 -> 秒: %d", MilliToSec(msec))

	// 微秒级时间戳转换
	t3 := MicroToTime(usec)
	logrus.Infof("微秒 -> Time: %v", t3)
	logrus.Infof("微秒 -> 毫秒: %d", MicroToMilli(usec))

	// Time转各种时间戳
	now := time.Now()
	logrus.Infof("Time -> 秒: %d", TimeToSec(now))
	logrus.Infof("Time -> 毫秒: %d", TimeToMilli(now))
	logrus.Infof("Time -> 微秒: %d", TimeToMicro(now))
}

// === 实用转换函数 ===

// 秒级时间戳转Time
func SecToTime(sec int64) time.Time {
	return time.Unix(sec, 0)
}

// 秒级时间戳转北京时间
func SecToBeijingTime(sec int64) time.Time {
	locationBeijing, _ := time.LoadLocation("Asia/Shanghai")
	return time.Unix(sec, 0).In(locationBeijing)
}

// 秒级时间戳转UTC时间
func SecToUTCTime(sec int64) time.Time {
	return time.Unix(sec, 0).UTC()
}

// 秒级时间戳转MySQL格式
func SecToMySQL(sec int64) string {
	return time.Unix(sec, 0).Format("2006-01-02 15:04:05")
}

// 秒级时间戳转ES格式
func SecToES(sec int64) string {
	return time.Unix(sec, 0).Format(time.RFC3339)
}

// 毫秒级时间戳转Time
func MilliToTime(msec int64) time.Time {
	return time.Unix(0, msec*int64(time.Millisecond))
}

// 毫秒转秒
func MilliToSec(msec int64) int64 {
	return msec / 1000
}

// 微秒级时间戳转Time
func MicroToTime(usec int64) time.Time {
	return time.Unix(0, usec*int64(time.Microsecond))
}

// 微秒转毫秒
func MicroToMilli(usec int64) int64 {
	return usec / 1000
}

// Time转秒级时间戳
func TimeToSec(t time.Time) int64 {
	return t.Unix()
}

// Time转毫秒级时间戳
func TimeToMilli(t time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond)
}

// Time转微秒级时间戳
func TimeToMicro(t time.Time) int64 {
	return t.UnixNano() / int64(time.Microsecond)
}

func TestTimeDemoSuite(t *testing.T) {
	suite.Run(t, new(TesttimeUtils_testSuite))
}
