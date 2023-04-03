package dtrack

import (
	"context"
	"net/http"
)

type PermissionService struct {
	client *Client
}

type Permission struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (ps PermissionService) GetAll(ctx context.Context, po PageOptions) (p Page[Permission], err error) {
	req, err := ps.client.newRequest(ctx, http.MethodGet, "/api/v1/permission", withPageOptions(po))
	if err != nil {
		return
	}

	res, err := ps.client.doRequest(req, &p.Items)
	if err != nil {
		return
	}

	p.TotalCount = res.TotalCount
	return
}
