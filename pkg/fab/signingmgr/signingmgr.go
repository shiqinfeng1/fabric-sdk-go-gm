/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package signingmgr

import (
	"github.com/shiqinfeng1/fabric-sdk-go-gm/pkg/common/providers/core"

	"github.com/pkg/errors"
	"github.com/shiqinfeng1/fabric-sdk-go-gm/pkg/core/cryptosuite"
)

// SigningManager is used for signing objects with private key
type SigningManager struct {
	cryptoProvider core.CryptoSuite
	hashOpts       core.HashOpts
	signerOpts     core.SignerOpts
}

// New Constructor for a signing manager.
// @param {BCCSP} cryptoProvider - crypto provider
// @param {Config} config - configuration provider
// @returns {SigningManager} new signing manager
func New(cryptoProvider core.CryptoSuite) (*SigningManager, error) {
	return &SigningManager{cryptoProvider: cryptoProvider, hashOpts: cryptosuite.GetGMSM3Opts()}, nil
}

// Sign will sign the given object using provided key
func (mgr *SigningManager) Sign(object []byte, key core.Key) ([]byte, error) {

	if len(object) == 0 {
		return nil, errors.New("object (to sign) required")
	}

	if key == nil {
		return nil, errors.New("key (for signing) required")
	}

	//digest, err := mgr.cryptoProvider.Hash(object, mgr.hashOpts)
	/*if err != nil {
		return nil, err
	}*/
	signature, err := mgr.cryptoProvider.Sign(key, object, mgr.signerOpts)
	if err != nil {
		return nil, err
	}
	return signature, nil
}
