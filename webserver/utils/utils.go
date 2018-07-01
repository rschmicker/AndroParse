package utils

import(
	"fmt"
	"os"
)

const InfoMsg string = `
No GET parameters found!

Here's some brief info on how to query:

fields=
     - Return only these fields
     - Example: fields=permissions,Apis,PackageName,Malicious
        - Returns only the fields Permissions, Apis, and PackageName

to=
     - Only return features up to this timestamp
     - Example: to=2018-02-01T10:00:00
	- Returns features parsed up to February 1, 2018 at 10AM UTC

from=
     - Only return features starting from this timestamp forward
     - Example: to=2018-02-01T10:00:00
	- Returns features parsed starting from February 1, 2018 at 10AM UTC

To query for everything in the data set: https://64.251.61.74/all

More detailed descriptions and examples can be seen on our wiki: https://github.com/rschmicker/AndroParse/wiki

`

func GetArg(name string, aMap map[string][]string) (arg string, err error) {
        var ok bool
        var arglist []string
        if arglist, ok = aMap[name]; !ok {
                return "", fmt.Errorf("unable to obtain arg from map with key %s", name)
        }
        return arglist[0], nil
}

func PrintUsage() {
        fmt.Println(`
Syntax:
        >webserver -key <Directory to HTTPS key> -cert <Directory to HTTPS certificate>

Example:
        >webserver -key ./keys/server.key -cert ./keys/server.crt
`)
        os.Exit(1)
}

func StringInSlice(a string, list []string) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}


