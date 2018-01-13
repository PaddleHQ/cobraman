package man

// TroffManTemplate generates a man page with only basic troff macros
const TroffManTemplate = `.TH "{{.CommandPath | dashify | backslashify | upper}}" "{{ .Section }}" "{{.CenterFooter}}" "{{.LeftFooter}}" "{{.CenterHeader}}" 
.\" disable hyphenation
.nh
.\" disable justification (adjust text to left margin only)
.ad l
." This file auto-generated by github.com/rayjohnson/cobra-man
.SH NAME
{{ .CommandPath | dashify | backslashify }}
{{- if .ShortDescription }} - {{ .ShortDescription }}
 {{- end }}
.SH SYNOPSIS
.sp
{{- if .SubCommands }}
{{- range .SubCommands }}
\fB{{ . }}\fR [ flags ]
.br{{ end }}
{{- else }}
\fB{{ .CommandPath }} \fR
{{- range .AllFlags -}}
[{{ if .Shorthand }}\fI{{ print "-" .Shorthand | backslashify }}\fP|{{ end -}}
\fI{{ print "--" .Name | backslashify }}\fP] {{ end }}
{{- if not .NoArgs }}[<args>]{{ end }}
{{- end }}
.SH DESCRIPTION
.PP
{{ .Description | simpleToTroff }}
{{- if .AllFlags }}
.SH OPTIONS
{{ range .AllFlags -}}
.TP
{{ if .Shorthand }}\fB{{ print "-" .Shorthand | backslashify }}\fP, {{ end -}}
\fB{{ print "--" .Name | backslashify }}\fP{{ if not .NoOptDefVal }} =
{{- if .ArgHint }} <{{ .ArgHint }}>{{ else }} {{ .DefValue }}{{ end }}{{ end }}
{{ .Usage | backslashify }}
{{ end }}
{{- end -}}
{{- if .Environment }}
.SH ENVIRONMENT
.PP
{{ .Environment | simpleToTroff }}
{{- end }}
{{- if .Files }}
.SH FILES
.PP
{{ .Files | simpleToTroff }}
{{- end }}
{{- if .Bugs }}
.SH BUGS
.PP
{{ .Bugs | simpleToTroff }}
{{- end }}
{{- if .Examples }}
.SH EXAMPLES
.PP
{{ .Examples | simpleToTroff }}
{{- end }}
.SH AUTHOR
{{- if .Author }}
{{ .Author }}
{{- end }}
.PP
.SM Page auto-generated by rayjohnson/cobra-man and spf13/cobra
{{- if .SeeAlsos }}
.SH SEE ALSO
{{- range .SeeAlsos }}
.BR {{ .CmdPath | dashify | backslashify }} ({{ .Section }})
{{- end }}
{{- end }}
." This file auto-generated by github.com/rayjohnson/cobra-man
`

// MdocManTemplate is a template what will use the mdoc macro package.
const MdocManTemplate = `.\" Man page for {{.CommandPath}}
.Dd {{ .Date.Format "January 2006"}}
{{ if .CenterHeader -}}
.Dt {{.CommandPath | dashify | backslashify | upper}} \&{{ .Section }} "{{.CenterHeader}}" 
{{- else -}}
.Dt {{.CommandPath | dashify | backslashify | upper}} {{ .Section }}
{{- end }}
./" TODO: The Dt macro can take one additonal arg - what does it do?
.Os
." This file auto-generated by github.com/rayjohnson/cobra-man
.Sh NAME
.Nm {{ .CommandPath | dashify | backslashify }}
{{- if .ShortDescription }}
.Nd {{ .ShortDescription }}
{{- end }}
.Sh SYNOPSIS
{{- if .SubCommands }}
{{- range .SubCommands }}
.Nm {{ . }} Op Fl flags Op args
{{- end }}
{{- else }}
.Nm {{ .CommandPath }}
{{- range .AllFlags }}
.Op Fl {{ if .Shorthand }}{{ .Shorthand | backslashify }} | {{ end -}}
{{ print "-" .Name | backslashify }}
{{- end }}
{{ if not .NoArgs }}.Op Fl <args>
{{- end }}
{{- end }}
.Ek
.Sh DESCRIPTION
.Nm
{{ .Description | simpleToMdoc }}
{{- if .AllFlags }}
.Pp
The options are as follows:
.Pp
.Bl -tag -width Ds -compact
{{ range .AllFlags -}}
.Pp
.It {{ if .Shorthand }}Fl {{ .Shorthand | backslashify }}, {{ end -}}
Fl {{ print "-" .Name | backslashify }}
{{- if not .NoOptDefVal }} Ar {{if .ArgHint }} {{ .ArgHint }}{{ else }} {{ .DefValue }}{{ end }}{{ end }}
{{ .Usage | backslashify }}
{{ end }}
.El
{{- end }}
{{- if .Environment }}
.Sh ENVIRONMENT
{{ .Environment | simpleToMdoc }}
{{- end }}
{{- if .Files }}
.Sh FILES
{{ .Files | simpleToMdoc }}
{{- end }}
{{- if .Bugs }}
.Sh BUGS
{{ .Bugs | simpleToMdoc }}
{{- end }}
{{- if .Examples }}
.Sh EXAMPLES
{{ .Examples | simpleToMdoc }}
{{- end }}
.Sh AUTHOR
{{- if .Author }}
{{ .Author }}
{{- end }}
.sp
Page auto-generated by rayjohnson/cobra-man and spf13/cobra
{{- if .SeeAlsos }}
.Sh SEE ALSO
{{- range $index, $element := .SeeAlsos}}
{{- if $index}} ,{{end}}
.Xr {{$element.CmdPath}} {{$element.Section}}
{{- end }}
{{- end }}
." This file auto-generated by github.com/rayjohnson/cobra-man
`
