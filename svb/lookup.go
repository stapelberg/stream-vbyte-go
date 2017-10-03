// Copyright 2017 Nelz
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

package svb

// lengths is the lookup table that is the shortcut to get from a control
// byte to knowing the 4 individual data byte lengths (in index 0 - 3), as
// well as the total data bytes needed for the given block (in index 4)
var lookup = map[byte][4]uint8{
	0x00: [4]uint8{1, 1, 1, 1},
	0x01: [4]uint8{1, 1, 1, 2},
	0x02: [4]uint8{1, 1, 1, 3},
	0x03: [4]uint8{1, 1, 1, 4},
	0x04: [4]uint8{1, 1, 2, 1},
	0x05: [4]uint8{1, 1, 2, 2},
	0x06: [4]uint8{1, 1, 2, 3},
	0x07: [4]uint8{1, 1, 2, 4},
	0x08: [4]uint8{1, 1, 3, 1},
	0x09: [4]uint8{1, 1, 3, 2},
	0x0a: [4]uint8{1, 1, 3, 3},
	0x0b: [4]uint8{1, 1, 3, 4},
	0x0c: [4]uint8{1, 1, 4, 1},
	0x0d: [4]uint8{1, 1, 4, 2},
	0x0e: [4]uint8{1, 1, 4, 3},
	0x0f: [4]uint8{1, 1, 4, 4},
	0x10: [4]uint8{1, 2, 1, 1},
	0x11: [4]uint8{1, 2, 1, 2},
	0x12: [4]uint8{1, 2, 1, 3},
	0x13: [4]uint8{1, 2, 1, 4},
	0x14: [4]uint8{1, 2, 2, 1},
	0x15: [4]uint8{1, 2, 2, 2},
	0x16: [4]uint8{1, 2, 2, 3},
	0x17: [4]uint8{1, 2, 2, 4},
	0x18: [4]uint8{1, 2, 3, 1},
	0x19: [4]uint8{1, 2, 3, 2},
	0x1a: [4]uint8{1, 2, 3, 3},
	0x1b: [4]uint8{1, 2, 3, 4},
	0x1c: [4]uint8{1, 2, 4, 1},
	0x1d: [4]uint8{1, 2, 4, 2},
	0x1e: [4]uint8{1, 2, 4, 3},
	0x1f: [4]uint8{1, 2, 4, 4},
	0x20: [4]uint8{1, 3, 1, 1},
	0x21: [4]uint8{1, 3, 1, 2},
	0x22: [4]uint8{1, 3, 1, 3},
	0x23: [4]uint8{1, 3, 1, 4},
	0x24: [4]uint8{1, 3, 2, 1},
	0x25: [4]uint8{1, 3, 2, 2},
	0x26: [4]uint8{1, 3, 2, 3},
	0x27: [4]uint8{1, 3, 2, 4},
	0x28: [4]uint8{1, 3, 3, 1},
	0x29: [4]uint8{1, 3, 3, 2},
	0x2a: [4]uint8{1, 3, 3, 3},
	0x2b: [4]uint8{1, 3, 3, 4},
	0x2c: [4]uint8{1, 3, 4, 1},
	0x2d: [4]uint8{1, 3, 4, 2},
	0x2e: [4]uint8{1, 3, 4, 3},
	0x2f: [4]uint8{1, 3, 4, 4},
	0x30: [4]uint8{1, 4, 1, 1},
	0x31: [4]uint8{1, 4, 1, 2},
	0x32: [4]uint8{1, 4, 1, 3},
	0x33: [4]uint8{1, 4, 1, 4},
	0x34: [4]uint8{1, 4, 2, 1},
	0x35: [4]uint8{1, 4, 2, 2},
	0x36: [4]uint8{1, 4, 2, 3},
	0x37: [4]uint8{1, 4, 2, 4},
	0x38: [4]uint8{1, 4, 3, 1},
	0x39: [4]uint8{1, 4, 3, 2},
	0x3a: [4]uint8{1, 4, 3, 3},
	0x3b: [4]uint8{1, 4, 3, 4},
	0x3c: [4]uint8{1, 4, 4, 1},
	0x3d: [4]uint8{1, 4, 4, 2},
	0x3e: [4]uint8{1, 4, 4, 3},
	0x3f: [4]uint8{1, 4, 4, 4},
	0x40: [4]uint8{2, 1, 1, 1},
	0x41: [4]uint8{2, 1, 1, 2},
	0x42: [4]uint8{2, 1, 1, 3},
	0x43: [4]uint8{2, 1, 1, 4},
	0x44: [4]uint8{2, 1, 2, 1},
	0x45: [4]uint8{2, 1, 2, 2},
	0x46: [4]uint8{2, 1, 2, 3},
	0x47: [4]uint8{2, 1, 2, 4},
	0x48: [4]uint8{2, 1, 3, 1},
	0x49: [4]uint8{2, 1, 3, 2},
	0x4a: [4]uint8{2, 1, 3, 3},
	0x4b: [4]uint8{2, 1, 3, 4},
	0x4c: [4]uint8{2, 1, 4, 1},
	0x4d: [4]uint8{2, 1, 4, 2},
	0x4e: [4]uint8{2, 1, 4, 3},
	0x4f: [4]uint8{2, 1, 4, 4},
	0x50: [4]uint8{2, 2, 1, 1},
	0x51: [4]uint8{2, 2, 1, 2},
	0x52: [4]uint8{2, 2, 1, 3},
	0x53: [4]uint8{2, 2, 1, 4},
	0x54: [4]uint8{2, 2, 2, 1},
	0x55: [4]uint8{2, 2, 2, 2},
	0x56: [4]uint8{2, 2, 2, 3},
	0x57: [4]uint8{2, 2, 2, 4},
	0x58: [4]uint8{2, 2, 3, 1},
	0x59: [4]uint8{2, 2, 3, 2},
	0x5a: [4]uint8{2, 2, 3, 3},
	0x5b: [4]uint8{2, 2, 3, 4},
	0x5c: [4]uint8{2, 2, 4, 1},
	0x5d: [4]uint8{2, 2, 4, 2},
	0x5e: [4]uint8{2, 2, 4, 3},
	0x5f: [4]uint8{2, 2, 4, 4},
	0x60: [4]uint8{2, 3, 1, 1},
	0x61: [4]uint8{2, 3, 1, 2},
	0x62: [4]uint8{2, 3, 1, 3},
	0x63: [4]uint8{2, 3, 1, 4},
	0x64: [4]uint8{2, 3, 2, 1},
	0x65: [4]uint8{2, 3, 2, 2},
	0x66: [4]uint8{2, 3, 2, 3},
	0x67: [4]uint8{2, 3, 2, 4},
	0x68: [4]uint8{2, 3, 3, 1},
	0x69: [4]uint8{2, 3, 3, 2},
	0x6a: [4]uint8{2, 3, 3, 3},
	0x6b: [4]uint8{2, 3, 3, 4},
	0x6c: [4]uint8{2, 3, 4, 1},
	0x6d: [4]uint8{2, 3, 4, 2},
	0x6e: [4]uint8{2, 3, 4, 3},
	0x6f: [4]uint8{2, 3, 4, 4},
	0x70: [4]uint8{2, 4, 1, 1},
	0x71: [4]uint8{2, 4, 1, 2},
	0x72: [4]uint8{2, 4, 1, 3},
	0x73: [4]uint8{2, 4, 1, 4},
	0x74: [4]uint8{2, 4, 2, 1},
	0x75: [4]uint8{2, 4, 2, 2},
	0x76: [4]uint8{2, 4, 2, 3},
	0x77: [4]uint8{2, 4, 2, 4},
	0x78: [4]uint8{2, 4, 3, 1},
	0x79: [4]uint8{2, 4, 3, 2},
	0x7a: [4]uint8{2, 4, 3, 3},
	0x7b: [4]uint8{2, 4, 3, 4},
	0x7c: [4]uint8{2, 4, 4, 1},
	0x7d: [4]uint8{2, 4, 4, 2},
	0x7e: [4]uint8{2, 4, 4, 3},
	0x7f: [4]uint8{2, 4, 4, 4},
	0x80: [4]uint8{3, 1, 1, 1},
	0x81: [4]uint8{3, 1, 1, 2},
	0x82: [4]uint8{3, 1, 1, 3},
	0x83: [4]uint8{3, 1, 1, 4},
	0x84: [4]uint8{3, 1, 2, 1},
	0x85: [4]uint8{3, 1, 2, 2},
	0x86: [4]uint8{3, 1, 2, 3},
	0x87: [4]uint8{3, 1, 2, 4},
	0x88: [4]uint8{3, 1, 3, 1},
	0x89: [4]uint8{3, 1, 3, 2},
	0x8a: [4]uint8{3, 1, 3, 3},
	0x8b: [4]uint8{3, 1, 3, 4},
	0x8c: [4]uint8{3, 1, 4, 1},
	0x8d: [4]uint8{3, 1, 4, 2},
	0x8e: [4]uint8{3, 1, 4, 3},
	0x8f: [4]uint8{3, 1, 4, 4},
	0x90: [4]uint8{3, 2, 1, 1},
	0x91: [4]uint8{3, 2, 1, 2},
	0x92: [4]uint8{3, 2, 1, 3},
	0x93: [4]uint8{3, 2, 1, 4},
	0x94: [4]uint8{3, 2, 2, 1},
	0x95: [4]uint8{3, 2, 2, 2},
	0x96: [4]uint8{3, 2, 2, 3},
	0x97: [4]uint8{3, 2, 2, 4},
	0x98: [4]uint8{3, 2, 3, 1},
	0x99: [4]uint8{3, 2, 3, 2},
	0x9a: [4]uint8{3, 2, 3, 3},
	0x9b: [4]uint8{3, 2, 3, 4},
	0x9c: [4]uint8{3, 2, 4, 1},
	0x9d: [4]uint8{3, 2, 4, 2},
	0x9e: [4]uint8{3, 2, 4, 3},
	0x9f: [4]uint8{3, 2, 4, 4},
	0xa0: [4]uint8{3, 3, 1, 1},
	0xa1: [4]uint8{3, 3, 1, 2},
	0xa2: [4]uint8{3, 3, 1, 3},
	0xa3: [4]uint8{3, 3, 1, 4},
	0xa4: [4]uint8{3, 3, 2, 1},
	0xa5: [4]uint8{3, 3, 2, 2},
	0xa6: [4]uint8{3, 3, 2, 3},
	0xa7: [4]uint8{3, 3, 2, 4},
	0xa8: [4]uint8{3, 3, 3, 1},
	0xa9: [4]uint8{3, 3, 3, 2},
	0xaa: [4]uint8{3, 3, 3, 3},
	0xab: [4]uint8{3, 3, 3, 4},
	0xac: [4]uint8{3, 3, 4, 1},
	0xad: [4]uint8{3, 3, 4, 2},
	0xae: [4]uint8{3, 3, 4, 3},
	0xaf: [4]uint8{3, 3, 4, 4},
	0xb0: [4]uint8{3, 4, 1, 1},
	0xb1: [4]uint8{3, 4, 1, 2},
	0xb2: [4]uint8{3, 4, 1, 3},
	0xb3: [4]uint8{3, 4, 1, 4},
	0xb4: [4]uint8{3, 4, 2, 1},
	0xb5: [4]uint8{3, 4, 2, 2},
	0xb6: [4]uint8{3, 4, 2, 3},
	0xb7: [4]uint8{3, 4, 2, 4},
	0xb8: [4]uint8{3, 4, 3, 1},
	0xb9: [4]uint8{3, 4, 3, 2},
	0xba: [4]uint8{3, 4, 3, 3},
	0xbb: [4]uint8{3, 4, 3, 4},
	0xbc: [4]uint8{3, 4, 4, 1},
	0xbd: [4]uint8{3, 4, 4, 2},
	0xbe: [4]uint8{3, 4, 4, 3},
	0xbf: [4]uint8{3, 4, 4, 4},
	0xc0: [4]uint8{4, 1, 1, 1},
	0xc1: [4]uint8{4, 1, 1, 2},
	0xc2: [4]uint8{4, 1, 1, 3},
	0xc3: [4]uint8{4, 1, 1, 4},
	0xc4: [4]uint8{4, 1, 2, 1},
	0xc5: [4]uint8{4, 1, 2, 2},
	0xc6: [4]uint8{4, 1, 2, 3},
	0xc7: [4]uint8{4, 1, 2, 4},
	0xc8: [4]uint8{4, 1, 3, 1},
	0xc9: [4]uint8{4, 1, 3, 2},
	0xca: [4]uint8{4, 1, 3, 3},
	0xcb: [4]uint8{4, 1, 3, 4},
	0xcc: [4]uint8{4, 1, 4, 1},
	0xcd: [4]uint8{4, 1, 4, 2},
	0xce: [4]uint8{4, 1, 4, 3},
	0xcf: [4]uint8{4, 1, 4, 4},
	0xd0: [4]uint8{4, 2, 1, 1},
	0xd1: [4]uint8{4, 2, 1, 2},
	0xd2: [4]uint8{4, 2, 1, 3},
	0xd3: [4]uint8{4, 2, 1, 4},
	0xd4: [4]uint8{4, 2, 2, 1},
	0xd5: [4]uint8{4, 2, 2, 2},
	0xd6: [4]uint8{4, 2, 2, 3},
	0xd7: [4]uint8{4, 2, 2, 4},
	0xd8: [4]uint8{4, 2, 3, 1},
	0xd9: [4]uint8{4, 2, 3, 2},
	0xda: [4]uint8{4, 2, 3, 3},
	0xdb: [4]uint8{4, 2, 3, 4},
	0xdc: [4]uint8{4, 2, 4, 1},
	0xdd: [4]uint8{4, 2, 4, 2},
	0xde: [4]uint8{4, 2, 4, 3},
	0xdf: [4]uint8{4, 2, 4, 4},
	0xe0: [4]uint8{4, 3, 1, 1},
	0xe1: [4]uint8{4, 3, 1, 2},
	0xe2: [4]uint8{4, 3, 1, 3},
	0xe3: [4]uint8{4, 3, 1, 4},
	0xe4: [4]uint8{4, 3, 2, 1},
	0xe5: [4]uint8{4, 3, 2, 2},
	0xe6: [4]uint8{4, 3, 2, 3},
	0xe7: [4]uint8{4, 3, 2, 4},
	0xe8: [4]uint8{4, 3, 3, 1},
	0xe9: [4]uint8{4, 3, 3, 2},
	0xea: [4]uint8{4, 3, 3, 3},
	0xeb: [4]uint8{4, 3, 3, 4},
	0xec: [4]uint8{4, 3, 4, 1},
	0xed: [4]uint8{4, 3, 4, 2},
	0xee: [4]uint8{4, 3, 4, 3},
	0xef: [4]uint8{4, 3, 4, 4},
	0xf0: [4]uint8{4, 4, 1, 1},
	0xf1: [4]uint8{4, 4, 1, 2},
	0xf2: [4]uint8{4, 4, 1, 3},
	0xf3: [4]uint8{4, 4, 1, 4},
	0xf4: [4]uint8{4, 4, 2, 1},
	0xf5: [4]uint8{4, 4, 2, 2},
	0xf6: [4]uint8{4, 4, 2, 3},
	0xf7: [4]uint8{4, 4, 2, 4},
	0xf8: [4]uint8{4, 4, 3, 1},
	0xf9: [4]uint8{4, 4, 3, 2},
	0xfa: [4]uint8{4, 4, 3, 3},
	0xfb: [4]uint8{4, 4, 3, 4},
	0xfc: [4]uint8{4, 4, 4, 1},
	0xfd: [4]uint8{4, 4, 4, 2},
	0xfe: [4]uint8{4, 4, 4, 3},
	0xff: [4]uint8{4, 4, 4, 4},
}
