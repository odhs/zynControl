package main

import (
    "fmt"
    "log"
    "os"
    "path/filepath"
    // tests
    "html/template"
    "strconv"
    "strings"
    // to minify
    "bufio"
    "bytes"
    "regexp"
)

type Instrument struct {
    Id      int
    Address int
    Name    string
    Path    string
}
type Bank struct {
    Id          int
    Name        string
    Path        string
    Instruments []Instrument
}

var bankPath string = "D:\\PROGRAMAS\\zyn-fusion-windows-64bit-3.0.3-demo\\banks\\"
//var bankPath string = "F:\\DOWNLOADS\\ZynAddSubFX-VDX_VST-2.4.1.505beta\\banks\\"


func ReadyInstrumentsBank() []Bank {

    var directory string
    var banks []Bank
    var b int = 0
    var instrument_id int = 0
    var address int
    var instrument_name string
    var bank_name string
    var skipRootDir bool = true

    err := filepath.Walk(bankPath,
        func(path string, info os.FileInfo, err error) error {
            if err != nil {
                return err
            }

            if info.IsDir() {
                if !skipRootDir {

                    directory = info.Name()
                    // new bank found
                    // resets the instrument counter
                    instrument_id = 0

                    bank_name = strings.ToUpper(strings.Replace(directory, "_", " ", -1))
                    banks = append(banks, Bank{b + 1, bank_name, bankPath + "\\" + directory + "\\", []Instrument{}})
                    b++
                }
                skipRootDir = false

            } else {

                instrument_name = info.Name()

                if instrument_name[len(instrument_name)-4:] == ".xiz" {

                    //fmt.Println(directory+"\\"+instrument_name)
                    instrument_id++

                    // take out '.xiz' of the file name and the numeration '0000'
                    address, err = strconv.Atoi(instrument_name[:4])
                    if err != nil {
                        address = -1
                    }

                    if len(instrument_name) > 9 {
                        instrument_name = instrument_name[5 : len(instrument_name)-4]
                    }

                    instrument_name = strings.Replace(instrument_name, "_", " ", -1)
                    banks[b-1].Instruments = append(banks[b-1].Instruments, Instrument{instrument_id, address, instrument_name, info.Name()})
                }
            }
            return nil
        })
    if err != nil {
        log.Println(err)
    }

    return banks
}

func main() {
    bankc := ReadyInstrumentsBank()

    /*output := strings.Replace(fmt.Sprintf("%#v\n", bankc), ", ", "\n", -1)
      output = strings.Replace(output, "{", "{\n", -1)
      fmt.Println(strings.Replace(output, "}", "}\n", -1))*/

    //banks_v := Banks{Banks: bankc}

    /*
       output := strings.Replace(fmt.Sprintf("%#v\n", banks_v), ", ", "\n", -1)
       output = strings.Replace(output, "{", "{\n", -1)
       fmt.Println(strings.Replace(output, "}", "}\n", -1))*/

    var out bytes.Buffer 

	var tmpl = template.Must(template.New("").Parse(index))
    var tScripts = template.HTML(RemoveJSComments(scripts))
    var tStyles = template.HTML(styles)

    err := tmpl.Execute(&out, map[string]interface{}{
		"Banks": bankc,
		"Scripts": tScripts,
		"Styles": tStyles,
	})
    //err := tmpl.Execute(&out, bankc)
    if err != nil {
        panic(err)
    }

    fmt.Println(MinifyHTML([]byte(out.String())))
    //fmt.Println(out.String())
}

func RemoveSpaceBetweenTags(content []byte) string {
    //text = text.replace/*removeHTMLComments*/(/<!-[\S\s]*?-->/gm, '');
    htmlspc := regexp.MustCompile(`/\>[\s]+\</g`)
    //return htmlspc.ReplaceAll(content, []byte("><"))
    return string(htmlspc.ReplaceAll(content, []byte("><")))
}

func RemoveJSComments(content string) string {
    htmlspc := regexp.MustCompile(`/\/\*[\S\s]*?\*\//gm`)
    return htmlspc.ReplaceAllString(content, "")
}

func RemoveHTMLComments(content []byte) []byte {
    // https://www.google.com/search?q=regex+html+comments
    // http://stackoverflow.com/a/1084759
    htmlcmt := regexp.MustCompile(`<!--[^>]*-->`)
    return htmlcmt.ReplaceAll(content, []byte(""))
}

func MinifyHTML(html []byte) string {
    // read line by line
    minifiedHTML := ""
    scanner := bufio.NewScanner(bytes.NewReader(RemoveHTMLComments(html)))

    for scanner.Scan() {
        // all leading and trailing white space of each line are removed
        lineTrimmed := strings.TrimSpace(scanner.Text())
        minifiedHTML += lineTrimmed
        if len(lineTrimmed) > 0 {
            // in case of following trimmed line:
            // <div id="foo"
            minifiedHTML += " "
        }
    }
    if err := scanner.Err(); err != nil {
        panic(err)
    }

    return minifiedHTML
}


/* DOC */

/*
type FileInfo interface {
        Name() string       // base name of the file
        Size() int64        // length in bytes for regular files; system-dependent for others
        Mode() FileMode     // file mode bits
        ModTime() time.Time // modification time
        IsDir() bool        // abbreviation for Mode().IsDir()
        Sys() interface{}   // underlying data source (can return nil)
}
*/
