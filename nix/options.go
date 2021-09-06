package nix

import (
	"log"
	"os/exec"
	"strings"
)

type Options struct {
	cache map[string]Attr
}

func NewOptions() Options {
	return Options{
		cache: make(map[string]Attr),
	}
}

func (o *Options) retrieve(attr string) Attr {
	if value, contains := o.cache[attr]; contains {
		return value
	} else {
		cmd := exec.Command("nixos-option", attr)
		output, err := cmd.Output()
		if err != nil {
			log.Panicf("nixos-option failed: %v", err) // FIXME
		}
		outputString := string(output)
		lines := strings.Split(outputString, "\n")
		if lines[0] == "This attribute set contains:" {
			o.cache[attr] = Attrset{lines[1:]}
		} else {
			o.cache[attr] = Value{outputString}
		}
		return o.cache[attr]
	}
}

func (o *Options) Show(attr string) error {
	return o.retrieve(attr).Show(attr, o)
}
