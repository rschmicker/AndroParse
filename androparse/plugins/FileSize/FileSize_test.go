package main

import (
	"github.com/AndroParse/androparse/utils"
	. "gopkg.in/check.v1"
	"path/filepath"
	"testing"
)

type FileSizeTestSuite struct{}

var _ = Suite(&FileSizeTestSuite{})

func Test(t *testing.T) { TestingT(t) }

func (s *FileSizeTestSuite) TestGetKey(c *C) {
	c.Check(GetKey(), Equals, "FileSize")
}

func (s *FileSizeTestSuite) TestNeedLock(c *C) {
	c.Check(NeedLock(), Equals, false)
}

func (s *FileSizeTestSuite) TestGetValue(c *C) {
	configPath, err := filepath.Abs("../../../test.yaml")
	c.Assert(err, IsNil)
	config := utils.ReadConfig(configPath)
	testLoc := config.ApkDir + "/Facebook Lite_v70.0.0.9.116_apkpure.com.apk"
	m, err := GetValue(testLoc, config)
	c.Assert(err, IsNil)
	res, ok := m.(int64)
	c.Check(ok, Equals, true)
	c.Check(res, Equals, int64(1699930))
}
