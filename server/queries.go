package server

import "context"

type queryResolver struct{ *rootResolver }

func (r *queryResolver) Test(ctx context.Context, arg *string) (*string, error) {
	return arg, nil
}
