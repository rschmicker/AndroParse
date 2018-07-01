package main

import (
	"github.com/AndroParse/androparse/utils"
	. "gopkg.in/check.v1"
	"path/filepath"
	"testing"
)

type Md5TestSuite struct{}

var _ = Suite(&Md5TestSuite{})

func Test(t *testing.T) { TestingT(t) }

func (s *Md5TestSuite) TestGetKey(c *C) {
	c.Check(GetKey(), Equals, "Md5")
}

func (s *Md5TestSuite) TestNeedLock(c *C) {
	c.Check(NeedLock(), Equals, false)
}

func (s *Md5TestSuite) TestGetValue(c *C) {
	configPath, err := filepath.Abs("../../../test.yaml")
	c.Assert(err, IsNil)
	config := utils.ReadConfig(configPath)
	testLoc := config.ApkDir + "/Facebook Lite_v70.0.0.9.116_apkpure.com.apk"
	m, err := GetValue(testLoc, config)
	c.Assert(err, IsNil)
	Md5, ok := m.(string)
	c.Check(ok, Equals, true)
	c.Check(Md5, Equals, "a1c88d70e6ffe6ed5167f75c8399af4e")
}
