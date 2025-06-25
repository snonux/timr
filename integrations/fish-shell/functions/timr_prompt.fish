function timr_prompt -d "Display timr timr_status in the prompt"
    if command -v timr >/dev/null
        set -l timr_status (timr prompt)
        if test -n "$timr_status"
            set -l icon (string sub -l 1 -- "$timr_status")
            set -l time (string sub -s 2 -- "$timr_status")
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
