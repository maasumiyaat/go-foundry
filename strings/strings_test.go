package strings

import (
	"testing"
)

type testCases struct {
	arg  string
	want bool
}

func TestString(t *testing.T) {

	t.Run("Is Only Letters tests: ", func(t *testing.T) {
		cases := []testCases{
			{"Golang", true},
			{"Golang542", false},
			{"Golang542", false},
		}

		for _, tc := range cases {
			got := IsOnlyLetter(tc.arg)
			if tc.want != got {
				t.Errorf("Expected '%t', but got '%t'", tc.want, got)
			}
		}
	})

	t.Run("String Trimming Test:", func(t *testing.T) {
		got := TrimString()
		want := "Hello, World"

		if got != want {
			t.Errorf("Got : %s, Wanted : %s", got, want)
		}
	})

	t.Run("Is All Upper case and only Letter test :", func(t *testing.T) {
		cases := []testCases{
			{"UPPERCASE", true},
			{"LowerCase", false},
		}

		for _, tc := range cases {
			got := AllUpperCase(tc.arg)

			if tc.want != got {
				t.Errorf("Expected '%t', but got '%t'", tc.want, got)
			}
		}
	})

	t.Run("If word contains: ", func(t *testing.T) {

		type tests struct {
			word   string
			letter string
			result bool
		}

		cases := []tests{
			{"abcd", "b", true},
			{"masum", "m", true},
			{"mishqat", "m", true},
			{"vscode", "w", false},
			{"vscode", "w", true},
			{"vscode", "w", true},
			{"vscode", "se", true},
		}

		for _, tc := range cases {
			got := IfContains(tc.word, tc.letter)

			if tc.result != got {
				t.Errorf("Expected '%t', but got '%t'", tc.result, got)
			}
		}
	})

}

func TestDefangIPaddr(t *testing.T) {
	t.Run("1108. Defanging an IP Address: ", func(t *testing.T) {
		type defangTestCases struct {
			param  string
			result string
		}

		cases := []defangTestCases{
			{"1.1.1.1", "1[.]1[.]1[.]1"},
			{"255.100.50.0", "255[.]100[.]50[.]0"},
		}

		for _, tc := range cases {
			got := DefangIPaddr(tc.param)

			t.Logf("Got : %s, Wanted : %s", got, tc.result)

			if tc.result != got {
				t.Errorf("Got : %s, Wanted : %s", got, tc.result)
			}
		}
	})
}

func TestFinalValue(t *testing.T) {
	t.Run("2011. Final Value of Variable After Performing Operations", func(t *testing.T) {

		type testCases struct {
			ops    []string
			result int32
		}

		cases := []testCases{
			{[]string{"++X", "++X", "X++"}, 3},
			{[]string{"--X", "X++", "X++"}, 1},
			{[]string{"X++", "++X", "--X", "X--"}, 0},
		}

		for _, tc := range cases {
			got := FinalValueAfterOperations(tc.ops)
			t.Log(got)

			if int(tc.result) != got {
				t.Errorf("Got : %d, Wanted : %d", got, int(tc.result))
			}
		}

	})
}

func TestStringContainsCount(t *testing.T) {
	t.Run("771. Jewels and Stones", func(t *testing.T) {
		type testCases struct {
			jewel  string
			stones string
			output int
		}

		cases := []testCases{
			{"aAb", "aAAbbbb", 7},
			{"z", "ZZ", 0},
			{"aA", "aAAbbbb", 3},
		}

		for _, tc := range cases {
			got := NumJewelsInStones(tc.jewel, tc.stones)

			if tc.output != got {
				t.Errorf("Got : %d, Wanted : %d", got, int(tc.output))
			}
		}
	})
}

func TestMostWordsFound(t *testing.T) {

	type testCases struct {
		sentences []string
		result    int
	}

	cases := []testCases{
		{[]string{"alice and bob love leetcode", "i think so too", "this is great thanks very much"}, 6},
		{[]string{"please wait", "continue to fight", "continue to win"}, 3},
	}

	for _, tc := range cases {
		got := MostWordsFound(tc.sentences)
		t.Log(got)

		if tc.result != got {
			t.Errorf("Got : %d, Wanted : %d", got, int(tc.result))
		}
	}

}

func TestInterpret(t *testing.T) {
	type testCases struct {
		command string
		output  string
	}

	cases := []testCases{
		{"G()(al)", "Goal"},
		{"G()()()()(al)", "Gooooal"},
		{"(al)G(al)()()G", "alGalooG"},
	}

	for _, tc := range cases {
		got := InterpretLoop(tc.command)

		if tc.output != got {
			t.Errorf("Got : %s, Wanted : %s", got, tc.output)
		}
	}
}

func TestValidParentheses(t *testing.T) {
	type testCases struct {
		input  string
		output bool
	}

	cases := []testCases{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"(({{()}}))(", false},
		{"(({{()}}))", true},
		{"(({{()}}))()", true},
	}

	for _, tc := range cases {
		got := IsValidParentheses(tc.input)

		if tc.output != got {
			t.Errorf("Test Case: %s, Got : %t, Wanted : %t", tc.input, got, tc.output)
		}
	}
}

func TestLC28(t *testing.T) {
	type testCases struct {
		haystack string
		needle   string
		output   int
	}

	cases := []testCases{
		{"hello", "ll", 2},
		{"aaaaa", "bba", -1},
		{"", "", 0},
		{"mississippi", "issip", 4},
	}

	for _, tc := range cases {
		got := LC28(tc.haystack, tc.needle)

		if tc.output != got {
			t.Errorf("Got : %d, Wanted : %d", got, tc.output)
		}
	}
}

func TestLongestCommonPrefix(t *testing.T) {
	type testcases struct {
		input  []string
		output string
	}

	cases := []testcases{
		{[]string{"flower", "flight", "flow"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
	}

	for _, tc := range cases {
		got := LongestCommonPrefix(tc.input)

		if tc.output != got {
			t.Errorf("Got : %s, Wanted : %s", got, tc.output)
		}
	}
}
