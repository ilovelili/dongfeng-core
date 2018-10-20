// Package handlers define the core behaviors of each API
package handlers

import (
	"context"

	api "github.com/ilovelili/dongfeng/core/services/proto"
)

// Healthcheck return 200
func (f *Facade) Healthcheck(ctx context.Context, req *api.HealthcheckRequest, rsp *api.HealthcheckResponse) error {
	rsp.Message = "OK"
	return nil
}
