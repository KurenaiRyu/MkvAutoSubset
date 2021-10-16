package mkvlib

import (
	"log"
	"os/exec"
)

const libName = "mkvlib"
const libVer = "1.0.0"

const LibFName = libName + " " + libVer

var _instance *mkvProcessor

func GetInstance() *mkvProcessor {
	ec := 0
	_, _ttx := exec.LookPath(ttx)
	_, _pyftsubset := exec.LookPath(pyftsubset)
	_, _mkvextract := exec.LookPath(mkvextract)
	_, _mkvmerge := exec.LookPath(mkvmerge)
	if _ttx != nil || _pyftsubset != nil {
		log.Printf(`Missing dependency: fonttools (need "%s" & "%s").`, ttx, pyftsubset)
		ec++
	}
	if _mkvextract != nil || _mkvmerge != nil {
		log.Printf(`Missing dependency: mkvtoolnix (need "%s" & "%s").`, mkvextract, mkvmerge)
		ec++
	}
	if ec > 0 {
		return nil
	}
	if _instance == nil {
		_instance = new(mkvProcessor)
	}
	return _instance
}
