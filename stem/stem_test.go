package stem

import (
	"strconv"
	"testing"
)

func TestDoStem(t *testing.T) {
	strs := []string{
		"caresses",
		"ponies",
		"ties",
		"caress",
		"cats",
		"feed",
		"agreed",
		"plastered",
		"bled",
		"motoring",
		"sing",
		"motoring",
		"conflated",
		"troubled",
		"sized",
		"hopping",
		"tanned",
		"failing",
		"filing",
		"sky",
		"happy",
		"relational",
		"conditional",
		"rational",
		"valenci",
		"hesitanci",
		"digitizer",
		"conformabli",
		"radicalli",
		"differentli",
		"vileli",
		"analogousli",
		"vietnamization",
		"predication",
		"operator",
		"feudalism",
		"decisiveness",
		"hopefulness",
		"callousness",
		"formaliti",
		"sensitiviti",
		"sensibiliti",
		"triplicate",
		"formative",
		"formalize",
		"electriciti",
		"electrical",
		"hopeful",
		"goodness",
		"revival",
		"allowance",
		"inference",
		"airliner",
		"gyroscopic",
		"adjustable",
		"defensible",
		"irritant",
		"replacement",
		"adjustment",
		"dependent",
		"adoption",
		"homologou",
		"communism",
		"activate",
		"angulariti",
		"homologous",
		"effective",
		"bowdlerize",
		"probate",
		"rate",
		"cease",
		"controll",
		"roll"}
	for _, s := range strs {
		news := DoStem(s)
		t.Log(s, "->", news)
	}
}

func TestAdd(t *testing.T) {
	stem := testInit("amazing")
	stem.Add('s')
	t.Log(string(stem.b))
	t.Logf("i  %d", stem.i)
	t.Logf("j  %d", stem.j)
	t.Logf("k0  %d", stem.k0)
	t.Logf("k  %d", stem.k)
}

func TestStep1(t *testing.T) {
	m := []string{
		"aaasses",
		"aaaies",
		"aaas",
		"aaaeed",
		"aaaated",
		"aaabled",
		"aaaized",
		"aaaazed",
		"aaazzed",
		"crime"}
	for _, s := range m {
		stem := testInit(s)
		stem.step1()
		t.Log(s, "->", string(stem.b), "  stem.k=", strconv.Itoa(stem.k))
	}
}
