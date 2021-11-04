package revenuecat

import (
	"context"
	"testing"
)

func TestGrantEntitlement(t *testing.T) {
	cl := newMockClient(t, 200, nil, nil)
	rc := New("apikey")
	rc.http = cl

	_, err := rc.GrantEntitlement(context.Background(), "123", "all", ThreeMonth, staticTime(t, "2020-01-15 23:54:17"))
	if err != nil {
		t.Errorf("error: %v", err)
	}

	cl.expectMethod(t, "POST")
	cl.expectPath(t, "/v1/subscribers/123/entitlements/all/promotional")
	cl.expectBody(t, `{"duration":"three_month","start_time_ms":1579132457000}`)
}

func TestRevokeEntitlement(t *testing.T) {
	cl := newMockClient(t, 200, nil, nil)
	rc := New("apikey")
	rc.http = cl

	_, err := rc.RevokeEntitlement(context.Background(), "123", "all")
	if err != nil {
		t.Errorf("error: %v", err)
	}

	cl.expectMethod(t, "POST")
	cl.expectPath(t, "/v1/subscribers/123/entitlements/all/revoke_promotionals")
}
