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
Create 8-Bit colors.
```
//colors.To256(fg,bg int)
c = c.To256(18,15)
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
