package templates

var Setter = `
func ({{.Receiver}} {{.Struct}}) {{.SetterMethod}}(val {{.Type}}) {
  {{- if ne .Lock "" }}
  {{.Receiver}}.{{.Lock}}.Lock()
  defer {{.Receiver}}.{{.Lock}}.Unlock()
  {{- end }}
  {{.Receiver}}.{{.Field}} = val
}
`
