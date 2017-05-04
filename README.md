# GoRemind

> `GoRemind` is a simple, [Slack-like](https://get.slack.help/hc/en-us/articles/208423427-Set-a-reminder) reminder command for the terminal.

## Prerequisites

`GoRemind` has been developed on Ubuntu 16.10 and requires `libnotify-bin` and `at` to be installed.

## Usage

Examples:

```bash
$ goremind me in 10 minutes to take a break
Okay, I will remind you at 2017-05-05 00:15 to take a break
$ goremind me tomorrow at 8am to make coffee
Okay, I will remind you at 2017-05-06 08:00 to make coffee
```

## Limitations

I've written `GoRemind` to scratch my own itch, i.e. to be able to set Slack-like reminders in my terminal. If it doesn't work for you, please open an issue.

`GoRemind` currently supports only English language input.