/*
Package ui provides a very simple set of functions that return Xterm-256color
compatible escape codes for terminal styling
*/
package ui

import (
    "fmt"
    "strconv"
)

const (
    // common
    reset         = "\033[0m"

    // special
    normal        = "\033[0;m"
    bold          = "\033[1;m"
    dim           = "\033[2;m"
    italic        = "\033[3;m"
    underline     = "\033[4;m"
    blink         = "\033[5;m"
    blinkfast     = "\033[6;m"
    reverse       = "\033[7;m"
    hidden        = "\033[8;m"
    strikethrough = "\033[9;m"

    // color
    black         = "\033[48;5;;38;5;0m"
    gray          = "\033[48;5;;38;5;8m"
    grey          = "\033[48;5;;38;5;8m"
    brown         = "\033[48;5;;38;5;94m"
    white         = "\033[48;5;;38;5;7m"
    whitelt       = "\033[48;5;;38;5;247m"
    whitebt       = "\033[48;5;;38;5;255m"
    whitedk       = "\033[48;5;;38;5;8m"
    indigo        = "\033[48;5;;38;5;21m"
    indigolt      = "\033[48;5;;38;5;27m"
    indigobt      = "\033[48;5;;38;5;4m"
    indigodk      = "\033[48;5;;38;5;18m"
    green         = "\033[48;5;;38;5;40m"
    greenlt       = "\033[48;5;;38;5;42m"
    greenbt       = "\033[48;5;;38;5;46m"
    greendk       = "\033[48;5;;38;5;22m"
    red           = "\033[48;5;;38;5;88m"
    redlt         = "\033[48;5;;38;5;124m"
    redbt         = "\033[48;5;;38;5;196m"
    reddk         = "\033[48;5;;38;5;52m"
    orange        = "\033[48;5;;38;5;172m"
    orangelt      = "\033[48;5;;38;5;178m"
    orangebt      = "\033[48;5;;38;5;184m"
    orangedk      = "\033[48;5;;38;5;166m"
    yellow        = "\033[48;5;;38;5;154m"
    yellowlt      = "\033[48;5;;38;5;156m"
    yellowbt      = "\033[48;5;;38;5;226m"
    yellowdk      = "\033[48;5;;38;5;142m"
    violet        = "\033[48;5;;38;5;165m"
    violetlt      = "\033[48;5;;38;5;177m"
    violetbt      = "\033[48;5;;38;5;201m"
    violetdk      = "\033[48;5;;38;5;93m"
    blue          = "\033[48;5;;38;5;75m"
    bluelt        = "\033[48;5;;38;5;45m"
    bluebt        = "\033[48;5;;38;5;14m"
    bluedk        = "\033[48;5;;38;5;33m"
)

/*Reset color code*/
func Reset(content string) (string) {
    return reset+content
}
/*Custom allows you to define any color*/
func Custom (foreground, background int, content string) (string) {
    return "\033[48;5;" + strconv.Itoa(background) + ";38;5;" + strconv.Itoa(foreground) + "m" + content + reset
}
/*Bl - blink*/
func Bl(content string) (string) {
    return blink+content+reset
}
/*B - bold*/
func B(content string) (string) {
    return bold+content+reset
}
/*U - underline*/
func U(content string) (string) {
    return underline+content+reset
}
/*I - italic*/
func I(content string) (string) {
    return italic+content+reset
}
/*Black color code*/
func Black(content string) (string) {
    return black+content+reset
}
/*Gray color code*/
func Gray(content string) (string) {
    return grey+content+reset
}
/*Grey color code*/
func Grey(content string) (string) {
    return grey+content+reset
}
/*Brown color code*/
func Brown(content string) (string) {
    return brown+content+reset
}
/*White color code*/
func White(content string) (string) {
    return white+content+reset
}
/*WhiteLt - light white color code*/
func WhiteLt(content string) (string) {
    return whitelt+content+reset
}
/*WhiteBt - bright white color code*/
func WhiteBt(content string) (string) {
    return whitebt+content+reset
}
/*WhiteDk - dark white color code*/
func WhiteDk(content string) (string) {
    return whitedk+content+reset
}
/*Indigo color code*/
func Indigo(content string) (string) {
    return indigo+content+reset
}
/*IndigoLt - light indigo color code*/
func IndigoLt(content string) (string) {
    return indigolt+content+reset
}
/*IndigoBt - bright indigo color code*/
func IndigoBt(content string) (string) {
    return indigobt+content+reset
}
/*IndigoDk - dark indigo color code*/
func IndigoDk(content string) (string) {
    return indigodk+content+reset
}
/*Green color code*/
func Green(content string) (string) {
    return green+content+reset
}
/*GreenLt - light green color code*/
func GreenLt(content string) (string) {
    return greenlt+content+reset
}
/*GreenBt - bright green color code*/
func GreenBt(content string) (string) {
    return greenbt+content+reset
}
/*GreenDk - dark green color code*/
func GreenDk(content string) (string) {
    return greendk+content+reset
}
/*Red color code*/
func Red(content string) (string) {
    return red+content+reset
}
/*RedLt - light red color code*/
func RedLt(content string) (string) {
    return redlt+content+reset
}
/*RedBt - bright red color code*/
func RedBt(content string) (string) {
    return redbt+content+reset
}
/*RedDk - dark red color code*/
func RedDk(content string) (string) {
    return reddk+content+reset
}
/*Orange color code*/
func Orange(content string) (string) {
    return orange+content+reset
}
/*OrangeLt - light orange color code*/
func OrangeLt(content string) (string) {
    return orangelt+content+reset
}
/*OrangeBt - bright orange color code*/
func OrangeBt(content string) (string) {
    return orangebt+content+reset
}
/*OrangeDk - dark orange color code*/
func OrangeDk(content string) (string) {
    return orangedk+content+reset
}
/*Yellow color code*/
func Yellow(content string) (string) {
    return yellow+content+reset
}
/*YellowLt - light yellow color code*/
func YellowLt(content string) (string) {
    return yellowlt+content+reset
}
/*YellowBt - bright yellow color code*/
func YellowBt(content string) (string) {
    return yellowbt+content+reset
}
/*YellowDk - dark yellow color code*/
func YellowDk(content string) (string) {
    return yellowdk+content+reset
}
/*Violet color code*/
func Violet(content string) (string) {
    return violet+content+reset
}
/*VioletLt - light violet color code*/
func VioletLt(content string) (string) {
    return violetlt+content+reset
}
/*VioletBt - bright violet color code*/
func VioletBt(content string) (string) {
    return violetbt+content+reset
}
/*VioletDk - dark violet color code*/
func VioletDk(content string) (string) {
    return violetdk+content+reset
}
/*Blue color code*/
func Blue(content string) (string) {
    return blue+content+reset
}
/*BlueLt - light blue color code*/
func BlueLt(content string) (string) {
    return bluelt+content+reset
}
/*BlueBt - bright blue color code*/
func BlueBt(content string) (string) {
    return bluebt+content+reset
}
/*BlueDk - dark blue color code*/
func BlueDk(content string) (string) {
    return bluedk+content+reset
}

/*
Test outputs all possible colors (256 total) and helper function usage
*/
func Test() {
    for a := 0; a < 256; a++ {
       fmt.Printf(Custom(a, 0, "Custom(%d, 0, \"The quick brown fox jumped over the lazy dog.\")\n"), a)
    }
    for a := 0; a < 256; a++ {
       fmt.Printf(Custom(0, a, "Custom(0, %d, \"The quick brown fox jumped over the lazy dog.\")\n"), a)
    }
       fmt.Printf(Black("   Black(The quick brown fox jumped over the lazy dog\n"))
        fmt.Printf(Gray("    Gray(The quick brown fox jumped over the lazy dog\n"))
        fmt.Printf(Grey("    Grey(The quick brown fox jumped over the lazy dog\n"))
       fmt.Printf(Brown("   Brown(The quick brown fox jumped over the lazy dog\n"))
       fmt.Printf(White("   White(The quick brown fox jumped over the lazy dog\n"))
     fmt.Printf(WhiteLt(" WhiteLt(The quick brown fox jumped over the lazy dog\n"))
     fmt.Printf(WhiteBt(" WhiteBt(The quick brown fox jumped over the lazy dog\n"))
     fmt.Printf(WhiteDk(" WhiteDk(The quick brown fox jumped over the lazy dog\n"))
         fmt.Printf(Red("     Red(The quick brown fox jumped over the lazy dog\n"))
       fmt.Printf(RedLt("   RedLt(The quick brown fox jumped over the lazy dog\n"))
       fmt.Printf(RedBt("   RedBt(The quick brown fox jumped over the lazy dog\n"))
       fmt.Printf(RedDk("   RedDk(The quick brown fox jumped over the lazy dog\n"))
      fmt.Printf(Orange("  Orange(The quick brown fox jumped over the lazy dog\n"))
    fmt.Printf(OrangeLt("OrangeLt(The quick brown fox jumped over the lazy dog\n"))
    fmt.Printf(OrangeBt("OrangeBt(The quick brown fox jumped over the lazy dog\n"))
    fmt.Printf(OrangeDk("OrangeDk(The quick brown fox jumped over the lazy dog\n"))
      fmt.Printf(Yellow("  Yellow(The quick brown fox jumped over the lazy dog\n"))
    fmt.Printf(YellowLt("YellowLt(The quick brown fox jumped over the lazy dog\n"))
    fmt.Printf(YellowBt("YellowBt(The quick brown fox jumped over the lazy dog\n"))
    fmt.Printf(YellowDk("YellowDk(The quick brown fox jumped over the lazy dog\n"))
       fmt.Printf(Green("   Green(The quick brown fox jumped over the lazy dog\n"))
     fmt.Printf(GreenLt(" GreenLt(The quick brown fox jumped over the lazy dog\n"))
     fmt.Printf(GreenBt(" GreenBt(The quick brown fox jumped over the lazy dog\n"))
     fmt.Printf(GreenDk(" GreenDk(The quick brown fox jumped over the lazy dog\n"))
        fmt.Printf(Blue("    Blue(The quick brown fox jumped over the lazy dog\n"))
      fmt.Printf(BlueLt("  BlueLt(The quick brown fox jumped over the lazy dog\n"))
      fmt.Printf(BlueBt("  BlueBt(The quick brown fox jumped over the lazy dog\n"))
      fmt.Printf(BlueDk("  BlueDk(The quick brown fox jumped over the lazy dog\n"))
      fmt.Printf(Indigo("  Indigo(The quick brown fox jumped over the lazy dog\n"))
    fmt.Printf(IndigoLt("IndigoLt(The quick brown fox jumped over the lazy dog\n"))
    fmt.Printf(IndigoBt("IndigoBt(The quick brown fox jumped over the lazy dog\n"))
    fmt.Printf(IndigoDk("IndigoDk(The quick brown fox jumped over the lazy dog\n"))
      fmt.Printf(Violet("  Violet(The quick brown fox jumped over the lazy dog\n"))
    fmt.Printf(VioletLt("VioletLt(The quick brown fox jumped over the lazy dog\n"))
    fmt.Printf(VioletBt("VioletBt(The quick brown fox jumped over the lazy dog\n"))
    fmt.Printf(VioletDk("VioletDk(The quick brown fox jumped over the lazy dog\n"))
          fmt.Printf(Bl("      Bl(The quick brown fox jumped over the lazy dog\n"))
           fmt.Printf(B("       B(The quick brown fox jumped over the lazy dog\n"))
           fmt.Printf(I("       I(The quick brown fox jumped over the lazy dog\n"))
           fmt.Printf(U("       U(The quick brown fox jumped over the lazy dog\n"))
       fmt.Printf(Reset("   Reset(The quick brown fox jumped over the lazy dog\n\n"))
}
