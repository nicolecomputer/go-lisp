package lisp

import (
	lisp "lisp/parse"
	parser "lisp/parse"
)

type (
	Environment struct{}
	Result      struct{}
)

func evaluate(exp parser.Expression, env Environment) (error, Result, Environment) {
	return nil, Result{}, Environment{}
}

func evaluateList(list lisp.List) (error, lisp.Expression) {
	operator := list.Children[0]

	switch operatorAtom := operator.(type) {
	case lisp.OperatorAtom:
		if operatorAtom.Value == "+" {
			thingsToAdd := list.Children[1:]
			var total float64 = 0
			for _, child := range thingsToAdd {
				switch childValue := child.(type) {
				case lisp.NumberAtom:
					total += childValue.Value
				}
			}

			return nil, lisp.NumberAtom{total}
		}
	}
	return nil, nil
}

func evaluateStr(expression string) (error, lisp.Expression) {
	error, value := lisp.Parse(expression)

	if error != nil {
		return error, nil
	}

	switch value := value.(type) {
	case lisp.List:
		return evaluateList(value)
	default:
		return nil, value
	}
}

// It would be nice to have
// save environment
// load environment
