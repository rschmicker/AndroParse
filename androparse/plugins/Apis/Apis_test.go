package main

import (
	"github.com/AndroParse/androparse/utils"
	. "gopkg.in/check.v1"
	"path/filepath"
	"testing"
)

type ApiTestSuite struct{}

var _ = Suite(&ApiTestSuite{})

func Test(t *testing.T) { TestingT(t) }

func (s *ApiTestSuite) TestGetKey(c *C) {
	c.Check(GetKey(), Equals, "Apis")
}

func (s *ApiTestSuite) TestNeedLock(c *C) {
	c.Check(NeedLock(), Equals, true)
}

func (s *ApiTestSuite) TestGetValue(c *C) {
	configPath, err := filepath.Abs("../../../test.yaml")
	c.Assert(err, IsNil)
	config := utils.ReadConfig(configPath)
	testLoc := config.ApkDir + "/Facebook Lite_v70.0.0.9.116_apkpure.com.apk"
	a, err := GetValue(testLoc, config)
	c.Assert(err, IsNil)
	apisInApk, ok := a.([]string)
	c.Check(ok, Equals, true)
	c.Check(len(apisInApk), Equals, 5003)
}
