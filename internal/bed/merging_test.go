package bed

import (
	"fmt"
	"testing"

	"github.com/go-test/deep"
)

var testMergeChrOnly = []Line{
	{
		Chr: "1", Start: 1, Stop: 4,
		Full: []string{"1", "1", "4", "1", "A"},
	},
	{
		Chr: "1", Start: 5, Stop: 8,
		Full: []string{"1", "5", "8", "1", "A"},
	},
	{
		Chr: "1", Start: 6, Stop: 8,
		Full: []string{"1", "6", "8", "1", "A"},
	},
	{
		Chr: "1", Start: 5, Stop: 8,
		Full: []string{"1", "5", "8", "-1", "A"},
	},
	{
		Chr: "2", Start: 6, Stop: 8,
		Full: []string{"2", "6", "8", "1", "A"},
	},
	{
		Chr: "1", Start: 5, Stop: 8,
		Full: []string{"1", "5", "8", "1", "B"},
	},
	{
		Chr: "1", Start: 20, Stop: 30,
		Full: []string{"1", "20", "30", "1", "A"},
	},
}

var testMergeChrStrand = []Line{
	{
		Chr: "1", Start: 1, Stop: 4,
		Strand: "1",
		Full:   []string{"1", "1", "4", "1", "A"},
	},
	{
		Chr: "1", Start: 5, Stop: 8,
		Strand: "1",
		Full:   []string{"1", "5", "8", "1", "A"},
	},
	{
		Chr: "1", Start: 6, Stop: 8,
		Strand: "1",
		Full:   []string{"1", "6", "8", "1", "A"},
	},
	{
		Chr: "1", Start: 5, Stop: 8,
		Strand: "-1",
		Full:   []string{"1", "5", "8", "-1", "A"},
	},
	{
		Chr: "2", Start: 6, Stop: 8,
		Strand: "1",
		Full:   []string{"2", "6", "8", "1", "A"},
	},
	{
		Chr: "1", Start: 5, Stop: 8,
		Strand: "1",
		Full:   []string{"1", "5", "8", "1", "B"},
	},
	{
		Chr: "1", Start: 20, Stop: 30,
		Strand: "1",
		Full:   []string{"1", "20", "30", "1", "A"},
	},
}

var testMergeChrFeat = []Line{

	{
		Chr: "1", Start: 1, Stop: 4,
		Feat: "A",
		Full: []string{"1", "1", "4", "1", "A"},
	},
	{
		Chr: "1", Start: 5, Stop: 8,
		Feat: "A",
		Full: []string{"1", "5", "8", "1", "A"},
	},
	{
		Chr: "1", Start: 6, Stop: 8,
		Feat: "A",
		Full: []string{"1", "6", "8", "1", "A"},
	},
	{
		Chr: "1", Start: 5, Stop: 8,
		Feat: "A",
		Full: []string{"1", "5", "8", "-1", "A"},
	},
	{
		Chr: "2", Start: 6, Stop: 8,
		Feat: "A",
		Full: []string{"2", "6", "8", "1", "A"},
	},
	{
		Chr: "1", Start: 5, Stop: 8,
		Feat: "B",
		Full: []string{"1", "5", "8", "1", "B"},
	},
	{
		Chr: "1", Start: 20, Stop: 30,
		Feat: "A",
		Full: []string{"1", "20", "30", "1", "A"},
	},
}

var testMergeFull = []Line{
	{
		Chr: "1", Start: 1, Stop: 4,
		Strand: "1", Feat: "A",
		Full: []string{"1", "1", "4", "1", "A"},
	},
	{
		Chr: "1", Start: 5, Stop: 8,
		Strand: "1", Feat: "A",
		Full: []string{"1", "5", "8", "1", "A"},
	},
	{
		Chr: "1", Start: 6, Stop: 8,
		Strand: "1", Feat: "A",
		Full: []string{"1", "6", "8", "1", "A"},
	},
	{
		Chr: "1", Start: 5, Stop: 8,
		Strand: "-1", Feat: "A",
		Full: []string{"1", "5", "8", "-1", "A"},
	},
	{
		Chr: "2", Start: 6, Stop: 8,
		Strand: "1", Feat: "A",
		Full: []string{"2", "6", "8", "1", "A"},
	},
	{
		Chr: "1", Start: 5, Stop: 8,
		Strand: "1", Feat: "B",
		Full: []string{"1", "5", "8", "1", "B"},
	},
	{
		Chr: "1", Start: 20, Stop: 30,
		Strand: "1", Feat: "A",
		Full: []string{"1", "20", "30", "1", "A"},
	},
}

func TestMergeAndPadLines(t *testing.T) {
	t.Parallel()
	type testCase struct {
		testing     string
		bed         Bedfile
		expectedBed Bedfile
		shouldFail  bool
	}
	testCases := []testCase{
		{
			testing: "testMergeChrOnly",
			bed: Bedfile{
				Lines: deepCopyLines(testMergeChrOnly),
			},
			expectedBed: Bedfile{
				Lines: []Line{
					{
						Chr: "1", Start: 1, Stop: 8,
						Full: []string{"1", "1", "8", "1,-1", "A,B"},
					},
					{
						Chr: "1", Start: 20, Stop: 30,
						Full: []string{"1", "20", "30", "1", "A"},
					},
					{
						Chr: "2", Start: 6, Stop: 8,
						Full: []string{"2", "6", "8", "1", "A"},
					},
				},
			},
		},
		{
			testing: "testMergeChrOnly, overlap -1",
			bed: Bedfile{
				Overlap: -1,
				Lines:   deepCopyLines(testMergeChrOnly),
			},
			expectedBed: Bedfile{
				Overlap: -1,
				Lines: []Line{
					{
						Chr: "1", Start: 1, Stop: 4,
						Full: []string{"1", "1", "4", "1", "A"},
					},
					{
						Chr: "1", Start: 5, Stop: 8,
						Full: []string{"1", "5", "8", "1,-1", "A,B"},
					},
					{
						Chr: "1", Start: 20, Stop: 30,
						Full: []string{"1", "20", "30", "1", "A"},
					},
					{
						Chr: "2", Start: 6, Stop: 8,
						Full: []string{"2", "6", "8", "1", "A"},
					},
				},
			},
		},
		{
			testing: "testMergeChrOnly, overlap 10",
			bed: Bedfile{
				Overlap: 10,
				Lines:   deepCopyLines(testMergeChrOnly),
			},
			expectedBed: Bedfile{
				Overlap: 10,
				Lines: []Line{
					{
						Chr: "1", Start: 1, Stop: 8,
						Full: []string{"1", "1", "8", "1,-1", "A,B"},
					},
					{
						Chr: "1", Start: 20, Stop: 30,
						Full: []string{"1", "20", "30", "1", "A"},
					},
					{
						Chr: "2", Start: 6, Stop: 8,
						Full: []string{"2", "6", "8", "1", "A"},
					},
				},
			},
		},
		{
			testing: "testMergeChrOnly, overlap 11",
			bed: Bedfile{
				Overlap: 11,
				Lines:   deepCopyLines(testMergeChrOnly),
			},
			expectedBed: Bedfile{
				Overlap: 11,
				Lines: []Line{
					{
						Chr: "1", Start: 1, Stop: 30,
						Full: []string{"1", "1", "30", "1,-1", "A,B"},
					},
					{
						Chr: "2", Start: 6, Stop: 8,
						Full: []string{"2", "6", "8", "1", "A"},
					},
				},
			},
		},
		{
			testing: "testMergeChrStrand",
			bed: Bedfile{
				StrandCol: 4 - 1,
				Lines:     deepCopyLines(testMergeChrStrand),
			},
			expectedBed: Bedfile{
				StrandCol: 4 - 1,
				Lines: []Line{
					{
						Chr: "1", Start: 5, Stop: 8,
						Strand: "-1",
						Full:   []string{"1", "5", "8", "-1", "A"},
					},
					{
						Chr: "1", Start: 1, Stop: 8,
						Strand: "1",
						Full:   []string{"1", "1", "8", "1", "A,B"},
					},
					{
						Chr: "1", Start: 20, Stop: 30,
						Strand: "1",
						Full:   []string{"1", "20", "30", "1", "A"},
					},
					{
						Chr: "2", Start: 6, Stop: 8,
						Strand: "1",
						Full:   []string{"2", "6", "8", "1", "A"},
					},
				},
			},
		},
		{
			testing: "testMergeChrFeat",
			bed: Bedfile{
				FeatCol: 5 - 1,
				Lines:   deepCopyLines(testMergeChrFeat),
			},
			expectedBed: Bedfile{
				FeatCol: 5 - 1,
				Lines: []Line{
					{
						Chr: "1", Start: 1, Stop: 8,
						Feat: "A",
						Full: []string{"1", "1", "8", "1,-1", "A"},
					},
					{
						Chr: "1", Start: 20, Stop: 30,
						Feat: "A",
						Full: []string{"1", "20", "30", "1", "A"},
					},
					{
						Chr: "2", Start: 6, Stop: 8,
						Feat: "A",
						Full: []string{"2", "6", "8", "1", "A"},
					},
					{
						Chr: "1", Start: 5, Stop: 8,
						Feat: "B",
						Full: []string{"1", "5", "8", "1", "B"},
					},
				},
			},
		},
		{
			testing: "testMergeFull",
			bed: Bedfile{
				StrandCol: 4 - 1,
				FeatCol:   5 - 1,
				Lines:     deepCopyLines(testMergeFull),
			},
			expectedBed: Bedfile{
				StrandCol: 4 - 1,
				FeatCol:   5 - 1,
				Lines: []Line{
					{
						Chr: "1", Start: 5, Stop: 8,
						Strand: "-1", Feat: "A",
						Full: []string{"1", "5", "8", "-1", "A"},
					},
					{
						Chr: "1", Start: 1, Stop: 8,
						Strand: "1", Feat: "A",
						Full: []string{"1", "1", "8", "1", "A"},
					},
					{
						Chr: "1", Start: 20, Stop: 30,
						Strand: "1", Feat: "A",
						Full: []string{"1", "20", "30", "1", "A"},
					},
					{
						Chr: "2", Start: 6, Stop: 8,
						Strand: "1", Feat: "A",
						Full: []string{"2", "6", "8", "1", "A"},
					},
					{
						Chr: "1", Start: 5, Stop: 8,
						Strand: "1", Feat: "B",
						Full: []string{"1", "5", "8", "1", "B"},
					},
				},
			},
		},
		{
			testing: "testMergeChrOnly, padding = 10, paddingType = safe, chr in chrLengthMap",
			bed: Bedfile{
				PaddingType:  SafePT,
				Padding:      10,
				FirstBase:    1,
				Lines:        deepCopyLines(testMergeChrOnly),
				chrLengthMap: testChrLengthMap,
			},
			expectedBed: Bedfile{
				PaddingType:  SafePT,
				Padding:      10,
				FirstBase:    1,
				chrLengthMap: testChrLengthMap,
				Lines: []Line{
					{
						Chr: "1", Start: 1, Stop: 40,
						Full: []string{"1", "1", "40", "1,-1", "A,B"},
					},
					{
						Chr: "2", Start: 1, Stop: 18,
						Full: []string{"2", "1", "18", "1", "A"},
					},
				},
			},
		},
		{
			testing: "testMergeChrOnly, padding = 10, paddingType = lax, chr in chrLengthMap",
			bed: Bedfile{
				PaddingType:  LaxPT,
				Padding:      10,
				FirstBase:    1,
				Lines:        deepCopyLines(testMergeChrOnly),
				chrLengthMap: testChrLengthMap,
			},
			expectedBed: Bedfile{
				PaddingType:  LaxPT,
				Padding:      10,
				FirstBase:    1,
				chrLengthMap: testChrLengthMap,
				Lines: []Line{
					{
						Chr: "1", Start: 1, Stop: 40,
						Full: []string{"1", "1", "40", "1,-1", "A,B"},
					},
					{
						Chr: "2", Start: 1, Stop: 18,
						Full: []string{"2", "1", "18", "1", "A"},
					},
				},
			},
		},
		{
			testing: "testMergeChrOnly, padding = 10, paddingType = force, chr in chrLengthMap",
			bed: Bedfile{
				PaddingType:  ForcePT,
				Padding:      10,
				FirstBase:    1,
				Lines:        deepCopyLines(testMergeChrOnly),
				chrLengthMap: testChrLengthMap,
			},
			expectedBed: Bedfile{
				PaddingType:  ForcePT,
				Padding:      10,
				FirstBase:    1,
				chrLengthMap: testChrLengthMap,
				Lines: []Line{
					{
						Chr: "1", Start: 1, Stop: 40,
						Full: []string{"1", "1", "40", "1,-1", "A,B"},
					},
					{
						Chr: "2", Start: 1, Stop: 18,
						Full: []string{"2", "1", "18", "1", "A"},
					},
				},
			},
		},
		{
			testing: "testMergeChrOnly, padding = 10, paddingType = safe, chr not in chrLengthMap",
			bed: Bedfile{
				PaddingType: SafePT,
				Padding:     10,
				FirstBase:   1,
				Lines:       deepCopyLines(testMergeChrOnly),
			},
			shouldFail: true,
		},
		{
			testing: "testMergeChrOnly, padding = 10, paddingType = lax, chr not in chrLengthMap",
			bed: Bedfile{
				PaddingType: LaxPT,
				Padding:     10,
				FirstBase:   1,
				Lines:       deepCopyLines(testMergeChrOnly),
			},
			expectedBed: Bedfile{
				PaddingType: LaxPT,
				Padding:     10,
				FirstBase:   1,
				Lines: []Line{
					{
						Chr: "1", Start: 1, Stop: 8,
						Full: []string{"1", "1", "8", "1,-1", "A,B"},
					},
					{
						Chr: "1", Start: 20, Stop: 30,
						Full: []string{"1", "20", "30", "1", "A"},
					},
					{
						Chr: "2", Start: 6, Stop: 8,
						Full: []string{"2", "6", "8", "1", "A"},
					},
				},
			},
		},
		{
			testing: "testMergeChrOnly, padding = 10, paddingType = force, chr not in chrLengthMap",
			bed: Bedfile{
				PaddingType: ForcePT,
				Padding:     10,
				FirstBase:   1,
				Lines:       deepCopyLines(testMergeChrOnly),
			},
			expectedBed: Bedfile{
				PaddingType: ForcePT,
				Padding:     10,
				FirstBase:   1,
				Lines: []Line{
					{
						Chr: "1", Start: 1, Stop: 40,
						Full: []string{"1", "1", "40", "1,-1", "A,B"},
					},
					{
						Chr: "2", Start: 1, Stop: 18,
						Full: []string{"2", "1", "18", "1", "A"},
					},
				},
			},
		},
		{
			testing: "padding=5 && overlap=-1",
			bed: Bedfile{
				PaddingType:  SafePT,
				Padding:      5,
				FirstBase:    1,
				Overlap:      -1,
				chrLengthMap: testChrLengthMap,
				Lines: []Line{
					{
						Chr: "1", Start: 1, Stop: 4,
						Full: []string{"1", "1", "4"},
					},
					{
						Chr: "1", Start: 5, Stop: 9,
						Full: []string{"1", "5", "9"},
					},
					{
						Chr: "1", Start: 20, Stop: 30,
						Full: []string{"1", "20", "30"},
					},
				},
			},
			expectedBed: Bedfile{
				PaddingType:  SafePT,
				Padding:      5,
				FirstBase:    1,
				Overlap:      -1,
				chrLengthMap: testChrLengthMap,
				Lines: []Line{
					{
						Chr: "1", Start: 1, Stop: 14,
						Full: []string{"1", "1", "14"},
					},
					{
						Chr: "1", Start: 15, Stop: 35,
						Full: []string{"1", "15", "35"},
					},
				},
			},
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.testing, func(t *testing.T) {
			t.Parallel()
			err := tc.bed.MergeAndPadLines()
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

func TestStringInSlice(t *testing.T) {
	t.Parallel()
	type testCase struct {
		testing        string
		slice          []string
		item           string
		expectedResult bool
	}
	testCases := []testCase{
		{
			testing:        "not in slice",
			slice:          []string{"10", "11", "1000"},
			item:           "1",
			expectedResult: false,
		},
		{
			testing:        "in slice",
			slice:          []string{"10", "11", "1000"},
			item:           "11",
			expectedResult: true,
		},
	}
	for _, tc := range testCases {
		tc := tc
		description := fmt.Sprintf("%s in %v", tc.item, tc.slice)
		t.Run(description, func(t *testing.T) {
			t.Parallel()
			result := stringInSlice(tc.slice, tc.item)
			if tc.expectedResult != result {
				t.Errorf("expected %t got %t", tc.expectedResult, result)
			}
		})
	}
}
