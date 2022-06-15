package main

import (
	"github.com/golang/protobuf/proto"
	protodata "golang_test/module/gRPC/protobuf/data"
	"log"
)

func main() {
	test := &protodata.Student{
		Name:   "geektutu",
		Male:   true,
		Scores: []int32{98, 85, 88},
	}
	data, err := proto.Marshal(test)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	log.Printf("data:%+v\n", data)

	newTest := &protodata.Student{}
	err = proto.Unmarshal(data, newTest)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}
	log.Printf("newTest:%+v\n", newTest)

	// Now test and newTest contain the same data.
	if test.GetName() != newTest.GetName() {
		log.Fatalf("data mismatch %q != %q", test.GetName(), newTest.GetName())
	}
}
