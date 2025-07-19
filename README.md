# timr

A simple command-line tool to track time spent on tasks. It has been primarily coded using Google Gemini CLI and Claude Code CLI.

## About

`timr` is a minimalist stopwatch-style timer that runs in your terminal. It helps you track the time you spend on your work by allowing you to start, stop, and check the status of the timer. The timer's state is saved to your system, so you can pause your work and resume tracking later.

It has been vibe coded!

## Installation

To build and install `timr`, you need to have Go installed on your system. You can then build the executable by running the following command in the project's root directory:

```bash
go build ./...
```

This will create a `timr` executable in the current directory. To make it accessible from anywhere, you can move it to a directory in your system's `PATH`, such as `/usr/local/bin`:

```bash
sudo mv timr /usr/local/bin/
```

## Usage

`timr` provides the following commands:

*   `timr start`: Starts the timer. If the timer was previously stopped, it will resume from where it left off.
*   `timr stop` or `timr pause`: Stops or pauses the timer. The elapsed time will be saved.
*   `timr status`: Shows the current status of the timer (running or stopped) and the total elapsed time.
*   `timr status raw`: Shows the current elapsed time in seconds, in a raw format.
*   `timr status rawm`: Shows the current elapsed time in minutes, in a raw format.
*   `timr reset`: Resets the timer. This will set the elapsed time to zero.
*   `timr live`: Shows a live, full-screen timer with keyboard controls (q: quit, s: start/stop, r: reset).

## Fish Shell Integration

`timr` can be integrated with the fish shell to display the current timer status in your prompt.

### Installation

Add this to your fish config:

```fish
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

complete -c timr -n __fish_use_subcommand -a start -d "Start the timer"
complete -c timr -n __fish_use_subcommand -a stop -d "Stop the timer"
complete -c timr -n __fish_use_subcommand -a pause -d "Pause the timer"
complete -c timr -n __fish_use_subcommand -a status -d "Show the timer status"
complete -c timr -n __fish_use_subcommand -a reset -d "Reset the timer"
complete -c timr -n __fish_use_subcommand -a live -d "Show the live timer"
complete -c timr -n __fish_use_subcommand -a prompt -d "Show the prompt status"
```

2.  Update your `fish_prompt` or `fish_right_prompt` function to include the `timr_prompt` function:

```fish
function fish_prompt
    # ... your existing prompt ...
    printf ' %s' (timr_prompt)
end
```
