//--------------------------------------------------------------------------
// Copyright 2018-2022 infinimesh
// www.infinimesh.io
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.
//--------------------------------------------------------------------------

package main

import (
	"errors"
	"fmt"
	"net"

	devpb "github.com/infinimesh/infinimesh/pkg/node/proto/devices"
	"github.com/slntopp/mqtt-go/packet"
	"go.uber.org/zap"
)

func HandleTCPConnections(tcp net.Listener) {
	log := log.Named("HandleTCPConnections")
	for {
		conn, _ := tcp.Accept() // nolint: gosec

		p, err := packet.ReadPacket(conn, 0)
		if err != nil {
			LogErrorAndClose(conn, fmt.Errorf("error while reading connect packet: %v", err))
			continue
		}
		log.Debug("ControlPacket", zap.Any("packet", p))

		connectPacket, ok := p.(*packet.ConnectControlPacket)
		if !ok {
			LogErrorAndClose(conn, errors.New("first packet isn't ConnectControlPacket"))
			continue
		}
		log.Debug("ConnectPacket", zap.Any("packet", p))

		var fingerprint []byte
		fingerprint, err = verifyBasicAuth(connectPacket)
		if err != nil {
			LogErrorAndClose(conn, fmt.Errorf("error verifying Basic Auth: %v", err))
			continue
		}

		log.Debug("Fingerprint", zap.ByteString("fingerprint", fingerprint))

		device, err := GetByFingerprintAndVerify(fingerprint, func(device *devpb.Device) (bool) {
			if device.Title != connectPacket.ConnectPayload.Username {
				log.Error("Failed to verify client as the device name doesn't match Basic Auth Username", zap.String("uuid", device.Uuid), zap.String("device", device.Title), zap.String("username", connectPacket.ConnectPayload.Username))
				return false
			} else if !device.BasicEnabled {
				log.Error("Failed to verify client as the device is not enabled for Basic Auth", zap.String("uuid", device.Uuid))
				return false
			} else if !device.Enabled {
				log.Error("Failed to verify client as the device is not enabled", zap.String("uuid", device.Uuid))
				return false
			} else {
				log.Info("Verified client as the device is enabled", zap.String("uuid", device.Uuid), zap.Strings("tags", device.Tags))
				return true
			}
		})
		if err != nil {
			LogErrorAndClose(conn, err)
			continue
		}

		log.Info("Client connected", zap.String("device", device.Uuid))

		go HandleConn(conn, connectPacket, device)
	}
}