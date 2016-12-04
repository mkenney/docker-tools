
##############################################################################
##############################################################################
##
##  UI - Common functions useful for terminal styling
##
##  _s() - Very simple styling function. Usage:
##
##      echo "$(_s STYLE [[MODIFIER] ... ])text$(_s STYLE [[MODIFIER] ... ])"
##
##  Available style arguments:
##
##      reset         - reset the terminal style to the default
##      blink
##      blinkfast     - doesn't work on most terminals
##      bold
##      italic
##      underline
##      strikethrough - doesn't work on most terminals
##
##  Available color arguments:
##
##      black   - No modifier
##      grey    - No modifier
##      brown   - No modifier
##      white
##      blue
##      green
##      red
##      orange
##      yellow
##      magenta
##      cyan
##
##  Available color modifiers (always optional):
##
##      lt - lighten
##      bt - brighten
##      dk - darken
##
##  Examples:
##
##      Bold, blinking, light red, underlined text
##          echo "$(_s bold blink red lt underline)text$(_s reset)"
##
##      Shortcuts are available for bold, italic, strikethrough, underline and reset
##          echo "$(_s b i s u)text$(_s r)"
##
##      Grey italic text
##          echo "$(_s grey i)text$(_s r)"
##
##############################################################################
##############################################################################

# Reset to default
declare _style_reset=$(tput sgr0)$'\e[0m'

# styles
declare _style_blink=$'\e[5m'
declare _style_blinkfast=$'\e[6m'
declare _style_bold=$(tput bold)
declare _style_italic=$'\e[3m'
declare _style_underline=$(tput smul)
declare _style_strikethrough=$'\e[9m' # doesn't work in most terminals

# Colors
declare _color_black=$(tput setaf 0)
declare _color_gray=$(tput setaf 8)
declare _color_grey=$(tput setaf 8)
declare _color_brown=$(tput setaf 94)
declare _color_white=$(tput setaf 7)
declare _color_white_lt=$(tput setaf 159)
declare _color_white_bt=$(tput setaf 255)
declare _color_white_dk=$(tput setaf 8)
declare _color_blue=$(tput setaf 21)
declare _color_blue_lt=$(tput setaf 27)
declare _color_blue_bt=$(tput setaf 4)
declare _color_blue_dk=$(tput setaf 18)
declare _color_green=$(tput setaf 40)
declare _color_green_lt=$(tput setaf 42)
declare _color_green_bt=$(tput setaf 46)
declare _color_green_dk=$(tput setaf 22)
declare _color_red=$(tput setaf 88)
declare _color_red_lt=$(tput setaf 124)
declare _color_red_bt=$(tput setaf 196)
declare _color_red_dk=$(tput setaf 52)
declare _color_orange=$(tput setaf 172)
declare _color_orange_lt=$(tput setaf 178)
declare _color_orange_bt=$(tput setaf 184)
declare _color_orange_dk=$(tput setaf 166)
declare _color_yellow=$(tput setaf 154)
declare _color_yellow_lt=$(tput setaf 156)
declare _color_yellow_bt=$(tput setaf 226)
declare _color_yellow_dk=$(tput setaf 142)
declare _color_magenta=$(tput setaf 165)
declare _color_magenta_lt=$(tput setaf 177)
declare _color_magenta_bt=$(tput setaf 201)
declare _color_magenta_dk=$(tput setaf 93)
declare _color_cyan=$(tput setaf 75)
declare _color_cyan_lt=$(tput setaf 45)
declare _color_cyan_bt=$(tput setaf 14)
declare _color_cyan_dk=$(tput setaf 33)

#
# Terminal output styling helper
#
function _s {
    local ret_val

    local prefix
    local style
    local suffix
    if [ "test" != "$1" ]; then
        while [ "$1" != "" ]; do
            style=$1
            shift
            case $style in

                # colors
                black|gray|grey|brown|white|blue|green|red|orange|yellow|magenta|cyan)
                    prefix="_color_"
                    suffix=""
                    if [ "lt" == "$1" ] || [ "bt" == "$1" ] || [ "dk" == "$1" ]; then
                        suffix="_${1}"
                        shift
                    fi
                ;;

                # styles
                blink|blinkfast|bold|italic|reset|strikethrough|underline|b|i|s|r|u)
                    prefix="_style_"
                    suffix=""
                    if [ "b" == "$style" ]; then style="bold"; fi
                    if [ "i" == "$style" ]; then style="italic"; fi
                    if [ "r" == "$style" ]; then style="reset"; fi
                    if [ "s" == "$style" ]; then style="strikethrough"; fi
                    if [ "u" == "$style" ]; then style="underline"; fi
                ;;

                # incorrect order somewhere
                lt|bt|dk)
                    >&2 echo "Unused color suffix '$style'"
                    exit 2
                ;;

                # invalid
                *)
                    >&2 echo "Invalid style '$style'"
                    exit 3
                ;;
            esac
            ret_val="${ret_val}$(eval "printf \"\$${prefix}${style}${suffix}\"")"
        done

    else
        IFS=$'\n'
        echo
        for i in $(seq 0 256); do
            printf "$(tput setaf $i)$i: The quick brown fox jumped over the lazy dog.\n";
        done
        echo "
    black         $(_s black)The quick brown fox jumped over the lazy dog$(_s reset)
    gray          $(_s gray)The quick brown fox jumped over the lazy dog$(_s reset)
    grey          $(_s grey)The quick brown fox jumped over the lazy dog$(_s reset)
    brown         $(_s brown)The quick brown fox jumped over the lazy dog$(_s reset)
    white         $(_s white)The quick brown fox jumped over the lazy dog$(_s reset)
    white lt      $(_s white lt)The quick brown fox jumped over the lazy dog$(_s reset)
    white bt      $(_s white bt)The quick brown fox jumped over the lazy dog$(_s reset)
    white dk      $(_s white dk)The quick brown fox jumped over the lazy dog$(_s reset)
    blue          $(_s blue)The quick brown fox jumped over the lazy dog$(_s reset)
    blue lt       $(_s blue lt)The quick brown fox jumped over the lazy dog$(_s reset)
    blue bt       $(_s blue bt)The quick brown fox jumped over the lazy dog$(_s reset)
    blue dk       $(_s blue dk)The quick brown fox jumped over the lazy dog$(_s reset)
    green         $(_s green)The quick brown fox jumped over the lazy dog$(_s reset)
    green lt      $(_s green lt)The quick brown fox jumped over the lazy dog$(_s reset)
    green bt      $(_s green bt)The quick brown fox jumped over the lazy dog$(_s reset)
    green dk      $(_s green dk)The quick brown fox jumped over the lazy dog$(_s reset)
    red           $(_s red)The quick brown fox jumped over the lazy dog$(_s reset)
    red lt        $(_s red lt)The quick brown fox jumped over the lazy dog$(_s reset)
    red bt        $(_s red bt)The quick brown fox jumped over the lazy dog$(_s reset)
    red dk        $(_s red dk)The quick brown fox jumped over the lazy dog$(_s reset)
    orange        $(_s orange)The quick brown fox jumped over the lazy dog$(_s reset)
    orange lt     $(_s orange lt)The quick brown fox jumped over the lazy dog$(_s reset)
    orange bt     $(_s orange bt)The quick brown fox jumped over the lazy dog$(_s reset)
    orange dk     $(_s orange dk)The quick brown fox jumped over the lazy dog$(_s reset)
    yellow        $(_s yellow)The quick brown fox jumped over the lazy dog$(_s reset)
    yellow lt     $(_s yellow lt)The quick brown fox jumped over the lazy dog$(_s reset)
    yellow bt     $(_s yellow bt)The quick brown fox jumped over the lazy dog$(_s reset)
    yellow dk     $(_s yellow dk)The quick brown fox jumped over the lazy dog$(_s reset)
    magenta       $(_s magenta)The quick brown fox jumped over the lazy dog$(_s reset)
    magenta lt    $(_s magenta lt)The quick brown fox jumped over the lazy dog$(_s reset)
    magenta bt    $(_s magenta bt)The quick brown fox jumped over the lazy dog$(_s reset)
    magenta dk    $(_s magenta dk)The quick brown fox jumped over the lazy dog$(_s reset)
    cyan          $(_s cyan)The quick brown fox jumped over the lazy dog$(_s reset)
    cyan lt       $(_s cyan lt)The quick brown fox jumped over the lazy dog$(_s reset)
    cyan bt       $(_s cyan bt)The quick brown fox jumped over the lazy dog$(_s reset)
    cyan dk       $(_s cyan dk)The quick brown fox jumped over the lazy dog$(_s reset)

    blink         $(_s blink)The quick brown fox jumped over the lazy dog$(_s reset)
    blinkfast     $(_s blinkfast)The quick brown fox jumped over the lazy dog$(_s reset)
    bold          $(_s b)The quick brown fox jumped over the lazy dog$(_s reset)
    italic        $(_s i)The quick brown fox jumped over the lazy dog$(_s reset)
    strikethrough $(_s s)The quick brown fox jumped over the lazy dog$(_s reset)
    underline     $(_s u)The quick brown fox jumped over the lazy dog$(_s reset)

    reset         $(_s reset)The quick brown fox jumped over the lazy dog$(_s reset)
"
        exit
    fi

    printf "$ret_val"
}
