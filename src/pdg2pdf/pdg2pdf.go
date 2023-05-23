package pdg2pdf

import (
	"fmt"
	"github.com/signintech/gopdf"
	"io/ioutil"
	"path/filepath"
	"sort"
	"strings"
)

type Pdg2pdf struct{}

func (Pdg2pdf) Convert(filename string, output string) (string, bool) {
	files, err := ioutil.ReadDir(filename)
	if err != nil {
		panic(err)
		return "", false
	}

	sort.Slice(files, func(i, j int) bool {
		iname := files[i].Name()
		jname := files[j].Name()

		if strings.Contains(iname, "cov001") {
			iname = "a_" + iname
		}
		if strings.Contains(jname, "cov001") {
			jname = "a_" + jname
		}

		if strings.Contains(iname, "bok") {
			iname = "b_" + iname
		}
		if strings.Contains(jname, "bok") {
			jname = "b_" + jname
		}

		if strings.Contains(iname, "leg") {
			iname = "c_" + iname
		}
		if strings.Contains(jname, "leg") {
			jname = "c_" + jname
		}

		if strings.Contains(iname, "fow") {
			iname = "d_" + iname
		}
		if strings.Contains(jname, "fow") {
			jname = "d_" + jname
		}
		if strings.Contains(iname, "!") {
			iname = "e_" + iname
		}
		if strings.Contains(jname, "!") {
			jname = "e_" + jname
		}
		if iname[0] >= '0' && iname[0] <= '9' {
			iname = "f_" + iname
		}
		if jname[0] >= '0' && jname[0] <= '9' {
			jname = "f_" + jname
		}
		if strings.Contains(iname, "cov002") {
			iname = "g_" + iname
		}
		if strings.Contains(jname, "cov002") {
			jname = "g_" + jname
		}
		return iname < jname
	})
	pdf := gopdf.GoPdf{}

	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	for _, file := range files {
		//fmt.Println(file.Name())
		pdf.AddPage()
		pdf.Image(filepath.Join(filename, file.Name()), 0, 0, &gopdf.Rect{W: 595.28, H: 840})
	}
	err = pdf.WritePdf(output)
	if err != nil {
		panic(err)
		return "", false
	}
	fmt.Println(output)
	return output, true
}
