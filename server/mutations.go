package server

import "context"

type mutationResolver struct{ *rootResolver }

func (r *mutationResolver) Test(ctx context.Context, arg *string) (*string, error) {
	return arg, nil
}
