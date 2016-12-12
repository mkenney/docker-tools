
package ui

import (
//    "bytes"
    "fmt"
    "log"
    "os"
    "os/exec"
    "strconv"
//    "io/ioutil"
//    "strings"
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
    black         = 0
    gray          = 8
    grey          = 8
    brown         = 94
    white         = 7
    white_lt      = 247
    white_bt      = 255
    white_dk      = 8
    indigo        = 21
    indigo_lt     = 27
    indigo_bt     = 4
    indigo_dk     = 18
    green         = 40
    green_lt      = 42
    green_bt      = 46
    green_dk      = 22
    red           = 88
    red_lt        = 124
    red_bt        = 196
    red_dk        = 52
    orange        = 172
    orange_lt     = 178
    orange_bt     = 184
    orange_dk     = 166
    yellow        = 154
    yellow_lt     = 156
    yellow_bt     = 226
    yellow_dk     = 142
    violet        = 165
    violet_lt     = 177
    violet_bt     = 201
    violet_dk     = 93
    blue          = 75
    blue_lt       = 45
    blue_bt       = 14
    blue_dk       = 33
)

//echo -e "\\033[48;5;95;38;5;214mhello world\\033[0m"

func render (fg int, bg int, style string, content string) (string) {
    return style + "\033[48;5;" + strconv.Itoa(bg) + ";38;5;" + strconv.Itoa(fg) + "m" + content + reset
}

// Blink
func Bl(content string) (string) {
    return render(normal, blink, normal, content)
}
// Bold
func B(content string) (string) {
    return render(normal, bold, content)
}
// Underline
func U(content string) (string) {
    return render(normal, bold, content)
}
// Italic
func I(content string) (string) {
    return render(normal, bold, content)
}

func Black(content string) (string) {
    return render(black, 0, normal)
}
func Gray(content string) (string) {
    return render(gray, 0, normal)
}
func Grey(content string) (string) {
    return render(grey, 0, normal)
}
func Brown(content string) (string) {
    return render(brown, 0, normal)
}
func White(content string) (string) {
    return render(white, 0, normal)
}
func WhiteLt(content string) (string) {
    return render(white_lt, 0, normal)
}
func WhiteBt(content string) (string) {
    return render(white_bt, 0, normal)
}
func WhiteDk(content string) (string) {
    return render(white_dk, 0, normal)
}
func Indigo(content string) (string) {
    return render(indigo, 0, normal)
}
func IndigoLt(content string) (string) {
    return render(indigo_lt, 0, normal)
}
func IndigoBt(content string) (string) {
    return render(indigo_bt, 0, normal)
}
func IndigoDk(content string) (string) {
    return render(indigo_dk, 0, normal)
}
func Green(content string) (string) {
    return render(green, 0, normal)
}
func GreenLt(content string) (string) {
    return render(green_lt, 0, normal)
}
func GreenBt(content string) (string) {
    return render(green_bt, 0, normal)
}
func GreenDk(content string) (string) {
    return render(green_dk, 0, normal)
}
func Red(content string) (string) {
    return render(red, 0, normal)
}
func RedLt(content string) (string) {
    return render(red_lt, 0, normal)
}
func RedBt(content string) (string) {
    return render(red_bt, 0, normal)
}
func RedDk(content string) (string) {
    return render(red_dk, 0, normal)
}
func Orange(content string) (string) {
    return render(orange, 0, normal)
}
func OrangeLt(content string) (string) {
    return render(orange_lt, 0, normal)
}
func OrangeBt(content string) (string) {
    return render(orange_bt, 0, normal)
}
func OrangeDk(content string) (string) {
    return render(orange_dk, 0, normal)
}
func Yellow(content string) (string) {
    return render(yellow, 0, normal)
}
func YellowLt(content string) (string) {
    return render(yellow_lt, 0, normal)
}
func YellowBt(content string) (string) {
    return render(yellow_bt, 0, normal)
}
func YellowDk(content string) (string) {
    return render(yellow_dk, 0, normal)
}
func Violet(content string) (string) {
    return render(violet, 0, normal)
}
func VioletLt(content string) (string) {
    return render(violet_lt, 0, normal)
}
func VioletBt(content string) (string) {
    return render(violet_bt, 0, normal)
}
func VioletDk(content string) (string) {
    return render(violet_dk, 0, normal)
}
func Blue(content string) (string) {
    return render(blue, 0, normal)
}
func BlueLt(content string) (string) {
    return render(blue_lt, 0, normal)
}
func BlueBt(content string) (string) {
    return render(blue_bt, 0, normal)
}
func BlueDk(content string) (string) {
    return render(blue_dk, 0, normal)
}

func Test() {

fmt.Println(Red(Bold("Some Text")))


//    for a := 0; a < 256; a++ {
//        //for b := 0; b < 256; b++ {
//        //style := fmt.Sprintf("\%03d[0m", a)
//           fmt.Println(a, render(a, 0, " The quick brown fox jumped over the lazy dog."))
////        fmt.Println(render(0, a, "The quick brown fox jumped over the lazy dog."))
//        //fmt.Printf("%s%d: The quick brown fox jumped over the lazy dog.\n"), style, a)
//        //if nil != err {
//        //    fmt.Printf("err: %v", err)
//        //} else {
//        //    fmt.Printf("out: %v", out)
//        //}
//        //}
//    }
}













var style_styles map[string]string


func Bold(str string) string {
    return "\033[1m"+str+"\033[0m"
}

func init() {
    Test()

    var style []byte
    var err error
    style_styles = make(map[string]string)

    out, err := exec.Command("date").Output()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("The date is %s\n", out)



    if out, err = exec.Command("tput", "setaf", "27").Output(); err != nil {
    //if out, err = exec.Command("/usr/bin/printf", "%s", "$(tput setaf 27)").Output(); err != nil {
    //if out, err = exec.Command("printf \\%s $(tput setaf 27)").Output(); err != nil {
        fmt.Printf("There was an error running the command: '%v' '%s'", err, out)
        os.Exit(1)
    } else {
        fmt.Printf("%sblue..kdf\n", out)
    }
    os.Exit(1)


//    if cmd_out, cmd_err := exec.Command("printf %s $(tput setaf 27)blue").Output(); cmd_err != nil {
//        fmt.Fprintln(os.Stderr, "There was an error running the command: ", cmd_err)
//        os.Exit(1)
//    } else {
//        fmt.Printf("ls: %s\n", cmd_out)
//    }
//
//
//
//    cmd := exec.Command("/bin/echo $PATH")
//    var tmp bytes.Buffer
//    cmd.Stdout = &tmp
//    err = cmd.Run()
//    if err != nil {
//        log.Fatal(err)
//    }
//    fmt.Printf("PATH: %s\n__________________\n\n", tmp, err)
//
//    style, err = exec.Command("printf \"$(tput setaf 27)\"").Output()
//    fmt.Printf("%sI SHOULD BLUE\n__________________\n%v\n__________________\n\n", string(style), err)






    style, err = exec.Command("printf \"$(tput sgr0)$'\\e[0m'\"").Output()
    style_styles["s_reset"] = ""
    if nil != err {style_styles["s_reset"] = string(style)}

    style, err = exec.Command("printf(\"$'\\e[5m'\")").Output()
    style_styles["s_blink"] = ""
    if nil != err {style_styles["s_blink"] = string(style)}

    style, err = exec.Command("printf(\"$'\\e[6m'\")").Output()
    style_styles["s_blinkfast"] = ""
    if nil != err {style_styles["s_blinkfast"] = string(style)}

    style, err = exec.Command("printf \"$(tput bold)\"").Output()
    style_styles["s_bold"] = ""
    if nil != err {style_styles["s_bold"] = string(style)}

    style, err = exec.Command("printf(\"$'\\e[3m'\")").Output()
    style_styles["s_italic"] = ""
    if nil != err {style_styles["s_italic"] = string(style)}

    style, err = exec.Command("printf \"$(tput smul)\"").Output()
    style_styles["s_underline"] = ""
    if nil != err {style_styles["s_underline"] = string(style)}

    style, err = exec.Command("printf(\"$'\\e[9m'\")").Output()
    style_styles["s_strikethrough"] = ""
    if nil != err {style_styles["s_strikethrough"] = string(style)}

    style, err = exec.Command("printf \"$(tput setaf 0)\"").Output()
    style_styles["c_black"] = ""
    if nil != err {style_styles["c_black"] = string(style)}

    style, err = exec.Command("printf \"$(tput setaf 8)\"").Output()
    style_styles["c_gray"] = ""
    if nil != err {style_styles["c_gray"] = string(style)}

    style, err = exec.Command("printf \"$(tput setaf 8)\"").Output()
    style_styles["c_grey"] = ""
    if nil != err {style_styles["c_grey"] = string(style)}

    style, err = exec.Command("printf \"$(tput setaf 94)\"").Output()
    style_styles["c_brown"] = ""
    if nil != err {style_styles["c_brown"] = string(style)}

    style, err = exec.Command("printf \"$(tput setaf 7)\"").Output()
    style_styles["c_white"] = ""
    if nil != err {style_styles["c_white"] = string(style)}

    style, err = exec.Command("printf \"$(tput setaf 247)\"").Output()
    style_styles["c_white_lt"] = ""
    if nil != err {style_styles["c_white_lt"] = string(style)}

    style, err = exec.Command("printf \"$(tput setaf 255)\"").Output()
    style_styles["c_white_bt"] = ""
    if nil != err {style_styles["c_white_bt"] = string(style)}

    style, err = exec.Command("printf \"$(tput setaf 8)\"").Output()
    style_styles["c_white_dk"] = ""
    if nil != err {style_styles["c_white_dk"] = string(style)}

    style, err = exec.Command("printf \"$(tput setaf 21)\"").Output()
    style_styles["c_indigo"] = ""
    if nil != err {style_styles["c_indigo"] = string(style)}

    style, err = exec.Command("printf \"$(tput setaf 27)\"").Output()
    style_styles["c_indigo_lt"] = ""
    if nil != err {style_styles["c_indigo_lt"] = string(style)}

    style, err = exec.Command("printf \"$(tput setaf 4)\"").Output()
    style_styles["c_indigo_bt"] = ""
    if nil != err {style_styles["c_indigo_bt"] = string(style)}

    style, err = exec.Command("printf \"$(tput setaf 18)\"").Output()
    style_styles["c_indigo_dk"] = ""
    if nil != err {style_styles["c_indigo_dk"] = string(style)}

    style, err = exec.Command("printf \"$(tput setaf 40)\"").Output()
    style_styles["c_green"] = ""
    if nil != err {style_styles["c_green"] = string(style)}

    style, err = exec.Command("printf \"$(tput setaf 42)\"").Output()
    style_styles["c_green_lt"] = ""
    if nil != err {style_styles["c_green_lt"] = string(style)}

    style, err = exec.Command("printf \"$(tput setaf 46)\"").Output()
    style_styles["c_green_bt"] = ""
    if nil != err {style_styles["c_green_bt"] = string(style)}

    style, err = exec.Command("printf \"$(tput setaf 22)\"").Output()
    style_styles["c_green_dk"] = ""
    if nil != err {style_styles["c_green_dk"] = string(style)}

    style, err = exec.Command("printf \"$(tput setaf 88)\"").Output()
    style_styles["c_red"] = ""
    if nil != err {style_styles["c_red"] = string(style)}

    style, err = exec.Command("printf \"$(tput setaf 124)\"").Output()
    style_styles["c_red_lt"] = ""
    if nil != err {style_styles["c_red_lt"] = string(style)}

    style, err = exec.Command("printf \"$(tput setaf 196)\"").Output()
    style_styles["c_red_bt"] = ""
    if nil != err {style_styles["c_red_bt"] = string(style)}

    style, err = exec.Command("printf \"$(tput setaf 52)\"").Output()
    style_styles["c_red_dk"] = ""
    if nil != err {style_styles["c_red_dk"] = string(style)}

    style, err = exec.Command("printf \"$(tput setaf 172)\"").Output()
    style_styles["c_orange"] = ""
    if nil != err {style_styles["c_orange"] = string(style)}

    style, err = exec.Command("printf \"$(tput setaf 178)\"").Output()
    style_styles["c_orange_lt"] = ""
    if nil != err {style_styles["c_orange_lt"] = string(style)}

    style, err = exec.Command("printf \"$(tput setaf 184)\"").Output()
    style_styles["c_orange_bt"] = ""
    if nil != err {style_styles["c_orange_bt"] = string(style)}

    style, err = exec.Command("printf \"$(tput setaf 166)\"").Output()
    style_styles["c_orange_dk"] = ""
    if nil != err {style_styles["c_orange_dk"] = string(style)}

    style, err = exec.Command("printf \"$(tput setaf 154)\"").Output()
    style_styles["c_yellow"] = ""
    if nil != err {style_styles["c_yellow"] = string(style)}

    style, err = exec.Command("printf \"$(tput setaf 156)\"").Output()
    style_styles["c_yellow_lt"] = ""
    if nil != err {style_styles["c_yellow_lt"] = string(style)}

    style, err = exec.Command("printf \"$(tput setaf 226)\"").Output()
    style_styles["c_yellow_bt"] = ""
    if nil != err {style_styles["c_yellow_bt"] = string(style)}

    style, err = exec.Command("printf \"$(tput setaf 142)\"").Output()
    style_styles["c_yellow_dk"] = ""
    if nil != err {style_styles["c_yellow_dk"] = string(style)}

    style, err = exec.Command("printf \"$(tput setaf 165)\"").Output()
    style_styles["c_violet"] = ""
    if nil != err {style_styles["c_violet"] = string(style)}

    style, err = exec.Command("printf \"$(tput setaf 177)\"").Output()
    style_styles["c_violet_lt"] = ""
    if nil != err {style_styles["c_violet_lt"] = string(style)}

    style, err = exec.Command("printf \"$(tput setaf 201)\"").Output()
    style_styles["c_violet_bt"] = ""
    if nil != err {style_styles["c_violet_bt"] = string(style)}

    style, err = exec.Command("printf \"$(tput setaf 93)\"").Output()
    style_styles["c_violet_dk"] = ""
    if nil != err {style_styles["c_violet_dk"] = string(style)}

    style, err = exec.Command("printf \"$(tput setaf 75)\"").Output()
    style_styles["c_blue"] = ""
    if nil != err {style_styles["c_blue"] = string(style)}

    style, err = exec.Command("printf \"$(tput setaf 45)\"").Output()
    style_styles["c_blue_lt"] = ""
    if nil != err {style_styles["c_blue_lt"] = string(style)}

    style, err = exec.Command("printf \"$(tput setaf 14)\"").Output()
    style_styles["c_blue_bt"] = ""
    if nil != err {style_styles["c_blue_bt"] = string(style)}

    style, err = exec.Command("printf \"$(tput setaf 33)\"").Output()
    if nil != err {style_styles["c_blue_dk"] = string(style)}

    S("test", "")
}

func S(style, modifier string) (string) {
    var ret_val string
    var prefix string
    var suffix string

fmt.Printf("style_styles: %v\n\n--------------------------------------------\n\n", style_styles)
    if  "test" != style {
        switch {
            case "black"  == style: fallthrough
            case "gray"   == style: fallthrough
            case "grey"   == style: fallthrough
            case "brown"  == style: fallthrough
            case "white"  == style: fallthrough
            case "indigo" == style: fallthrough
            case "green"  == style: fallthrough
            case "red"    == style: fallthrough
            case "orange" == style: fallthrough
            case "yellow" == style: fallthrough
            case "violet" == style: fallthrough
            case "blue"   == style:
                prefix = "c_"
                suffix = ""
                if "lt" == modifier {
                    suffix = "_"+modifier
                } else if "bt" == modifier {
                    suffix = "_"+modifier
                } else if "dk" == modifier {
                    suffix = "_"+modifier
                }

            case "blink"         == style: fallthrough
            case "blinkfast"     == style: fallthrough
            case "bold"          == style: fallthrough
            case "italic"        == style: fallthrough
            case "reset"         == style: fallthrough
            case "strikethrough" == style: fallthrough
            case "underline"     == style: fallthrough
            case "b"             == style: fallthrough
            case "i"             == style: fallthrough
            case "s"             == style: fallthrough
            case "r"             == style: fallthrough
            case "u"             == style:
                prefix = "s_"
                suffix = ""
                if "b" == style {style="bold"}
                if "i" == style {style="italic"}
                if "r" == style {style="reset"}
                if "s" == style {style="strikethrough"}
                if "u" == style {style="underline"}
        }
        ret_val = string(style_styles[prefix+style+suffix])

//    } else {
        for a := 0; a < 256; a++ {
            out, err := exec.Command(fmt.Sprintf("printf(\"$(tput setaf %v)%v: The quick brown fox jumped over the lazy dog.\\n\")", a, a)).Output()
            if nil != err {
                fmt.Printf("%v", out)
            }
        }
//
//        fmt.Printf("black         %vThe quick brown fox jumped over the lazy dog %v\n", S("black", ""), S("r", ""))
//        fmt.Printf("gray          %vThe quick brown fox jumped over the lazy dog %v\n", S("gray",  ""), S("r", ""))
//        fmt.Printf("grey          %vThe quick brown fox jumped over the lazy dog %v\n", S("grey",  ""), S("r", ""))
//        fmt.Printf("brown         %vThe quick brown fox jumped over the lazy dog %v\n", S("brown", ""), S("r", ""))
//
//        fmt.Printf("white         %vThe quick brown fox jumped over the lazy dog %v\n", S("white",  ""),   S("r", ""))
//        fmt.Printf("white lt      %vThe quick brown fox jumped over the lazy dog %v\n", S("white",  "lt"), S("r", ""))
//        fmt.Printf("white bt      %vThe quick brown fox jumped over the lazy dog %v\n", S("white",  "bt"), S("r", ""))
//        fmt.Printf("white dk      %vThe quick brown fox jumped over the lazy dog %v\n", S("white",  "dk"), S("r", ""))
//        fmt.Printf("red           %vThe quick brown fox jumped over the lazy dog %v\n", S("red",    ""),   S("r", ""))
//        fmt.Printf("red lt        %vThe quick brown fox jumped over the lazy dog %v\n", S("red",    "lt"), S("r", ""))
//        fmt.Printf("red bt        %vThe quick brown fox jumped over the lazy dog %v\n", S("red",    "bt"), S("r", ""))
//        fmt.Printf("red dk        %vThe quick brown fox jumped over the lazy dog %v\n", S("red",    "dk"), S("r", ""))
//        fmt.Printf("orange        %vThe quick brown fox jumped over the lazy dog %v\n", S("orange", ""),   S("r", ""))
//        fmt.Printf("orange lt     %vThe quick brown fox jumped over the lazy dog %v\n", S("orange", "lt"), S("r", ""))
//        fmt.Printf("orange bt     %vThe quick brown fox jumped over the lazy dog %v\n", S("orange", "bt"), S("r", ""))
//        fmt.Printf("orange dk     %vThe quick brown fox jumped over the lazy dog %v\n", S("orange", "dk"), S("r", ""))
//        fmt.Printf("yellow        %vThe quick brown fox jumped over the lazy dog %v\n", S("yellow", ""),   S("r", ""))
//        fmt.Printf("yellow lt     %vThe quick brown fox jumped over the lazy dog %v\n", S("yellow", "lt"), S("r", ""))
//        fmt.Printf("yellow bt     %vThe quick brown fox jumped over the lazy dog %v\n", S("yellow", "bt"), S("r", ""))
//        fmt.Printf("yellow dk     %vThe quick brown fox jumped over the lazy dog %v\n", S("yellow", "dk"), S("r", ""))
//        fmt.Printf("green         %vThe quick brown fox jumped over the lazy dog %v\n", S("green",  ""),   S("r", ""))
//        fmt.Printf("green lt      %vThe quick brown fox jumped over the lazy dog %v\n", S("green",  "lt"), S("r", ""))
//        fmt.Printf("green bt      %vThe quick brown fox jumped over the lazy dog %v\n", S("green",  "bt"), S("r", ""))
//        fmt.Printf("green dk      %vThe quick brown fox jumped over the lazy dog %v\n", S("green",  "dk"), S("r", ""))
//        fmt.Printf("blue          %vThe quick brown fox jumped over the lazy dog %v\n", S("blue",   ""),   S("r", ""))
//        fmt.Printf("blue lt       %vThe quick brown fox jumped over the lazy dog %v\n", S("blue",   "lt"), S("r", ""))
//        fmt.Printf("blue bt       %vThe quick brown fox jumped over the lazy dog %v\n", S("blue",   "bt"), S("r", ""))
//        fmt.Printf("blue dk       %vThe quick brown fox jumped over the lazy dog %v\n", S("blue",   "dk"), S("r", ""))
//        fmt.Printf("indigo        %vThe quick brown fox jumped over the lazy dog %v\n", S("indigo", ""),   S("r", ""))
//        fmt.Printf("indigo lt     %vThe quick brown fox jumped over the lazy dog %v\n", S("indigo", "lt"), S("r", ""))
//        fmt.Printf("indigo bt     %vThe quick brown fox jumped over the lazy dog %v\n", S("indigo", "bt"), S("r", ""))
//        fmt.Printf("indigo dk     %vThe quick brown fox jumped over the lazy dog %v\n", S("indigo", "dk"), S("r", ""))
//        fmt.Printf("violet        %vThe quick brown fox jumped over the lazy dog %v\n", S("violet", ""),   S("r", ""))
//        fmt.Printf("violet lt     %vThe quick brown fox jumped over the lazy dog %v\n", S("violet", "lt"), S("r", ""))
//        fmt.Printf("violet bt     %vThe quick brown fox jumped over the lazy dog %v\n", S("violet", "bt"), S("r", ""))
//        fmt.Printf("violet dk     %vThe quick brown fox jumped over the lazy dog %v\n", S("violet", "dk"), S("r", ""))
//
//        fmt.Printf("blink         %vThe quick brown fox jumped over the lazy dog %v\n", S("blink", ""),         S("r", ""))
//        fmt.Printf("blinkfast     %vThe quick brown fox jumped over the lazy dog %v\n", S("blinkfast", ""),     S("r", ""))
//        fmt.Printf("bold          %vThe quick brown fox jumped over the lazy dog %v\n", S("bold", ""),          S("r", ""))
//        fmt.Printf("italic        %vThe quick brown fox jumped over the lazy dog %v\n", S("italic", ""),        S("r", ""))
//        fmt.Printf("strikethrough %vThe quick brown fox jumped over the lazy dog %v\n", S("strikethrough", ""), S("r", ""))
//        fmt.Printf("underline     %vThe quick brown fox jumped over the lazy dog %v\n", S("underline", ""),     S("r", ""))
//        fmt.Printf("reset         %vThe quick brown fox jumped over the lazy dog %v\n", S("reset", ""),         S("r", ""))
    }
    return ret_val
}
