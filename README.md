# Anker Mouse Tool

A set of tools to set up the Anker 8200 DPI Programmable Gaming Mouse.

Released under the terms of [MIT Licence][licence].

## The Device

The Anker 8200 DPI Programmable Gaming mouse is a device with 9
programmable buttons and a scroll wheel. The provided software (for
Windows only) allows programming macros and single keys aon all of the
buttons, as well as allow one of the button to switch between two
profiles and four preconfigured resoluion settings.

The device also has a configurable light, which can be configured to
blink at different speeds and intensities.

## Tools

The repository contains two tools: `anker-mouse-replayer` and
`anker-mouse-light`.

The former is an exploratory tool developed while trying to imitate
the commands sent by the original Windows tool.

The latter is a tool that allows setting the current profile's device
light without having to reconnect the device, and can possibly be used
to signal to the user.

## Author

Diego Elio Pettenò <flameeyes@flameeyes.eu>

[licence]: https://opensource.org/licenses/mit-license.php
