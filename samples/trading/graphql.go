package main

import (
	"fmt"

	"github.com/graphql-go/graphql"
	"github.com/zhs007/ankadb/graphqlext"

	pb "github.com/zhs007/ankadb/samples/trading/proto"
)

var c = pb.Candle{
	Open:         1,
	Close:        2,
	High:         3,
	Low:          4,
	Volume:       5,
	OpenInterest: 6,
	Curtime:      7,
}

var candleInputType = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "CandleInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"curtime": &graphql.InputObjectFieldConfig{
				Type: graphqlext.Timestamp,
			},
			"open": &graphql.InputObjectFieldConfig{
				Type: graphqlext.Int64,
			},
			"close": &graphql.InputObjectFieldConfig{
				Type: graphqlext.Int64,
			},
			"high": &graphql.InputObjectFieldConfig{
				Type: graphqlext.Int64,
			},
			"low": &graphql.InputObjectFieldConfig{
				Type: graphqlext.Int64,
			},
			"volume": &graphql.InputObjectFieldConfig{
				Type: graphqlext.Int64,
			},
			"openInterest": &graphql.InputObjectFieldConfig{
				Type: graphqlext.Int64,
			},
		},
	},
)

// candleType - Candle
var candleType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Candle",
		Fields: graphql.Fields{
			"curtime": &graphql.Field{
				Type: graphqlext.Timestamp,
			},
			"open": &graphql.Field{
				Type: graphqlext.Int64,
			},
			"close": &graphql.Field{
				Type: graphqlext.Int64,
			},
			"high": &graphql.Field{
				Type: graphqlext.Int64,
			},
			"low": &graphql.Field{
				Type: graphqlext.Int64,
			},
			"volume": &graphql.Field{
				Type: graphqlext.Int64,
			},
			"openInterest": &graphql.Field{
				Type: graphqlext.Int64,
			},
		},
	},
)

// candleChunkType - CandleChunk
var candleChunkType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "CandleChunk",
		Fields: graphql.Fields{
			"code": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"startTime": &graphql.Field{
				Type: graphqlext.Timestamp,
			},
			"endTime": &graphql.Field{
				Type: graphqlext.Timestamp,
			},
			"candles": &graphql.Field{
				Type: graphql.NewList(candleType),
			},
		},
	},
)

var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"candleChunks": &graphql.Field{
				Type: candleType,
				Args: graphql.FieldConfigArgument{
					"code": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"name": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"startTime": &graphql.ArgumentConfig{
						Type: graphqlext.Timestamp,
					},
					"endTime": &graphql.ArgumentConfig{
						Type: graphqlext.Timestamp,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					// idQuery, isOK := p.Args["id"].(string)
					// if isOK {
					// return data[idQuery], nil
					// }
					return c, nil
				},
			},
		},
	},
)

var mutationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"insertCandles": &graphql.Field{
			Type:        candleChunkType,
			Description: "insert candles",
			Args: graphql.FieldConfigArgument{
				"code": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"endTime": &graphql.ArgumentConfig{
					Type: graphqlext.Timestamp,
				},
				"candle": &graphql.ArgumentConfig{
					Type: candleInputType,
				},
				// "candles": &graphql.ArgumentConfig{
				// 	Type: graphql.NewList(candleType),
				// },
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				// rand.Seed(time.Now().UnixNano())
				// product := Product{
				// 	ID:    int64(rand.Intn(100000)), // generate random ID
				// 	Name:  params.Args["name"].(string),
				// 	Info:  params.Args["info"].(string),
				// 	Price: params.Args["price"].(float64),
				// }
				// products = append(products, product)
				return nil, nil
			},
		},
	},
})

var curTypes = []graphql.Type{candleType}

var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:    queryType,
		Mutation: mutationType,
		Types:    curTypes,
	},
)

func executeQuery(query string, mapvar map[string]interface{}, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:         schema,
		RequestString:  query,
		VariableValues: mapvar,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	}
	return result
}
