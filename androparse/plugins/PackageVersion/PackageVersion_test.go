package main

import (
	"github.com/AndroParse/androparse/utils"
	. "gopkg.in/check.v1"
	"path/filepath"
	"testing"
)

type PackageVersionTestSuite struct{}

var _ = Suite(&PackageVersionTestSuite{})

func Test(t *testing.T) { TestingT(t) }

func (s *PackageVersionTestSuite) TestGetKey(c *C) {
	c.Check(GetKey(), Equals, "Version")
}

func (s *PackageVersionTestSuite) TestNeedLock(c *C) {
	c.Check(NeedLock(), Equals, false)
}

func (s *PackageVersionTestSuite) TestGetValue(c *C) {
	configPath, err := filepath.Abs("../../../test.yaml")
	c.Assert(err, IsNil)
	config := utils.ReadConfig(configPath)
	testLoc := config.ApkDir + "/Facebook Lite_v70.0.0.9.116_apkpure.com.apk"
	i, err := GetValue(testLoc, config)
	c.Assert(err, IsNil)
	PackageVersion, ok := i.(string)
	c.Check(ok, Equals, true)
	c.Check(PackageVersion, Equals, "70.0.0.9.116")
}
