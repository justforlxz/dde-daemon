package power

import (
	"bufio"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"pkg.deepin.io/lib/dbus1"
	"pkg.deepin.io/lib/dbusutil"
)

// laptop mode tools config file
const lmtConfigFile = "/etc/laptop-mode/laptop-mode.conf"
const laptopModeBin = "/usr/sbin/laptop_mode"

const (
	lmtConfigAuto     = 1
	lmtConfigEnabled  = 2
	lmtConfigDisabled = 3
)

func setLMTConfig(mode int) error {
	_, err := os.Stat(laptopModeBin)
	if err != nil {
		return err
	}

	lines, err := loadLmtConfig()
	if err != nil {
		logger.Warning(err)
	}

	dict := make(map[string]string)
	switch mode {
	case lmtConfigAuto:
		dict["ENABLE_LAPTOP_MODE_TOOLS"] = "1"
		dict["ENABLE_LAPTOP_MODE_ON_BATTERY"] = "1"
		dict["ENABLE_LAPTOP_MODE_ON_AC"] = "0"
	case lmtConfigEnabled:
		dict["ENABLE_LAPTOP_MODE_TOOLS"] = "1"
		dict["ENABLE_LAPTOP_MODE_ON_BATTERY"] = "1"
		dict["ENABLE_LAPTOP_MODE_ON_AC"] = "1"
	case lmtConfigDisabled:
		dict["ENABLE_LAPTOP_MODE_TOOLS"] = "1"
		dict["ENABLE_LAPTOP_MODE_ON_BATTERY"] = "0"
		dict["ENABLE_LAPTOP_MODE_ON_AC"] = "0"
	}
	lines, changed := modifyLMTConfig(lines, dict)
	if changed {
		logger.Debug("write LMT Config")
		err = writeLmtConfig(lines)
		if err != nil {
			return err
		}
		err = reloadLaptopModeService()
		if err != nil {
			return err
		}
	}
	return nil
}

func reloadLaptopModeService() error {
	systemBus, err := dbus.SystemBus()
	if err != nil {
		return err
	}
	systemdObj := systemBus.Object("org.freedesktop.systemd1", "/org/freedesktop/systemd1")
	return systemdObj.Call("org.freedesktop.systemd1.Manager.ReloadUnit",
		dbus.FlagNoAutoStart, "laptop-mode.service", "replace").Err
}

func modifyLMTConfig(lines []string, dict map[string]string) ([]string, bool) {
	var changed bool
	for idx := range lines {
		line := lines[idx]
		for key, value := range dict {
			if strings.HasPrefix(line, key) {
				newLine := key + "=" + value
				if line != newLine {
					changed = true
					lines[idx] = newLine
				}
				delete(dict, key)
			}
		}
		if len(dict) == 0 {
			break
		}
	}
	if len(dict) > 0 {
		for key, value := range dict {
			newLine := key + "=" + value
			lines = append(lines, newLine)
		}
		changed = true
	}
	return lines, changed
}

func loadLmtConfig() ([]string, error) {
	f, err := os.Open(lmtConfigFile)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(bufio.NewReader(f))
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	return lines, nil
}

func writeLmtConfig(lines []string) error {
	tempFile, err := writeLmtConfigTemp(lines)
	if err != nil {
		if tempFile != "" {
			os.Remove(tempFile)
		}
		return err
	}
	return os.Rename(tempFile, lmtConfigFile)
}

func writeLmtConfigTemp(lines []string) (string, error) {
	dir := filepath.Dir(lmtConfigFile)
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return "", err
	}

	f, err := ioutil.TempFile(dir, "laptop-mode.conf")
	logger.Debug("writeLmtConfig temp file", f.Name())
	if err != nil {
		return "", err
	}
	defer f.Close()
	err = f.Chmod(0644)
	if err != nil {
		return f.Name(), err
	}

	bufWriter := bufio.NewWriter(f)
	for _, line := range lines {
		bufWriter.WriteString(line)
		bufWriter.WriteByte('\n')
	}
	return f.Name(), bufWriter.Flush()
}

func (m *Manager) writePowerSavingModeEnabledCb(write *dbusutil.PropertyWrite) *dbus.Error {
	logger.Debug("set laptop mode enabled", write.Value)

	m.PropsMu.Lock()
	if m.PowerSavingModeAuto {
		m.PowerSavingModeAuto = false
		m.emitPropChangedPowerSavingModeAuto(false)
	}
	m.PropsMu.Unlock()

	enabled := write.Value.(bool)
	var err error
	if enabled {
		err = setLMTConfig(lmtConfigEnabled)
	} else {
		err = setLMTConfig(lmtConfigDisabled)
	}

	if err != nil {
		logger.Warning("failed to set LMT config:", err)
	}

	return dbusutil.ToError(err)
}

func (m *Manager) writePowerSavingModeAutoCb(write *dbusutil.PropertyWrite) *dbus.Error {
	logger.Debug("set laptop mode auto switch", write.Value)

	autoSwitch := write.Value.(bool)
	var err error
	if autoSwitch {
		m.PropsMu.Lock()
		m.setPropPowerSavingModeEnabled(m.OnBattery)
		m.PropsMu.Unlock()

		err = setLMTConfig(lmtConfigAuto)

	} else {
		m.PropsMu.RLock()
		enabled := m.PowerSavingModeEnabled
		m.PropsMu.RUnlock()

		if enabled {
			err = setLMTConfig(lmtConfigEnabled)
		} else {
			err = setLMTConfig(lmtConfigDisabled)
		}
	}

	if err != nil {
		logger.Warning("failed to set LMT config:", err)
	}

	return dbusutil.ToError(err)
}
