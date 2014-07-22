package gobuggery

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// Debugger is a gateway to the remote debugger methods
type Debugger struct {
	endpoint string
	buf      bytes.Buffer
}

// Request will be marshalled to JSON to make a request to the remote debugger
type Request struct {
	Method string        `json:"method"`
	Args   []interface{} `json:"args"`
}

// NewDebugger returns a Debugger configured to interact with the given
// endpoint URI
func NewDebugger(ep string) Debugger {
	return Debugger{endpoint: ep}
}

// RunMethod runs a remote method via a POST/JSON API
func (d *Debugger) RunMethod(meth string, args []interface{}) (interface{}, error) {

	d.buf.Reset()
	err := json.NewEncoder(&d.buf).Encode(
		Request{
			Method: meth,
			Args:   args,
		},
	)

	if err != nil {
		return []interface{}{}, err
	}

	resp, err := http.Post(d.endpoint, "application/json", &d.buf)
	if err != nil {
		return []interface{}{}, err
	}

	defer resp.Body.Close()

	var result interface{}
	e := json.NewDecoder(resp.Body).Decode(&result)
	return result, e

}

// Execute runs an arbitrary command string (as if typed into windbg) and
// returns the result
func (d *Debugger) Execute(command string) (string, error) {

	res, err := d.RunMethod("execute", []interface{}{command})
	if err != nil {
		return "", err
	}

	return res.(string), err
}

// AttachLocalKernel tells the remote to attach as a local kernel debugger (
// like lkd )
func (d *Debugger) AttachLocalKernel() (bool, error) {

	res, err := d.RunMethod("attach_local_kernel", []interface{}{})
	if err != nil {
		return false, err
	}
	return res.(bool), nil
}

// WaitForEvent asks the remote to wait for the next debugger event. Blocking.
func (d *Debugger) WaitForEvent(timeout int) error {

	_, err := d.RunMethod("wait_for_event", []interface{}{timeout})
	return err
}
