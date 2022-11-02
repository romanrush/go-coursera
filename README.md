## Course 1

### Module 4 Protocols and Formats

#### Objectives:

- Describe remote function calls (RFCs).
- Describe JavaScript Object Notation (JSON).
- Describe methods for accessing data in files.
- Write and implement a Go program that employs RFCs and JSON.
- Write and implement a Go program that reads structured text data from a file and then performs some - manipulations with those data.

assessments: makejson.go and read.go

_Notes:_
HTML is a protocol, a standardized language that all clients and servers understand.
URI as well.
HTTP also.

> > > In go you have protocol packages providing functions to help encode and decode data in the particular format. (ex: net/http or tcp/ip)

JSON.
JSON is a commonly used data format. It represents data as structured information as key-value pairs (struct or map in golang). ex of JSON object> {"name":"joe", "phone":"123"}
JSON Properties: all unicode, human readable, fairly compact, types can be combined recursively (array of structs etc.)
Marshalling:
p1 being a go struct.

```
p1:= Person(name:"Joe", phone:"123")
barr, err := json.Marshal(p1)
```

Unmarshalling:

```
var p2 Person
arr := json.Unmarshall(barr, &p2)
```

unpacked the barr into p2
p2 got to have the same attribute as barr.
It must "fit" the JSON.

Files.
Open, read, Write, Close, Seek are basic operations with files.
**io/ioutil** package has some basic functions.
ioutil.ReadFile and WriteFile take care of open & close.
**os** package file access. more granular.
os.Open() return a File struct.
os.Close()
os.Read() reads from file to byte array but you can control the amount
os.Write() behave same as os.Read()
The below code will read and fill barr, read returns the # of bytes read.

```
f, err := os.Open("file.txt")
barr := make([]byte, 10)
nb, err := f.Read(barr)
f.Close()
```

**Glossary:**
RFC request for comments
JSON Marshalling : generate a JSON representation of an object (ex: from a go struct object)
Unmarshalling : from a JSON object generate a go Object.
