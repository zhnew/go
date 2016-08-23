package stem

type Stemmer interface {
	Add(byte)
	GetResultBuffer() []byte
	GetResultLength() int
	Reset()
	Stem(string) string
	ToString() string
}

const (
	INC   = 50
	EXTRA = 1
)

//implementStemmer interface
type Stem struct {
	b           []byte
	i, j, k, k0 int
	dirty       bool
}

//constructor
func DoStem(s string) string {
	sm := Stem{b: make([]byte, INC), i: 0}
	//runtime.setFinalizer(sm,(*Stem) .stop)
	//gosm.backend()
	return sm.Stem(s)
}

func (stem *Stem) Add(ch byte) {
	if len(stem.b) <= stem.i+EXTRA {
		new_b := make([]byte, len(stem.b)+INC)
		copy(new_b, stem.b[0:])
		stem.b = new_b
	}
	stem.b[stem.i] = ch
	stem.i++
}

func (stem *Stem) ToString() string {
	str := string(stem.b[:stem.i+1])
	return str
}

func (stem *Stem) GetResultLength() int {
	return stem.i
}

func (stem *Stem) GetResultBuffer() []byte {
	return stem.b
}

func (stem *Stem) cons(i int) bool {
	switch stem.b[i] {
	case 'a':
		return false
	case 'e':
		return false
	case 'i':
		return false
	case 'o':
		return false
	case 'u':
		return false
	case 'y':
		f := func() bool {
			if i == stem.k0 {
				return true
			} else {
				return !stem.cons(i - 1)
			}
		}
		return f()
	}
	return true
}

/*m() measures the number of consonant sequences between k0 and j. if c is
a consonant sequence and v a vowel sequence, and <..> indicates arbitrary
presence,

<c><v>gives0
<c>vc<v>gives1
<c>vcvc<v>gives2
<c>vcvcvc<v>gives3
....
*/
func (stem *Stem) m() int {
	n := 0
	i := stem.k0
	for {
		if i > stem.j {
			return n
		}
		if !stem.cons(i) { //stem.b[i]isavowel
			break
		}
		i++
	}
	i++
	for {
		for { //meetaconsona,thenn++
			if i > stem.j {
				return n
			}
			if stem.cons(i) { //stem.b[i]isaconsonant
				break
			}
			i++
		}
		i++
		n++
		for { //skipavowel,stopwhenmeetavowel
			if i > stem.j {
				return n
			}
			if !stem.cons(i) {
				break
			}
			i++
		}
		i++
	}
}

/*vowelinstem() istrue<=>k0,...jcontainsavowel*/
func (stem *Stem) vowelinstem() bool {
	for i := stem.k0; i <= stem.j; i++ {
		if !stem.cons(i) {
			return true
		}
	}
	return false
}

/*doublec(j) istrue<=>j,(j-1) containadoubleconsonant.*/
func (stem *Stem) doublec(j int) bool {
	if j < stem.k0+1 {
		return false
	}
	if stem.b[j] != stem.b[j-1] {
		return false
	}
	return stem.cons(j)
}

/*cvc(i) istrue<=>i-2,i-1,ihastheformconsonant-vowel-consonant
andalsoif thesecondcisnotw,xory.thisisusedwhentryingto
restoreaneattheendofashortword.e.g.

cav(e) ,lov(e) ,hop(e) ,crim(e) ,but
snow,box,tray.

*/
func (stem *Stem) cvc(i int) bool {
	if i < stem.k0+1 || !stem.cons(i) || stem.cons(i-1) || !stem.cons(i-2) {
		return false
	} else {
		ch := stem.b[i]
		if ch == 'w' || ch == 'x' || ch == 'y' {
			return false
		}
	}
	return true
}

/*stem.b[?:k]==s[:]*/
func (stem *Stem) ends(s string) bool {
	l := len(s)
	o := stem.k - l + 1
	if o < stem.k0 {
		return false
	}
	for i := 0; i < l; i++ {
		if stem.b[o+i] != s[i] {
			return false
		}
	}
	stem.j = stem.k - l //stem is changed
	return true
}

/*setto(s) sets(j+1) ,...ktothecharacters inthestrings,readjusting
k.*/
func (stem *Stem) setto(s string) {
	l := len(s)
	o := stem.j + 1
	for i := 0; i < l; i++ {
		stem.b[o+i] = s[i]
	}
	stem.k = stem.j + l
	stem.dirty = true
}

/*r(s) isusedfurtherdown.*/
func (stem *Stem) r(s string) {
	if stem.m() > 0 {
		stem.setto(s)
	}
}

func (stem *Stem) step1() {
	if stem.b[stem.k] == 's' {
		if stem.ends("sses") {
			stem.k -= 2
		} else if stem.ends("ies") {
			stem.setto("i")
		} else if stem.b[stem.k-1] != 's' {
			stem.k--
		}
	}
	if stem.ends("eed") {
		if stem.m() > 0 {
			stem.k--
		}
	} else if (stem.ends("ed") || stem.ends("ing")) && stem.vowelinstem() {
		stem.k = stem.j
		if stem.ends("at") {
			stem.setto("ate")
		} else if stem.ends("bl") {
			stem.setto("ble")
		} else if stem.ends("iz") {
			stem.setto("ize")
		} else if stem.doublec(stem.k) {
			ch := stem.b[stem.k]
			stem.k--
			if ch == 'l' || ch == 's' || ch == 'z' {
				stem.k++
			}
		} else if stem.m() == 1 && stem.cvc(stem.k) {
			stem.setto("e")
		}
	}
}

/*step2() turnsterminalytoiwhenthereisanothervowel inthestem.*/
func (stem *Stem) step2() {
	if stem.ends("y") && stem.vowelinstem() {
		stem.b[stem.k] = 'i'
		stem.dirty = true
	}
}

/*step3() mapsdoublesufficestosingleones.so-ization(=-izeplus
-ation) mapsto-izeetc.notethatthestringbeforethesuffixmustgive
m() >0.*/

func (stem *Stem) step3() {
	if stem.k == stem.k0 {
		return /*ForBug1*/
	}
	switch stem.b[stem.k-1] {
	case 'a':
		if stem.ends("ational") {
			stem.r("ate")
		}
		if stem.ends("tional") {
			stem.r("tion")
		}
	case 'c':
		if stem.ends("enci") {
			stem.r("ence")
		}
		if stem.ends("anci") {
			stem.r("ance")
			break
		}
	case 'e':
		if stem.ends("izer") {
			stem.r("ize")
		}
	case 'l':
		if stem.ends("bli") {
			stem.r("ble")
			break
		}
		if stem.ends("alli") {
			stem.r("al")
			break
		}
		if stem.ends("entli") {
			stem.r("ent")
			break
		}
		if stem.ends("eli") {
			stem.r("e")
			break
		}
		if stem.ends("ousli") {
			stem.r("ous")
			break
		}
	case 'o':
		if stem.ends("ization") {
			stem.r("ize")
			break
		}
		if stem.ends("ation") {
			stem.r("ate")
			break
		}
		if stem.ends("ator") {
			stem.r("ate")
			break
		}
	case 's':
		if stem.ends("alism") {
			stem.r("al")
			break
		}
		if stem.ends("iveness") {
			stem.r("ive")
			break
		}
		if stem.ends("fulness") {
			stem.r("ful")
			break
		}
		if stem.ends("ousness") {
			stem.r("ous")
			break
		}
	case 't':
		if stem.ends("aliti") {
			stem.r("al")
			break
		}
		if stem.ends("iviti") {
			stem.r("ive")
			break
		}
		if stem.ends("biliti") {
			stem.r("ble")
			break
		}
	case 'g':
		if stem.ends("logi") {
			stem.r("log")
			break
		}
	}
}

/*step4() dealswith-ic-,-full,-nessetc.similarstrategytostep3.*/

func (stem *Stem) step4() {
	switch stem.b[stem.k] {
	case 'e':
		if stem.ends("icate") {
			stem.r("ic")
			break
		}
		if stem.ends("ative") {
			stem.r("")
			break
		}
		if stem.ends("alize") {
			stem.r("al")
			break
		}
	case 'i':
		if stem.ends("iciti") {
			stem.r("ic")
			break
		}
	case 'l':
		if stem.ends("ical") {
			stem.r("ic")
			break
		}
		if stem.ends("ful") {
			stem.r("")
			break
		}
	case 's':
		if stem.ends("ness") {
			stem.r("")
			break
		}
	}
}

/*step5() takesoff-ant,-enceetc.,incontext<c>vcvc<v>.*/

func (stem *Stem) step5() {
	if stem.k == stem.k0 {
		return /*forBug1*/
	}
	switch stem.b[stem.k-1] {
	case 'a':
		if stem.ends("al") {
			break
		}
		return
	case 'c':
		if stem.ends("ance") {
			break
		}
		if stem.ends("ence") {
			break
		}
		return
	case 'e':
		if stem.ends("er") {
			break
		}
		return
	case 'i':
		if stem.ends("ic") {
			break
		}
		return
	case 'l':
		if stem.ends("able") {
			break
		}
		if stem.ends("ible") {
			break
		}
		return
	case 'n':
		if stem.ends("ant") {
			break
		}
		if stem.ends("ement") {
			break
		}
		if stem.ends("ment") {
			break
		}
		/*elementetc.notstrippedbeforethem*/
		if stem.ends("ent") {
			break
		}
		return
	case 'o':
		if stem.ends("ion") && stem.j >= 0 && (stem.b[stem.j] == 's' || stem.b[stem.j] == 't') {
			break
		}
		/*j>=0fixesBug2*/
		if stem.ends("ou") {
			break
		}
		return
	/*takescareof-ous*/
	case 's':
		if stem.ends("ism") {
			break
		}
		return
	case 't':
		if stem.ends("ate") {
			break
		}
		if stem.ends("iti") {
			break
		}
		return
	case 'u':
		if stem.ends("ous") {
			break
		}
		return
	case 'v':
		if stem.ends("ive") {
			break
		}
		return
	case 'z':
		if stem.ends("ize") {
			break
		}
		return
	default:
		return
	}
	if stem.m() > 1 {
		stem.k = stem.j
	}
}

func (stem *Stem) step6() {
	stem.j = stem.k
	if stem.b[stem.k] == 'e' {
		a := stem.m()
		if a > 1 || a == 1 && !stem.cvc(stem.k-1) {
			stem.k--
		}
	}
	if stem.b[stem.k] == 'l' && stem.doublec(stem.k) && stem.m() > 1 {
		stem.k--
	}
}

func (stem *Stem) stem(wordBuffer []byte, offset int, wordLen int) bool {
	stem.i = 0
	stem.dirty = false
	if len(stem.b) < wordLen {
		new_b := make([]byte, wordLen+EXTRA)
		stem.b = new_b
	}
	copy(stem.b, wordBuffer[offset:wordLen])
	stem.i = wordLen
	//stem(0)
	stem.k = stem.i - 1
	stem.k0 = 0
	if stem.k > stem.k0+1 {
		stem.step1()
		stem.step2()
		stem.step3()
		stem.step4()
		stem.step5()
		stem.step6()
	}
	if stem.i != stem.k+1 {
		stem.dirty = true
	}
	stem.i = stem.k + 1
	return stem.dirty
}

func (stem *Stem) Stem(s string) string {
	if stem.stem([]byte(s), 0, len(s)) {
		return stem.ToString()
	} else {
		return s
	}
}

func testInit(s string) *Stem {
	stem := Stem{b: make([]byte, INC), i: 0}
	stem.i = 0
	stem.dirty = false
	wordLen := len(s)
	wordBuffer := []byte(s)
	if len(stem.b) < wordLen {
		new_b := make([]byte, wordLen+EXTRA)
		stem.b = new_b
	}
	copy(stem.b, wordBuffer[0:wordLen])
	stem.i = wordLen
	//stem(0)
	stem.k = stem.i - 1
	stem.k0 = 0
	return &stem
}
