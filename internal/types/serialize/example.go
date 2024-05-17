package serialize

import (
	"github.com/ldd27/go-starter-kit/internal/model"
	"github.com/ldd27/go-starter-kit/internal/types"
)

func Example2Api(item model.Example) types.Example {
	return types.Example{
		ID:   item.ID,
		Name: item.Name,
	}
}

func Example2Apis(items []model.Example) []types.Example {
	res := make([]types.Example, 0, len(items))
	for _, item := range items {
		res = append(res, Example2Api(item))
	}
	return res
}
