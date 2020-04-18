package cases

import (
	"fmt"
	t "time"

	shell "github.com/ipfs/go-ipfs-api"
)

type Case struct {
	Country   string `json:"Country"`
	Total     string `json:"ID"`
	Active    string `json:"Name"`
	Death     string `json:"Salary"`
	Timestamp t.Time `json:"Time"`
}

type Hash string

type User struct {
	Address uint8
}

type Addressable struct {
	CIDs map[Hash]CID `json:"CIDs"`
}

type CID struct {
	Timestamp t.Time `json:"Time"`
}

// Global variable to handle all the IPFS API client calls
var sh *shell.Shell

func main() {
	// Map of the Cases
	s := StoreConfigure()

	// Infura
	sh = ShellConfigure()

	// Inputs
	c := ScanInput()

	// Store
	StoreCase(s, c)
	fmt.Println(s)

	// Push to IPFS
}

func StoreCase(s *Store, c Case) {
	s.Pending[c.Country] = c
}

func ScanInput() Case {
	// scanner := bufio.NewScanner(os.Stdin)
	//
	// fmt.Println("Enter the total number of cases: ")
	// scanner.Scan()
	// total := scanner.Text()
	//
	// fmt.Println("Enter the active number of cases: ")
	// scanner.Scan()
	// active := scanner.Text()
	//
	// fmt.Println("Enter the death number: ")
	// scanner.Scan()
	// death := scanner.Text()

	covid := Case{
		Country:   "Australia",
		Total:     "10",
		Active:    "3",
		Death:     "2",
		Timestamp: t.Now(), // todo: remove t.Now(), instead should be concrete time
	}

	return covid
}
