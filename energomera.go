// Package energomera provides primitives for decoding and encoding
// GOST-R MEK 61107-2001 (ГОСТ Р МЭК 61107-2001) protocol messages.
package energomera

import (
	"fmt"
)
// Energomera data decoder is used to decode message received from bus
func DataDecode(sdata []byte) (msg map[string]string) {

	msg = map[string]string{
		"head": "",
		"body": "",
		"lrc":  "False",
	}

	lrc := uint8(0)

	if len(sdata) <= 1 {
		msg["body"] = string(sdata)
		msg["lrc"] = "True"
	} else {
		headAdd := "False"
		bodyAdd := "False"
		lrcAdd := "False"

		for i := 0; i < (len(sdata) - 1); i++ {
			// Catch SOH
			if sdata[i] == 1 {
				headAdd = "True"
				lrcAdd = "True"
			} else if sdata[i] == 2 {
				headAdd = "False"
				bodyAdd = "True"
				if lrcAdd == "True" {
					lrc = (lrc + sdata[i]) & 0x7f
				} else {
					lrcAdd = "True"
				}
			} else if sdata[i] == 3 {
				headAdd = "False"
				bodyAdd = "False"
				lrcAdd = "False"
				lrc = (lrc + sdata[i]) & 0x7f
			} else {
				if headAdd == "True" {
					msg["head"] += string(sdata[i])
				} else if bodyAdd == "True" {
					msg["body"] += string(sdata[i])
				}
				if lrcAdd == "True" {
					lrc = (lrc + uint8(sdata[i])) & 0x7f
				}
			}
		}
	}

	if lrc == sdata[len(sdata)-1]&0x7f {
		msg["lrc"] = "True"
	} else {
		msg["lrc"] = "False"
	}

	fmt.Printf("%q", msg)

	return
}

// Energomera data encoder is used to encode data which will be sent over bus
func DataEncode(msg map[string]string) (sdata []byte) {
	if len(msg["head"]) > 0 {
		sdata = append(sdata, []byte{0x01}...)
		sdata = append(sdata, msg["head"]...)
	}
	if len(msg["body"]) > 0 {
		sdata = append(sdata, []byte{0x02}...)
		sdata = append(sdata, msg["body"]...)
	}
	sdata = append(sdata, []byte{0x03}...)

	lrc := uint8(0)
	lrcAdd := "False"

	for i := 0; i < len(sdata); i++ {
		if sdata[i] == 1 {
			lrcAdd = "True"
		} else if sdata[i] == 2 {
			if lrcAdd == "True" {
				lrc = (lrc + uint8(sdata[i])) & 0x7f
			} else {
				lrcAdd = "True"
			}
		} else if sdata[i] == 2 {
			lrcAdd = "False"
			lrc = (lrc + uint8(sdata[i])) & 0x7f
		} else {
			if lrcAdd == "True" {
				lrc = (lrc + uint8(sdata[i])) & 0x7f
			}
		}
	}

	sdata = append(sdata, lrc)

	return
}
