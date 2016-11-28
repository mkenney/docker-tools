declare c_reset=$(tput sgr0)$'\e[0m'
declare c_black=$(tput setaf 0)
declare c_grey=$(tput setaf 8)
declare c_brown=$(tput setaf 94)
declare c_white=$(tput setaf 7)
declare c_white_lt=$(tput setaf 159)
declare c_white_bt=$(tput setaf 255)
declare c_white_dk=$(tput setaf 8)
declare c_blue=$(tput setaf 21)
declare c_blue_lt=$(tput setaf 27)
declare c_blue_bt=$(tput setaf 4)
declare c_blue_dk=$(tput setaf 18)
declare c_green=$(tput setaf 40)
declare c_green_lt=$(tput setaf 42)
declare c_green_bt=$(tput setaf 46)
declare c_green_dk=$(tput setaf 22)
declare c_red=$(tput setaf 88)
declare c_red_lt=$(tput setaf 124)
declare c_red_bt=$(tput setaf 196)
declare c_red_dk=$(tput setaf 52)
declare c_orange=$(tput setaf 172)
declare c_orange_lt=$(tput setaf 178)
declare c_orange_bt=$(tput setaf 184)
declare c_orange_dk=$(tput setaf 166)
declare c_yellow=$(tput setaf 154)
declare c_yellow_lt=$(tput setaf 156)
declare c_yellow_bt=$(tput setaf 226)
declare c_yellow_dk=$(tput setaf 142)
declare c_magenta=$(tput setaf 165)
declare c_magenta_lt=$(tput setaf 177)
declare c_magenta_bt=$(tput setaf 201)
declare c_magenta_dk=$(tput setaf 93)
declare c_cyan=$(tput setaf 75)
declare c_cyan_lt=$(tput setaf 45)
declare c_cyan_bt=$(tput setaf 14)
declare c_cyan_dk=$(tput setaf 33)

declare s_blink=$'\e[5m'
declare s_blinkfast=$'\e[6m'
declare s_bold=$(tput bold)
declare s_italic=$'\e[3m'
declare s_underline=$(tput smul)
declare s_strikethrough=$'\e[9m' # doesn't work in most terminals


#
# Terminal output styling helper
#
function _s {
    local ret_val

    local prefix
    local style
    local suffix
    while [ "$1" != "" ]; do
        style=$1
        shift
        case $style in

            # colors
            reset|r|black|grey|brown|white|blue|green|red|orange|yellow|magenta|cyan)
                prefix="c_"
                suffix=""
                if [ "r" == "$style" ]; then style="reset"; fi
                if [ "lt" == "$1" ] || [ "bt" == "$1" ] || [ "dk" == "$1" ]; then
                    suffix="_${1}"
                    shift
                fi
            ;;

            # styles
            blink|blinkfast|bold|italic|strikethrough|underline|b|i|s|u)
                prefix="s_"
                suffix=""
                if [ "b" == "$style" ]; then style="bold"; fi
                if [ "i" == "$style" ]; then style="italic"; fi
                if [ "s" == "$style" ]; then style="strikethrough"; fi
                if [ "u" == "$style" ]; then style="underline"; fi
            ;;

            # incorrect order somewhere
            lt|bt|dk)
                >&2 echo "Unused color suffix '$style'"
                exit 1
            ;;

            # invalid
            *)
                >&2 echo "Invalid style '$style'"
                exit 1
            ;;
        esac
        ret_val="${ret_val}$(eval "printf \"\$${prefix}${style}${suffix}\"")"
    done

    printf "$ret_val"
}

#IFS=$'\n'
#echo
#for i in $(seq 0 256); do
#    printf "$(tput setaf $i)$i: The quick brown fox jumped over the lazy dog.\n";
#done
#echo "
#    black         $(_s black)The quick brown fox jumped over the lazy dog$(_s reset)
#    grey          $(_s grey)The quick brown fox jumped over the lazy dog$(_s reset)
#    brown         $(_s brown)The quick brown fox jumped over the lazy dog$(_s reset)
#    white         $(_s white)The quick brown fox jumped over the lazy dog$(_s reset)
#    white lt      $(_s white lt)The quick brown fox jumped over the lazy dog$(_s reset)
#    white bt      $(_s white bt)The quick brown fox jumped over the lazy dog$(_s reset)
#    white dk      $(_s white dk)The quick brown fox jumped over the lazy dog$(_s reset)
#    blue          $(_s blue)The quick brown fox jumped over the lazy dog$(_s reset)
#    blue lt       $(_s blue lt)The quick brown fox jumped over the lazy dog$(_s reset)
#    blue bt       $(_s blue bt)The quick brown fox jumped over the lazy dog$(_s reset)
#    blue dk       $(_s blue dk)The quick brown fox jumped over the lazy dog$(_s reset)
#    green         $(_s green)The quick brown fox jumped over the lazy dog$(_s reset)
#    green lt      $(_s green lt)The quick brown fox jumped over the lazy dog$(_s reset)
#    green bt      $(_s green bt)The quick brown fox jumped over the lazy dog$(_s reset)
#    green dk      $(_s green dk)The quick brown fox jumped over the lazy dog$(_s reset)
#    red           $(_s red)The quick brown fox jumped over the lazy dog$(_s reset)
#    red lt        $(_s red lt)The quick brown fox jumped over the lazy dog$(_s reset)
#    red bt        $(_s red bt)The quick brown fox jumped over the lazy dog$(_s reset)
#    red dk        $(_s red dk)The quick brown fox jumped over the lazy dog$(_s reset)
#    orange        $(_s orange)The quick brown fox jumped over the lazy dog$(_s reset)
#    orange lt     $(_s orange lt)The quick brown fox jumped over the lazy dog$(_s reset)
#    orange bt     $(_s orange bt)The quick brown fox jumped over the lazy dog$(_s reset)
#    orange dk     $(_s orange dk)The quick brown fox jumped over the lazy dog$(_s reset)
#    yellow        $(_s yellow)The quick brown fox jumped over the lazy dog$(_s reset)
#    yellow lt     $(_s yellow lt)The quick brown fox jumped over the lazy dog$(_s reset)
#    yellow bt     $(_s yellow bt)The quick brown fox jumped over the lazy dog$(_s reset)
#    yellow dk     $(_s yellow dk)The quick brown fox jumped over the lazy dog$(_s reset)
#    magenta       $(_s magenta)The quick brown fox jumped over the lazy dog$(_s reset)
#    magenta lt    $(_s magenta lt)The quick brown fox jumped over the lazy dog$(_s reset)
#    magenta bt    $(_s magenta bt)The quick brown fox jumped over the lazy dog$(_s reset)
#    magenta dk    $(_s magenta dk)The quick brown fox jumped over the lazy dog$(_s reset)
#    cyan          $(_s cyan)The quick brown fox jumped over the lazy dog$(_s reset)
#    cyan lt       $(_s cyan lt)The quick brown fox jumped over the lazy dog$(_s reset)
#    cyan bt       $(_s cyan bt)The quick brown fox jumped over the lazy dog$(_s reset)
#    cyan dk       $(_s cyan dk)The quick brown fox jumped over the lazy dog$(_s reset)
#
#    blink         $(_s blink)The quick brown fox jumped over the lazy dog$(_s reset)
#    blinkfast     $(_s blinkfast)The quick brown fox jumped over the lazy dog$(_s reset)
#    bold          $(_s b)The quick brown fox jumped over the lazy dog$(_s reset)
#    italic        $(_s i)The quick brown fox jumped over the lazy dog$(_s reset)
#    strikethrough $(_s s)The quick brown fox jumped over the lazy dog$(_s reset)
#    underline     $(_s u)The quick brown fox jumped over the lazy dog$(_s reset)
#
#    reset         $(_s reset)The quick brown fox jumped over the lazy dog$(_s reset)
#"
#exit
