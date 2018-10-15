package junitgo

import (
	"encoding/xml"
	"io"
	"os"
)

type Property struct {
	Name  string `xml:"name,attr,omitempty"`
	Value string `xml:"value,attr,omitempty"`
}

type Properties []Property

type Failure struct {
	Message string `xml:"message,attr,omitempty"`
	Type    string `xml:"type,attr,omitempty"`
	Data    string `xml:",chardata"`
}

type TestCase struct {
	ClassName   string       `xml:"classname,attr,omitempty"`
	Name        string       `xml:"name,attr"`
	Time        string       `xml:"time,attr"`
	Failure     *Failure     `xml:"failure"`
	SkipMessage *SkipMessage `xml:"skipped,omitempty"`
	Systemout   string       `xml:"system-out,omitempty"`
	Systemerr   string       `xml:"system-err,omitempty"`
}

type SkipMessage struct {
	Message string `xml:"message,attr"`
}

type TestSuite struct {
	XMLName    xml.Name   `xml:"testsuite"`
	Errors     int        `xml:"errors,attr"`
	Failures   int        `xml:"failures,attr"`
	Hostname   string     `xml:"hostname,attr"`
	Name       string     `xml:"name,attr"`
	Tests      int        `xml:"tests,attr"`
	Time       float64    `xml:"time,attr"`
	Timestamp  string     `xml:"timestamp,attr"`
	Properties Properties `xml:"properties>property,omitempty"`
	Testcases  []TestCase `xml:"testcase"`
	Package    string     `xml:"package"`
	Id         string     `xml:"id"`
}

type TestRun struct {
	XMLName    xml.Name `xml:"testsuites"`
	Name       string   `xml:"name,attr,omitempty"`
	TestSuites []*TestSuite
}

func (tr *TestRun) AddTestSuite(props Properties, hostname string, name string, pkg string, id string) *TestSuite {
	ts := TestSuite{}
	ts.Hostname = hostname
	ts.Name = name
	ts.Package = pkg
	ts.Id = id
	tr.TestSuites = append(tr.TestSuites, &ts)
	ts.Properties = props
	return &ts
}

func (ts *TestSuite) AddTestCase(class string, name string, time string, failure *Failure, stdout string, stderr string) {

	tc := TestCase{}
	tc.ClassName = class
	tc.Name = name
	tc.Time = time
	tc.Systemerr = stderr
	tc.Systemout = stdout
	tc.Failure = failure
	if failure != nil {
		ts.Failures += 1
	}
	ts.Tests += 1
	ts.Testcases = append(ts.Testcases, tc)
}

func (tr TestRun) WriteResults(writer io.Writer) error {
	b, err := xml.MarshalIndent(tr, "", "  ")
	if err != nil {
		return err
	}
	_, err = writer.Write(b)
	if err != nil {
		return err
	}
	return nil
}

func (tr TestRun) WriteToFile(path string) error {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	err = tr.WriteResults(f)
	if err != nil {
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}
	return nil
}
