package please

import (
    "fmt"
    "os"
    "reflect"
    "strconv"
    "strings"
)

func FormatYAML(in interface{}, path string) (out string) {
    return formatYAML(in, path, 0)[1:]
}

func formatYAML(in interface{}, path string, level int) (out string) {
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

    indent := "\n" + strings.Repeat("  ", level)

    switch val.Kind() {
    case reflect.Map:
        vv := in.(map[string]interface{})

        if this_path != "" {
            if _, ok := vv[this_path]; !ok {
                fmt.Fprintf(os.Stderr, "Key does not exist: %s\n", this_path)
                os.Exit(1)
            }

            return formatYAML(vv[this_path], next_path, level)
        }

        parts := make([]string, len(vv))

        i := 0

        for key, value := range vv {
            parts[i] = fmt.Sprintf("%s'%s': %s", indent, key, formatYAML(value, next_path, level + 1))
            i++
        }

        return fmt.Sprint(strings.Join(parts, ""))
    case reflect.Array, reflect.Slice:
        if this_path != "" {
            index, err := strconv.Atoi(this_path)

            if err != nil || index < 0 || index >= val.Len() {
                fmt.Fprintf(os.Stderr, "Key does not exist: %s\n", this_path)
                os.Exit(1)
            }

            return formatYAML(val.Index(index).Interface(), next_path, level)
        }

        parts := make([]string, val.Len())

        i := 0

        for index := 0; index < val.Len(); index++ {
            value := val.Index(index).Interface()

            parts[i] = fmt.Sprintf("%s- %s", indent, formatYAML(value, next_path, level + 1))
            i++
        }

        return fmt.Sprint(strings.Join(parts, ""))
    default:
        if this_path != "" {
            fmt.Fprintf(os.Stderr, "Key does not exist: %s\n", this_path)
            os.Exit(1)
        }

        out := fmt.Sprint(in)

        out = strings.Replace(out, "\\", "\\\\", -1)
        out = strings.Replace(out, "\"", "\\\"", -1)
        out = strings.Replace(out, "\n", "\\n", -1)
        out = fmt.Sprintf("\"%s\"", out)

        return out
    }
}
