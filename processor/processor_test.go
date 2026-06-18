package processor

import "testing"

func TestHandleHexBin(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"1E (hex) files were added", "30 files were added"},
		{"It has been 10 (bin) years", "It has been 2 years"},
		{"Simply add 42 (hex) and 10 (bin)", "Simply add 66 and 2"},
		{"FF (hex) and 111 (bin)", "255 and 7"},
	}
	for _, tt := range tests {
		result := HandleHexBin(tt.input)
		if result != tt.expected {
			t.Errorf("expected %q got %q", tt.expected, result)
		}
	}
}

func TestCaseModifiers(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"go (up)", "GO"},
		{"HELLO (low)", "hello"},
		{"brooklyn (cap)", "Brooklyn"},
		{"This is so exciting (up, 2)", "This is SO EXCITING"},
		{"THIS IS VERY LOUD (low, 3)", "THIS is very loud"},
		{"welcome to the age of foolishness (cap, 6)", "Welcome To The Age Of Foolishness"},
	}
	for _, tt := range tests {
		result := HandleCaseModifiers(tt.input)
		if result != tt.expected {
			t.Errorf("expected %q got %q", tt.expected, result)
		}
	}
}

func TestFixArticles(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"a apple", "an apple"},
		{"A amazing rock", "An amazing rock"},
		{"an dog", "a dog"},
		{"An cat", "A cat"},
		{"a house", "an house"},
	}
	for _, tt := range tests {
		result := FixArticles(tt.input)
		if result != tt.expected {
			t.Errorf("expected %q got %q", tt.expected, result)
		}
	}
}

func TestFixPunctuation(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello , world", "hello, world"},
		{"hello,world", "hello, world"},
		{"What ?", "What?"},
		{"I was thinking ... You were right", "I was thinking... You were right"},
		{"I am ' awesome '", "I am 'awesome'"},
	}
	for _, tt := range tests {
		result := FixPunctuation(tt.input)
		if result != tt.expected {
			t.Errorf("expected %q got %q", tt.expected, result)
		}
	}
}

func TestProcess(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			"Simply add 42 (hex) and 10 (bin) , wow !",
			"Simply add 66 and 2, wow!",
		},
		{
			"it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair.",
			"It was the best of times, it was the worst of TIMES, it was the age of wisdom, It Was The Age Of Foolishness, it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, it was the winter of despair.",
		},
		{
			"There is no greater agony than bearing a untold story inside you.",
			"There is no greater agony than bearing an untold story inside you.",
		},
		{
			"I am exactly how they describe me: ' awesome '",
			"I am exactly how they describe me: 'awesome'",
		},
		{
			"FF (hex) and 111 (bin)",
			"255 and 7",
		},
	}
	for _, tt := range tests {
		result := Process(tt.input)
		if result != tt.expected {
			t.Errorf("expected %q got %q", tt.expected, result)
		}
	}
}
