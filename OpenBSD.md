# Proton Mail Bridge on OpenBSD

This document describes how to build and run Proton Mail Bridge on OpenBSD.

## Building

Install the required build dependencies:

```sh
doas pkg_add bash gmake go password-store
````

Build the command-line version without the desktop GUI:

```sh
gmake build-nogui
```

The resulting binary will be:

```text
proton-bridge
```

## Running Proton Mail Bridge

The OpenBSD version of Proton Mail Bridge uses **`pass`** as its keyring backend.

> **Note:** `gnome-keyring` support has been removed from the OpenBSD build.

### Install `pass`

Install the OpenBSD `password-store` package:

```sh
doas pkg_add password-store
```

This is the recommended and tested password-store backend for OpenBSD.

### Initialize the password store

Generate a GPG key using the parameters provided with the source:

```sh
gpg --generate-key --batch --quiet utils-bsd/gpgparams
```

Initialize `pass` for Proton Mail Bridge:

```sh
pass init proton-bridge
```

### Start Proton Mail Bridge

Start Bridge in command-line mode:

```sh
./bridge -c
```

The `-c` option starts Bridge without the desktop interface.

## FIDO2 / Hardware Security Keys

FIDO2 hardware security-key authentication is currently **not supported on OpenBSD**.

The OpenBSD build does not include the FIDO2/libfido2 functionality. If FIDO authentication is requested, Bridge will return an error indicating that FIDO2 hardware-key authentication is unsupported.

## Known Limitations

The OpenBSD build currently has the following limitations:

* No desktop GUI.
* `pass` is used as the keyring backend.
* `gnome-keyring` is not supported.
* FIDO2 hardware security keys are not supported.

## Tested Environment

This build has been tested on:

* OpenBSD 7.9
* amd64

## Build Command

In short:

```sh
doas pkg_add bash gmake go password-store

gmake build-nogui

gpg --generate-key --batch --quiet utils-bsd/gpgparams
pass init proton-bridge

./bridge -c
```
## Development Note
AI tools were used as an assisting tool during development, primarily for OpenBSD- and Go-specific research, troubleshooting, and suggestions.
The changes were reviewed, adapted, and tested manually. I have extensive software development experience, particularly in web development managing Linux server, but less experience with OpenBSD and Go.