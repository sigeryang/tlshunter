// Code generated by "stringer -type=RiskType"; DO NOT EDIT.

package main

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[RiskNSCMissing-0]
	_ = x[RiskCleartext-1]
	_ = x[RiskUserAnchors-2]
	_ = x[RiskAnchorsOverridePinning-3]
	_ = x[RiskUnpinned-4]
	_ = x[RiskPinningExpiration-5]
	_ = x[RiskProxyAnchors-6]
	_ = x[RiskMalformedNSC-7]
}

const _RiskType_name = "RiskNSCMissingRiskCleartextRiskUserAnchorsRiskAnchorsOverridePinningRiskUnpinnedRiskPinningExpirationRiskProxyAnchorsRiskMalformedNSC"

var _RiskType_index = [...]uint8{0, 14, 27, 42, 68, 80, 101, 117, 133}

func (i RiskType) String() string {
	if i < 0 || i >= RiskType(len(_RiskType_index)-1) {
		return "RiskType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _RiskType_name[_RiskType_index[i]:_RiskType_index[i+1]]
}
