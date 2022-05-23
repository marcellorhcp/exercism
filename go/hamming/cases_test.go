package hamming

import "errors"

// Source: exercism/problem-specifications
// Commit: 4119671 Hamming: Add a tests to avoid wrong recursion solution (#1450)
// Problem Specifications Version: 2.3.0

var testCases = []struct {
	s1          string
	s2          string
	want        int
	expectError error
}{
	{ // empty strands
		"",
		"",
		0,
		nil,
	},
	{ // single letter identical strands
		"A",
		"A",
		0,
		nil,
	},
	{ // single letter different strands
		"G",
		"T",
		1,
		nil,
	},
	{ // long identical strands
		"GGACTGAAATCTG",
		"GGACTGAAATCTG",
		0,
		nil,
	},
	{ // long different strands
		"GGACGGATTCTG",
		"AGGACGGATTCT",
		9,
		nil,
	},
	{ // disallow first strand longer
		"AATG",
		"AAA",
		0,
		errors.New("a and b have different lengths"),
	},
	{ // disallow second strand longer
		"ATA",
		"AGTG",
		0,
		errors.New("a and b have different lengths"),
	},
	{ // disallow left empty strand
		"",
		"G",
		0,
		errors.New("a and b have different lengths"),
	},
	{ // disallow right empty strand
		"G",
		"",
		0,
		errors.New("a and b have different lengths"),
	},
}
