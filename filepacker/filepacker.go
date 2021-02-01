package filepacker

import (
	"bytes"
	"compress/gzip"
	"encoding/hex"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

type Content struct {
	Tool      string
	Timestamp string
	Package   string
	Func      string
	Data      string
}

const (
	toolName = "filepacker"
)

// https://gist.github.com/alex-ant/aeaaf497055590dacba760af24839b8d
func Pack(infile, outfile, templateFile, timestamp, packageName, funcName string) {
	data, err := readData(infile)
	if err != nil {
		panic(err)
	}

	compressedData, err := compress(data)
	if err != nil {
		panic(err)
	}
	hexifiedData := bytesToHexString(compressedData)

	templateBytes, err := ioutil.ReadFile(templateFile)
	if err != nil {
		panic(err)
	}

	t, err := template.New("test").Parse(string(templateBytes))
	if err != nil {
		panic(err)
	}

	content := Content{toolName, timestamp, packageName, funcName, hexifiedData}

	var buf bytes.Buffer
	err = t.Execute(&buf, content)

	writeData(outfile, buf.String())
}

func Unpack(content string) ([]byte, error) {
	values := strings.Split(content, "\\x")
	joined := strings.Join(values, "")
	data, err := decompress([]byte(joined))
	if err != nil {
		return nil, err
	}
	return data, nil
}

func Noop() {}

func readData(filename string) ([]byte, error) {
	input, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer input.Close()

	fileStats, err := input.Stat()
	if err != nil {
		return nil, err
	}

	fileSize := fileStats.Size()
	data := make([]byte, fileSize)

	bytesRead, err := input.Read(data)
	if err != nil {
		return nil, err
	}

	if int64(bytesRead) != fileSize {
		return nil, fmt.Errorf("Bytes read to filesize mismatch")
	}

	return data, nil
}

func writeData(filename, content string) error {
	output, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer output.Close()
	output.WriteString(content)
	return nil
}

// this is probably feasible in a more efficient way, but hey, it's something
func bytesToHexString(data []byte) string {
	str := ""
	var single []byte = make([]byte, 1, 1)
	for _, b := range data {
		single[0] = b
		str += fmt.Sprintf("\\x%v", hex.EncodeToString(single))
	}
	return str
}

func compress(data []byte) ([]byte, error) {
	var buf bytes.Buffer
	w, err := gzip.NewWriterLevel(&buf, gzip.BestCompression)
	if err != nil {
		return nil, err
	}

	_, err = w.Write(data)
	if err != nil {
		return nil, err
	}

	if err = w.Flush(); err != nil {
		return nil, err
	}

	if err = w.Close(); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func decompress(data []byte) ([]byte, error) {
	buf := bytes.NewBuffer(data)

	var r io.Reader
	r, err := gzip.NewReader(buf)
	if err != nil {
		return nil, err
	}

	var res bytes.Buffer
	_, err = res.ReadFrom(r)
	if err != nil {
		return nil, err
	}

	return res.Bytes(), nil
}
