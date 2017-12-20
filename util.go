package faker

import (
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"../../tr"
)

func replaceSymbols(s string) string {
	alpha := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	str := ""

	for _, v := range alpha {
		if strings.Compare(v, "#") == 0 {
			str += strconv.Itoa(random(10))
		} else if strings.Compare(v, "?") == 0 {
			str += alpha[random(len(alpha))]
		} else if strings.Compare(v, "*") == 0 {
			rn := random(2)
			if rn == 0 {
				str += strconv.Itoa(random(10))
			} else {
				str += alpha[random(len(alpha))]
			}
		} else {
			str += v
		}

	}
	return str
}

func random(i int) int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(i)
}

func randomFloat(i int) float64 {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return float64(r.Intn(i)) + r.Float64()
}

func directoryExists(dir string) bool {
	dir, _ = filepath.Abs(dir)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return false
	}
	return true
}

func getList(n *tr.Engine, q, d string) ([]string, error) {
	w, err := n.Tr(q)
	if err != nil {
		w, err = n.Lang(d).Tr(q)
	}
	if err != nil {
		return nil, err
	}
	return strings.Split(w, "\n"), nil
}

