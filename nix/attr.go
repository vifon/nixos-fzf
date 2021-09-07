package nix

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

// An abstract Nix attribute, possibly nested.
type Attr interface {
	Browse() error
}


// A Nix scalar value.
type Value struct {
	Path string
	Documentation string
}

func (v Value) Browse() error {
	less := exec.Command("less")
	less.Stdin = strings.NewReader(v.Documentation)
	less.Stderr = os.Stderr
	less.Stdout = os.Stdout
	return less.Run()
}


// A Nix attribute set.
type Attrset struct {
	Path string
	Attrs []string
	cache map[string]Attr
}

func RootAttrset() Attrset {
	a := Attrset{"", []string{}, make(map[string]Attr)}
	return a.GetAttr("").(Attrset)
}

func (a *Attrset) GetAttr(attr string) Attr {
	if value, contains := a.cache[attr]; contains {
		return value
	} else {
		var attrPath string
		if len(a.Path) > 0 {
			attrPath = a.Path + "." + attr
		} else {
			attrPath = attr
		}

		cmd := exec.Command("nixos-option", attrPath)
		output, err := cmd.Output()
		if err != nil {
			log.Panicf("nixos-option failed: %v", err) // FIXME
		}
		outputString := string(output)
		lines := strings.Split(outputString, "\n")
		if lines[0] == "This attribute set contains:" {
			a.cache[attr] = Attrset{attrPath, lines[1:], make(map[string]Attr)}
		} else {
			a.cache[attr] = Value{attrPath, outputString}
		}
		return a.cache[attr]
	}
}

func (a Attrset) Browse() error {
	var prompt string
	if len(a.Path) == 0 {
		prompt = "> "
	} else {
		prompt = a.Path + "."
	}

	for {
		fzf := exec.Command("fzf", "--prompt", prompt, "--expect=.")
		fzf.Stdin = strings.NewReader(strings.Join(a.Attrs, "\n"))
		fzf.Stderr = os.Stderr
		output, err := fzf.Output()
		if err != nil {
			switch v := err.(type) {
			case *exec.ExitError:
				switch v.ExitCode() {
				case 1:
				// No match, do nothing and repeat the prompt.
				default:
					return err
				}
			default:
				return err
			}
		}

		newAttr := strings.Split(string(output), "\n")[1]
		a.GetAttr(newAttr).Browse()
	}
}
