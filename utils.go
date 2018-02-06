// Copyright 2016 Factom Foundation
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package evmdbapi

import (
  "fmt"
  "math/big"  
  "github.com/tonnerre/golang-go.crypto/sha3"  
)

func ParseBigInt(q string) *big.Int {
  n, ok := new(big.Int).SetString(q, 0)
  if !ok {
    panic("invalid")
  }
  return n
}

func MethodHash(method string) string {
  h := sha3.NewKeccak256()
  h.Write([]byte(method))
  var digest [32]byte
  h.Sum(digest[:0])
  return fmt.Sprintf("%x", digest[:4])
}
