package graph

import "starktrix/ql/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	links []*model.Link
	count int32
}
