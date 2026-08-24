# auto-license
a nice little single-file auto-licensing machine written in Go under the Apache license. 
absurdly fast, able to do the entire Linux kernel [its c, rust, GNU Asm, Python, and Shell [with extension] files] in ~7sec (max) using the default config as provided.
WILL choke on large sourcecode files, as i didnt implement buffer-on-read. (will fix later.)

<p align="center">
  <img src="edited.mp4" alt="auto-license demo run" width="100%" style="max-width: 800px; border-radius: 8px;" />
</p>

## Why create this?
* I wanted a tool that was dead simple to use and integrate into my projects.
* Also, I am too lazy to implement a tool such as add-license, and tools as such tend to be slow.

## Advantages
* Won't re-run itself on files that it detects that are already licensed [assuming you didn't switch license].
* Absurdly fast, up to 10k files/sec on a Intel i5-14400F [albeit with a SSD as well.]
* Zero external dependencies other than the Go stdlib.
* Can be configured to any programming language that supports one-liner comments that use 'prepend' syntax. [basically all of them]

## Disadvantages
* The filename auto-license uses for config can conflict with other files in the project, and it can't be changed without sourcecode modification.
* No TUI/CLI, only run it and it goes off.

## Usage

### 1. License text.
Put a file named 'license.txt' into the same directory as the binary.

This should hold whatever license you wish to apply to your code.

### 2. Config
The repo has a example config file named "auto_license-config-EXAMPLE.json"

You can use it [after renaming it to auto_license-config.json], however I recommend configuring other programming languages as it doesn't even have Go.

The filename must be 'auto_license-config.json', and the config has a JSON structure:

`
{
  "EXTENSION": "INSERT_COMMENT_FORMAT"
}
`

The extension must retain the dot, otherwise it will not work. [checking uses filepath.Ext]

You can add more types by simply adding a new JSON field, as it is parsed like a map[string]string.

The comment format refers to how you write 'prepend' comments, e.g using "//" for C-style languages.

### 3. Run the file.
If you are using binaries or even compiled it yourself, the binary goes into the working directory of whatever project you are working on.

(Note: it will crawl all subdirectories as well.(

Then, simply run the binary, ensuring that both the license and config are in the same directory as the binary. Done!


### License
Apache license 2.0. See LICENSE for details.
