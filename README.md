# Colors
![colors](./colors.png)

Terminal colors ansi escape code write in golang.

### Install
Get package.
```
go get github.com/informeai/colors
```
### Usage
Create new color.
```
c := colors.NewColors()
```
Add foreground and background.
```
c = c.Add(colors.FgGreen,colors.BgWhite)
```
Return formated string colors.
```
s := c.Sprintf("%v", "any text.")
```
Print text formated with colors.
```
_,err := c.Println("more text here :)")
```

made by wellington gadelha :star:
