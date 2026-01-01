# ReTag 🎵

**ReTag** is a simple, cross-platform **CLI MP3 ID3 tag editor** written in **Go**.
It’s built to be fast, minimal, and easy to use. 

---

## ✨ Features

* Cross-platform (Windows, macOS, Linux)
* Edit MP3 ID3 tags from the command line
* Set title, artist, album, and cover art
* Clean and simple CLI interface
* Written in Go → single binary

---

## 📦 Installation

### Build from source

```bash
git clone https://github.com/ReRixV1/ReTag.git
cd retag
go build -o retag
```

This produces a `retag` binary you can move anywhere in your `$PATH`.
**Move to /usr/local/bin on MacOS / Linux**

---

## 🚀 Usage

Basic command structure:

```bash
retag <file.mp3> [flags]
```

### Example

```bash
retag extra.mp3 -a "Juice WRLD" -t "Extra" -c "extra.png"
```

This will:

* Set the **artist** to `Juice WRLD`
* Set the **title** to `Extra`
* Set the **cover art** to `extra.png`

> ⚠️ Currently, only **PNG** files are supported for cover art.

---

## 🏷️ Flags

| Alias | Name   | Description       |
| ----: | ------ | ----------------- |
|  `-t` | title  | Track title       |
|  `-a` | artist | Artist name       |
| `-al` | album  | Album name        |
|  `-c` | cover  | Cover image (PNG) |

You can combine any flags in a single command.

---

## 🛠️ Notes

* Only **MP3** files are supported
* Cover images must be **PNG**
* Existing tags will be overwritten if provided

---

## 📄 License

MIT License — use it however you want.
