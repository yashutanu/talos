// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

// Code generated by "deep-copy -type MemoryModuleSpec -type PCIDeviceSpec -type PCIDriverRebindConfigSpec -type PCIDriverRebindStatusSpec -type PCRStatusSpec -type ProcessorSpec -type SystemInformationSpec -header-file ../../../../hack/boilerplate.txt -o deep_copy.generated.go ."; DO NOT EDIT.

package hardware

// DeepCopy generates a deep copy of MemoryModuleSpec.
func (o MemoryModuleSpec) DeepCopy() MemoryModuleSpec {
	var cp MemoryModuleSpec = o
	return cp
}

// DeepCopy generates a deep copy of PCIDeviceSpec.
func (o PCIDeviceSpec) DeepCopy() PCIDeviceSpec {
	var cp PCIDeviceSpec = o
	return cp
}

// DeepCopy generates a deep copy of PCIDriverRebindConfigSpec.
func (o PCIDriverRebindConfigSpec) DeepCopy() PCIDriverRebindConfigSpec {
	var cp PCIDriverRebindConfigSpec = o
	return cp
}

// DeepCopy generates a deep copy of PCIDriverRebindStatusSpec.
func (o PCIDriverRebindStatusSpec) DeepCopy() PCIDriverRebindStatusSpec {
	var cp PCIDriverRebindStatusSpec = o
	return cp
}

// DeepCopy generates a deep copy of PCRStatusSpec.
func (o PCRStatusSpec) DeepCopy() PCRStatusSpec {
	var cp PCRStatusSpec = o
	return cp
}

// DeepCopy generates a deep copy of ProcessorSpec.
func (o ProcessorSpec) DeepCopy() ProcessorSpec {
	var cp ProcessorSpec = o
	return cp
}

// DeepCopy generates a deep copy of SystemInformationSpec.
func (o SystemInformationSpec) DeepCopy() SystemInformationSpec {
	var cp SystemInformationSpec = o
	return cp
}
