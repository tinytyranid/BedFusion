package bed

import (
	"testing"

	"github.com/go-test/deep"
)

var testLinesToPad = []Line{
	{
		Chr: "1", Start: 50, Stop: 51,
		Full: []string{"1", "50", "51"},
	},
	{
		Chr: "2", Start: 150, Stop: 151,
		Full: []string{"2", "150", "151"},
	},
	{
		Chr: "3", Start: 250, Stop: 251,
		Full: []string{"3", "250", "251"},
	},
	{
		Chr: "4", Start: 350, Stop: 351,
		Full: []string{"4", "350", "351"},
	},
}

var testChrLengthMap = map[string]int{
	"1": 100,
	"2": 200,
	"3": 300,
	"4": 400,
}

func TestPadLines(t *testing.T) {
	t.Parallel()
	type testCase struct {
		testing     string
		bed         Bedfile
		expectedBed Bedfile
		shouldFail  bool
	}
	testCases := []testCase{
		{
			testing: "padding within chromosome, no missing, paddingType=safe",
			bed: Bedfile{
				PaddingType:  SafePT,
				Padding:      10,
				FirstBase:    1,
				Lines:        deepCopyLines(testLinesToPad),
				chrLengthMap: testChrLengthMap,
			},
			expectedBed: Bedfile{
				PaddingType: SafePT,
				Padding:     10,
				FirstBase:   1,
				Lines: []Line{
					{
						Chr: "1", Start: 40, Stop: 61,
						Full: []string{"1", "40", "61"},
					},
					{
						Chr: "2", Start: 140, Stop: 161,
						Full: []string{"2", "140", "161"},
					},
					{
						Chr: "3", Start: 240, Stop: 261,
						Full: []string{"3", "240", "261"},
					},
					{
						Chr: "4", Start: 340, Stop: 361,
						Full: []string{"4", "340", "361"},
					},
				},
				chrLengthMap: testChrLengthMap,
			},
		},
		{
			testing: "padding within chromosome, no missing, paddingType=lax",
			bed: Bedfile{
				PaddingType:  LaxPT,
				Padding:      10,
				FirstBase:    1,
				Lines:        deepCopyLines(testLinesToPad),
				chrLengthMap: testChrLengthMap,
			},
			expectedBed: Bedfile{
				PaddingType: LaxPT,
				Padding:     10,
				FirstBase:   1,
				Lines: []Line{
					{
						Chr: "1", Start: 40, Stop: 61,
						Full: []string{"1", "40", "61"},
					},
					{
						Chr: "2", Start: 140, Stop: 161,
						Full: []string{"2", "140", "161"},
					},
					{
						Chr: "3", Start: 240, Stop: 261,
						Full: []string{"3", "240", "261"},
					},
					{
						Chr: "4", Start: 340, Stop: 361,
						Full: []string{"4", "340", "361"},
					},
				},
				chrLengthMap: testChrLengthMap,
			},
		},
		{
			testing: "padding within chromosome, no missing, paddingType=force",
			bed: Bedfile{
				PaddingType:  ForcePT,
				Padding:      10,
				FirstBase:    1,
				Lines:        deepCopyLines(testLinesToPad),
				chrLengthMap: testChrLengthMap,
			},
			expectedBed: Bedfile{
				PaddingType: ForcePT,
				Padding:     10,
				FirstBase:   1,
				Lines: []Line{
					{
						Chr: "1", Start: 40, Stop: 61,
						Full: []string{"1", "40", "61"},
					},
					{
						Chr: "2", Start: 140, Stop: 161,
						Full: []string{"2", "140", "161"},
					},
					{
						Chr: "3", Start: 240, Stop: 261,
						Full: []string{"3", "240", "261"},
					},
					{
						Chr: "4", Start: 340, Stop: 361,
						Full: []string{"4", "340", "361"},
					},
				},
				chrLengthMap: testChrLengthMap,
			},
		},
		{
			testing: "padding beyond chromosome, no missing, paddingType=safe",
			bed: Bedfile{
				PaddingType:  SafePT,
				Padding:      1000,
				FirstBase:    1,
				Lines:        deepCopyLines(testLinesToPad),
				chrLengthMap: testChrLengthMap,
			},
			expectedBed: Bedfile{
				PaddingType: SafePT,
				Padding:     1000,
				FirstBase:   1,
				Lines: []Line{
					{
						Chr: "1", Start: 1, Stop: 100,
						Full: []string{"1", "1", "100"},
					},
					{
						Chr: "2", Start: 1, Stop: 200,
						Full: []string{"2", "1", "200"},
					},
					{
						Chr: "3", Start: 1, Stop: 300,
						Full: []string{"3", "1", "300"},
					},
					{
						Chr: "4", Start: 1, Stop: 400,
						Full: []string{"4", "1", "400"},
					},
				},
				chrLengthMap: testChrLengthMap,
			},
		},
		{
			testing: "padding beyond chromosome, no missing, paddingType=lax",
			bed: Bedfile{
				PaddingType:  LaxPT,
				Padding:      1000,
				FirstBase:    1,
				Lines:        deepCopyLines(testLinesToPad),
				chrLengthMap: testChrLengthMap,
			},
			expectedBed: Bedfile{
				PaddingType: LaxPT,
				Padding:     1000,
				FirstBase:   1,
				Lines: []Line{
					{
						Chr: "1", Start: 1, Stop: 100,
						Full: []string{"1", "1", "100"},
					},
					{
						Chr: "2", Start: 1, Stop: 200,
						Full: []string{"2", "1", "200"},
					},
					{
						Chr: "3", Start: 1, Stop: 300,
						Full: []string{"3", "1", "300"},
					},
					{
						Chr: "4", Start: 1, Stop: 400,
						Full: []string{"4", "1", "400"},
					},
				},
				chrLengthMap: testChrLengthMap,
			},
		},
		{
			testing: "padding beyond chromosome, no missing, paddingType=force",
			bed: Bedfile{
				PaddingType:  ForcePT,
				Padding:      1000,
				FirstBase:    1,
				Lines:        deepCopyLines(testLinesToPad),
				chrLengthMap: testChrLengthMap,
			},
			expectedBed: Bedfile{
				PaddingType: ForcePT,
				Padding:     1000,
				FirstBase:   1,
				Lines: []Line{
					{
						Chr: "1", Start: 1, Stop: 100,
						Full: []string{"1", "1", "100"},
					},
					{
						Chr: "2", Start: 1, Stop: 200,
						Full: []string{"2", "1", "200"},
					},
					{
						Chr: "3", Start: 1, Stop: 300,
						Full: []string{"3", "1", "300"},
					},
					{
						Chr: "4", Start: 1, Stop: 400,
						Full: []string{"4", "1", "400"},
					},
				},
				chrLengthMap: testChrLengthMap,
			},
		},
		{
			testing: "padding beyond chromosome, all missing, paddingType=safe",
			bed: Bedfile{
				PaddingType: SafePT,
				Padding:     1000,
				FirstBase:   1,
				Lines:       deepCopyLines(testLinesToPad),
			},
			shouldFail: true,
		},
		{
			testing: "padding beyond chromosome, all missing, paddingType=lax",
			bed: Bedfile{
				PaddingType: LaxPT,
				Padding:     1000,
				FirstBase:   1,
				Lines:       deepCopyLines(testLinesToPad),
			},
			expectedBed: Bedfile{
				PaddingType: LaxPT,
				Padding:     1000,
				FirstBase:   1,
				Lines:       deepCopyLines(testLinesToPad),
			},
		},
		{
			testing: "padding beyond chromosome, all missing, paddingType=force",
			bed: Bedfile{
				PaddingType: ForcePT,
				Padding:     1000,
				FirstBase:   1,
				Lines:       deepCopyLines(testLinesToPad),
			},
			expectedBed: Bedfile{
				PaddingType: ForcePT,
				Padding:     1000,
				FirstBase:   1,
				Lines: []Line{
					{
						Chr: "1", Start: 1, Stop: 1051,
						Full: []string{"1", "1", "1051"},
					},
					{
						Chr: "2", Start: 1, Stop: 1151,
						Full: []string{"2", "1", "1151"},
					},
					{
						Chr: "3", Start: 1, Stop: 1251,
						Full: []string{"3", "1", "1251"},
					},
					{
						Chr: "4", Start: 1, Stop: 1351,
						Full: []string{"4", "1", "1351"},
					},
				},
			},
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.testing, func(t *testing.T) {
			t.Parallel()
			err := tc.bed.PadLines()
			if (!tc.shouldFail && err != nil) || (tc.shouldFail && err == nil) {
				t.Fatalf("shouldFail is %t, but err is %q", tc.shouldFail, err)
			}
			if !tc.shouldFail {
				if diff := deep.Equal(tc.expectedBed, tc.bed); diff != nil {
					t.Error("expected VS received bed", diff)
				}
			}
		})
	}
}

func TestPadLineAccordingToPaddingType(t *testing.T) {
	t.Parallel()
	type testCase struct {
		testing            string
		bed                Bedfile
		line               Line
		missChrMap         []string
		expectedPaddedLine Line
		expectedMisschrMap []string
		shouldFail         bool
	}
	testCases := []testCase{
		{
			testing: "padding within chromosome, no missing, paddingType=safe",
			bed: Bedfile{
				PaddingType:  SafePT,
				Padding:      10,
				FirstBase:    1,
				chrLengthMap: testChrLengthMap,
			},
			line: deepCopyLine(testLinesToPad[1]),
			expectedPaddedLine: Line{
				Chr: "2", Start: 140, Stop: 161,
				Full: []string{"2", "140", "161"},
			},
		},
		{
			testing: "padding within chromosome, no missing, paddingType=lax",
			bed: Bedfile{
				PaddingType:  LaxPT,
				Padding:      10,
				FirstBase:    1,
				chrLengthMap: testChrLengthMap,
			},
			line: deepCopyLine(testLinesToPad[1]),
			expectedPaddedLine: Line{
				Chr: "2", Start: 140, Stop: 161,
				Full: []string{"2", "140", "161"},
			},
		},
		{
			testing: "padding within chromosome, no missing, paddingType=force",
			bed: Bedfile{
				PaddingType:  ForcePT,
				Padding:      10,
				FirstBase:    1,
				chrLengthMap: testChrLengthMap,
			},
			line: deepCopyLine(testLinesToPad[1]),
			expectedPaddedLine: Line{
				Chr: "2", Start: 140, Stop: 161,
				Full: []string{"2", "140", "161"},
			},
		},
		{
			testing: "padding beyond chromosome, no missing, paddingType=safe",
			bed: Bedfile{
				PaddingType:  SafePT,
				Padding:      1000,
				FirstBase:    1,
				chrLengthMap: testChrLengthMap,
			},
			line: deepCopyLine(testLinesToPad[1]),
			expectedPaddedLine: Line{
				Chr: "2", Start: 1, Stop: 200,
				Full: []string{"2", "1", "200"},
			},
		},
		{
			testing: "padding beyond chromosome, no missing, paddingType=lax",
			bed: Bedfile{
				PaddingType:  LaxPT,
				Padding:      1000,
				FirstBase:    1,
				chrLengthMap: testChrLengthMap,
			},
			line: deepCopyLine(testLinesToPad[1]),
			expectedPaddedLine: Line{
				Chr: "2", Start: 1, Stop: 200,
				Full: []string{"2", "1", "200"},
			},
		},
		{
			testing: "padding beyond chromosome, no missing, paddingType=force",
			bed: Bedfile{
				PaddingType:  ForcePT,
				Padding:      1000,
				FirstBase:    1,
				chrLengthMap: testChrLengthMap,
			},
			line: deepCopyLine(testLinesToPad[1]),
			expectedPaddedLine: Line{
				Chr: "2", Start: 1, Stop: 200,
				Full: []string{"2", "1", "200"},
			},
		},
		{
			testing: "padding beyond chromosome, all missing, paddingType=safe",
			bed: Bedfile{
				PaddingType: SafePT,
				Padding:     1000,
				FirstBase:   1,
			},
			line:       deepCopyLine(testLinesToPad[1]),
			shouldFail: true,
		},
		{
			testing: "padding beyond chromosome, all missing, paddingType=lax",
			bed: Bedfile{
				PaddingType: LaxPT,
				Padding:     1000,
				FirstBase:   1,
			},
			line:       deepCopyLine(testLinesToPad[1]),
			missChrMap: []string{"1"},
			expectedPaddedLine: Line{
				Chr: "2", Start: 150, Stop: 151,
				Full: []string{"2", "150", "151"},
			},
			expectedMisschrMap: []string{"1", "2"},
		},
		{
			testing: "padding beyond chromosome, all missing, paddingType=force",
			bed: Bedfile{
				PaddingType: ForcePT,
				Padding:     1000,
				FirstBase:   1,
			},
			line:       deepCopyLine(testLinesToPad[1]),
			missChrMap: []string{"1"},
			expectedPaddedLine: Line{
				Chr: "2", Start: 1, Stop: 1151,
				Full: []string{"2", "1", "1151"},
			},
			expectedMisschrMap: []string{"1", "2"},
		},
		{
			testing:    "padding type does not exist",
			bed:        Bedfile{PaddingType: "test"},
			shouldFail: true,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.testing, func(t *testing.T) {
			t.Parallel()
			paddedLines, chrNotInLengthMap, err := tc.bed.padAccordingToPaddingType(tc.line, tc.missChrMap)
			if (!tc.shouldFail && err != nil) || (tc.shouldFail && err == nil) {
				t.Fatalf("shouldFail is %t, but err is %q", tc.shouldFail, err)
			}
			if !tc.shouldFail {
				if diff := deep.Equal(tc.expectedPaddedLine, paddedLines); diff != nil {
					t.Error("expected VS received paddedLine", diff)
				}
				if diff := deep.Equal(tc.expectedMisschrMap, chrNotInLengthMap); diff != nil {
					t.Error("expected VS received chrNotInLengthMap", diff)
				}
			}
		})
	}
}

func TestPadLine(t *testing.T) {
	t.Parallel()
	type testCase struct {
		testing          string
		bed              Bedfile
		line             Line
		expectedLine     Line
		expectedChrInMap bool
		shouldFail       bool
	}
	testCases := []testCase{
		{
			testing: "padding within chromosome, first base 0",
			bed: Bedfile{
				Padding:   10,
				FirstBase: 0,
				chrLengthMap: map[string]int{
					"1": 100,
				},
			},
			line: Line{
				Chr: "1", Start: 50, Stop: 51,
				Full: []string{"1", "50", "51"},
			},
			expectedLine: Line{
				Chr: "1", Start: 40, Stop: 61,
				Full: []string{"1", "40", "61"},
			},
			expectedChrInMap: true,
		},
		{
			testing: "padding within chromosome, first base 1",
			bed: Bedfile{
				Padding:   10,
				FirstBase: 1,
				chrLengthMap: map[string]int{
					"1": 100,
				},
			},
			line: Line{
				Chr: "1", Start: 50, Stop: 51,
				Full: []string{"1", "50", "51"},
			},
			expectedLine: Line{
				Chr: "1", Start: 40, Stop: 61,
				Full: []string{"1", "40", "61"},
			},
			expectedChrInMap: true,
		},
		{
			testing: "padding beyond chromosome, first base 0",
			bed: Bedfile{
				Padding:   1000,
				FirstBase: 0,
				chrLengthMap: map[string]int{
					"1": 100,
				},
			},
			line: Line{
				Chr: "1", Start: 50, Stop: 51,
				Full: []string{"1", "50", "51"},
			},
			expectedLine: Line{
				Chr: "1", Start: 0, Stop: 100,
				Full: []string{"1", "0", "100"},
			},
			expectedChrInMap: true,
		},
		{
			testing: "padding beyond chromosome, first base 1",
			bed: Bedfile{
				Padding:   1000,
				FirstBase: 1,
				chrLengthMap: map[string]int{
					"1": 100,
				},
			},
			line: Line{
				Chr: "1", Start: 50, Stop: 51,
				Full: []string{"1", "50", "51"},
			},
			expectedLine: Line{
				Chr: "1", Start: 1, Stop: 100,
				Full: []string{"1", "1", "100"},
			},
			expectedChrInMap: true,
		},
		{
			testing: "chromosome not part of chrLengthMap, first base 0",
			bed: Bedfile{
				Padding:   10,
				FirstBase: 0,
				chrLengthMap: map[string]int{
					"2": 100,
				},
			},
			line: Line{
				Chr: "1", Start: 50, Stop: 51,
				Full: []string{"1", "50", "51"},
			},
			expectedLine: Line{
				Chr: "1", Start: 40, Stop: 61,
				Full: []string{"1", "40", "61"},
			},
			expectedChrInMap: false,
		},
		{
			testing: "chromosome not part of chrLengthMap, first base 1",
			bed: Bedfile{
				Padding:   10,
				FirstBase: 1,
				chrLengthMap: map[string]int{
					"2": 100,
				},
			},
			line: Line{
				Chr: "1", Start: 50, Stop: 51,
				Full: []string{"1", "50", "51"},
			},
			expectedLine: Line{
				Chr: "1", Start: 40, Stop: 61,
				Full: []string{"1", "40", "61"},
			},
			expectedChrInMap: false,
		},
		{
			testing: "negative padding, first base 1",
			bed: Bedfile{
				Padding:   -10,
				FirstBase: 1,
				chrLengthMap: map[string]int{
					"1": 100,
				},
			},
			line: Line{
				Chr: "1", Start: 40, Stop: 70,
				Full: []string{"1", "40", "70"},
			},
			expectedLine: Line{
				Chr: "1", Start: 50, Stop: 60,
				Full: []string{"1", "50", "60"},
			},
			expectedChrInMap: true,
		},
		{
			testing: "negative padding where region would be flipped",
			bed: Bedfile{
				Padding:   -100,
				FirstBase: 0,
				chrLengthMap: map[string]int{
					"1": 100,
				},
			},
			line: Line{
				Chr: "1", Start: 40, Stop: 70,
				Full: []string{"1", "40", "70"},
			},
			shouldFail: true,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.testing, func(t *testing.T) {
			t.Parallel()
			deepCopiedLine := deepCopyLine(tc.line)
			paddedLine, chrInMap, err := tc.bed.padLine(tc.line)
			if (!tc.shouldFail && err != nil) || (tc.shouldFail && err == nil) {
				t.Fatalf("shouldFail is %t, but err is %q", tc.shouldFail, err)
			}
			if !tc.shouldFail {
				if diff := deep.Equal(tc.expectedLine, paddedLine); diff != nil {
					t.Error("expected VS received line", diff)
				}
				if (!tc.expectedChrInMap && chrInMap) || (tc.expectedChrInMap && !chrInMap) {
					t.Fatalf("expectedChrInMap is %t, but chrInMap is %t", tc.expectedChrInMap, chrInMap)
				}
				if diff := deep.Equal(deepCopiedLine, tc.line); diff != nil {
					t.Error("deep copy test, expected VS received line", diff)
				}
			}
		})
	}
}
