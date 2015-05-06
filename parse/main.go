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

func main() {
    var in interface{}
    var err error

    input, _ := ioutil.ReadAll(os.Stdin)

    // JSON
    err = json.Unmarshal(input, &in)
    if err == nil {
        fmt.Println(parseObj(in))
        return
    }

    // XML
    xin := make(map[string]interface{})
    err = x2j.Unmarshal(input, &xin)
    if err == nil {
        fmt.Println(parseObj(xin))
        return
    }

    // CSV
    in, err = csv.NewReader(bytes.NewReader(input)).ReadAll()
    if err == nil {
        fmt.Println(parseObj(in))
        return
    }

    fmt.Fprintln(os.Stderr, "Input could not be parsed")
    os.Exit(1)
}
