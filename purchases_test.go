package revenuecat

import (
	"context"
	"testing"
)

func TestCreatePurchaseWithoutOpts(t *testing.T) {
	cl := newMockClient(t, 200, nil, nil)
	rc := New("apikey")
	rc.http = cl

	_, err := rc.CreatePurchase(context.Background(), "123", "testreceipt", nil)
	if err != nil {
		t.Errorf("error: %v", err)
	}

	cl.expectMethod(t, "POST")
	cl.expectPath(t, "/v1/receipts")
	cl.expectXPlatform(t, "")
	cl.expectBody(t, `{"app_user_id":"123","fetch_token":"testreceipt"}`)
}

func TestCreatePurchaseWithOpts(t *testing.T) {
	cl := newMockClient(t, 200, nil, nil)
	rc := New("apikey")
	rc.http = cl

	opt := &CreatePurchaseOptions{
		Platform: "ios",

		ProductID: "product_sku",
		IsRestore: true,
	}

	_, err := rc.CreatePurchase(context.Background(), "123", "testreceipt", opt)
	if err != nil {
		t.Errorf("error: %v", err)
	}

	cl.expectMethod(t, "POST")
	cl.expectPath(t, "/v1/receipts")
	cl.expectXPlatform(t, "ios")
	cl.expectBody(t, `{"app_user_id":"123","fetch_token":"testreceipt","product_id":"product_sku","is_restore":true}`)
}
