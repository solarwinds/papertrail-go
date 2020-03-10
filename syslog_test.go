package papertrail_go

import (
	"testing"
)

func TestNewPapertailShipper(t *testing.T) {
	type args struct {
		paperTrailProtocol string
		paperTrailHost     string
		paperTrailPort     int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		wantNil bool
	}{
		{
			name: "valid case 1",
			args: args{
				paperTrailHost:     "logs.papertrailapp.com",
				paperTrailPort:     34513,
				paperTrailProtocol: "udp",
			},
			wantErr: false,
			wantNil: false,
		},
		{
			name: "valid case 2",
			args: args{
				paperTrailHost:     "logs.papertrailapp.com",
				paperTrailPort:     34513,
				paperTrailProtocol: "tcp",
			},
			wantErr: false,
			wantNil: false,
		},
		{
			name: "valid case 3",
			args: args{
				paperTrailHost:     "logs.papertrailapp.com",
				paperTrailPort:     34513,
				paperTrailProtocol: "tls",
			},
			wantErr: false,
			wantNil: false,
		},
		{
			name: "empty host case",
			args: args{
				paperTrailHost:     "",
				paperTrailPort:     34513,
				paperTrailProtocol: "udp",
			},
			wantErr: true,
			wantNil: true,
		},
		{
			name: "invalid protocol",
			args: args{
				paperTrailHost:     "logs.papertrailapp.com",
				paperTrailPort:     34513,
				paperTrailProtocol: "udp4",
			},
			wantErr: true,
			wantNil: true,
		},
		{
			name: "valid port",
			args: args{
				paperTrailHost:     "logs.papertrailapp.com",
				paperTrailProtocol: "udp",
			},
			wantErr: true,
			wantNil: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewPapertailShipper(tt.args.paperTrailProtocol, tt.args.paperTrailHost, tt.args.paperTrailPort)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewPapertailShipper() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == nil && !tt.wantNil {
				t.Errorf("expected a non-nil value")
			}
		})
	}
}

func Test_getCerts(t *testing.T) {
	got, err := getCerts()
	if err != nil {
		t.Errorf("unable to fetch certs, error: %v", err)
		t.FailNow()
	}
	if len(got) == 0 {
		t.Error("cert is empty")
		t.FailNow()
	}
}

func Test_getRootCAs(t *testing.T) {
	got, err := getRootCAs()
	if err != nil {
		t.Errorf("unable to fetch root CAs, error: %v", err)
		t.FailNow()
	}
	if got == nil {
		t.Error("cert is empty")
		t.FailNow()
	}
}
