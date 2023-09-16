# introduction

Convert a srt file to a fcpxml file. And then you can import it into Final Cut Pro. [File->Import->XML]

It use a Motion effect "background-caption" under the `res` directory. 

# build

```bash
## compile
make 

## rebuild and install to /usr/local/bin/
make install

```

# usage

```bash
$ srt2fcpxml
  -fd int
    	Frame rate is currently supported 23.98、24、25、29.97、30、50、59.94、60 (default 25)
  -srt string
    	srt Subtitle files 
  -width int
        width resolution default 1920
  -height int
        high resolution default 1080
```