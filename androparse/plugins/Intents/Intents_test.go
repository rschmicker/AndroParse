package main

import (
	"github.com/AndroParse/androparse/utils"
	. "gopkg.in/check.v1"
	"path/filepath"
	"testing"
)

type IntentTestSuite struct{}

var _ = Suite(&IntentTestSuite{})

func Test(t *testing.T) { TestingT(t) }

func (s *IntentTestSuite) TestGetKey(c *C) {
	c.Check(GetKey(), Equals, "Intents")
}

func (s *IntentTestSuite) TestNeedLock(c *C) {
	c.Check(NeedLock(), Equals, false)
}

func (s *IntentTestSuite) TestGetValue(c *C) {
	configPath, err := filepath.Abs("../../../test.yaml")
	c.Assert(err, IsNil)
	config := utils.ReadConfig(configPath)
	testLoc := config.ApkDir + "/Facebook Lite_v70.0.0.9.116_apkpure.com.apk"
	i, err := GetValue(testLoc, config)
	c.Assert(err, IsNil)
	intents, ok := i.([]string)
	c.Check(ok, Equals, true)
	c.Check(len(intents), Equals, 60)
}
