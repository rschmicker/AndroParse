package main

import (
	"github.com/AndroParse/androparse/utils"
	. "gopkg.in/check.v1"
	"path/filepath"
	"testing"
)

type Sha256TestSuite struct{}

var _ = Suite(&Sha256TestSuite{})

func Test(t *testing.T) { TestingT(t) }

func (s *Sha256TestSuite) TestGetKey(c *C) {
	c.Check(GetKey(), Equals, "Sha256")
}

func (s *Sha256TestSuite) TestNeedLock(c *C) {
	c.Check(NeedLock(), Equals, false)
}

func (s *Sha256TestSuite) TestGetValue(c *C) {
	configPath, err := filepath.Abs("../../../test.yaml")
	c.Assert(err, IsNil)
	config := utils.ReadConfig(configPath)
	testLoc := config.ApkDir + "/Facebook Lite_v70.0.0.9.116_apkpure.com.apk"
	m, err := GetValue(testLoc, config)
	c.Assert(err, IsNil)
	Sha256, ok := m.(string)
	c.Check(ok, Equals, true)
	c.Check(Sha256, Equals, "8fc218d35790b7c363b7423f9bd6faa71b2adcc59e55444431eced0cf0e60a4d")
}
