package main

import (
	"github.com/AndroParse/androparse/utils"
	. "gopkg.in/check.v1"
	"path/filepath"
	"testing"
)

type IsMaliciousTestSuite struct{}

var _ = Suite(&IsMaliciousTestSuite{})

func Test(t *testing.T) { TestingT(t) }

func (s *IsMaliciousTestSuite) TestGetKey(c *C) {
	c.Check(GetKey(), Equals, "Malicious")
}

func (s *IsMaliciousTestSuite) TestNeedLock(c *C) {
	c.Check(NeedLock(), Equals, false)
}

func (s *IsMaliciousTestSuite) TestGetValue(c *C) {
	configPath, err := filepath.Abs("../../../test.yaml")
	c.Assert(err, IsNil)
	config := utils.ReadConfig(configPath)
	testLoc := config.ApkDir + "/8fc218d35790b7c363b7423f9bd6faa71b2adcc59e55444431eced0cf0e60a4d.apk"
	config.VtApiCheck = true
	m, err := GetValue(testLoc, config)
	c.Assert(err, IsNil)
	malicious, ok := m.(string)
	c.Check(ok, Equals, true)
	c.Check(malicious, Equals, "false")
}
