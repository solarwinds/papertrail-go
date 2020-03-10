// Copyright 2020 Solarwinds Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package papertrailgo

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/sirupsen/logrus"
)

func TestNewPapertrailLogger(t *testing.T) {
	logrus.SetLevel(logrus.DebugLevel)
	type args struct {
		ctx                context.Context
		paperTrailProtocol string
		paperTrailHost     string
		paperTrailPort     int
		dbLocation         string
		retention          time.Duration
		workerCount        int
		maxDiskUsage       float64
	}
	tests := []struct {
		name    string
		args    args
		want    *Logger
		wantErr bool
	}{
		{
			name: "valid case",
			args: args{
				ctx:                context.Background(),
				paperTrailProtocol: "udp",
				paperTrailHost:     "logs5.papertrailapp.com",
				paperTrailPort:     29646,
				dbLocation:         os.TempDir(),
				retention:          time.Minute * 5,
				workerCount:        4,
				maxDiskUsage:       50,
			},
			wantErr: false,
		},
	}
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewPapertrailLogger(tt.args.ctx, tt.args.paperTrailProtocol, tt.args.paperTrailHost, tt.args.paperTrailPort, tt.args.dbLocation, tt.args.retention, tt.args.workerCount, tt.args.maxDiskUsage)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewPapertrailLogger() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.Log(&Payload{
				Hostname: fmt.Sprintf("blah%d/blah/asdfd-fdasfad-fasdadf", i),
				Tag:      fmt.Sprintf("container%d/fdaf/asdf/das/fadsf-sdf", i),
				Log:      fmt.Sprintf("log line %d", i),
				LogTime:  ptypes.TimestampNow(),
			}) != nil {
				t.Errorf("unexpected error while logging")
				return
			}
			time.Sleep(time.Second * 30)
			if got.Close() != nil {
				t.Errorf("unexpected error while closing")
				return
			}
			// if !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("NewPapertrailLogger() = %v, want %v", got, tt.want)
			// }
		})
	}
}
