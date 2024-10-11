package main

import (
	. "github.com/wii-tools/powerpc"
)

var NegateECTitle = PatchSet{
	Name: "Negate EC Title Check",

	Patches: []Patch{
		{
			Name:     "Permit downloading all titles",
			AtOffset: 619648,

			// Generic function prolog
			Before: Instructions{
				STWU(R1, R1, 0xffe0),
				MFSPR(R0, LR),
			}.Bytes(),

			// Immediately return true
			After: Instructions{
				LI(R3, 1),
				BLR(),
			}.Bytes(),
		},
		{
			Name:     "Nullify ec::removeAllTitles",
			AtOffset: 588368,
			Before: Instructions{
				STWU(R1, R1, 0xffc0),
				MFSPR(R0, LR),
			}.Bytes(),
			After: Instructions{
				LI(R3, 0),
				BLR(),
			}.Bytes(),
		},
	},
}
