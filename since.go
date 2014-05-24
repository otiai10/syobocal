package animapi

import "time"
import "errors"
import "regexp"
import "strconv"
import "strings"

var exp_whole = regexp.MustCompile("(-*)([0-9\\.]+)([a-z]{1})")
var exp_units = regexp.MustCompile("[w|d|h]")
var build_map = map[string]func(string) (string, string){
	"w": func(val string) (string, string) {
		i, _ := strconv.Atoi(val)
		return strconv.Itoa(i * 7 * 24), "h"
	},
	"d": func(val string) (string, string) {
		i, _ := strconv.Atoi(val)
		return strconv.Itoa(i * 24), "h"
	},
	"h": func(val string) (string, string) {
		return val, "h"
	},
}

type since struct {
	val     string
	matches []string
}

func (s *since) Parse() (dur time.Duration, e error) {
	if e = s.validate(); e != nil {
		return
	}
	s.convert()
	dur, e = s.build()
	return
}
func (s *since) validate() error {
	if e := s.validateWhole(); e != nil {
		return e
	}
	if e := s.validateUnits(); e != nil {
		return e
	}
	return nil
}
func (s *since) validateWhole() error {
	if !exp_whole.MatchString(s.val) {
		return errors.New("Given `since` format is not valid")
	}
	s.matches = exp_whole.FindStringSubmatch(s.val)
	return nil
}
func (s *since) validateUnits() error {
	if !exp_units.MatchString(s.matches[len(s.matches)-1]) {
		return errors.New("Allowd units are `w`, `d`, `h`")
	}
	s.matches = s.matches[1:]
	return nil
}
func (s *since) convert() {
	v := len(s.matches) - 2
	u := len(s.matches) - 1
	s.matches[v], s.matches[u] = build_map[s.matches[u]](s.matches[v])
}
func (s *since) build() (time.Duration, error) {
	return time.ParseDuration(strings.Join(s.matches, ""))
}
