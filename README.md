femtoterm
===

A bare-bones serial terminal that reads bytes from STDIN and outputs bytes on STDOUT. Useful to connect to Micropython devices, and other things.

```
Usage: femtoterm [<port-name>] [flags]

Arguments:
  [<port-name>]

Flags:
  -h, --help               Show context-sensitive help.
      --baudRate=115200    Baud rate for the connection
```

If run without a `<port-name>` argument it will attempt to list available serial port devices on your system.

Download
---
For a Windows x64 binary, get the latest nightly build from the last [GitHub Actions run](https://nightly.link/jstarpl/femtoterm/workflows/go/main).