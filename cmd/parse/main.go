package main

import (
    "bytes"
    "encoding/csv"
    "encoding/json"
    "fmt"
    "github.com/clbanning/x2j"
    "golang.org/x/net/html"
    "io/ioutil"
    "os"
    "reflect"
    "strconv"
    "strings"
)

type node struct {
    node interface{} `xml:",any"`
    list []interface{} `xml:",any"`
    value interface{} `xml:",any"`
}

func wrapObj(in interface{}, path string) string {
    out := parseObj(in, path)
    out = strings.Replace(out, "\\", "\\\\", -1)
    out = strings.Replace(out, "\"", "\\\"", -1)
    out = strings.Replace(out, "\n", "\\n", -1)
    out = strings.Replace(out, "$", "\\$", -1)
    out = fmt.Sprintf("\"%s\"", out)

    return out
}

func parseObj(in interface{}, path string) (out string) {

    if in == nil {
        return ""
    }

    val := reflect.ValueOf(in)

    split_path := strings.SplitN(path, ".", 2)

    this_path := split_path[0]
    var next_path string

    if len(split_path) > 1{
        next_path = split_path[1]
    }

    switch val.Kind() {
    case reflect.Map:
        vv := in.(map[string]interface{})

        if this_path != "" {
            if _, ok := vv[this_path]; !ok {
                fmt.Fprintf(os.Stderr, "Key does not exist: %s\n", this_path)
                os.Exit(1)
            }

            return parseObj(vv[this_path], next_path)
        }

        parts := make([]string, len(vv))

        i := 0

        for key, value := range vv {
            parts[i] = fmt.Sprintf("[%s]=%s", key, wrapObj(value, next_path))
            i++
        }

        return fmt.Sprintf("(%s)", strings.Join(parts, " "))
    case reflect.Array, reflect.Slice:
        if this_path != "" {
            index, err := strconv.Atoi(this_path)

            if err != nil || index < 0 || index >= val.Len() {
                fmt.Fprintf(os.Stderr, "Key does not exist: %s\n", this_path)
                os.Exit(1)
            }

            return parseObj(val.Index(index).Interface(), next_path)
        }

        parts := make([]string, val.Len())

        i := 0

        for index := 0; index < val.Len(); index++ {
            value := val.Index(index).Interface()

            parts[i] = fmt.Sprintf("[%d]=%s", index, wrapObj(value, next_path))
            i++
        }

        return fmt.Sprintf("(%s)", strings.Join(parts, " "))
    default:
        if this_path != "" {
            fmt.Fprintf(os.Stderr, "Key does not exist: %s\n", this_path)
            os.Exit(1)
        }

        return fmt.Sprint(in)
    }
}

func tryJSON(input []byte, path string) {
    var in interface{}

    err := json.Unmarshal(input, &in)

    if err == nil {
        fmt.Println(parseObj(in, path))
        os.Exit(0)
    }
}

func tryXML(input []byte, path string) {
    in := make(map[string]interface{})

    err := x2j.Unmarshal(input, &in)

    if err == nil {
        fmt.Println(parseObj(in, path))
        os.Exit(0)
    }
}

func tryCSV(input []byte, path string) {
    in, err := csv.NewReader(bytes.NewReader(input)).ReadAll()

    if err == nil {
        fmt.Println(parseObj(in, path))
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

func tryHTML(input []byte, path string) {
    doc, err := html.Parse(bytes.NewReader(input))

    if err == nil {
        in := formatHTML(doc)
        fmt.Println(parseObj(in, path))
        os.Exit(0)
    }
}

func main() {
    input, _ := ioutil.ReadAll(os.Stdin)
    var path string

    if len(os.Args) > 1 {
        path = os.Args[1]
    }

    tryJSON(input, path)
    tryXML(input, path)
    tryCSV(input, path)
    tryHTML(input, path)

    fmt.Fprintln(os.Stderr, "Input could not be parsed")
    os.Exit(1)
}
