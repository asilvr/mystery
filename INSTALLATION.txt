# Mystery

## Installation

Thanks for downloading Mystery! The installation instructions are OS-specific.
If you are reading this but don't have the tool, download it from 
https://github.com/asilvr/mystery/releases.

### macOS

on macOS, you will want to place the `mystery` tool in your `PATH`. This is
where all of your command line tools live. If you're not familiar with `PATH`,
run this command:

```
$ echo $PATH
```

This outputs your path. You can see that the command `ls`, for example, is in
the `PATH` by running this:

```
$ which ls
```

In this guide, we'll install `mystery` into `/usr/local/bin`. Note that you
may be asked to enter your password to perform this action:

```
$ cp mystery /usr/local/bin/mystery
$ source ~/.bash_profile
```

Now try to use `mystery`:

```
$ mystery --help
```

You're up and running! See the help command for usage. The most likely command
you'll run is `mystery generate`. Enjoy!