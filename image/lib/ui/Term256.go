
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
    white_lt      = "\033[48;5;;38;5;247m"
    white_bt      = "\033[48;5;;38;5;255m"
    white_dk      = "\033[48;5;;38;5;8m"
    indigo        = "\033[48;5;;38;5;21m"
    indigo_lt     = "\033[48;5;;38;5;27m"
    indigo_bt     = "\033[48;5;;38;5;4m"
    indigo_dk     = "\033[48;5;;38;5;18m"
    green         = "\033[48;5;;38;5;40m"
    green_lt      = "\033[48;5;;38;5;42m"
    green_bt      = "\033[48;5;;38;5;46m"
    green_dk      = "\033[48;5;;38;5;22m"
    red           = "\033[48;5;;38;5;88m"
    red_lt        = "\033[48;5;;38;5;124m"
    red_bt        = "\033[48;5;;38;5;196m"
    red_dk        = "\033[48;5;;38;5;52m"
    orange        = "\033[48;5;;38;5;172m"
    orange_lt     = "\033[48;5;;38;5;178m"
    orange_bt     = "\033[48;5;;38;5;184m"
    orange_dk     = "\033[48;5;;38;5;166m"
    yellow        = "\033[48;5;;38;5;154m"
    yellow_lt     = "\033[48;5;;38;5;156m"
    yellow_bt     = "\033[48;5;;38;5;226m"
    yellow_dk     = "\033[48;5;;38;5;142m"
    violet        = "\033[48;5;;38;5;165m"
    violet_lt     = "\033[48;5;;38;5;177m"
    violet_bt     = "\033[48;5;;38;5;201m"
    violet_dk     = "\033[48;5;;38;5;93m"
    blue          = "\033[48;5;;38;5;75m"
    blue_lt       = "\033[48;5;;38;5;45m"
    blue_bt       = "\033[48;5;;38;5;14m"
    blue_dk       = "\033[48;5;;38;5;33m"
)

func Reset(content string) (string) {
    return reset+content
}

/**
 * @param int     foreground  0 - 255
 * @param int     background  0 - 255
 * @param string  content     The content to style
 * @return string
 */
func Custom (foreground, background int, content string) (string) {
    return "\033[48;5;" + strconv.Itoa(background) + ";38;5;" + strconv.Itoa(foreground) + "m" + content + reset
}

func Bl(content string) (string) {
    return blink+content+reset
}
func B(content string) (string) {
    return bold+content+reset
}
func U(content string) (string) {
    return underline+content+reset
}
func I(content string) (string) {
    return italic+content+reset
}
func Black(content string) (string) {
    return black+content+reset
}
func Gray(content string) (string) {
    return grey+content+reset
}
func Grey(content string) (string) {
    return grey+content+reset
}
func Brown(content string) (string) {
    return brown+content+reset
}
func White(content string) (string) {
    return white+content+reset
}
func WhiteLt(content string) (string) {
    return white_lt+content+reset
}
func WhiteBt(content string) (string) {
    return white_bt+content+reset
}
func WhiteDk(content string) (string) {
    return white_dk+content+reset
}
func Indigo(content string) (string) {
    return indigo+content+reset
}
func IndigoLt(content string) (string) {
    return indigo_lt+content+reset
}
func IndigoBt(content string) (string) {
    return indigo_bt+content+reset
}
func IndigoDk(content string) (string) {
    return indigo_dk+content+reset
}
func Green(content string) (string) {
    return green+content+reset
}
func GreenLt(content string) (string) {
    return green_lt+content+reset
}
func GreenBt(content string) (string) {
    return green_bt+content+reset
}
func GreenDk(content string) (string) {
    return green_dk+content+reset
}
func Red(content string) (string) {
    return red+content+reset
}
func RedLt(content string) (string) {
    return red_lt+content+reset
}
func RedBt(content string) (string) {
    return red_bt+content+reset
}
func RedDk(content string) (string) {
    return red_dk+content+reset
}
func Orange(content string) (string) {
    return orange+content+reset
}
func OrangeLt(content string) (string) {
    return orange_lt+content+reset
}
func OrangeBt(content string) (string) {
    return orange_bt+content+reset
}
func OrangeDk(content string) (string) {
    return orange_dk+content+reset
}
func Yellow(content string) (string) {
    return yellow+content+reset
}
func YellowLt(content string) (string) {
    return yellow_lt+content+reset
}
func YellowBt(content string) (string) {
    return yellow_bt+content+reset
}
func YellowDk(content string) (string) {
    return yellow_dk+content+reset
}
func Violet(content string) (string) {
    return violet+content+reset
}
func VioletLt(content string) (string) {
    return violet_lt+content+reset
}
func VioletBt(content string) (string) {
    return violet_bt+content+reset
}
func VioletDk(content string) (string) {
    return violet_dk+content+reset
}
func Blue(content string) (string) {
    return blue+content+reset
}
func BlueLt(content string) (string) {
    return blue_lt+content+reset
}
func BlueBt(content string) (string) {
    return blue_bt+content+reset
}
func BlueDk(content string) (string) {
    return blue_dk+content+reset
}

/**
 * Run through all colored output to see what a terminal supports
 */
func Test() {
    for a := 0; a < 256; a++ {
       fmt.Printf(Custom(a, 0, "Custom(%d, 0, \"The quick brown fox jumped over the lazy dog.\")\n"), a)
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
