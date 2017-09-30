package sillypass

import (
	"crypto/rand"
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
)

const bucket = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_.,<>?:!@#$%&*()-+[]^=/'\"\\`"
const bucketLen = len(bucket)

func generate(n int) string {
	if n < 1 {
		return ""
	}

	return makePassword(n)
}

func makePassword(n int) string {
	if n < 1 {
		return ""
	}

	b := make([]byte, n)

	// The slice will now contain `n` random bytes.
	rand.Read(b)

	s := make([]string, n)

	for i, j := range b {
		// Are these conversions considered bad practice?
		s[i] = string(bucket[int(j)%bucketLen])
	}

	return strings.Join(s, "")
}

func usage() {
	fmt.Printf(`Usage of %s:
    -n int
          Length of password (default 12)
`, os.Args[0])
}

/**
 * The default size of the generated password.
 */
func main() {
	nPtr := flag.Int("n", 12, "Length of password")
	flag.Parse()

	args := flag.Args()
	hasArgs := len(args) > 0

	if hasArgs && flag.Arg(0) != "generate" {
		fmt.Printf("%s", errors.New("oo"))
	}

	fmt.Println(generate(*nPtr))

	/*
		if len(os.Args) == 1 {
			usage()
			return
		}
	*/
}
