package main

import (
	"github.com/AndroParse/androparse/utils"
	. "gopkg.in/check.v1"
	"path/filepath"
	"testing"
)

type StringsTestSuite struct{}

var _ = Suite(&StringsTestSuite{})

func Test(t *testing.T) { TestingT(t) }

func (s *StringsTestSuite) TestGetKey(c *C) {
	c.Check(GetKey(), Equals, "Strings")
}

func (s *StringsTestSuite) TestNeedLock(c *C) {
	c.Check(NeedLock(), Equals, true)
}

func (s *StringsTestSuite) TestGetValue(c *C) {
	configPath, err := filepath.Abs("../../../test.yaml")
	c.Assert(err, IsNil)
	config := utils.ReadConfig(configPath)
	testLoc := config.ApkDir + "/Facebook Lite_v70.0.0.9.116_apkpure.com.apk"
	a, err := GetValue(testLoc, config)
	c.Assert(err, IsNil)
	StringssInApk, ok := a.([]string)
	c.Check(ok, Equals, true)
	c.Check(len(StringssInApk), Equals, 14961)
}
