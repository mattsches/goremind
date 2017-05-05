# GoRemind

> `GoRemind` is a simple, [Slack-like](https://get.slack.help/hc/en-us/articles/208423427-Set-a-reminder) reminder command for the terminal.

## Prerequisites

`GoRemind` has been developed on Ubuntu 16.10 and requires `libnotify-bin` and `at` to be installed.

## Usage

Examples:

```bash
$ goremind me to take a break in 10 minutes
Okay, I will remind you to take a break at 2017-05-05 00:15
$ goremind me tomorrow at 8am to make coffee
Okay, I will remind you to make coffee at 2017-05-06 08:00
```

## Limitations

I've written `GoRemind` to scratch my own itch, i.e. to be able to set Slack-like reminders in my terminal. If it doesn't work for you, please open an issue.

`GoRemind` currently supports only English language input.