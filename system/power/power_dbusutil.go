// Code generated by "dbusutil-gen -type Manager,Battery -import pkg.deepin.io/dde/api/powersupply/battery manager.go battery.go"; DO NOT EDIT.

package power

import (
	"pkg.deepin.io/dde/api/powersupply/battery"
)

func (v *Manager) setPropOnBattery(value bool) (changed bool) {
	if v.OnBattery != value {
		v.OnBattery = value
		v.emitPropChangedOnBattery(value)
		return true
	}
	return false
}

func (v *Manager) emitPropChangedOnBattery(value bool) error {
	return v.service.EmitPropertyChanged(v, "OnBattery", value)
}

func (v *Manager) setPropHasLidSwitch(value bool) (changed bool) {
	if v.HasLidSwitch != value {
		v.HasLidSwitch = value
		v.emitPropChangedHasLidSwitch(value)
		return true
	}
	return false
}

func (v *Manager) emitPropChangedHasLidSwitch(value bool) error {
	return v.service.EmitPropertyChanged(v, "HasLidSwitch", value)
}

func (v *Manager) setPropHasBattery(value bool) (changed bool) {
	if v.HasBattery != value {
		v.HasBattery = value
		v.emitPropChangedHasBattery(value)
		return true
	}
	return false
}

func (v *Manager) emitPropChangedHasBattery(value bool) error {
	return v.service.EmitPropertyChanged(v, "HasBattery", value)
}

func (v *Manager) setPropBatteryPercentage(value float64) (changed bool) {
	if v.BatteryPercentage != value {
		v.BatteryPercentage = value
		v.emitPropChangedBatteryPercentage(value)
		return true
	}
	return false
}

func (v *Manager) emitPropChangedBatteryPercentage(value float64) error {
	return v.service.EmitPropertyChanged(v, "BatteryPercentage", value)
}

func (v *Manager) setPropBatteryStatus(value battery.Status) (changed bool) {
	if v.BatteryStatus != value {
		v.BatteryStatus = value
		v.emitPropChangedBatteryStatus(value)
		return true
	}
	return false
}

func (v *Manager) emitPropChangedBatteryStatus(value battery.Status) error {
	return v.service.EmitPropertyChanged(v, "BatteryStatus", value)
}

func (v *Manager) setPropBatteryTimeToEmpty(value uint64) (changed bool) {
	if v.BatteryTimeToEmpty != value {
		v.BatteryTimeToEmpty = value
		v.emitPropChangedBatteryTimeToEmpty(value)
		return true
	}
	return false
}

func (v *Manager) emitPropChangedBatteryTimeToEmpty(value uint64) error {
	return v.service.EmitPropertyChanged(v, "BatteryTimeToEmpty", value)
}

func (v *Manager) setPropBatteryTimeToFull(value uint64) (changed bool) {
	if v.BatteryTimeToFull != value {
		v.BatteryTimeToFull = value
		v.emitPropChangedBatteryTimeToFull(value)
		return true
	}
	return false
}

func (v *Manager) emitPropChangedBatteryTimeToFull(value uint64) error {
	return v.service.EmitPropertyChanged(v, "BatteryTimeToFull", value)
}

func (v *Manager) setPropPowerSavingModeEnabled(value bool) (changed bool) {
	if v.PowerSavingModeEnabled != value {
		v.PowerSavingModeEnabled = value
		v.emitPropChangedPowerSavingModeEnabled(value)
		return true
	}
	return false
}

func (v *Manager) emitPropChangedPowerSavingModeEnabled(value bool) error {
	return v.service.EmitPropertyChanged(v, "PowerSavingModeEnabled", value)
}

func (v *Manager) setPropPowerSavingModeAuto(value bool) (changed bool) {
	if v.PowerSavingModeAuto != value {
		v.PowerSavingModeAuto = value
		v.emitPropChangedPowerSavingModeAuto(value)
		return true
	}
	return false
}

func (v *Manager) emitPropChangedPowerSavingModeAuto(value bool) error {
	return v.service.EmitPropertyChanged(v, "PowerSavingModeAuto", value)
}

func (v *Battery) setPropSysfsPath(value string) (changed bool) {
	if v.SysfsPath != value {
		v.SysfsPath = value
		v.emitPropChangedSysfsPath(value)
		return true
	}
	return false
}

func (v *Battery) emitPropChangedSysfsPath(value string) error {
	return v.service.EmitPropertyChanged(v, "SysfsPath", value)
}

func (v *Battery) setPropIsPresent(value bool) (changed bool) {
	if v.IsPresent != value {
		v.IsPresent = value
		v.emitPropChangedIsPresent(value)
		return true
	}
	return false
}

func (v *Battery) emitPropChangedIsPresent(value bool) error {
	return v.service.EmitPropertyChanged(v, "IsPresent", value)
}

func (v *Battery) setPropManufacturer(value string) (changed bool) {
	if v.Manufacturer != value {
		v.Manufacturer = value
		v.emitPropChangedManufacturer(value)
		return true
	}
	return false
}

func (v *Battery) emitPropChangedManufacturer(value string) error {
	return v.service.EmitPropertyChanged(v, "Manufacturer", value)
}

func (v *Battery) setPropModelName(value string) (changed bool) {
	if v.ModelName != value {
		v.ModelName = value
		v.emitPropChangedModelName(value)
		return true
	}
	return false
}

func (v *Battery) emitPropChangedModelName(value string) error {
	return v.service.EmitPropertyChanged(v, "ModelName", value)
}

func (v *Battery) setPropSerialNumber(value string) (changed bool) {
	if v.SerialNumber != value {
		v.SerialNumber = value
		v.emitPropChangedSerialNumber(value)
		return true
	}
	return false
}

func (v *Battery) emitPropChangedSerialNumber(value string) error {
	return v.service.EmitPropertyChanged(v, "SerialNumber", value)
}

func (v *Battery) setPropName(value string) (changed bool) {
	if v.Name != value {
		v.Name = value
		v.emitPropChangedName(value)
		return true
	}
	return false
}

func (v *Battery) emitPropChangedName(value string) error {
	return v.service.EmitPropertyChanged(v, "Name", value)
}

func (v *Battery) setPropTechnology(value string) (changed bool) {
	if v.Technology != value {
		v.Technology = value
		v.emitPropChangedTechnology(value)
		return true
	}
	return false
}

func (v *Battery) emitPropChangedTechnology(value string) error {
	return v.service.EmitPropertyChanged(v, "Technology", value)
}

func (v *Battery) setPropEnergy(value float64) (changed bool) {
	if v.Energy != value {
		v.Energy = value
		v.emitPropChangedEnergy(value)
		return true
	}
	return false
}

func (v *Battery) emitPropChangedEnergy(value float64) error {
	return v.service.EmitPropertyChanged(v, "Energy", value)
}

func (v *Battery) setPropEnergyFull(value float64) (changed bool) {
	if v.EnergyFull != value {
		v.EnergyFull = value
		v.emitPropChangedEnergyFull(value)
		return true
	}
	return false
}

func (v *Battery) emitPropChangedEnergyFull(value float64) error {
	return v.service.EmitPropertyChanged(v, "EnergyFull", value)
}

func (v *Battery) setPropEnergyFullDesign(value float64) (changed bool) {
	if v.EnergyFullDesign != value {
		v.EnergyFullDesign = value
		v.emitPropChangedEnergyFullDesign(value)
		return true
	}
	return false
}

func (v *Battery) emitPropChangedEnergyFullDesign(value float64) error {
	return v.service.EmitPropertyChanged(v, "EnergyFullDesign", value)
}

func (v *Battery) setPropEnergyRate(value float64) (changed bool) {
	if v.EnergyRate != value {
		v.EnergyRate = value
		v.emitPropChangedEnergyRate(value)
		return true
	}
	return false
}

func (v *Battery) emitPropChangedEnergyRate(value float64) error {
	return v.service.EmitPropertyChanged(v, "EnergyRate", value)
}

func (v *Battery) setPropVoltage(value float64) (changed bool) {
	if v.Voltage != value {
		v.Voltage = value
		v.emitPropChangedVoltage(value)
		return true
	}
	return false
}

func (v *Battery) emitPropChangedVoltage(value float64) error {
	return v.service.EmitPropertyChanged(v, "Voltage", value)
}

func (v *Battery) setPropPercentage(value float64) (changed bool) {
	if v.Percentage != value {
		v.Percentage = value
		v.emitPropChangedPercentage(value)
		return true
	}
	return false
}

func (v *Battery) emitPropChangedPercentage(value float64) error {
	return v.service.EmitPropertyChanged(v, "Percentage", value)
}

func (v *Battery) setPropCapacity(value float64) (changed bool) {
	if v.Capacity != value {
		v.Capacity = value
		v.emitPropChangedCapacity(value)
		return true
	}
	return false
}

func (v *Battery) emitPropChangedCapacity(value float64) error {
	return v.service.EmitPropertyChanged(v, "Capacity", value)
}

func (v *Battery) setPropStatus(value battery.Status) (changed bool) {
	if v.Status != value {
		v.Status = value
		v.emitPropChangedStatus(value)
		return true
	}
	return false
}

func (v *Battery) emitPropChangedStatus(value battery.Status) error {
	return v.service.EmitPropertyChanged(v, "Status", value)
}

func (v *Battery) setPropTimeToEmpty(value uint64) (changed bool) {
	if v.TimeToEmpty != value {
		v.TimeToEmpty = value
		v.emitPropChangedTimeToEmpty(value)
		return true
	}
	return false
}

func (v *Battery) emitPropChangedTimeToEmpty(value uint64) error {
	return v.service.EmitPropertyChanged(v, "TimeToEmpty", value)
}

func (v *Battery) setPropTimeToFull(value uint64) (changed bool) {
	if v.TimeToFull != value {
		v.TimeToFull = value
		v.emitPropChangedTimeToFull(value)
		return true
	}
	return false
}

func (v *Battery) emitPropChangedTimeToFull(value uint64) error {
	return v.service.EmitPropertyChanged(v, "TimeToFull", value)
}

func (v *Battery) setPropUpdateTime(value int64) (changed bool) {
	if v.UpdateTime != value {
		v.UpdateTime = value
		v.emitPropChangedUpdateTime(value)
		return true
	}
	return false
}

func (v *Battery) emitPropChangedUpdateTime(value int64) error {
	return v.service.EmitPropertyChanged(v, "UpdateTime", value)
}
