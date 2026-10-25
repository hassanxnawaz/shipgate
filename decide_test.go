package main

import "testing"

func TestFixtures(t *testing.T) {
	cases := []struct {
		name   string
		row    Row
		want   string
		fp     string
	}{
		{"qwen_rerun", Row{11, -7, 30, true, false, false}, "RERUN", "ddd8c80582b8021d5abe3449732197753368c78cb19275e04192f0ee1157ee29"},
		{"clean_ship", Row{12, 2, 22, false, false, true}, "SHIP", "8b6de9a54db8847def7d502627c5c862bd2c807e9c81a6fd35056dc5fc504bca"},
		{"gain_but_guardrail", Row{12, 2, 22, true, false, true}, "CHANGE", "13fe8c9521a483afd7a44214ea53c22144a7b0ee6e2c3bb0fc62f43c322d19c4"},
		{"real_loss", Row{-8, -14, -2, false, true, false}, "STOP", "ef19fb28ecc72c37e9d296b41b92776c51ad13e39baae1e704c6c269d3057f98"},
		{"flat", Row{-1, -6, 4, false, false, false}, "STOP", "f15db289e44c90332827719aca6af9e01a85c6e8ef545cbadf5a977d1554fa50"},
		{"held_and_better", Row{2, -3, 7, false, false, true}, "SHIP", "2ee6464ccc3ebbdecd7bebcc0b35f8dfdb19656361654923fa4cf9cd10d03fbe"},
		{"small_gain_guardrail", Row{2, -3, 7, true, false, true}, "RERUN", "38575bd287988fc0f16b5e2848b542c1c464b4beb7b52087386ac867ab5f10cf"},
	}
	for _, c := range cases {
		if got := Decide(c.row); got != c.want {
			t.Errorf("%s decision %s, want %s", c.name, got, c.want)
		}
		if got := Fingerprint(c.row); got != c.fp {
			t.Errorf("%s fingerprint %s, want %s", c.name, got, c.fp)
		}
	}
}
