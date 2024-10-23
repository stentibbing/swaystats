# Swaybar status line data source 

This project will become a good companion to sway and it's family of components. As it is, it's not yet packing a lot
of features and is under active development. So it is in very early pre-aplha state.

## Motivation
I love Sway and all of it's companioning components. Swaybar is really fast but it does not provide enough data for my
personal taste. So this project is there to fill that gap. The idea of this project is to become fully configurable
status line on swaybar, implementing the same configuring language as sway and it's family of components. 

## Implemented features
Not a lot :) 

### Renderer
Renderer gathers all the blocks, marshals them and then sends the string to stdout. If any update is sent to channel,
the thing does this again.

### Modules
It has small small structure for implementing modules. Modules or Blocks (visible pieces of information on statusline)
must implement Block interface. Block interface provides a function to update the block (Update) and a function to 
construct information for the block. Each Block will run on it's own Go Routine. Even tough Update function itself 
determines when to update a block, polling should be the last reserve for things unable to fire events, like clock.

## Implemented Modules
- Time

## TODO:

### Features
- Configuration
  Read config from the same locations as sway does (eg. "~/.config/swaystats" etc.). Parse the config, set general config
  based on it and send module config to modules while registering them. Has to support multiple modules of the same kind.
  Always leave as much of design and configuration (eg. block separators) to swaybar as possible.

### Modules
- Volume (in progress, currently going to support Pulse Audio's pactl only).
- Battery status
- Wifi status
- Memory
- CPU
- Temps

