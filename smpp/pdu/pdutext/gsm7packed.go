// Copyright 2015 go-smpp authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package pdutext

import (
	"github.com/fatihorhan/go-smpp/smpp/encoding"
	"golang.org/x/text/transform"
)

// GSM 7-bit (packed)
type GSM7Packed struct {
	bytes []byte
	nli   encoding.NLI
}

// Type implements the Codec interface.
func (s GSM7Packed) Type() DataCoding {
	return DefaultType
}

// Encode to GSM 7-bit (packed)
func (s GSM7Packed) Encode() []byte {
	e := encoding.GSM7(true, s.nli).NewEncoder()
	es, _, err := transform.Bytes(e, s.bytes)
	if err != nil {
		return s.bytes
	}
	return es
}

// Decode from GSM 7-bit (packed)
func (s GSM7Packed) Decode() []byte {
	e := encoding.GSM7(true, s.nli).NewDecoder()
	es, _, err := transform.Bytes(e, s.bytes)
	if err != nil {
		return s.bytes
	}
	return es
}
