package main

import (
	"github.com/AndroParse/androparse/utils"
	. "gopkg.in/check.v1"
	"path/filepath"
	"testing"
)

type PermissionTestSuite struct{}

var _ = Suite(&PermissionTestSuite{})

func Test(t *testing.T) { TestingT(t) }

func (s *PermissionTestSuite) TestGetKey(c *C) {
	c.Check(GetKey(), Equals, "Permissions")
}

func (s *PermissionTestSuite) TestNeedLock(c *C) {
	c.Check(NeedLock(), Equals, false)
}

func (s *PermissionTestSuite) TestGetValue(c *C) {
	configPath, err := filepath.Abs("../../../test.yaml")
	c.Assert(err, IsNil)
	config := utils.ReadConfig(configPath)
	testLoc := config.ApkDir + "/Facebook Lite_v70.0.0.9.116_apkpure.com.apk"
	i, err := GetValue(testLoc, config)
	c.Assert(err, IsNil)
	Permissions, ok := i.([]string)
	c.Check(ok, Equals, true)
	c.Check(len(Permissions), Equals, 45)
}
