package chronolib

import (
	"bytes"
	"fmt"
	humanize "github.com/dustin/go-humanize"
	"github.com/fatih/color"
	"strings"
	"text/template"
	"time"
)

type colorFormat func(string) string

var cyan = color.New(color.FgCyan).SprintFunc()
var green = color.New(color.FgGreen).SprintFunc()
var blue = color.New(color.FgBlue).SprintFunc()
var red = color.New(color.FgRed).SprintFunc()
var magenta = color.New(color.FgMagenta).SprintFunc()
var boldWhite = color.New(color.FgWhite).Add(color.Bold).SprintFunc()
var funcMap = template.FuncMap{
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

// FormatDateHeader returns a time formated with Tuesday 2 January 2006
func FormatDateHeader(date time.Time) string {
	return cyan(date.Format("Monday _2 January 2006"))
}

// FormatFrameLine returns a frame line (all the metadata about a frame except for its notes)
func FormatFrameLine(frame Frame) string {
	tags := ""
	shorthex := GetShortHex(frame.UUID)
	start := frame.StartedAt.Format("15:04")
	end := frame.EndedAt.Format("15:04")
	if len(frame.Tags) != 0 {
		tags = FormatTags(frame.Tags)
	}
	hours, minutes, seconds := GetTimeElapsed(frame.StartedAt, frame.EndedAt)

	return fmt.Sprintf("\t(ID: %s) %s to %s %4dh %02dm %02ds  %-12s%s", shorthex, green(start), green(end), hours, minutes, seconds, magenta(frame.Project), tags)
}

// FormatReportProjectTotal returns the total time spent on a project
func FormatReportProjectTotal(project string, total time.Duration) string {
	hours, minutes, seconds := GetTimeElapsedForDuration(total)
	time := green(fmt.Sprintf("%dh %02dm %02ds", hours, minutes, seconds))
	return fmt.Sprintf("\n%s - %s", magenta(project), time)
}

// FormatReportProjectTagTotal return the total time spent on a tag for a project
func FormatReportProjectTagTotal(tag string, total time.Duration) string {
	hours, minutes, seconds := GetTimeElapsedForDuration(total)
	time := green(fmt.Sprintf("%dh %02dm %02ds", hours, minutes, seconds))
	return fmt.Sprintf("\t[%s %s]", blue(tag), time)
}

// FormatNoteLine formats a single frame note
func FormatNoteLine(note string) string {
	return fmt.Sprintf("\t\t- %s", note)
}

// FormatNoteShowLine formats a single frame note for the notes show command
func FormatNoteShowLine(index int, note string) string {
	return fmt.Sprintf("[%s]: %s", cyan(index), note)
}

// FormatStartFrame returns the output when a new frame is created
func FormatStartFrame(frame Frame) string {
	startTime := frame.StartedAt.Format("15:04")
	endTime := frame.StartedAt.Format("15:04")
	tags := ""
	if len(frame.Tags) != 0 {
		tags = FormatTags(frame.Tags)
	}
	return fmt.Sprintf("Added project %s%s, started at %s and ended at %s", magenta(frame.Project), tags, green(startTime), green(endTime))
}

// FormatStartCurrentFrame returns the output when a new frame is created
func FormatStartCurrentFrame(currentFrame CurrentFrame) string {
	startTime := currentFrame.StartedAt.Format("15:04")
	tags := ""
	if len(currentFrame.Tags) != 0 {
		tags = FormatTags(currentFrame.Tags)
	}
	return fmt.Sprintf("Starting project %s%s at %s", magenta(currentFrame.Project), tags, green(startTime))
}

// FormatNewFrameMessage returns the output when a new frame is created
func FormatNewFrameMessage(frame Frame) string {
	startTime := frame.StartedAt.Format("15:04")
	tags := ""
	if len(frame.Tags) != 0 {
		tags = FormatTags(frame.Tags)
	}
	return fmt.Sprintf("Starting project %s%s at %s", magenta(frame.Project), tags, green(startTime))
}

// FormatEditFrameMessage returns the output when a frame is modified
func FormatEditFrameMessage(frame Frame) string {
	project := magenta(frame.Project)
	started := green(frame.StartedAt.Format("15:04:05"))
	ended := green(frame.EndedAt.Format("15:04:05"))
	hours, minutes, seconds := GetTimeElapsed(frame.StartedAt, frame.EndedAt)
	return fmt.Sprintf("Edited frame for project %s, from %s to %s (%dh %02dm %02ds)", project, started, ended, hours, minutes, seconds)
}

// FormatTags joins tags together and color them blue
func FormatTags(tags []string) string {
	for index, tag := range tags {
		tags[index] = blue(tag)
	}
	return " [" + strings.Join(tags, ", ") + "]"
}

// FormatStopFrameMessage returns the output when time tracking for a frame is stopped
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

// FormatStatusFrameMessage returns the output the current status for active frame
func FormatStatusFrameMessage(frame Frame) string {
	project := magenta(frame.Project)
	started := green(humanize.Time(frame.StartedAt))
	tags := ""
	if len(frame.Tags) != 0 {
		tags = FormatTags(frame.Tags)
	}
	return fmt.Sprintf("Project %s%s started %s.", project, tags, started)
}

// FormatNoProjectMessage returns the output when there is no current frame
func FormatNoProjectMessage() string {
	return "No project started"
}

// FormatNoNotesMessage returns the output if there are no notes in the current frame
func FormatNoNotesMessage() string {
	return "No notes have been added"
}

// FormatNoFramesMessage returns the output when there is no current frame
func FormatNoFramesMessage() string {
	return "No time has been logged"
}

// RenderCurrentFrameStatus returns the status output using Go template string
func RenderCurrentFrameStatus(currentFrame CurrentFrame, format string) string {
	tmpl := template.New("format")
	tmpl.Funcs(funcMap)
	tmpl, err := tmpl.Parse(format)
	if err != nil {
		panic(err)
	}

	var tpl bytes.Buffer
	if err := tmpl.Execute(&tpl, currentFrame); err != nil {
		panic(err)
	}

	return tpl.String()
}

// RenderStatusFormatString returns the status output using Go template string
func RenderStatusFormatString(frame Frame, format string) string {
	tmpl := template.New("format")
	tmpl.Funcs(funcMap)
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

// PrettyDate returns the date using format 2 January 2006 15:04
func PrettyDate(t *time.Time) string {
	return t.Format("_2 January 2006 15:04")
}

// FormatFrameDescribe shows information about a frame in a consise manner
func FormatFrameDescribe(frame Frame) string {
	shortHex := GetShortHex(frame.UUID)
	tags := ""
	if len(frame.Tags) != 0 {
		tags = FormatTags(frame.Tags)
	}
	startTime := green(frame.StartedAt.Format("Jan 2 15:04:05"))
	endTime := green(frame.EndedAt.Format("Jan 2 15:04:05"))
	return fmt.Sprintf("(%s) %s%s: %s to %s", shortHex, magenta(frame.Project), tags, startTime, endTime)
}

// FormatCancelMessage shows a message if the current state is stopped but not saved
func FormatCancelMessage(currentFrame CurrentFrame) string {
	cancelTime := time.Now().Format("15:04")
	tags := ""
	if len(currentFrame.Tags) != 0 {
		tags = FormatTags(currentFrame.Tags)
	}
	return fmt.Sprintf("Cancelled project %s%s at %s", magenta(currentFrame.Project), tags, green(cancelTime))
}

// FormatTimeStringNotValid returns a message when a time string was unable to be parsed
func FormatTimeStringNotValid() string {
	return red("Invalid time format")
}

// FormatReportDurationDate returns the date using format Mon 02 January 2006
func FormatReportDurationDate(t time.Time) string {
	return t.Format("Mon 02 January 2006")
}

// FormatStartError returns a message formated red that says a project is already started
func FormatStartError(frame Frame) string {
	return red(fmt.Sprintf("Project %s is already started.", frame.Project))
}

// FormatReportDuration returns the duration currently being viewed in the report command
func FormatReportDuration(timeStart time.Time) string {
	today := time.Now()
	return cyan(fmt.Sprintf("%s -> %s", FormatReportDurationDate(timeStart), FormatReportDurationDate(today)))
}
