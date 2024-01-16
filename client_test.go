package unisat_sdk_go

import (
	"encoding/json"
	"testing"
)

var client *Client

func init() {
	client = NewClient("", "")
}

func TestClient_GetBRC20Status(t *testing.T) {
	type args struct {
		start int
		limit int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "",
			args: args{
				start: 30,
				limit: 500,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if brc20Status, err := client.GetBRC20Status(tt.args.start, tt.args.limit); (err != nil) != tt.wantErr {
				t.Errorf("GetBRC20Status() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				t.Logf("brc20Status: %v", brc20Status)
			}
		})
	}
}

func TestClient_GetBrc20TickerInfo(t *testing.T) {
	type args struct {
		ticker string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "",
			args: args{
				ticker: "ordi",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := client.GetBrc20TickerInfo(tt.args.ticker)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBrc20TickerInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			marshal, _ := json.Marshal(got)
			t.Logf(string(marshal))
		})
	}
}

func TestClient_GetBrc20Holders(t *testing.T) {
	type args struct {
		ticker string
		start  int
		limit  int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "",
			args: args{
				ticker: "ordi",
				start:  0,
				limit:  0,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := client.GetBrc20Holders(tt.args.ticker, tt.args.start, tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBrc20Holders() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			marshal, _ := json.Marshal(got)
			t.Logf(string(marshal))
		})
	}
}

func TestClient_GetAddressSummary(t *testing.T) {
	type args struct {
		address string
		start   int
		limit   int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "",
			args: args{
				address: "bc1qggf48ykykz996uv5vsp5p9m9zwetzq9run6s64hm6uqfn33nhq0ql9t85q",
				start:   0,
				limit:   0,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := client.GetAddressSummary(tt.args.address, tt.args.start, tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAddressSummary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			marshal, _ := json.Marshal(got)
			t.Logf(string(marshal))
		})
	}
}
