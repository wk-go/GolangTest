--
-- Created by IntelliJ IDEA.
-- User: walkskyer
-- Date: 2021/4/15
-- Time: 11:28
-- To change this template use File | Settings | File Templates.
--
local m = require("mymodule")
m.myfunc()
print(m.name)
print(m.field1)

local m2 = require("mymodule2")
m2.myfunc()
print(m2.name)
print(m2.field1)
print("number sum:1+2="..m2.sum_lua(1,2))
print("number sum:1+2+3+4+5="..m2.sum_lua(1,2,3,4,5))