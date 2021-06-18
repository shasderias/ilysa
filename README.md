# Ilysa
Ilysa is a Go library that generates lighting events for Beat Saber beatmaps. In
particular, Ilysa is designed for generating Chroma events.

Mapping knowledge is required. Ilysa may not be the tool for you if:
- you have never placed a Chroma event in ChroMapper;
- ```_eventType``` and ```_eventValue``` for foreign to you;
- 

## Ilysa is/does not ...
* generate any other beatmap elements (for walls, you probably want spookyGh0st's [Beatwalls](https://github.com/spookyGh0st/beatwalls#readme]))
* 

# Getting Started
## Requirements
* a working Go installation
* a working Git installation
* a code editor (these instructions are tested with Visual Studio Code and the author uses Goland)
* a beatmap with *all* required BPM blocks placed

### Go Installation
Follow the instructions at:
* https://golang.org/doc/install
* https://golang.org/doc/tutorial/getting-started

### Git Installation

### Code Editor

### Beatmap
Ilysa works in BPM adjusted beats - the beat numbers displayed in MMA2 or ChroMapper. If you do not
place all the required BPM blocks before starting, your lighting events will probably be mistimed if
any BPM blocks are added after.

# Walkthrough
**Ilysa will replace all lighting events in the selected difficulty. Please dedicate a copy for use
with Ilysa and make backups!**

## Preliminaries 
Create a new empty directory to hold your Ilysa project.




# Convention

T - (0 to 1)

B - scaled beats

LightID - 1-indexed

# Tips and Tricks
## Visual Studio Code Keyboard Shortcuts
Ctrl-Shift-Space
Ctrl-Space

# Resources
https://www.desmos.com/calculator
