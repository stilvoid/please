package formatter

import (
    "fmt"
    "os"
    "reflect"
    "strconv"
    "strings"
)

func wrapObj(in interface{}, path string) string {
    out := Bash(in, path)
    out = strings.Replace(out, "\\", "\\\\", -1)
    out = strings.Replace(out, "\"", "\\\"", -1)
    out = strings.Replace(out, "\n", "\\n", -1)
    out = strings.Replace(out, "$", "\\$", -1)
    out = fmt.Sprintf("\"%s\"", out)

    return out
}

func Bash(in interface{}, path string) (out string) {

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

            return Bash(vv[this_path], next_path)
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

            return Bash(val.Index(index).Interface(), next_path)
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
