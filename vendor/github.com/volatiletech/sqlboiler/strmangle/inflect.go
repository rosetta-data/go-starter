package strmangle

import "github.com/volatiletech/inflect"

var boilRuleset *inflect.Ruleset

// create a new ruleset and load it with the default
// set of common English pluralization rules
func newBoilRuleset() *inflect.Ruleset {
	rs := inflect.NewRuleset()
	rs.AddPlural("s", "s")
	rs.AddPlural("testis", "testes")
	rs.AddPlural("axis", "axes")
	rs.AddPlural("octopus", "octopi")
	rs.AddPlural("virus", "viri")
	rs.AddPlural("octopi", "octopi")
	rs.AddPlural("viri", "viri")
	rs.AddPlural("alias", "aliases")
	rs.AddPlural("status", "statuses")
	rs.AddPlural("bus", "buses")
	rs.AddPlural("buffalo", "buffaloes")
	rs.AddPlural("tomato", "tomatoes")
	rs.AddPlural("tum", "ta")
	rs.AddPlural("ium", "ia")
	rs.AddPlural("ta", "ta")
	rs.AddPlural("ia", "ia")
	rs.AddPlural("sis", "ses")
	rs.AddPlural("lf", "lves")
	rs.AddPlural("rf", "rves")
	rs.AddPlural("afe", "aves")
	rs.AddPlural("bfe", "bves")
	rs.AddPlural("cfe", "cves")
	rs.AddPlural("dfe", "dves")
	rs.AddPlural("efe", "eves")
	rs.AddPlural("gfe", "gves")
	rs.AddPlural("hfe", "hves")
	rs.AddPlural("ife", "ives")
	rs.AddPlural("jfe", "jves")
	rs.AddPlural("kfe", "kves")
	rs.AddPlural("lfe", "lves")
	rs.AddPlural("mfe", "mves")
	rs.AddPlural("nfe", "nves")
	rs.AddPlural("ofe", "oves")
	rs.AddPlural("pfe", "pves")
	rs.AddPlural("qfe", "qves")
	rs.AddPlural("rfe", "rves")
	rs.AddPlural("sfe", "sves")
	rs.AddPlural("tfe", "tves")
	rs.AddPlural("ufe", "uves")
	rs.AddPlural("vfe", "vves")
	rs.AddPlural("wfe", "wves")
	rs.AddPlural("xfe", "xves")
	rs.AddPlural("yfe", "yves")
	rs.AddPlural("zfe", "zves")
	rs.AddPlural("hive", "hives")
	rs.AddPlural("quy", "quies")
	rs.AddPlural("by", "bies")
	rs.AddPlural("cy", "cies")
	rs.AddPlural("dy", "dies")
	rs.AddPlural("fy", "fies")
	rs.AddPlural("gy", "gies")
	rs.AddPlural("hy", "hies")
	rs.AddPlural("jy", "jies")
	rs.AddPlural("ky", "kies")
	rs.AddPlural("ly", "lies")
	rs.AddPlural("my", "mies")
	rs.AddPlural("ny", "nies")
	rs.AddPlural("py", "pies")
	rs.AddPlural("qy", "qies")
	rs.AddPlural("ry", "ries")
	rs.AddPlural("sy", "sies")
	rs.AddPlural("ty", "ties")
	rs.AddPlural("vy", "vies")
	rs.AddPlural("wy", "wies")
	rs.AddPlural("xy", "xies")
	rs.AddPlural("zy", "zies")
	rs.AddPlural("x", "xes")
	rs.AddPlural("ch", "ches")
	rs.AddPlural("ss", "sses")
	rs.AddPlural("sh", "shes")
	rs.AddPlural("matrix", "matrices")
	rs.AddPlural("vertix", "vertices")
	rs.AddPlural("indix", "indices")
	rs.AddPlural("matrex", "matrices")
	rs.AddPlural("vertex", "vertices")
	rs.AddPlural("index", "indices")
	rs.AddPlural("mouse", "mice")
	rs.AddPlural("louse", "lice")
	rs.AddPlural("mice", "mice")
	rs.AddPlural("lice", "lice")
	rs.AddPluralExact("ox", "oxen", true)
	rs.AddPluralExact("oxen", "oxen", true)
	rs.AddPluralExact("quiz", "quizzes", true)
	rs.AddSingular("s", "")
	rs.AddSingular("ss", "ss")
	rs.AddSingular("as", "as")
	rs.AddSingular("us", "us")
	rs.AddSingular("is", "is")
	rs.AddSingular("news", "news")
	rs.AddSingular("ta", "tum")
	rs.AddSingular("ia", "ium")
	rs.AddSingular("analyses", "analysis")
	rs.AddSingular("bases", "basis")
	rs.AddSingular("diagnoses", "diagnosis")
	rs.AddSingular("parentheses", "parenthesis")
	rs.AddSingular("prognoses", "prognosis")
	rs.AddSingular("synopses", "synopsis")
	rs.AddSingular("theses", "thesis")
	rs.AddSingular("analyses", "analysis")
	rs.AddSingular("aves", "afe")
	rs.AddSingular("bves", "bfe")
	rs.AddSingular("cves", "cfe")
	rs.AddSingular("dves", "dfe")
	rs.AddSingular("eves", "efe")
	rs.AddSingular("gves", "gfe")
	rs.AddSingular("hves", "hfe")
	rs.AddSingular("ives", "ife")
	rs.AddSingular("jves", "jfe")
	rs.AddSingular("kves", "kfe")
	rs.AddSingular("lves", "lfe")
	rs.AddSingular("mves", "mfe")
	rs.AddSingular("nves", "nfe")
	rs.AddSingular("oves", "ofe")
	rs.AddSingular("pves", "pfe")
	rs.AddSingular("qves", "qfe")
	rs.AddSingular("rves", "rfe")
	rs.AddSingular("sves", "sfe")
	rs.AddSingular("tves", "tfe")
	rs.AddSingular("uves", "ufe")
	rs.AddSingular("vves", "vfe")
	rs.AddSingular("wves", "wfe")
	rs.AddSingular("xves", "xfe")
	rs.AddSingular("yves", "yfe")
	rs.AddSingular("zves", "zfe")
	rs.AddSingular("hives", "hive")
	rs.AddSingular("tives", "tive")
	rs.AddSingular("lves", "lf")
	rs.AddSingular("rves", "rf")
	rs.AddSingular("quies", "quy")
	rs.AddSingular("bies", "by")
	rs.AddSingular("cies", "cy")
	rs.AddSingular("dies", "dy")
	rs.AddSingular("fies", "fy")
	rs.AddSingular("gies", "gy")
	rs.AddSingular("hies", "hy")
	rs.AddSingular("jies", "jy")
	rs.AddSingular("kies", "ky")
	rs.AddSingular("lies", "ly")
	rs.AddSingular("mies", "my")
	rs.AddSingular("nies", "ny")
	rs.AddSingular("pies", "py")
	rs.AddSingular("qies", "qy")
	rs.AddSingular("ries", "ry")
	rs.AddSingular("sies", "sy")
	rs.AddSingular("ties", "ty")
	rs.AddSingular("vies", "vy")
	rs.AddSingular("wies", "wy")
	rs.AddSingular("xies", "xy")
	rs.AddSingular("zies", "zy")
	rs.AddSingular("series", "series")
	rs.AddSingular("movies", "movie")
	rs.AddSingular("xes", "x")
	rs.AddSingular("ches", "ch")
	rs.AddSingular("sses", "ss")
	rs.AddSingular("shes", "sh")
	rs.AddSingular("mice", "mouse")
	rs.AddSingular("lice", "louse")
	rs.AddSingular("buses", "bus")
	rs.AddSingular("oes", "o")
	rs.AddSingular("shoes", "shoe")
	rs.AddSingular("crises", "crisis")
	rs.AddSingular("axes", "axis")
	rs.AddSingular("testes", "testis")
	rs.AddSingular("octopi", "octopus")
	rs.AddSingular("viri", "virus")
	rs.AddSingular("statuses", "status")
	rs.AddSingular("aliases", "alias")
	rs.AddSingularExact("oxen", "ox", true)
	rs.AddSingular("vertices", "vertex")
	rs.AddSingular("indices", "index")
	rs.AddSingular("matrices", "matrix")
	rs.AddSingularExact("quizzes", "quiz", true)
	rs.AddSingular("databases", "database")
	rs.AddSingular("menus", "menu")
	rs.AddIrregular("person", "people")
	rs.AddIrregular("man", "men")
	rs.AddIrregular("child", "children")
	rs.AddIrregular("sex", "sexes")
	rs.AddIrregular("move", "moves")
	rs.AddIrregular("zombie", "zombies")
	rs.AddIrregular("cookie", "cookies")
	rs.AddSingularExact("a", "a", true)
	rs.AddSingularExact("i", "i", true)
	rs.AddSingularExact("is", "is", true)
	rs.AddSingularExact("us", "us", true)
	rs.AddSingularExact("as", "as", true)
	rs.AddSingularExact("areas", "area", true)
	rs.AddPluralExact("a", "a", true)
	rs.AddPluralExact("i", "i", true)
	rs.AddPluralExact("is", "is", true)
	rs.AddPluralExact("us", "us", true)
	rs.AddPluralExact("as", "as", true)
	return rs
}
