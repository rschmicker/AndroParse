package main

import (
	"github.com/AndroParse/androparse/utils"
	. "gopkg.in/check.v1"
	"path/filepath"
	"testing"
)

type PackageNameTestSuite struct{}

var _ = Suite(&PackageNameTestSuite{})

func Test(t *testing.T) { TestingT(t) }

func (s *PackageNameTestSuite) TestGetKey(c *C) {
	c.Check(GetKey(), Equals, "PackageName")
}

func (s *PackageNameTestSuite) TestNeedLock(c *C) {
	c.Check(NeedLock(), Equals, false)
}

func (s *PackageNameTestSuite) TestGetValue(c *C) {
	configPath, err := filepath.Abs("../../../test.yaml")
	c.Assert(err, IsNil)
	config := utils.ReadConfig(configPath)
	testLoc := config.ApkDir + "/Facebook Lite_v70.0.0.9.116_apkpure.com.apk"
	i, err := GetValue(testLoc, config)
	c.Assert(err, IsNil)
	PackageName, ok := i.(string)
	c.Check(ok, Equals, true)
	c.Check(PackageName, Equals, "com.facebook.lite")
}
