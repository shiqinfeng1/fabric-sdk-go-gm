/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package discovery

import (
	"path/filepath"
	"testing"

	"github.com/shiqinfeng1/fabric-sdk-go-gm/pkg/client/common/discovery/staticdiscovery"
	"github.com/shiqinfeng1/fabric-sdk-go-gm/pkg/common/providers/fab"
	"github.com/shiqinfeng1/fabric-sdk-go-gm/pkg/core/config"
	fabImpl "github.com/shiqinfeng1/fabric-sdk-go-gm/pkg/fab"
	mocks "github.com/shiqinfeng1/fabric-sdk-go-gm/pkg/fab/mocks"
	"github.com/shiqinfeng1/fabric-sdk-go-gm/pkg/msp/test/mockmsp"
	"github.com/shiqinfeng1/fabric-sdk-go-gm/test/metadata"
)

const configFile = "config_test.yaml"

type mockFilter struct {
	called bool
}

// Accept returns true if this peer is to be included in the target list
func (df *mockFilter) Accept(peer fab.Peer) bool {
	df.called = true
	return true
}

func TestDiscoveryFilter(t *testing.T) {

	configBackend, err := config.FromFile(filepath.Join(metadata.GetProjectPath(), metadata.SDKConfigPath, configFile))()
	if err != nil {
		t.Fatalf(err.Error())
	}

	config1, err := fabImpl.ConfigFromBackend(configBackend...)
	if err != nil {
		t.Fatalf(err.Error())
	}

	discoveryService, err := staticdiscovery.NewService(config1, mocks.NewMockContext(mockmsp.NewMockSigningIdentity("user1", "Org1MSP")).InfraProvider(), "mychannel")
	if err != nil {
		t.Fatalf("Failed to setup discovery service: %s", err)
	}

	discoveryFilter := &mockFilter{called: false}

	filteredService := NewDiscoveryFilterService(discoveryService, discoveryFilter)

	peers, err := filteredService.GetPeers()
	if err != nil {
		t.Fatalf("Failed to get peers from discovery service: %s", err)
	}

	// One peer is configured for "mychannel"
	expectedNumOfPeers := 1
	if len(peers) != expectedNumOfPeers {
		t.Fatalf("Expecting %d, got %d peers", expectedNumOfPeers, len(peers))
	}

	if !discoveryFilter.called {
		t.Fatal("Expecting true, got false")
	}

}
