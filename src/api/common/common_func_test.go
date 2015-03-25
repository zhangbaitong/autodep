package common

import (
	"testing"
)

func TestConfig(t *testing.T) {

	// result is string "http://www.example.com/some/path"
	if ret, _ := Config().String("service-1", "url"); ret != "http://www.example.com/some/path" {
		t.Errorf("service-1 url is ", ret)
	}

	// result is int 200
	if ret1, _ := Config().Int("service-1", "maxclients"); ret1 != 200 {
		t.Errorf("service-1 maxclients is ", ret1)
	}

	// result is bool true
	if ret2, _ := Config().Bool("service-1", "delegation"); ret2 != true {
		t.Errorf("service-1 delegation is ", ret2)
	}

	// result is string "This is a multi-line\nentry"
	if ret3, _ := Config().String("service-1", "comments"); ret3 != "This is a multi-line\nentry" {
		t.Errorf("service-1 comments is ", ret3)
	}

	// result is string "10.122.75.228"
	if ret4, _ := Config().String("autodep", "ip"); ret4 != "10.122.75.228" {
		t.Errorf("autodep ip is ", ret4)
	}

}
