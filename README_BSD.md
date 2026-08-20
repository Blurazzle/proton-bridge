
# Proton Mail Bridge on OpenBSD and FreeBSD

This document explains how to build and run the CLI version of Proton Mail Bridge on OpenBSD and FreeBSD.
> **Note:** This is an unofficial, independently maintained fork of Proton Mail Bridge with modifications for OpenBSD and FreeBSD, including BSD-specific build changes and keyring support. This is a hobby project and is not affiliated with or endorsed by Proton AG. **Your mileage may vary.**
## Limitations

The BSD build has the following limitations:

-   **CLI only:** The desktop GUI is not built.
    
-   **`pass` only:** `pass` is the only supported keyring backend. `gnome-keyring` is not supported.
    
-   **No FIDO2:** FIDO2 hardware security keys are not supported.
    

## OpenBSD

### Install dependencies

```sh
doas pkg_add bash gmake go password-store
```

### Build Bridge

Build the CLI version without the desktop GUI:

```sh
gmake build-nogui
```

The resulting binary is:

```text
bridge
```

## FreeBSD

### Install dependencies

```sh
pkg install git go password-store gmake bash
```

### Build Bridge

Build the CLI version without the desktop GUI:

```sh
gmake build-nogui
```

The resulting binary is:

```text
bridge
```

## Configure `pass`

The BSD build uses [`pass`](https://www.passwordstore.org/) as its only keyring backend.

`password-store` is the only runtime dependency. The other packages listed above are required for building.

### Generate a GPG key

Generate a GPG key using the parameters included with the source:

```sh
gpg --generate-key --batch --quiet utils-bsd/gpgparams
```

### Initialize `pass`

Initialize the password store for Proton Mail Bridge:

```sh
pass init proton-bridge
```

## Start Proton Mail Bridge
Start Bridge in its interactive shell:
```sh
./bridge -c
```
The `-c` option starts Bridge in interactive mode. It is required because the GUI is not built.

## Source Changes

The BSD build is based on the official Proton Mail Bridge source code, with a small number of BSD-specific changes:

-   Desktop GUI support is disabled.
    
-   `pass` is used as the only keyring backend.
    
-   `gnome-keyring` support is removed.
    
-   FIDO2 hardware security key support is removed.
    

## Tested On

This build has been tested on:

-   OpenBSD 7.9
    
-   FreeBSD 15.1
    

## Development Note

AI tools were used during development for OpenBSD- and Go-specific research, troubleshooting, suggestions, and README editing. AI was also used to generate empty FIDO stubs:

-   `internal/fido/fido_freebsd.go`
    
-   `internal/fido/fido_openbsd.go`
    

All changes were reviewed, adapted, and tested manually. I have extensive software development experience, particularly in web development and complex cloud environments, but less experience with OpenBSD and Go.
