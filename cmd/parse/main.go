package main

import (
    "bytes"
    "encoding/csv"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
    "reflect"
    "strings"
    "github.com/clbanning/x2j"
    "golang.org/x/net/html"
)

type node struct {
    node interface{} `xml:",any"`
    list []interface{} `xml:",any"`
    value interface{} `xml:",any"`
}

func wrapObj(in interface{}) string {
    out := parseObj(in)
    out = strings.Replace(out, "\\", "\\\\", -1)
    out = strings.Replace(out, "\"", "\\\"", -1)
    out = strings.Replace(out, "\n", "\\n", -1)
    out = strings.Replace(out, "$", "\\$", -1)
    out = fmt.Sprintf("\"%s\"", out)

    return out
}

func parseObj(in interface{}) (out string) {

    if in == nil {
        return ""
    }

    val := reflect.ValueOf(in)

    switch val.Kind() {
    case reflect.Map:
        vv := in.(map[string]interface{})

        parts := make([]string, len(vv))

        i := 0

        for key, value := range vv {
            parts[i] = fmt.Sprintf("[%s]=%s", key, wrapObj(value))
            i++
        }

        out = fmt.Sprintf("(%s)", strings.Join(parts, " "))
    case reflect.Array, reflect.Slice:
        parts := make([]string, val.Len())

        i := 0

        for index := 0; index < val.Len(); index++ {
            value := val.Index(index).Interface()

            parts[i] = fmt.Sprintf("[%d]=%s", index, wrapObj(value))
            i++
        }

        out = fmt.Sprintf("(%s)", strings.Join(parts, " "))
    default:
        out = fmt.Sprint(in)
    }

    return
}

func tryJSON(input []byte) {
    var in interface{}

    err := json.Unmarshal(input, &in)

    if err == nil {
        fmt.Println(parseObj(in))
        os.Exit(0)
    }
}

func tryXML(input []byte) {
    in := make(map[string]interface{})

    err := x2j.Unmarshal(input, &in)

    if err == nil {
        fmt.Println(parseObj(in))
        os.Exit(0)
    }
}

func tryCSV(input []byte) {
    in, err := csv.NewReader(bytes.NewReader(input)).ReadAll()

    if err == nil {
        fmt.Println(parseObj(in))
        os.Exit(0)
    }
}

func formatHTML(n *html.Node) map[string]interface{} {
    out := make(map[string]interface{})

    for _, a := range n.Attr {
        out[fmt.Sprintf("-%s", a.Key)] = a.Val
    }

    for c := n.FirstChild; c != nil; c = c.NextSibling {
        if c.Type == html.TextNode {
            text := strings.TrimSpace(c.Data)

            if text != "" {
                out["#text"] = c.Data
            }
        } else {
            // FIXME - Deal with multiples of the same node type
            out[c.Data] = formatHTML(c)
        }
    }

    return out
}

func tryHTML(input []byte) {
    doc, err := html.Parse(bytes.NewReader(input))

    if err == nil {
        in := formatHTML(doc)
        fmt.Println(parseObj(in))
        os.Exit(0)
    }
}

func main() {
    input, _ := ioutil.ReadAll(os.Stdin)

    tryJSON(input)
    tryXML(input)
    tryCSV(input)
    tryHTML(input)

    fmt.Fprintln(os.Stderr, "Input could not be parsed")
    os.Exit(1)
}
