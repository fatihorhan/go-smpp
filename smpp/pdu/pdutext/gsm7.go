// Copyright 2015 go-smpp authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package pdutext

import (
	"github.com/fatihorhan/go-smpp/smpp/encoding"
	"golang.org/x/text/transform"
)

// GSM7 is 7-bit (unpacked)
type GSM7 struct {
	bytes []byte
	nli   encoding.NLI
}

func NewGSM7(bytes []byte, nli encoding.NLI) GSM7 {
	return GSM7{
		bytes: bytes,
		nli:   nli,
	}
}

// Type implements the Codec interface.
func (s GSM7) Type() DataCoding {
	return DefaultType
}

// Nli returns NLI
func (s GSM7) Nli() encoding.NLI {
	return s.nli
}

// Encode to GSM 7-bit (unpacked)
func (s GSM7) Encode() []byte {
	e := encoding.GSM7(false, s.nli).NewEncoder()
	es, _, err := transform.Bytes(e, s.bytes)
	if err != nil {
		return s.bytes
	}

	return es
}

// Decode from GSM 7-bit (unpacked)
func (s GSM7) Decode() []byte {
	e := encoding.GSM7(false, s.nli).NewDecoder()
	es, _, err := transform.Bytes(e, s.bytes)
	if err != nil {
		return s.bytes
	}
	return es
}
