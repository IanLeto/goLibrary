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

}

func TestTimeDemoSuite(t *testing.T) {
	suite.Run(t, new(TesttimeUtils_testSuite))
}
