// -*- Mode: Go; indent-tabs-mode: t -*-

/*
 * Copyright (C) 2019 Canonical Ltd
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License version 3 as
 * published by the Free Software Foundation.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 */

package secboot

import (
	"github.com/canonical/go-tpm2"
)

const (
	lockNVHandle     tpm2.Handle = 0x01801100 // Global NV handle for locking access to sealed key objects
	lockNVDataHandle tpm2.Handle = 0x01801101 // NV index containing policy data for lockNVHandle

	// SHA-256 is mandatory to exist on every PC-Client TPM
	// XXX: Maybe dynamically select algorithms based on what's available on the device?
	defaultSessionHashAlgorithm tpm2.HashAlgorithmId = tpm2.HashAlgorithmSHA256
)
