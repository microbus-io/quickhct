/*
Copyright 2023 Microbus LLC and various contributors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package quickhct

import (
	"embed"
)

var (
	//go:embed tonedata.bin
	embeddedFS embed.FS
)

/*
WithTone takes in an RGB color and a Tone to produce the RGB color that
maps to the same Hue and Chroma of the original color but with the given tone.
The algorithm uses a precomputed tonal map with a granularity of 0x08 so results
are an approximation.

R, G and B must be in the range of 0 to 255. Tone must in the range of 0 to 100.
*/
func WithTone(r, g, b int, tone int) (rr int, gg int, bb int) {
	if tone <= 0 || tone%10 == 0 || tone == 95 || tone == 99 || tone >= 100 {
		rr, gg, bb = computeTone(r, g, b, tone)
		return rr, gg, bb
	}

	var lower int
	var upper int
	if tone <= 90 {
		lower = tone / 10 * 10
		upper = lower + 10
	} else if tone <= 95 {
		lower = 90
		upper = 95
	} else if tone <= 99 {
		lower = 95
		upper = 99
	}

	r1, g1, b1 := computeTone(r, g, b, lower)
	r2, g2, b2 := computeTone(r, g, b, upper)

	rr = (r2*(tone-lower) + r1*(upper-tone)) / (upper - lower)
	gg = (g2*(tone-lower) + g1*(upper-tone)) / (upper - lower)
	bb = (b2*(tone-lower) + b1*(upper-tone)) / (upper - lower)
	return rr, gg, bb
}

/*
computeTone returns an approximation of the HCT tonal level for a given color.
The algorithm uses a precomputed tonal map with a granularity of 0x08 so results
are an approximation.

R, G and B must be in the range of 0 to 255.
Tone must be 0, 10, 20, 30, 40, 50, 60, 70, 80, 90, 95, 99 or 100.
*/
func computeTone(r, g, b int, tone int) (rr int, gg int, bb int) {
	if tone <= 0 {
		return 0, 0, 0
	}
	if tone >= 100 {
		return 0xff, 0xff, 0xff
	}

	r = r & 0xff
	g = g & 0xff
	b = b & 0xff
	ri := (r + 0x4) / 0x8
	gi := (g + 0x4) / 0x8
	bi := (b + 0x4) / 0x8

	pos := ri*33*33 + gi*33 + bi
	pos *= 3 * 11
	if tone <= 90 && tone%10 == 0 {
		// 10, 20, ... 90
		pos += +3 * (tone/10 - 1)
	} else if tone == 95 {
		// 95
		pos += 3 * 9
	} else if tone == 99 {
		// 99
		pos += 3 * 10
	} else {
		return 0, 0, 0 // Invalid tone
	}
	bin, _ := embeddedFS.ReadFile("tonedata.bin")
	rgbResult := bin[pos : pos+3]
	return int(rgbResult[0]), int(rgbResult[1]), int(rgbResult[2])
}
