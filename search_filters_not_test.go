// Copyright 2012-2015 Oliver Eilhard. All rights reserved.
// Use of this source code is governed by a MIT-license.
// See http://olivere.mit-license.org/license.txt for details.

package elastic

import (
	"encoding/json"
	"testing"
)

func TestNotFilter(t *testing.T) {
	f := NewNotFilter(NewTermFilter("user", "olivere"))
	src, err := f.Source()
	if err != nil {
		t.Fatal(err)
	}
	data, err := json.Marshal(src)
	if err != nil {
		t.Fatalf("marshaling to JSON failed: %v", err)
	}
	got := string(data)
	expected := `{"not":{"filter":{"term":{"user":"olivere"}}}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}

func TestNotFilterWithParams(t *testing.T) {
	postDateFilter := NewRangeFilter("postDate").From("2010-03-01").To("2010-04-01")
	f := NewNotFilter(postDateFilter)
	f = f.FilterName("MyFilterName")
	src, err := f.Source()
	if err != nil {
		t.Fatal(err)
	}
	data, err := json.Marshal(src)
	if err != nil {
		t.Fatalf("marshaling to JSON failed: %v", err)
	}
	got := string(data)
	expected := `{"not":{"_name":"MyFilterName","filter":{"range":{"postDate":{"from":"2010-03-01","include_lower":true,"include_upper":true,"to":"2010-04-01"}}}}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}
