package main

import (
	"github.com/AndroParse/androparse/utils"
	. "gopkg.in/check.v1"
	"path/filepath"
	"testing"
)

type Sha1TestSuite struct{}

var _ = Suite(&Sha1TestSuite{})

func Test(t *testing.T) { TestingT(t) }

func (s *Sha1TestSuite) TestGetKey(c *C) {
	c.Check(GetKey(), Equals, "Sha1")
}

func (s *Sha1TestSuite) TestNeedLock(c *C) {
	c.Check(NeedLock(), Equals, false)
}

func (s *Sha1TestSuite) TestGetValue(c *C) {
	configPath, err := filepath.Abs("../../../test.yaml")
	c.Assert(err, IsNil)
	config := utils.ReadConfig(configPath)
	testLoc := config.ApkDir + "/Facebook Lite_v70.0.0.9.116_apkpure.com.apk"
	m, err := GetValue(testLoc, config)
	c.Assert(err, IsNil)
	Sha1, ok := m.(string)
	c.Check(ok, Equals, true)
	c.Check(Sha1, Equals, "3ce20472e647d0194fd11518c302236934d5f605")
}
