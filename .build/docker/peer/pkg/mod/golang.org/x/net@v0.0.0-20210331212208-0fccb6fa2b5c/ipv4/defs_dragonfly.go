// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ignore
// +build ignore

// +godefs map struct_in_addr [4]byte /* in_addr */

package ipv4

/*
#include <netinet/in.h>
*/
import "C"

const (
	sysIP_RECVDSTADDR = C.IP_RECVDSTADDR
	sysIP_RECVIF      = C.IP_RECVIF
	sysIP_RECVTTL     = C.IP_RECVTTL

	sizeofIPMreq = C.sizeof_struct_ip_mreq
)

type ipMreq C.struct_ip_mreq
