package main

import (
	"github.com/AndroParse/androparse/utils"
	. "gopkg.in/check.v1"
	"path/filepath"
	"testing"
)

type DateTestSuite struct{}

var _ = Suite(&DateTestSuite{})

func Test(t *testing.T) { TestingT(t) }

func (s *DateTestSuite) TestGetKey(c *C) {
	c.Check(GetKey(), Equals, "Date")
}

func (s *DateTestSuite) TestNeedLock(c *C) {
	c.Check(NeedLock(), Equals, false)
}

func (s *DateTestSuite) TestGetValue(c *C) {
	configPath, err := filepath.Abs("../../../test.yaml")
	c.Assert(err, IsNil)
	config := utils.ReadConfig(configPath)
	testLoc := config.ApkDir + "/Facebook Lite_v70.0.0.9.116_apkpure.com.apk"
	m, err := GetValue(testLoc, config)
	c.Assert(err, IsNil)
	_, ok := m.(string)
	c.Check(ok, Equals, true)
}
