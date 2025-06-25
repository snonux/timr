function timr_prompt -d "Display timr status in the prompt"
    if command -v timr >/dev/null
        set -l status (timr prompt)
        if test -n "$status"
            set -l icon (string sub -l 1 -- "$status")
            set -l time (string sub -s 2 -- "$status")
            if test "$icon" = "▶"
                set_color green
            else
                set_color yellow
            end
            printf '%s' "$icon"
            set_color normal
            printf ' %s' "$time"
        end
    end
end
