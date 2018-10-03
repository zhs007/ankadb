package graphqlext

import (
	"strconv"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
)

const (
	// MaxInt53 2^53 - 1
	MaxInt53 = 1<<53 - 1
	// MinInt53 -(2^53 - 1)
	MinInt53 = -((1 << 53) - 1)
)

// As per the GraphQL Spec, Integers are only treated as valid when a valid
// 64-bit signed integer, providing the broadest support across platforms.
//
// n.b. JavaScript's integers are safe between -(2^53 - 1) and 2^53 - 1 because
// they are internally represented as IEEE 754 doubles.
func coerceInt(value interface{}) interface{} {
	switch value := value.(type) {
	case bool:
		if value == true {
			return 1
		}
		return 0
	case *bool:
		if value == nil {
			return nil
		}
		return coerceInt(*value)
	case int:
		// if value < int64(MinInt53) || value > int64(MaxInt53) {
		// 	return nil
		// }
		return value
	case *int:
		if value == nil {
			return nil
		}
		return coerceInt(*value)
	case int8:
		return int64(value)
	case *int8:
		if value == nil {
			return nil
		}
		return int64(*value)
	case int16:
		return int64(value)
	case *int16:
		if value == nil {
			return nil
		}
		return int64(*value)
	case int32:
		return int64(value)
	case *int32:
		if value == nil {
			return nil
		}
		return int64(*value)
	case int64:
		if value < int64(MinInt53) || value > int64(MaxInt53) {
			return nil
		}
		return int64(value)
	case *int64:
		if value == nil {
			return nil
		}
		return coerceInt(*value)
	case uint:
		// if value > uint64(MaxInt53) {
		// 	return nil
		// }
		return int64(value)
	case *uint:
		if value == nil {
			return nil
		}
		return coerceInt(*value)
	case uint8:
		return int64(value)
	case *uint8:
		if value == nil {
			return nil
		}
		return int64(*value)
	case uint16:
		return int64(value)
	case *uint16:
		if value == nil {
			return nil
		}
		return int64(*value)
	case uint32:
		// if value > uint32(MaxInt53) {
		// 	return nil
		// }
		return int64(value)
	case *uint32:
		if value == nil {
			return nil
		}
		return coerceInt(*value)
	case uint64:
		if value > uint64(MaxInt53) {
			return nil
		}
		return int64(value)
	case *uint64:
		if value == nil {
			return nil
		}
		return coerceInt(*value)
	case float32:
		// if value < float64(MinInt53) || value > float64(MaxInt53) {
		// 	return nil
		// }
		return int64(value)
	case *float32:
		if value == nil {
			return nil
		}
		return coerceInt(*value)
	case float64:
		if value < float64(MinInt53) || value > float64(MaxInt53) {
			return nil
		}
		return int64(value)
	case *float64:
		if value == nil {
			return nil
		}
		return coerceInt(*value)
	case string:
		val, err := strconv.ParseFloat(value, 0)
		if err != nil {
			return nil
		}
		return coerceInt(val)
	case *string:
		if value == nil {
			return nil
		}
		return coerceInt(*value)
	}

	// If the value cannot be transformed into an int, return nil instead of '0'
	// to denote 'no integer found'
	return nil
}

// Int64 is the GraphQL Integer type definition.
var Int64 = graphql.NewScalar(graphql.ScalarConfig{
	Name: "Int64",
	Description: "The `Int64` scalar type represents non-fractional signed whole numeric " +
		"values. Int can represent values between -(2^53 - 1) and 2^53 - 1. ",
	Serialize:  coerceInt,
	ParseValue: coerceInt,
	ParseLiteral: func(valueAST ast.Value) interface{} {
		switch valueAST := valueAST.(type) {
		case *ast.IntValue:
			if intValue, err := strconv.Atoi(valueAST.Value); err == nil {
				return intValue
			}
		}
		return nil
	},
})

// Timestamp is the GraphQL Integer type definition.
var Timestamp = graphql.NewScalar(graphql.ScalarConfig{
	Name: "Timestamp",
	Description: "The `Timestamp` scalar type represents non-fractional signed whole numeric " +
		"values. Int can represent values between -(2^53 - 1) and 2^53 - 1. ",
	Serialize:  coerceInt,
	ParseValue: coerceInt,
	ParseLiteral: func(valueAST ast.Value) interface{} {
		switch valueAST := valueAST.(type) {
		case *ast.IntValue:
			if intValue, err := strconv.Atoi(valueAST.Value); err == nil {
				return intValue
			}
		}
		return nil
	},
})
