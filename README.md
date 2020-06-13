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

### `anker-mouse-replayer`

An exploratory tool developed while trying to imitate the commands
sent by the original Windows tool. It is not designed for cleanliness,
but rather for an ease of changing the data in the reports.

### `anker-mouse-light`

Allows setting the current device light parameters (color, brightness,
breath speed).

These settings are temporary and not saved onto the device profile,
and will be reset ot the profile value once the device is
disconnected.

This can be used for signalling information to the user.

### `anker-mouse-profile`

Allows switching between the two configured profiles in the device.

## Author

Diego Elio Pettenò <flameeyes@flameeyes.com>

[licence]: https://opensource.org/licenses/mit-license.php
