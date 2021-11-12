package main

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/util"
)

func main() {
	e, _ := casbin.NewEnforcer("model.conf", "policy.csv")

	//default:rbac_with_resource_roles
	fmt.Println("##########default:rbac_with_resource_roles################")
	testEnforce(e, true, "alice", "data1", "read")
	testEnforce(e, true, "alice", "data1", "write")
	testEnforce(e, false, "alice", "data2", "read")
	testEnforce(e, true, "alice", "data2", "write")
	testEnforce(e, false, "bob", "data1", "read")
	testEnforce(e, false, "bob", "data1", "write")
	testEnforce(e, false, "bob", "data2", "read")
	testEnforce(e, true, "bob", "data2", "write")
	fmt.Println("##########default:rbac_with_resource_roles################")

	//2:abac
	fmt.Println("\n\n##########2:abac################")
	eCtx2 := casbin.EnforceContext{"r2", "p2", "e2", "m2"}
	obj := struct{ Owner string }{Owner: "alice"}
	testEnforceEx(e, []string{}, eCtx2, "alice", obj, "write")
	fmt.Println("##########2:abac################")

	//3:restful
	fmt.Println("\n\n##########3:restful################")
	eCtx3 := casbin.EnforceContext{"r3", "p3", "e3", "m3"}
	testEnforce(e, true, eCtx3, "alice", "/alice_data/resource1", "GET")
	testEnforce(e, true, eCtx3, "alice", "/alice_data/resource1", "POST")
	testEnforce(e, true, eCtx3, "alice", "/alice_data/resource2", "GET")
	testEnforce(e, false, eCtx3, "alice", "/alice_data/resource2", "POST")
	testEnforce(e, false, eCtx3, "alice", "/bob_data/resource1", "GET")
	testEnforce(e, false, eCtx3, "alice", "/bob_data/resource1", "POST")
	testEnforce(e, false, eCtx3, "alice", "/bob_data/resource2", "GET")
	testEnforce(e, false, eCtx3, "alice", "/bob_data/resource2", "POST")

	testEnforce(e, false, eCtx3, "bob", "/alice_data/resource1", "GET")
	testEnforce(e, false, eCtx3, "bob", "/alice_data/resource1", "POST")
	testEnforce(e, true, eCtx3, "bob", "/alice_data/resource2", "GET")
	testEnforce(e, false, eCtx3, "bob", "/alice_data/resource2", "POST")
	testEnforce(e, false, eCtx3, "bob", "/bob_data/resource1", "GET")
	testEnforce(e, true, eCtx3, "bob", "/bob_data/resource1", "POST")
	testEnforce(e, false, eCtx3, "bob", "/bob_data/resource2", "GET")
	testEnforce(e, true, eCtx3, "bob", "/bob_data/resource2", "POST")
	fmt.Println("##########3:restful################")

	//4:abac_model
	fmt.Println("\n\n##########4:abac_model################")
	m := e.GetModel()
	for sec, ast := range m {
		fmt.Println(sec)
		for ptype, p := range ast {
			fmt.Println(ptype, p)
		}
	}
	eCtx4 := casbin.EnforceContext{"r4", "p4", "e4", "m4"}
	sub1 := newTestSubject("alice", 16)
	sub2 := newTestSubject("alice", 20)
	sub3 := newTestSubject("alice", 65)
	sub4 := newTestSubject("admin", 70)

	testEnforce(e, false, eCtx4, sub1, "/data1", "read")
	testEnforce(e, false, eCtx4, sub1, "/data2", "read")
	testEnforce(e, false, eCtx4, sub1, "/data1", "write")
	testEnforce(e, true, eCtx4, sub1, "/data2", "write")
	testEnforce(e, true, eCtx4, sub2, "/data1", "read")
	testEnforce(e, false, eCtx4, sub2, "/data2", "read")
	testEnforce(e, false, eCtx4, sub2, "/data1", "write")
	testEnforce(e, true, eCtx4, sub2, "/data2", "write")
	testEnforce(e, true, eCtx4, sub3, "/data1", "read")
	testEnforce(e, false, eCtx4, sub3, "/data2", "read")
	testEnforce(e, false, eCtx4, sub3, "/data1", "write")
	testEnforce(e, false, eCtx4, sub3, "/data2", "write")
	testEnforce(e, true, eCtx4, sub4, "/data1", "read")
	testEnforce(e, true, eCtx4, sub4, "/data2", "write")
	fmt.Println("##########4:abac_model################")
}

type testSub struct {
	Name string
	Age  int
}

func newTestSubject(name string, age int) testSub {
	s := testSub{}
	s.Name = name
	s.Age = age
	return s
}

func testEnforce(e *casbin.Enforcer, result bool, data ...interface{}) {
	if myRes, _ := e.Enforce(data...); myRes != result {
		fmt.Printf("[Fail]%+v: %t upposed to be %t\n", data, myRes, result)
	} else {
		fmt.Printf("[Sucs]%+v: %t upposed to be %t\n", data, myRes, result)
	}
}

func testEnforceEx(e *casbin.Enforcer, res []string, data ...interface{}) {
	_, myRes, _ := e.EnforceEx(data...)

	if ok := util.ArrayEquals(res, myRes); !ok {
		fmt.Println("[Fail]Key: ", myRes, ", supposed to be ", res)
	}
	fmt.Println("[Sucs]Key: ", myRes, ", supposed to be ", res)
}
