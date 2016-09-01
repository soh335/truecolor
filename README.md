[![wercker status](https://app.wercker.com/status/acada82918999d19aa45088dc7adcf59/s/master "wercker status")](https://app.wercker.com/project/byKey/acada82918999d19aa45088dc7adcf59)

# truecolor

Output with true color in temrinal. (require true color support temrinals)

![](https://raw.githubusercontent.com/soh335/truecolor/master/_example/screencapture.png)

## INSTALL

```
$ go get github.com/soh335/truecolor
```

## USAGE

```go
tc := truecolor.New()
tc.Add(truecolor.NewBackgrond(color.RGBA{255, 0, 0, 255}))
tc.Add(truecolor.NewForeground(color.RGBA{100, 100, 100, 255}))
tc.Add(truecolor.Italic)
tc.Fprintln(os.Stdout, "truecolor")
```

## LICENSE

* MIT

## SEA ALSO

* https://gist.github.com/XVilka/8346728

## TODO

* document
