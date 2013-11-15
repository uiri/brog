package brogger

import (
	"io/ioutil"
	"os"
	"testing"
)

var config *Config

func TestSelfValidate(t *testing.T) {
	config = newDefaultConfig()
	err := config.selfValidate()
	if err != nil {
		t.Error("Error validating default conig")
	}
	config.ProdPortNumber = -1
	err = config.selfValidate()
	if err == nil {
		t.Error("-1 is not a valid production port number")
	}
	config.ProdPortNumber = 65600
	err = config.selfValidate()
	if err == nil {
		t.Error("656000 is not a valid production port number")
	}
	config.ProdPortNumber = DefaultProdPortNumber
	config.DevelPortNumber = -1
	err = config.selfValidate()
	if err == nil {
		t.Error("-1 is not a valid development port number")
	}
	config.DevelPortNumber = 65600
	err = config.selfValidate()
	if err == nil {
		t.Error("65600 is not a valid development port number")
	}
	config.DevelPortNumber = DefaultDevelPortNumber
	config.MaxCPUs = -1
	err = config.selfValidate()
	if err == nil {
		t.Error("-1 is not a valid number of threads")
	}
	config.MaxCPUs = DefaultMaxCPUs
	config.PostFileExt = ""
	err = config.selfValidate()
	if err == nil {
		t.Error("Post file extension cannot be empty")
	}
}

func TestJsonConfigStruct(t *testing.T) {
	os.Chdir("base")
	defer os.Chdir("..")
	config, _ = loadConfig()
	config.persistToFile("test_config.json")
	defer os.Remove("test_config.json")
	origfile, _ := os.Open("brog_config.json")
	testfile, _ := os.Open("test_config.json")
	origbuf, origerr := ioutil.ReadAll(origfile)
	testbuf, testerr := ioutil.ReadAll(testfile)
	origfile.Close()
	testfile.Close()
	if origerr != nil || testerr != nil {
		t.Error("Error reading original config:", origerr, "or reading test config:", testerr)
	}
	if len(origbuf) != len(testbuf) {
		t.Error("One config file ended before the other")
	}
	for i := range origbuf {
		if origbuf[i] != testbuf[i] {
			t.Error("Config files are not byte-by-byte equal")
		}
	}
}
