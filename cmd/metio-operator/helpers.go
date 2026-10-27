// SPDX-License-Identifier: Apache-2.0

package main

import "os"

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}