/**
 * Copyright (C) 2014 Deepin Technology Co., Ltd.
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 3 of the License, or
 * (at your option) any later version.
 **/

package dock

import (
	"gir/gio-2.0"
	"path/filepath"
	"pkg.deepin.io/lib/dbus"
	"pkg.deepin.io/lib/dbus/property"
)

func (m *DockManager) initHideStateManager() error {
	m.hideStateManager = NewHideStateManager()
	m.hideStateManager.dockRect = m.dockRect
	logger.Debug("initHideStateManager dockRect", m.hideStateManager.dockRect)
	m.hideStateManager.mode = m.hideMode
	m.hideStateManager.initHideState()
	err := dbus.InstallOnSession(m.hideStateManager)
	return err
}

func (m *DockManager) initDockProperty() error {
	m.dockProperty = NewDockProperty(m)
	err := dbus.InstallOnSession(m.dockProperty)
	return err
}

func (m *DockManager) loadCache() error {
	var err error
	m.desktopWindowsMapCacheManager, err = newDesktopWindowsMapCacheManager(filepath.Join(cacheDir, "desktopWindowsMapCache.gob"))
	if err != nil {
		return err
	}
	m.desktopHashFileMapCacheManager, err = newDesktopHashFileMapCacheManager(filepath.Join(cacheDir, "desktopHashFileMapCache.gob"))
	if err != nil {
		return err
	}
	return nil
}

func (m *DockManager) initEntries() {
	// init entries
	m.desktopWindowsMapCacheManager.SetAutoSaveEnabled(false)
	m.desktopHashFileMapCacheManager.SetAutoSaveEnabled(false)

	m.initDockedApps()
	m.initClientList()

	m.desktopWindowsMapCacheManager.SetAutoSaveEnabled(true)
	m.desktopWindowsMapCacheManager.AutoSave()
	m.desktopHashFileMapCacheManager.SetAutoSaveEnabled(true)
	m.desktopHashFileMapCacheManager.AutoSave()
}

func (m *DockManager) connectSettingKeyChanged(key string, handler func(*gio.Settings, string)) {
	m.settings.Connect("changed::"+key, handler)
}

func (m *DockManager) listenSettingsChanged() {
	// listen hide mode change
	m.connectSettingKeyChanged(HideModeKey, func(g *gio.Settings, key string) {
		mode := HideModeType(g.GetEnum(key))
		logger.Debug(key, "changed to", mode)

		m.hideStateManager.updateHideMode(mode)
	})

	// listen display mode change
	m.connectSettingKeyChanged(DisplayModeKey, func(g *gio.Settings, key string) {
		mode := DisplayModeType(g.GetEnum(key))
		logger.Debug(key, "changed to", mode)

		m.dockHeight = getDockHeightByDisplayMode(mode)
		m.updateDockRect()
	})
}

func (m *DockManager) init() error {
	var err error

	m.settings = gio.NewSettings(dockSchema)

	m.HideMode = property.NewGSettingsEnumProperty(m, "HideMode", m.settings, HideModeKey)
	m.DisplayMode = property.NewGSettingsEnumProperty(m, "DisplayMode", m.settings, DisplayModeKey)

	// ensure init display after init setting
	err = m.initDisplay()
	if err != nil {
		return err
	}
	logger.Info("initialize display done")

	err = m.initHideStateManager()
	if err != nil {
		return err
	}
	logger.Info("initialize hide state manager done")

	err = m.initDockProperty()
	if err != nil {
		return err
	}
	logger.Info("initialize dock property done")

	m.listenSettingsChanged()

	m.appIdFilterGroup = NewAppIdFilterGroup()
	err = m.loadCache()
	if err != nil {
		return err
	}
	m.dockedAppManager = NewDockedAppManager(m)
	m.initEntries()
	err = dbus.InstallOnSession(m.dockedAppManager)
	if err != nil {
		return err
	}

	err = dbus.InstallOnSession(m)
	if err != nil {
		return err
	}
	return nil
}