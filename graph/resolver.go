package graph

import "github.com/minhkstn/go-graphql-tutorial/graph/model"

type Resolver struct {
	CharactorStore map[string]model.Character
}
