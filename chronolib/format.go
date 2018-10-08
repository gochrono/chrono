package chronolib

import (
    "bytes"
    "text/template"
    "strings"
    "fmt"
    "github.com/fatih/color"
    "time"
    humanize "github.com/dustin/go-humanize"
)


var cyan = color.New(color.FgCyan).SprintFunc()
var green = color.New(color.FgGreen).SprintFunc()
var blue = color.New(color.FgBlue).SprintFunc()
var magenta = color.New(color.FgMagenta).SprintFunc()
var boldWhite = color.New(color.FgWhite).Add(color.Bold).SprintFunc()

func FormatDateHeader(date time.Time) string {
    return cyan(date.Format("Tuesday _2 January 2006"))
}

func FormatFrameLine(frame Frame) string {
    tags := ""
    shorthex := GetShortHex(frame.UUID)
    start := frame.StartedAt.Format("15:04")
    end := frame.EndedAt.Format("15:04")
    if len(frame.Tags) != 0 {
        tags = FormatTags(frame.Tags)
    }
    hours, minutes, seconds := GetTimeElapsed(frame.StartedAt, frame.EndedAt)

    return fmt.Sprintf("\t(ID: %s) %s to %s %4dh %02dm %02ds  %-12s%s", shorthex, green(start), green(end), hours, minutes, seconds, magenta(frame.Project), blue(tags))
}


func FormatNoteLine(note string) string {
    return fmt.Sprintf("\t\t- %s", note)
}


func FormatNewFrameMessage(frame Frame) string {
    startTime := frame.StartedAt.Format("15:02")
    tags := ""
    if len(frame.Tags) != 0 {
        tags = FormatTags(frame.Tags)
    }
    return fmt.Sprintf("Starting project %s%s at %s", magenta(frame.Project), tags,  green(startTime))
}

func FormatEditFrameMessage(frame Frame) string {
    project := magenta(frame.Project)
    started := green(frame.StartedAt.Format("15:04:05"))
    ended := green(frame.EndedAt.Format("15:04:05"))
    hours, minutes, seconds := GetTimeElapsed(frame.StartedAt, frame.EndedAt)
    return fmt.Sprintf("Edited frame for project %s, from %s to %s (%dh %02dm %02ds)", project, started, ended, hours, minutes, seconds)
}

func FormatTags(tags []string) string {
    for index, tag := range tags {
        tags[index] = blue(tag)
    }
    return " [" + strings.Join(tags, ", ") + "]"
}

func FormatStopFrameMessage(frame Frame) string {
    project := magenta(frame.Project)
    started := green(humanize.Time(frame.StartedAt))
    tags := ""
    if len(frame.Tags) != 0 {
        tags = FormatTags(frame.Tags)
    }
    uuid := boldWhite(GetShortHex(frame.UUID))
    return fmt.Sprintf("Stopping project %s%s, started %s (id: %s)", project, tags, started, uuid)
}

func FormatStatusFrameMessage(frame Frame) string {
    project := magenta(frame.Project)
    started := green(humanize.Time(frame.StartedAt))
    tags := ""
    if len(frame.Tags) != 0 {
        tags = FormatTags(frame.Tags)
    }
    return fmt.Sprintf("Project %s%s started %s.", project, tags, started)
}

type colorFormat func(string) string

var FuncMap = template.FuncMap{
        "cyan": func(input string) string {
            return cyan(input)
        },
        "magenta": func(input string) string {
            return magenta(input)
        },
        "green": func(input string) string {
            return green(input)
        },
        "blue": func(input string) string {
            return blue(input)
        },
        "boldWhite": func(input string) string {
            return boldWhite(input)
        },
        "joinTags": func(input []string) string {
            return strings.Join(input, ", ")
        },
        "humanize": func(t time.Time) string {
            return humanize.Time(t)
        },
    }

func RenderStatusFormatString(frame Frame, format string) string {
    tmpl := template.New("format")
    tmpl.Funcs(FuncMap)
    tmpl, err := tmpl.Parse(format)
    if err != nil {
        panic(err)
    }

    var tpl bytes.Buffer
    if err := tmpl.Execute(&tpl, frame); err != nil {
        panic(err)
    }

    return tpl.String()
}
