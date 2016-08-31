# truecolor

Output with true color in temrinal. (require true color support temrinals)

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
