package myconf

type Myconf struct {
	AddSfile string `translate:"IP"`
}

var myconf *Myconf = new(Myconf)

func ParseConf(filename string) {
}
