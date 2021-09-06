# gostatus
<img src="https://i.imgur.com/p9HO9Cu.png" width=60%/>

gostatus is an ergonomic status monitor alternative to [slstatus](https://github.com/drkhsh/slstatus).

## Features
* Single configuration file (config.go)
* Simple component-based organization 
* Core functionality in less than 80 lines of code
* Extremely extensible
* Native [dwm extrabar](https://dwm.suckless.org/patches/extrabar/) support

## Installation
`go install` 

..or..

`go build && sudo mv ./gostatus /usr/local/bin/`


## Usage
Simply run `gostatus` on startup, just like slstatus (for instance, in your `.profile`)

## Configuration
All configuration is extracted to `config.go`. Change `UPDATES_PER_SECOND` to speed up
or slow down the refresh rate and alter `buildTopBar` and `buildBottomBar` to customize
layout and drawing.

## Extension
`model.Component` is a single-method interface, allowing any object implementing `Render() string`
to be included in a bar. If you write a new component or improve one of the few existing ones,
be sure to submit a PR and help this project grow!

*NOTE: model.Bar implements the Component interface, allowing nesting of bars for organization.*

## TODO
- [ ] Support single-bar usage
- [ ] Support other RSS readers ([RSS Guard](https://github.com/martinrotter/rssguard) support is currently hard-coded)
- [ ] Add multi-core support to `CPU`
- [ ] Replace `os/exec` usage with direct library call
- [ ] Port all applicable [slstatus components](https://github.com/drkhsh/slstatus/tree/master/components)

## Support
<a href='https://ko-fi.com/U7U84VTAW' target='_blank'><img height='36' style='border:0px;height:36px;' src='https://cdn.ko-fi.com/cdn/kofi2.png?v=2' border='0' alt='Buy Me a Coffee at ko-fi.com' /></a>
