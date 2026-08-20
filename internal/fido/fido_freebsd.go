// Copyright (c) 2026 Proton AG
//
// This file is part of Proton Mail Bridge.
//
// Proton Mail Bridge is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// Proton Mail Bridge is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with Proton Mail Bridge. If not, see <https://www.gnu.org/licenses/>.

//go:build freebsd

package fido

import (
	"context"
	"errors"

	"github.com/ProtonMail/go-proton-api"
)

var errFidoUnsupported = errors.New("FIDO2 hardware key authentication is not supported on OpenBSD")

func IsPinSupported() (bool, error) {
	return false, errFidoUnsupported
}

func AuthWithHardwareKeyGUI(
	_ context.Context,
	_ *proton.Client,
	_ proton.Auth,
	_ chan struct{},
	_ chan struct{},
	_ string,
) error {
	return errFidoUnsupported
}

func AuthWithHardwareKeyCLI(
	_ CLIProvider,
	_ *proton.Client,
	_ proton.Auth,
) error {
	return errFidoUnsupported
}
