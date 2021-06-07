# gostatus
![image](https://i.imgur.com/BHhU7yK.png)

gostatus is an ergonmic status monitor alternative to [slstatus](https://github.com/drkhsh/slstatus).

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
- [ ] Port all applicable [slstatus components](https://github.com/drkhsh/slstatus/tree/master/components)

## Support
<script type='text/javascript' src='https://storage.ko-fi.com/cdn/widget/Widget_2.js'></script><script type='text/javascript'>kofiwidget2.init('Buy me a coffee', '#29abe0', 'U7U84VTAW');kofiwidget2.draw();</script>
