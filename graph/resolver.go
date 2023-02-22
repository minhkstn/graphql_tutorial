package graph

import "github.com/minhkstn/graphql_tutorial/graph/model"

type Resolver struct {
	CharactorStore map[string]model.Character
}
