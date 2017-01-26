package jsonstore

import (
	"encoding/json"
	"fmt"
	"testing"
)

func BenchmarkGet(b *testing.B) {
	var fs JSONStore
	fs.Init()
	fs.Set("data", 1234)
	fs.Set("name", "bob")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fs.Get("data")
	}
}

func BenchmarkGetMany(b *testing.B) {
	var fs JSONStore
	fs.Init()
	fs.Set("name:1", "bob")
	fs.Set("name:2", "zack")
	fs.Set("name:3", "bill")
	fs.Set("country:1", "usa")
	fs.Set("country:2", "usa")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fs.Get("name*")
	}
}

func BenchmarkSet(b *testing.B) {
	var fs JSONStore
	fs.Init()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fs.Set("data", 1234)
	}
}

func BenchmarkSetMem(b *testing.B) {
	var fs JSONStore
	fs.Init()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fs.SetMem("data", 1234)
	}
}

func TestGetMany(t *testing.T) {
	var fs JSONStore
	fs.Init()
	fs.Set("name:1", "bob")
	fs.Set("name:2", "zack")
	fs.Set("name:3", "bill")
	fs.Set("country:1", "usa")
	fs.Set("country:2", "usa")
	val, err := fs.Get("name*")
	if err != nil {
		t.Errorf("Got %+v, and %s", val, err.Error())
	}
}

func TestSet(t *testing.T) {
	// Test simple saving and getting
	var fs JSONStore
	fs.Init()
	fs.Set("name", "zack")
	fmt.Println(string(fs.Data["name"]))

	var fs2 JSONStore
	fs2.Init()
	fs2.Load()
	s, err := fs2.Get("name")
	var val string
	json.Unmarshal([]byte(s), &val)
	fmt.Println(s)
	fmt.Println(string(fs2.Data["name"]))
	if err != nil || val != "zack" {
		t.Errorf("Got %+v, and %v", val, err)
	}
	//
	// // Test saving a different place
	// os.Remove("test.json")
	// fs.SetLocation("test.json")
	// fs.SetMem("name2", "zack2") // doesn't persist
	// fs.Save()                   // now its saved
	//
	// fs2.SetLocation("test.json")
	// fs2.Load()
	// byteJSON, err = fs2.Get("name2")
	// json.Unmarshal(byteJSON, &val)
	// if err != nil || val != "zack2" {
	// 	t.Errorf("Got %+v, and %s", val, err.Error())
	// }
}

//
// func TestSetNoCompress(t *testing.T) {
// 	// Test simple saving and getting
// 	var fs JSONStore
// 	fs.Init()
// 	fs.SetLocation("nocompress.json")
// 	fs.SetGzip(false)
// 	fs.Set("name", "zack")
//
// 	var fs2 JSONStore
// 	fs2.Init()
// 	fs2.SetLocation("nocompress.json")
// 	fs2.SetGzip(false)
// 	fs2.Load()
// 	var val string
// 	byteJSON, err := fs2.Get("name2")
// 	json.Unmarshal(byteJSON, &val)
// 	if err != nil || val != "zack" {
// 		t.Errorf("Got %+v, and %s", val, err.Error())
// 	}
//
// }
