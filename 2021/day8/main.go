package main

import (
	"fmt"
	"strconv"
	"strings"
)


var input []display


func main(){
	Init()
	fmt.Printf("part 1: %v\n", partOne())
	fmt.Printf("part 2: %v\n", partTwo())
}

func partOne() (count int) {
	for _, display := range input {
		for _, d := range display.digits {
			// count all digits that are 1,4,7,8
			switch digit(d).value() {
			case 1, 4, 7, 8:
				count++
			}
		}
	}
	return
}

func partTwo() (count int) {
	for _, display := range input {
		count += display.DecodeValue()
	}
	return
}

func Init() {
	ins := strings.Split(_input, "\n")
	input = make([]display, len(ins))
	for i, v := range ins {
		input[i] = parseDisplay(v)
	}
}

type display struct {
	digits []string

	signals map[int]string
	sig2seg map[rune]rune
	seg2sig map[rune]rune

}


func strip(in string, remove ...string) string {
	for _, r := range remove {
		for _, v := range r {
			in = strings.Replace(in, string(v), "", -1)
		}
	}

	return in
}

type digit string

func (d digit) value() int{
	switch len(d) {
	case 2:
		return 1
	case 3:
		return 7
	case 4:
		return 4
	case 7:
		return 8
	}

	return 0
}

func parseDisplay(in string) (d display) {
	parts := strings.Split(in, " | ")

	// parse signal sequence
	_sigs := strings.Split(parts[0], " ")
	sigs := make(map[string]bool)
	for _, v := range _sigs {
		sigs[v] = true
	}

	// parse display digits
	d.digits = strings.Split(parts[1], " ")

	// learn the display
	signals := map[int]string{}
	sig2seg := map[rune]rune{}
	seg2sig := map[rune]rune{}

	// seed what we already know
	for v, _ := range sigs {
		switch len(v) {
		case 2:
			signals[1] = v
			delete(sigs, v)
		case 3:
			signals[7] = v
			delete(sigs, v)
		case 4:
			signals[4] = v
			delete(sigs, v)
		case 7:
			signals[8] = v
			delete(sigs, v)
		}
	}

	// compute 'a' (7 - 1)
	a := strip(signals[7], signals[1])
	sig2seg[rune(a[0])] = 'a'
	seg2sig['a'] = rune(a[0])

	// compute 'g': '9' (length-6 that intersect '1' == '1') - '1' - '4' - '7')
find9:
	for v, _ := range sigs {
		if len(v) != 6 {
			continue
		}
		for _, v1 := range signals[4] {
			if !strings.Contains(v, string(v1)){
				continue find9
			}
		}
		signals[9] = v
		delete(sigs, v)
	}
	g := strip(signals[9], signals[1], signals[4], signals[7])
	sig2seg[rune(g[0])] = 'g'
	seg2sig['g'] = rune(g[0])


	// compute 'e': '8' - '9'
	e := strip(signals[8], signals[9])
	sig2seg[rune(e[0])] = 'e'
	seg2sig['e'] = rune(e[0])


	// compute 'b': '0' (len(6) V '7' = '7' && !'9') -'7' - `e` - `g`
find0:
	for v, _ := range sigs {
		if len(v) != 6 {
			continue
		}
		for _, v1 := range signals[7] {
			if !strings.Contains(v, string(v1)){
				continue find0
			}
		}
		signals[0] = v
		delete(sigs, v)
	}
	b := strip(signals[0], signals[7], string(seg2sig['e']), string(seg2sig['g']))
	sig2seg[rune(b[0])] = 'b'
	seg2sig['b'] = rune(b[0])

	// compute 'd': '8' - '0'
	_d := strip(signals[8], signals[0])
	sig2seg[rune(_d[0])] = 'd'
	seg2sig['d'] = rune(_d[0])

	// find '6'
	for v, _ := range sigs {
		if len(v) != 6 {
			continue
		}
		signals[6] = v
		delete(sigs, v)
	}

	// compute 'c'
	c := strip(signals[8], signals[6])
	sig2seg[rune(c[0])] = 'c'
	seg2sig['c'] = rune(c[0])

	// compute 'f'
	f := strip(signals[1], string(seg2sig['c']))
	sig2seg[rune(f[0])] = 'f'
	seg2sig['f'] = rune(f[0])

	// find 3
	for v, _ := range sigs {
		if len(strip(v, signals[1])) != 3 {
			continue
		}
		signals[3] = v
		delete(sigs, v)
	}

	// find 2
	for v, _ := range sigs {
		if strings.Contains(v, string(seg2sig['e'])){
			signals[2] = v
			delete(sigs, v)
		}
	}

	// find 5
	for v, _ := range sigs {
		if !strings.Contains(v, string(seg2sig['b'])) || len(v) != 5 {
			fmt.Println("CRAP")
		} else {
			signals[5] = v
		}
	}

	d.seg2sig = seg2sig
	d.sig2seg = sig2seg
	d.signals = signals

	return d
}

func (d *display) DecodeValue() int {
	var chars string
	for _, dig := range d.digits {
		for k, sig := range d.signals {
			if len(sig) != len(dig) || len(strip(dig, sig)) != 0{
				continue
			}
			chars += fmt.Sprintf("%d", k)
			break
		}
	}
	val, _ := strconv.Atoi(chars)
	return val
}


var _input_simple = `acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cdfeb fcadb cdfeb cdbaf`

var _input = `ec cabfe afebd dbagef afbcg feabcd cdef eafdcbg ecb caegdb | ceb ecb febac ec
acfdg ea baec dgbafec cfebg efa cfebgd dfbaeg egcaf abecfg | bgcdfe ebac cadgf dceagfb
gfed caefd dcegaf ge daegc bcdag dfagbce cfbead gce afbegc | ceadg edfg gdbca fcead
ea dgfcae feba cae efbcagd cagfb gedcb gbfcad gebca afbceg | facbgde cegdbaf cgabe degcb
gcbfed bdf cbdfgea fgeabc dcfab caefb db caebdf abde fgdca | fdb fabce dbacf afceb
bacf afg egfab dcbgef bcegdaf efgbac efcgda fa gbfce gebad | edgba bfacged af fdacbeg
adgef bagcf cdea adfebg dfgca adfgec cd fcdgbe cdf cdfegab | afgdbe facdg gdfbeac dcea
cdgeb cd cdb gfecbda adgcbe bcdfag aecd bfceg edbfga agdbe | bdc dbeacg dceabg gcedb
ae cdbeg gefcdb eca ebgfac dcbafeg deag abcde facbd becadg | aec cbaedg aec deag
gabfc cdfbe ea dbcfeag aebfgc befagd ebfac aeb gaec bagdfc | ea gbadfc dgbfca decbf
bcad cd cfbge cegabdf fcdeb defabc gebadf fdc adfbe ecadfg | dc fcd egdcfa dc
begcda cbf egacb fc ebcaf baefd gabdfc gfce gaecfb acgdfeb | bdegfca cf fc acbdeg
fbgce dagcbe eafcbg bcgdfe geacdfb efdb dcgaf dge de gfdec | cgbfe bagced daecbgf bfde
fdgae fce cdeab gcfd fc efdca eadfgb fcgdae efdcbag gaefcb | efbdag fgeda adbfeg befagc
cfag cadgfb cf cfd cdbgf gaebdc bgfde cgbfead befacd cadbg | gcbda fcd bfged dcf
gebf edbcga edcabgf cfabed aeg agcfeb fgaec ge cgadf ebafc | abcgdfe gafbce cfagbe bgfe
aegbfcd gbacef bfea bfg efcag cfbge bf fcdega bedgc gbacfd | bcagedf bfea acdbefg faceg
dcf df aedf ebacf egabfcd fcdbge acfbd adbgc dfabec afgceb | caebf dcf dcf ebcagf
bgad abcge dbcfgae bcfae gecdab eagcfd ag cgdbe gae bcedfg | adbg gdba ga aeg
cega begfdca bdcfea febcg bec fbadcg fgbeac ec bfgac fgbed | ec gadbfc cbgfa bgcadef
fegbac debg bgc cadbg gdefca efbagcd dbgcea dfbca caegd bg | gb edgbcaf gbc cbg
gbefc badecfg acgeb ecfgba cfbgd bef ef acbged caebdf afeg | gefa ecdbafg gbdfc cegfb
gbefca gcdba dbfeag ec gbfae becf gdafce fegadcb aegcb aec | gebafc befcga gfadbe cgeab
cegadf cafg gdeca fdbcae fadcgbe gdc gc cbefdg ebgad fcdea | acdef gbcdef agcf cgbdaef
gcdbef eagdbc ecbdg afecgb bfc fbecd geabfdc fb fbgd fdcae | bcf cfdegba fb dfcbe
gafed badfce db dcgb dab dbaeg dgafceb acegb cabedg aebgcf | afdge db bcgaed cbaefd
bgcda gdfb fecbda gb cbg bafcgd gacbdef bfeagc cbafd dagce | gadec cgb efgdacb gdacb
cbefd dfegba bcfead db aebcfdg deacf bdf dbac fcbeg geadcf | gbefda efacd bd dfb
agdfbe gfcbeda egcad cegbd bgafce gcfead gfade acfd ac cea | afdge eac gaefd cdega
ebfad fde edbca gbfdce feag gaebfd ef adgfb bgcfad cbdfgae | gfabdc edf fage cdabe
bdce cefabdg bafdge afced cd cad badfce gdbfac aefcg ebdaf | ecagf bedc dac dc
cgebf gba ecadb ebcdfa gcad adgcbfe dbecga ga gbeca eagbdf | ag bfgec ag cdbeaf
dbcf dgecf ebcgd gbcea gebdcf bed dfacbge gceadf ebafgd bd | db dbe faecdbg dbe
gdea fed aedbf efadgb de fbega bgaecf gaecbdf egbcdf cdfab | ed cfeabg aged dgecabf
afebc fdeacb egf cfgb daceg egbcaf cafge gf ebadcfg egfbda | fge afceb eacfg agecd
ecfad deabcf cagfe gc degfca dgcf dabcge gac gabcefd ebfag | cga afgce gdfc agc
eac gbaec cfeb acefgb afebg ce dacbg bdegaf egafdc bgceafd | eadcfgb aec ace aec
gefba afb af gacbe ecbgaf dagfcb feca deabgc fbdceag dfbge | befdcga aegfbc efbcag baf
bdg ecdgf fbagc ebdafg db cbgfd cebd baegcfd fedcag dcegbf | db eacfdbg bdce bd
cfeab ebdgca agb ga aecbg cgda eafgdb ebcgdf cbfaegd dgbce | cbdfega ecafb gdac bga
caed gcdfa ac bfcedg edagfc eabcfg caf dagcfbe gbfda dfcge | degfacb fbcaged cfa dgafbce
bgcefda eagdbf acdbe gdfbe abgf aedcfg cedfbg fea fdeab fa | aef bcead fgbdae abfdeg
cbd cgdaeb gaebd dfceg cgbde cbefgda bfcgda eafbgd cbae bc | baegd bc aceb gbdeca
fbaegc dcbaf dbcae abcfg adbcegf bcdfga dfa cdfg fgdaeb df | fcgaedb fd fd begcaf
bcdeg ea fbaegdc eag eafgbc dgaeb bgafdc feda adbfg gdabfe | gdbea eag ea aeg
beacfd bead dafce cea edfagcb ea gadfc bfecd aecfgb fcdbeg | ae cgefab adcef aebd
dcefgb agfdce bgdfe agfdbe fad bdafe ceadb fa fbag dgcafbe | abefd ebfdgac fa dbgef
cbdfg cgeabfd fe fge bfgdca gfacde becgdf bdef ebcga egbfc | ef bgfcaed egbfc gaebcfd
adcge fcad bagced cf cef fegcd edacgbf cafegd febgd aecgbf | fec agefdc dacf egcabdf
fagbd bgec fdcbea be eacgf eba dceafg faebg adbecfg baefgc | gfcaedb ebfga bea fabgd
dbecfa bfcga begdfa cgdebf ge fge gbcef cegd dcfeabg fdceb | cbegfad eg gdec bgacf
cae ebcf ce gcfad bdefag fcaeg cgfdbae abdegc acbgef faegb | fegac ec cdgaf bcaedgf
bcfae cbgedf agdfcbe gdca fgc gcdfba bagfd gbafed gc bgafc | fcg edfbcga feadgbc adgc
cfdgab fdg dbcfa gf efgdac afgdb decagfb edfbca cgbf abged | fagdceb bdagf cebdgaf edcbaf
bcefg bafge ebadgfc cfe ec dcgfb bafgec edbgfa caeg dcebfa | afgeb cefbga dcbgf fegba
fbc gbced edcbfg fb dbgf befgc fgace degabc bcafdeg fbaced | aefgcbd fedagbc bfc fb
gd gbd eabfcdg egad cgbda eagbc cdgabe feabgc dcebgf adbfc | cadfbeg daeg bcgefa gd
edafb dagfb agd ecbgda dgbfc edgbcf fcga geafbdc cfdbga ag | gcfa gda gfdebca gda
edabg bcafgd gacfd dgfcae bfa fb dafbecg edcfab fgbc dafgb | fba fdcgeab fb badge
ebfd gbdaef gafcbe efadg agfdb bf gcadbfe fab cbdga cfgead | fab aebgfcd fb debf
geadcf cadgfb ab cbagf acgdf dfcegab adbcge fbgec abfd abc | acfdeg acgfdb bcfga dgcfa
dacfbe cbga ebdfga bc ecbfadg agfeb ecgfd bfc eagcfb bfegc | fecdab cfegb bfecg decgf
cgdabf ecdfab ecgf ecfdb bgf febgcd gf gdbef abged gdefbac | afgcdb gefc gf adfcgbe
cgda gcfedb cbdage aedfbgc cd cgaebf eadbc bceag dbc defba | abgedc cdb cdb fgcebda
dafbgce agd fcdage gd afecg edgc egabfc efdab dabgfc afedg | beafgc dgec gad gecd
dbgecf faedbg dgfea dacfge fcaed dagc dc dfc febca abfgced | fedag dfeagc dgfacbe cd
fbgc gcbadf dafcg gdaeb gfdace bf bafgd defbca bfa gfbecda | fba fb cgfb cfdabg
abfceg adgfbe cba cfea cbega abgef dbfgca ca cebgd ebgdafc | ac ca acb bca
bafg decfa feabdg cafbged ga gea fgedb dbcage fadge fedbgc | eacdgb ega bfga cefbdg
cbafg fe debga bgfea ecfa efb dfbegc afcbdg fbecga dcgbfea | bgdfcae fcbga bgaecf fe
gdfa eagbc gdc bdfeca cdfbga dg afcdebg cagbd bfdceg fbcad | dgaf dg dgc cgd
egbcdf gdbeafc fcd cafde abfdge adfebc gecda fbead fc bcfa | dagce gdaefbc cf fdc
eca cgafdb afeb gecba febdgca gadfce fcagb ea dbceg cgbfea | cafegd cae ecbdfag cgbea
ge cfbead abefgc acfbe ceafdg fbge abcge badgc gec caefbgd | efbca egc dbgafec dafbgec
dc gaefc deafgb fdc dfeba adbc efdca dfgacbe fdegcb dcbafe | fcd afecg efcgbda feadb
dcbfe bfe efdacb afcebg cebda gcefd afgbcde bdfa fb edbcag | gecdf fb bfe bf
ecbfa febagc bf dbagfe edgcab fcgdeab afb eacgb ecafd gcfb | acbefg gaebdc cgbf fbcage
begcda fcedba gadefbc cgeaf dce bfdc dacfe cd afebd abedgf | ecdabg gdfeba afbcged afgcbde
ae ebcgd dcfgaeb becagd dbeacf afbgc ecbag eca gcbdfe edag | cea becga gbecad ae
cgbafed geafdc ebacfg cde bdcfg gfced dafe acfeg ed agdbce | daef eafd edc cfegd
fdceag ecdabgf bfgcd fgdba fag daebg beaf af egcabd daefgb | afdcebg beaf af dfegbca
bc caeb adgfcb gfbde aedgcf gefabc adcfgeb bgc fagce cfegb | bc cb fegca bc
agecbd gaefbc gefdc bgfdac bdgca ae age edabgcf decga edab | dacefgb dafgecb ega ae
egbad debfa dfag bgaecf abgdce fba fa cgbefad agedfb cdbef | fgbead fagd becgda af
bgc edgbfca bfgce agdcfe acgefb efcbd cafeg bg bage bagfdc | gb bgea bgcef dcbef
gab cfeab gdcea facgbe adbfcg cafdebg dfaceb begac egfb gb | dcgfba bg gb bag
bgfceda gcafed egfd eadcg abgdce gfaec bdfagc eabfc gf fgc | egdf agdce cgf cafged
bacfegd egbad adgce adcb dbagec acgfed bag ba ebdfg gafebc | bacdeg badc efgbdac bdac
bfcdg bdgcfe cfb cb becd gfdec eacgdf bfecgad fbgda cfegab | cbf bfc cb cafbegd
cebgf gbeda fa fedcab cdegfb cgbeaf bafdegc fgeab eaf gcfa | gcebfd fagcebd af fcga
cefba eg fecbag adebcg aebcdgf dbagf cgef abecdf egb fgaeb | fadceb bafcegd bfeac gcfe
acgd fadgb edcfbga afgcbd cdfbe gc bfgdc cfg gfaecb ebgfad | cbdgefa bfaecgd bfcaedg fcdeb
dgbe aedbcfg bfegdc agcbdf dcgfe fcdgb fed ed efbcad efgca | edgb fgcbd bdge edgb
gfadc geacf gdf fd dgafcb gaecbd bdcag cgdebf bafd fdbaceg | bedgacf gdebfc fbagced fdba
dcbefa eadfbcg ab eadcg fegbda dba fgebd gabf gbcdfe bdage | afbg bfcdae bgaf bda
cbf cf edgcba fdcabg ebgdf gacdb bgfcea cgbdf fbgecda acdf | egdbf fc fbc fdbagc
ebdcgf edgba ef afce aefgdc edf edfag bcagdf gfaebdc gfdca | def edcgaf fdeagc fe
bdaef fbcgae ae dcae dgcfab fbcda bae dbgfe badcgef dfeacb | aeb cbgfae bafced ae
ebacf gebfa gbfaed edfga bg ebdg edgacf bdcafg gba cgabefd | abg edgb dgbe fecabdg
faged ba gaefb abcf abfecg gefbc edagcb dafbcge gab cbfegd | ba gfbae gbadfce cabdgfe
dbag cgdbe ebg dbaefc gebcfa geabcd daebgfc gfdce dbeca gb | abdg bg geb gebadc
bcegda egcd ecfgdba edagb dgfbea abcfd cg bgacd ebgfca gcb | efgabc bgeda degc egadbc
bagecfd geadbf debaf dfge abegf fecbga fbd df abfgcd edabc | cbefgad faegdb dfebga df
ecdabf gbdfa eabfc bedc cdf gcfaebd cgbfea cd dfagce abfcd | dcaegf becd cd dbfcea
egac fdabce faedg fdceag gfe dgafb defac daecfbg ge bgdecf | geafd fcdea cedgaf gbfceda
facebgd ebcgaf afecdb eba egaf ae cdegb agfbc bdfcag bgeac | ecgbd egfa ea fgae
dgfbce dfa fbedga bgadf cdgab af gefa aecdbf ebfdg bafgcde | febdg fega acbfdge abefdg
bcdae aebcg dacf gefbda cbefd badcegf begfdc adb ad efcbad | ad adcf acfdbe bad
fgbcda dgcfeab gdcae fabge dgb bdfgea bd adbeg bceagf fbed | bgd db db bd
dgb cbfgaed bdeafc bfacdg gbdef gd agde adfeb dbgafe fgceb | egbafdc bagefd egda agde
fegbdc eabfgd edfbgac ad fdgbe bade agcfe gadef dfa dacfgb | da eadb dgfae ad
dbfag fbgac adfeg cfdb gdbeca dab fcabdg db ebgfca gecabfd | adb bgceadf dfaeg cgaefbd
ebgca ecdbfa bcafg ge cgbefad agde edbca geb ebdcgf dgabce | bdcfgae ebg ge bge
gcadbef fgbd gcd gd aecgb bcdefg afdebc cbfde afcdeg dcebg | gcd dgc gdc befacd
fdabge ab gdeca fagcbed gba dgbefc eafb bgcfda gdbae fgdeb | adceg abef efab eafb
cfgdba cd abecd cfbadge dcb adbfe agfbde bcafde egabc dcfe | aecdb defbag eadfbc dc
dgcaeb acfdbg ge dgecfb gdbecaf gace aedbf adbcg dge gbdea | edabf cgdfbe eg dgaeb
egbfac ceg geacd ecbda ebgd dbceag dbgcfea dbacef gadfc ge | egc dcbaegf acfgd aegbdc
gd fcgbad abegdf gbed feabcd acdfgbe efacg gfead feadb fdg | dg egcfa aedgf ebdcaf
gf defcg gefb gfecdb fgc ebacgfd gbdcfa facedb ebfcd degca | efbg eadbfgc fgc fgecbd
acgfb febgda dgfabec fa fab gdeacb eafc cbgdf fbegac ebgca | af acef beagfc fcbga
fb fbec dgacbef fagecd fabdeg fbdcg dgecf fegbdc bdf bagcd | cebf fb efcb fb
ebgfc cbefdg adgfbc edbfg gadfbce fbd adebg fd gbfeac efcd | fd dfegbca bgecfa df
ceafd abdfc dbgcf gdaecb eabcdf cegfda abfe abd ba cbefadg | ab edfac fbea cegdfab
bgdce aedfbc agecb egfcad ab abe cafge gbefca bfcgaed abgf | acbeg cgeabf ecafgd fdgabec
faegd eb fabe efbdag egb fgbed cbgdea dcfgb cbgefad gfecad | adbgcef agfed ebg fbegda
agbcd cfgabe gca dfcegab agdbf gc gdfc dbfcga eacdb faedgb | cbafdg abdce fdbga cabfgde
bdaec bfgde ga agdbce gdca gceafbd abdeg gae cfabeg edacfb | deacb adgbcfe becda ecagdbf
efagdcb edgbac bfcea cfagb dfecb bae fagcbe aefg ea cdafbg | afge afge dcebgfa fega
bgaec geadb fgbdec gefdba fbacg cbdage aedc gce dfaegcb ec | ecg adec acfdgbe cdae
egad dgcbf fcbeag eg cgaedf ceafd cbfead cegfd acfgbde cge | fgdec afdbce fbecda gefbcad
cdegb dbea gbdcae dgfecab febcga cbaeg ed fdbgc cde cadfeg | dce ecd adgbec gbacef
efbgdca cbafg gfdbce egbafd ecbaf ag gadc abgfdc bag gcfbd | bcgfd ag gebafd gabfed
gfac gfdbca gba degbc bcfad agdbfe ebdgcfa efcadb ag bdgca | aebgfcd afecbd gab acgbedf
dfeagc ecfd abegdc bcdfga eca geacf fgacebd gfeab ce dgcfa | bgcfda egcaf dfbcaeg agefb
efbdag ecgbaf dcgafb dacbe fdgacbe dfgc df fda acbdf cgbfa | fda afbgc cdabgf df
edcagb bdgac eagb ga bcgfd ecdab acgefd cafedb abgefdc dag | ag gbdca ga bgcda
dafbecg cgead fcbd agbedf db bde becad cbeaf bfgeac dfbeca | abgefd fecba fbcd cbdf
edg bfdace agebc ebdcf dgfabe cbged dg bgdefc gdcf bdgceaf | fcgd fdcg fdcg ged
fcabd bcde dfcbaeg efdbac cfbage bdfag fadec afecdg cb bca | cfdea dabfec cbafeg dafce
gbeafd agcb cefbg gdaecf bg gbf bcdfe fceag bagcef afecdgb | gb efgadb cgab caefg
cegbadf gaed dcafeg feacg gd fcbagd dgfec bfegca gcd cebdf | gd gd aecfdg cgd
gfecd gedba dceagbf fdage gedafc afd fcbgad becdfg feac af | cgdeafb fgcebda edgaf gfedca
fdbgcea fedbgc agdcfe debcag deba bcgde dga bfcga dcabg ad | dcagb cfgbdae dcgebfa dbae
abd debcfg caegdb da bcafgd acdf ebgfa fagcedb gabfd fbdcg | bgefa da cgdefb edcgabf
ea fgace egafbc dfegabc fbae fgcbe cegbad ega dbcfge gdacf | bfea cgdfa ea bdacfge
ecba eadcg eb edb fbagd gdcebf gbceda dacfge gbaed fagdbec | aefgcd bdceafg gbfdeac afdgb
ed dgeb dfacbg dbcgf gbdcfae efbdgc dec deafcg bfaec efdbc | fbgdce gadbfec fdgcbe gadfce
dgec gafbcde bcgdfa dacefg efcda cegaf acd cegafb cd eafdb | gaecdfb fcdgae dc gdcebaf
fdbgca abfc cbadg cf fcg edfgb cdgbae cgaedf abcefgd dcgfb | badegc cf adgceb fc
eb fgcde bafcgd cbfeda cegdba becgd bce baeg dagcb ceabfgd | adgceb debgafc be gacdb
ebcafd dgfeb befdgca gce egabdc adecb egdcb cbegaf gc cgda | bfaedc gbecaf egc dgbecfa
gafdbc cfaeg aecbd efbd fcegadb fcb bacfe cagebd bf fcbade | fb bcf bdef gdeacb
ce gedc afbdgce abdfc dcbgea cae abcde afgbec edgbfa bgade | dfaegb baced cfadebg cea
bagf bge dfgce fdbae egadbf dgbcae fdgbe fabcgde gb facbed | bfga gfab bg gbe
cb fabc cfedb fdcge fdbeca gedbca bce abgfdce adfeb ebfagd | cabfde abfc ebc ebcfd
edbf bafeg cadfgb cgeab fba bf agfed cgbfaed fgacde gfbaed | cedbagf bf bf fba
bceadgf begfac gfedab gabdcf cegf caf efgab bcade cbaef fc | efgc fbcgdae gbcfda bdgafe
bafdg agedf dbgecaf fge dcefa bafcge ge eafbcd gcdaef egdc | eg fdcebga ecgd egf
afbdeg abfceg ecg afgeb bfce fgcae fcadg gcdebaf ce dbcega | befag dabcgfe eagbdc egc
fbga dabeg debcf dgafbe fea bfdea gdbfcea agcfde fa egbacd | edgfac eaf fa fae
bafec cf cfadgb faebd bcf beacg eadgbf bdceafg cdfe dfabec | fcb afcdgb cfb bfcae
ebfacd geafbc cgdfbe gabfd ceag cg cabfe bgcfa bgc fdeagcb | febdgc gbc gc dcefgab
egacf cadb efdcab cfedb gfabde dea feadc decbgf fbadecg da | gcfbed acbgdfe bcedf becfad
bfgcad cebaf fabgc afgbec egfb fe abdce efc fdagce baedcfg | fec gbfca fe bdefacg
gbfedca bega bcdef fgeadc bgcad ea bcfgda gdcabe acdeb ace | ea dgbcae afbdcge ae
fcadebg gbecd adc fdbag bdagc aedbfc ecgfbd ac bdceag geca | eacg cgae adc ca
afdgec cfgdab fbedag eabcf agb cgfdeba dgeb gb afged abegf | bg fbgea badfge gdeb
edgab adfcg cedabgf ce edbgcf daecbg ecd dgcae ebac fbgaed | ec ec ecba eagbd
cebf eadbg bc degfbca cbg gcfbda agfce bgcae cbgfae efcdga | cbef badcfg gcb bcfe
cafbde cbfg abgef gacedf gf edbag faedgcb bfeac gbefca fga | gcfb fceab gaf beadg
edfab feag dag bgdfc adgbf gbfaed ga egdcabf cadgbe bacdef | gefabcd aegf afge eagf
cdfeab cdg efdbc fcbdga dcgaebf dbge eagfc debfcg gd cgfde | fabcgde dgc egbd cgd
cagbef gcb fbgead cfdg cbdea bdegc cg ebgdfc fgedacb bdgef | bfegda gc gc cfegdab
edcgf be feagcd dbcfa ecbagdf bfgcea cbfged fedcb feb gbed | bef fdceg egbd bfe
egbf fdgcab dbage edafc degabf fgd gaefd cbdage afedbcg fg | gdf acebdgf gdf gafdeb
acfegb cabed fbdacg afb gafec gadecbf cbfae efbg bf egdfac | fegb abf cebda egfb
caegfd faebcgd dfeag cg cafbge fbdega gca ceadg caebd fgcd | gc agc gafebc fdega
fbgda dfebg bde eb fgdce dfegabc edcfba egbc cdbgfe faegcd | be eb dcbgfae gadbecf
fbgadc efac ce caedb badfce bce ebadg bdgfce cdbfa fbcagde | ec ebc efac ecb
edagcf ecdga facbeg fcdge ca begad gca bdgfcea cgdebf afdc | adcf fdgaec edcgf acg
feacg cdfb eadfbg adcbfg bf bfg aecgdb fdgceab gcdab bcafg | gfb bcfag fb afceg
fbdgea fcae abgecdf deafg gfceda gec cgdef gbdfc ec aecgbd | caef afec dfcgae ecaf
geadfc afbcegd eadc abdfeg ac agdfe cgefa gac gfebc cgbfad | faedg gabfde ca eadc
dcaeg afc ecgf edbcag fcdeba gcfdea abgfd adgcf cf dcebfga | dfegac fceg facdg afc
fbacged egc ecfgb cegadb efdg eg efbdc bcfga deabfc bfcegd | ceg fgedbca fged eg
bfceag egbcad debgaf dbegf febga edfgabc cdfeb gd fdag gde | egfdcba gaecfb ecdfb dagfbe
eadfgc gc bgdfcae bgcdf fbdeac cadgbf cgd gebfd badfc cagb | gc agcfed efdcga gefadbc
gcbe feagb caegf adfce agc cg egfbca dafbeg acfdgb acdfbge | beagf gcdafb egabf abfegd
abfce dafcbe eg gbe fecgab agec eacgbdf gcbdf fcegb gdbaef | bgedfa geca eg bgfec
bfadg befdgc defgb cafgeb de fced edg adgceb gacedbf bfcge | efcd bgceda de de
efdcba efgcdb efcbd ba bae cbaf geadc afgbde decab bgecdaf | cfba daegbcf dbafec bcedgaf
dfeab eb cgfbaed aecfgd efadgb gcdefb agbe cdbaf bfe edfag | dgeacfb cgdebf be bgae
dcfbea dagfeb dfgceba agfdc gcef gecfad abgcd deafc gdf gf | bgacd cefg efcg dfaceg
egfdb ecbgd fade cdfbega bfaeg gfbacd gdf df becfag dfabge | dgf geadbcf gacefdb df
fcdbga egdabf ca bcdfa cdbef befdgac gfbda acgf adc bdgeac | aebdgc gadbfce cda dagfb
dgce adfbecg eafcdb cfgda cfage gbafd ceafgb fcd dc aedcgf | egdcfba fcd fcadg gdbcfae
dae dacgfe fceab fcabge cbdeg daecb ad bcadef dbfa gefcdab | geacdfb egacfb fdab da
gcebf cdebaf gafc agbecf daebgc gcfabde cg ecg efacb edbgf | cgafdeb cegbda cg dcgeba
befcag cefdga ebaf ba bgfcdae fcaeg ecbdg cab fdbacg aecgb | bafe cfgbda geacb acebg`