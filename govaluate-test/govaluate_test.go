package govaluate_test

import (
	"fmt"
	"github.com/Knetic/govaluate"
)

func main()  {
	aexpression, _ := govaluate.NewEvaluableExpression("10 > 0")
	aresult, aerr := aexpression.Evaluate(nil)
	// result is now set to "true", the bool value.
	if aerr != nil {
		panic(aerr.Error())
	} else {
		fmt.Println(aresult)
	}
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	bexpression, berr := govaluate.NewEvaluableExpression("foo > 0")

	bparameters := make(map[string]interface{}, 8)
	bparameters["foo"] = -1

	bresult, berr := bexpression.Evaluate(bparameters)
	// result is now set to "false", the bool value.
	if aerr != nil {
		panic(berr.Error())
	} else {
		fmt.Println(bresult)
	}
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	cexpression, cerr := govaluate.NewEvaluableExpression("(requests_made * requests_succeeded / 100) >= 80 && name == 'test' && isBool == true && floatTest == 1.35")

	cparameters := make(map[string]interface{}, 8)
	cparameters["requests_made"] = 100
	cparameters["requests_succeeded"] = 80
	cparameters["name"] = "test"
	cparameters["test"] = "test"
	cparameters["isBool"] = true
	cparameters["floatTest"] = 1.35

	fmt.Println(cparameters)

	cresult, cerr := cexpression.Evaluate(cparameters)
	// result is now set to "false", the bool value.
	if cerr != nil {
		panic(cerr.Error())
	} else {
		fmt.Println(cresult)
	}
}
