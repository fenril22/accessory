package templates

var Getter = `
func ({{.Receiver}} {{.Struct}}) {{.GetterMethod}}() {{.Type}} {
  {{- if ne .Lock "" }}
  {{.Receiver}}.{{.Lock}}.Lock()
  defer {{.Receiver}}.{{.Lock}}.Unlock()
  {{- end }}
  return {{.Receiver}}.{{.Field}}
}`
