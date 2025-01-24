package querycompass

import (
    "encoding/json"
    "fmt"
    "net/url"
    "os"
    "strings"
    "io/ioutil"
)

func Run(args []string) {
    if len(args) == 0 {
        fmt.Println("Please provide input file.")
        return
    }

    inputFile := args[0]
    content, err := ioutil.ReadFile(inputFile)
    if err != nil {
        fmt.Println("Error reading input file:", err)
        os.Exit(1)
    }

    lines := strings.Split(string(content), "\n")
    result := make(map[string]map[string][]string)

    for _, line := range lines {
        line = strings.TrimSpace(line)
        if line == "" {
            continue
        }

        parsedURL, err := url.Parse(line)
        if err != nil {
            fmt.Println("Error parsing URL:", err)
            continue
        }

        domain := parsedURL.Host
        if _, exists := result[domain]; !exists {
            result[domain] = make(map[string][]string)
        }

        params := parsedURL.Query()
        for key, values := range params {
            result[domain][key] = append(result[domain][key], values...)
        }
    }

    outputJSON, err := json.MarshalIndent(result, "", "    ")
    if err != nil {
        fmt.Println("Error generating JSON:", err)
        os.Exit(1)
    }

    fmt.Println(string(outputJSON))
}
