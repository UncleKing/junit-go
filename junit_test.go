package main

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestAbsolutelyNothing(t *testing.T) {
	assert.Nil(t, nil, "Just don't FAIL !")
}

func TestValidateXmlOutputEmptyTestSuite(t *testing.T) {
	tr := TestRun{}
	var props Properties
	tr.AddTestSuite(props, "host", "name", "package", "id")
	var b bytes.Buffer

	tr.WriteResults(&b)

	str := b.String()
	//fmt.Println(str)

	expected :=
		`<testsuites>
  <testsuite errors="0" failures="0" hostname="host" name="name" tests="0" time="0" timestamp="">
    <properties></properties>
    <package>package</package>
    <id>id</id>
  </testsuite>
</testsuites>`
	assert.Equal(t, expected, str, "Empty testsuite should match")

}

func TestValidateXmlSingleTest(t *testing.T) {
	tr := TestRun{}
	var props Properties
	ts := tr.AddTestSuite(props, "host", "name", "package", "id")

	ts.AddTestCase("test1", "validity", "time", nil, "a\nb\nc", "")
	var b bytes.Buffer

	tr.WriteResults(&b)

	str := b.String()

	//fmt.Println(str)

	expected :=
		`<testsuites>
  <testsuite errors="0" failures="0" hostname="host" name="name" tests="1" time="0" timestamp="">
    <properties></properties>
    <testcase classname="test1" name="validity" time="time">
      <system-out>a&#xA;b&#xA;c</system-out>
    </testcase>
    <package>package</package>
    <id>id</id>
  </testsuite>
</testsuites>`
	assert.Equal(t, expected, str, "Single tc in testsuite should match")
	assert.Equal(t, 1, ts.Tests, "No of tests should be 1")
}

func TestValidateXmlMultipleTestSuite(t *testing.T) {
	tr := TestRun{}
	var props Properties
	ts1 := tr.AddTestSuite(props, "host", "ts1", "package", "id")
	ts1.AddTestCase("test1", "validity", "time", nil, "a\nb\nc", "")

	ts2 := tr.AddTestSuite(props, "host", "ts2", "package", "id")
	ts2.AddTestCase("test1", "validity", "time", nil, "a\nb\nc", "")

	var b bytes.Buffer
	tr.WriteResults(&b)

	str := b.String()
	expected :=
		`<testsuites>
  <testsuite errors="0" failures="0" hostname="host" name="ts1" tests="1" time="0" timestamp="">
    <properties></properties>
    <testcase classname="test1" name="validity" time="time">
      <system-out>a&#xA;b&#xA;c</system-out>
    </testcase>
    <package>package</package>
    <id>id</id>
  </testsuite>
  <testsuite errors="0" failures="0" hostname="host" name="ts2" tests="1" time="0" timestamp="">
    <properties></properties>
    <testcase classname="test1" name="validity" time="time">
      <system-out>a&#xA;b&#xA;c</system-out>
    </testcase>
    <package>package</package>
    <id>id</id>
  </testsuite>
</testsuites>`
	assert.Equal(t, expected, str, "Single tc in testsuite should match")

	//fmt.Println(str)
}

func TestValidateXml1FailedTestSuite(t *testing.T) {
	tr := TestRun{}
	var props Properties

	ts1Fail := Failure{}
	ts1Fail.Type = "Error"
	ts1Fail.Message = "Never Add 1 with 1"

	ts1 := tr.AddTestSuite(props, "host", "ts1", "package", "id")
	ts1.AddTestCase("test1", "validity", "time", &ts1Fail, "a\nb\nc", "")

	ts2 := tr.AddTestSuite(props, "host", "ts2", "package", "id")
	ts2.AddTestCase("test1", "validity", "time", nil, "a\nb\nc", "")

	var b bytes.Buffer
	tr.WriteResults(&b)

	str := b.String()
	expected :=
		`<testsuites>
  <testsuite errors="0" failures="1" hostname="host" name="ts1" tests="1" time="0" timestamp="">
    <properties></properties>
    <testcase classname="test1" name="validity" time="time">
      <failure message="Never Add 1 with 1" type="Error"></failure>
      <system-out>a&#xA;b&#xA;c</system-out>
    </testcase>
    <package>package</package>
    <id>id</id>
  </testsuite>
  <testsuite errors="0" failures="0" hostname="host" name="ts2" tests="1" time="0" timestamp="">
    <properties></properties>
    <testcase classname="test1" name="validity" time="time">
      <system-out>a&#xA;b&#xA;c</system-out>
    </testcase>
    <package>package</package>
    <id>id</id>
  </testsuite>
</testsuites>`
	assert.Equal(t, expected, str, "testsuite output should match")
	//fmt.Println(str)
}

func TestValidateXmlMoreThan1FailedTestSuite(t *testing.T) {
	tr := TestRun{}
	var props Properties

	ts1Fail := Failure{}
	ts1Fail.Type = "Error"
	ts1Fail.Message = "Never Add 1 with 1"

	ts1 := tr.AddTestSuite(props, "host", "ts1", "package", "id")
	ts1.AddTestCase("test1", "validity", "time", &ts1Fail, "a\nb\nc", "")

	ts2Fail := Failure{}
	ts2Fail.Type = "Error"
	ts2Fail.Message = "Never Add 1 with 1"

	ts2 := tr.AddTestSuite(props, "host", "ts2", "package", "id")
	ts2.AddTestCase("test1", "validity", "time", &ts2Fail, "a\nb\nc", "")

	ts2.AddTestCase("test1", "validity", "time", nil, "a\nb\nc", "")

	var b bytes.Buffer
	tr.WriteResults(&b)

	str := b.String()
	expected :=
		`<testsuites>
  <testsuite errors="0" failures="1" hostname="host" name="ts1" tests="1" time="0" timestamp="">
    <properties></properties>
    <testcase classname="test1" name="validity" time="time">
      <failure message="Never Add 1 with 1" type="Error"></failure>
      <system-out>a&#xA;b&#xA;c</system-out>
    </testcase>
    <package>package</package>
    <id>id</id>
  </testsuite>
  <testsuite errors="0" failures="1" hostname="host" name="ts2" tests="2" time="0" timestamp="">
    <properties></properties>
    <testcase classname="test1" name="validity" time="time">
      <failure message="Never Add 1 with 1" type="Error"></failure>
      <system-out>a&#xA;b&#xA;c</system-out>
    </testcase>
    <testcase classname="test1" name="validity" time="time">
      <system-out>a&#xA;b&#xA;c</system-out>
    </testcase>
    <package>package</package>
    <id>id</id>
  </testsuite>
</testsuites>`
	assert.Equal(t, expected, str, "testsuite output should match")

}

func TestValidateXmlWithProp(t *testing.T) {
	tr := TestRun{}
	var props Properties
	p1 := Property{}
	p1.Name = "GOPATH"
	p1.Value = "somePath"
	props = append(props, p1)

	ts1Fail := Failure{}
	ts1Fail.Type = "Error"
	ts1Fail.Message = "Never Add 1 with 1"

	ts1 := tr.AddTestSuite(props, "host", "ts1", "package", "id")
	ts1.AddTestCase("test1", "validity", "time", &ts1Fail, "a\nb\nc", "")

	ts2Fail := Failure{}
	ts2Fail.Type = "Error"
	ts2Fail.Message = "Never Add 1 with 1"

	ts2 := tr.AddTestSuite(props, "host", "ts2", "package", "id")
	ts2.AddTestCase("test1", "validity", "time", &ts2Fail, "a\nb\nc", "")

	ts2.AddTestCase("test1", "validity", "time", nil, "a\nb\nc", "")

	var b bytes.Buffer
	tr.WriteResults(&b)

	str := b.String()
	expected :=
		`<testsuites>
  <testsuite errors="0" failures="1" hostname="host" name="ts1" tests="1" time="0" timestamp="">
    <properties>
      <property name="GOPATH" value="somePath"></property>
    </properties>
    <testcase classname="test1" name="validity" time="time">
      <failure message="Never Add 1 with 1" type="Error"></failure>
      <system-out>a&#xA;b&#xA;c</system-out>
    </testcase>
    <package>package</package>
    <id>id</id>
  </testsuite>
  <testsuite errors="0" failures="1" hostname="host" name="ts2" tests="2" time="0" timestamp="">
    <properties>
      <property name="GOPATH" value="somePath"></property>
    </properties>
    <testcase classname="test1" name="validity" time="time">
      <failure message="Never Add 1 with 1" type="Error"></failure>
      <system-out>a&#xA;b&#xA;c</system-out>
    </testcase>
    <testcase classname="test1" name="validity" time="time">
      <system-out>a&#xA;b&#xA;c</system-out>
    </testcase>
    <package>package</package>
    <id>id</id>
  </testsuite>
</testsuites>`
	assert.Equal(t, expected, str, "testsuite output should match")

	//fmt.Printf(str)
}


func TestValidateXmlFile(t *testing.T) {
	tr := TestRun{}
	var props Properties
	p1 := Property{}
	p1.Name = "GOPATH"
	p1.Value = "somePath"
	props = append(props, p1)

	ts1Fail := Failure{}
	ts1Fail.Type = "Error"
	ts1Fail.Message = "Never Add 1 with 1"

	ts1 := tr.AddTestSuite(props, "host", "ts1", "package", "id")
	ts1.AddTestCase("test1", "validity", "time", &ts1Fail, "a\nb\nc", "")

	ts2Fail := Failure{}
	ts2Fail.Type = "Error"
	ts2Fail.Message = "Never Add 1 with 1"

	ts2 := tr.AddTestSuite(props, "host", "ts2", "package", "id")
	ts2.AddTestCase("test1", "validity", "time", &ts2Fail, "a\nb\nc", "")

	ts2.AddTestCase("test1", "validity", "time", nil, "a\nb\nc", "")


	err := tr.WriteToFile("out.xml")
	assert.NoError(t, err, "Save to file shouldn't fail")
	b, err := ioutil.ReadFile("out.xml")
	assert.NoError(t, err, "Save to file shouldn't fail")
	assert.NotNil(t, b, "File cannot be nil.")

	str := string(b)
	expected :=
		`<testsuites>
  <testsuite errors="0" failures="1" hostname="host" name="ts1" tests="1" time="0" timestamp="">
    <properties>
      <property name="GOPATH" value="somePath"></property>
    </properties>
    <testcase classname="test1" name="validity" time="time">
      <failure message="Never Add 1 with 1" type="Error"></failure>
      <system-out>a&#xA;b&#xA;c</system-out>
    </testcase>
    <package>package</package>
    <id>id</id>
  </testsuite>
  <testsuite errors="0" failures="1" hostname="host" name="ts2" tests="2" time="0" timestamp="">
    <properties>
      <property name="GOPATH" value="somePath"></property>
    </properties>
    <testcase classname="test1" name="validity" time="time">
      <failure message="Never Add 1 with 1" type="Error"></failure>
      <system-out>a&#xA;b&#xA;c</system-out>
    </testcase>
    <testcase classname="test1" name="validity" time="time">
      <system-out>a&#xA;b&#xA;c</system-out>
    </testcase>
    <package>package</package>
    <id>id</id>
  </testsuite>
</testsuites>`
	assert.Equal(t, expected, str, "testsuite output should match")

}