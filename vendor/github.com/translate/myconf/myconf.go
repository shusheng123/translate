package myconf

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

type Myconf struct {
	AddSfile string `translate:"IP"`
}

var myconf *Myconf = new(Myconf)

type Gconf struct {
	File string
	Gcf  map[string]map[string]string
}

func NewGconf(filename string) (*Gconf, error) {
	gconf := new(Gconf)
	gconf.File = filename
	gconf.Gcf = make(map[string]map[string]string)

	return gconf, nil

}

func ParseConf(filename string) (*Gconf, error) {

	gconf := new(Gconf)
	gconf.File = filename
	gconf.Gcf = make(map[string]map[string]string)

	fi, err := os.Open(filename)
	if err != nil {
		fmt.Println("open conf file error")
		return nil, nil
	}
	readler := bufio.NewReader(fi)
	mkey_reg := `^\[.*\]$`
	ikey_reg := `^(.*)\=(.*)$`
	mreg := regexp.MustCompile(mkey_reg)
	ireg := regexp.MustCompile(ikey_reg)

	var imap map[string]string

	for {
		line, _, err := readler.ReadLine()
		if err == io.EOF {
			break
		}
		sline := string(line)

		if mreg.MatchString(sline) {
			if len(sline) < 5 || sline[0] == '#' || sline[0] == ';' {
				imap = nil
				continue
			}
			item := sline[1 : len(sline)-1]
			mk := strings.Trim(item, " ")
			imap = make(map[string]string)
			gconf.Gcf[mk] = imap

		} else if ireg.MatchString(sline) {
			if imap == nil {
				continue
			}
			ik, iv, error := getkey(sline)
			if error != nil {
				fmt.Println("line error")
				continue
			}
			imap[ik] = iv
		} else {
			//fmt.Println("sline:", sline)
			//fmt.Println("no match continue")
			continue
		}

	}
	return gconf, nil

}

func getkey(line string) (string, string, error) {
	if len(line) < 2 {
		return "", "", errors.New(fmt.Sprintf("%s len %d error", line, len(line)))
	}
	index := strings.IndexAny(line, "=")
	var ik string
	var iv string
	//like config f=
	if len(line) == 2 {
		//ik = gcf.StripBlank(line[0:index])
		ik = strings.Trim(line[0:index], " ")
		iv = ""
	} else {
		//ik = gcf.StripBlank(line[0:index])
		ik = strings.Trim(line[0:index], " ")
		//iv = gcf.StripBlank(line[index+1:])
		iv = strings.Trim(line[index+1:], " ")
	}
	//注释掉的行
	if ik[0] == '#' || ik[0] == ';' {
		return "", "", nil
	}
	return ik, iv, nil

}
