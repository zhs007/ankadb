package main

import (
	"fmt"

	"github.com/graphql-go/graphql"
	"github.com/zhs007/ankadb/graphqlext"
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
				Type: graphql.NewList(candleChunkType),
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
					return nil, nil
				},
			},
		},
	},
)

var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query: queryType,
	},
)

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	}
	return result
}
