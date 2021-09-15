package resolvers

//Unused package
//
import (
	"go/src/question-answer/helpers"
	. "go/src/question-answer/models"
	"github.com/graphql-go/graphql"
)

var db, err = helpers.GetDb()

func Resolver(p graphql.ResolveParams) (interface{}, error) {
	var req User
	db.Find(&req)
	return req, err
}
