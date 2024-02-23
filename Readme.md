# Klaviyo's Unofficial Golang implementation

## Filter Builder

```go

fb := commons.NewFilerBuilder()

//Simple operations

// equals
fb.Contains("name","test name") 
fb.Build()// "contains(name, test name)"

//Chain operators
fb.Equals("field1","value1").LessThan("field2","value2") 
fb.Build()// "equals(field1,value1),less-than(field2,value2)"

//Boolean Operators

op1 := commons.NewFilerBuilder()
api.Equals("field1","value1")
op2 := commons.NewFilerBuilder()
api.LessThan("field2","value2")

fb.And(op1,op2)

fb.Build() // and(equals(field1,value1),less-than(field2,value2))

```
