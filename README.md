# Quick HCT

`Quick HCT` uses a precomputed tonal map to quickly approximate transformations of RGB colors to a given tonal value. Tone is one of the dimensions of the [HCT color space](https://m3.material.io/styles/color/the-color-system/key-colors-tones) introduced by [Material Design 3](https://m3.material.io). A tone value of 100 is equivalent to the idea of light at its maximum and results in white. Every tone value between 0 and 100 expresses the amount of light present in the color.

Example tonal range for RGB color <span style="color:#426915">#426915</span>:

<div style="display:flex; line-height: 2.5;">
<span style="width:4ch; background:#000000; color: white; text-align: center;">0</span>
<span style="width:4ch; background:#0e2000; color: white; text-align: center;">10</span>
<span style="width:4ch; background:#1b3700; color: white; text-align: center;">20</span>
<span style="width:4ch; background:#2a5000; color: white; text-align: center;">30</span>
<span style="width:4ch; background:#416919; color: white; text-align: center;">40</span>
<span style="width:4ch; background:#598330; color: white; text-align: center;">50</span>
<span style="width:4ch; background:#719d47; color: black; text-align: center;">60</span>
<span style="width:4ch; background:#8bb85f; color: black; text-align: center;">70</span>
<span style="width:4ch; background:#a5d477; color: black; text-align: center;">80</span>
<span style="width:4ch; background:#c0f191; color: black; text-align: center;">90</span>
<span style="width:4ch; background:#d1ffa2; color: black; text-align: center;">95</span>
<span style="width:4ch; background:#eeffd8; color: black; text-align: center;">98</span>
<span style="width:4ch; background:#ffffff; color: black; text-align: center;">100</span>
</div>
<p>

Example tonal range for RGB color <span style="color:#715573">#715573</span>:

<div style="display:flex; line-height: 2.5;">
<span style="width:4ch; background:#000000; color: white; text-align: center;">0</span>
<span style="width:4ch; background:#271529; color: white; text-align: center;">10</span>
<span style="width:4ch; background:#3e293f; color: white; text-align: center;">20</span>
<span style="width:4ch; background:#563f56; color: white; text-align: center;">30</span>
<span style="width:4ch; background:#6f576f; color: white; text-align: center;">40</span>
<span style="width:4ch; background:#886f88; color: white; text-align: center;">50</span>
<span style="width:4ch; background:#a388a2; color: black; text-align: center;">60</span>
<span style="width:4ch; background:#bfa2bd; color: black; text-align: center;">70</span>
<span style="width:4ch; background:#dbbdd9; color: black; text-align: center;">80</span>
<span style="width:4ch; background:#f8d9f6; color: black; text-align: center;">90</span>
<span style="width:4ch; background:#ffebfb; color: black; text-align: center;">95</span>
<span style="width:4ch; background:#fff7fe; color: black; text-align: center;">98</span>
<span style="width:4ch; background:#ffffff; color: black; text-align: center;">100</span>
</div>
<p>

Usage:

```go
r, g, b := quickhct.WithTone(0x42, 0x69, 0x15, 40)
```

`Quick HCT` is released by `Microbus LLC` under the [Apache 2.0 license](http://www.apache.org/licenses/LICENSE-2.0).
