/**
 * Copyright (C) 2014 Deepin Technology Co., Ltd.
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 3 of the License, or
 * (at your option) any later version.
 **/

package main

import (
	"os"

	// modules:
	_ "pkg.deepin.io/dde/daemon/accounts"
	_ "pkg.deepin.io/dde/daemon/apps"
	_ "pkg.deepin.io/dde/daemon/system/gesture"
	_ "pkg.deepin.io/dde/daemon/system/power"
	_ "pkg.deepin.io/dde/daemon/system/swapsched"

	"gir/glib-2.0"
	"pkg.deepin.io/dde/daemon/loader"
	"pkg.deepin.io/lib"
	"pkg.deepin.io/lib/dbus"
	. "pkg.deepin.io/lib/gettext"
	"pkg.deepin.io/lib/log"
)

var logger = log.NewLogger("daemon/dde-system-daemon")

func main() {
	logger.BeginTracing()
	defer logger.EndTracing()

	if !lib.UniqueOnSystem("com.deepin.daemon") {
		logger.Warning("There already has an dde daemon running.")
		return
	}

	// fix no PATH when was launched by dbus
	if os.Getenv("PATH") == "" {
		logger.Warning("No PATH found, manual special")
		os.Setenv("PATH", "/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin")
	}

	InitI18n()
	Textdomain("dde-daemon")

	logger.SetRestartCommand("/usr/lib/deepin-daemon/dde-system-daemon")

	loader.StartAll()
	defer loader.StopAll()

	dbus.DealWithUnhandledMessage()
	// NOTE: system/power module requires glib loop
	go glib.StartLoop()

	if err := dbus.Wait(); err != nil {
		logger.Errorf("Lost dbus: %v", err)
		os.Exit(-1)
	} else {
		os.Exit(0)
	}
}
