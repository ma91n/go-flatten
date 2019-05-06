package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"testing"
)

func TestJSONHasObjectiveElement(t *testing.T) {
	log.Println("start", t.Name())

	input := `{"A":123, "B":{"X":456}}`

	// change working directory
	prev, _ := filepath.Abs(".")
	defer func() {
		_ = os.Chdir(prev)
	}()
	dir, _ := os.Getwd()
	if err := os.Chdir(dir); err != nil {
		t.Fatal(err)
	}

	fmt.Println(input)

	c1 := exec.Command("cmd", "/c", "echo", input)
	c2 := exec.Command("go", "run", "main.go")

	r, w := io.Pipe()
	c1.Stdout = w
	c2.Stdin = r

	var out bytes.Buffer
	var stderr bytes.Buffer
	c2.Stdout = &out
	c2.Stderr = &stderr

	if err := c1.Start(); err != nil {
		t.Fatal(err)
	}

	if err := c2.Start(); err != nil {
		t.Fatal(err)
	}

	if err := c1.Wait(); err != nil {
		t.Fatal(err)
	}

	if err := w.Close(); err != nil {
		t.Fatal(err)
	}

	if err := c2.Wait(); err != nil {
		fmt.Println(string(stderr.Bytes()))
		t.Fatal(err)
	}

	//out, err := exec.Command("echo " + input + " | go version").CombinedOutput()
	//if err != nil {
	//	t.Fatal(err)
	//}

	fmt.Println(input, string(out.Bytes()))

	if err := mapEqual([]byte(input), out.Bytes()); err != nil {
		t.Fatal(err)
	}
	log.Println("complete", t.Name())
}

func mapEqual(expected, actual []byte) error {

	expectedMap := map[string]interface{}{}
	if err := json.Unmarshal(expected, &expectedMap); err != nil {
		return err
	}

	actualMap := map[string]interface{}{}
	if err := json.Unmarshal(actual, &actualMap); err != nil {
		return err
	}

	eq := reflect.DeepEqual(expectedMap, actualMap)
	if eq {
		return nil
	}
	return errors.New(fmt.Sprintf("unmatch expected %v and actual %v\n", expectedMap, actualMap))
}
