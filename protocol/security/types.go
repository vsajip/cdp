// Code generated by cdpgen. DO NOT EDIT.

package security

import (
	"encoding/json"
	"errors"
	"fmt"
)

// CertificateID An internal certificate ID value.
type CertificateID int

// MixedContentType A description of mixed content (HTTP resources on HTTPS pages), as defined by https://www.w3.org/TR/mixed-content/#categories
type MixedContentType int

// MixedContentType as enums.
const (
	MixedContentTypeNotSet MixedContentType = iota
	MixedContentTypeBlockable
	MixedContentTypeOptionallyBlockable
	MixedContentTypeNone
)

// Valid returns true if enum is set.
func (e MixedContentType) Valid() bool {
	return e >= 1 && e <= 3
}

func (e MixedContentType) String() string {
	switch e {
	case 0:
		return "MixedContentTypeNotSet"
	case 1:
		return "blockable"
	case 2:
		return "optionally-blockable"
	case 3:
		return "none"
	}
	return fmt.Sprintf("MixedContentType(%d)", e)
}

// MarshalJSON encodes enum into a string or null when not set.
func (e MixedContentType) MarshalJSON() ([]byte, error) {
	if e == 0 {
		return []byte("null"), nil
	}
	if !e.Valid() {
		return nil, errors.New("security.MixedContentType: MarshalJSON on bad enum value: " + e.String())
	}
	return json.Marshal(e.String())
}

// UnmarshalJSON decodes a string value into a enum.
func (e *MixedContentType) UnmarshalJSON(data []byte) error {
	switch string(data) {
	case "null":
		*e = 0
	case "\"blockable\"":
		*e = 1
	case "\"optionally-blockable\"":
		*e = 2
	case "\"none\"":
		*e = 3
	default:
		return fmt.Errorf("security.MixedContentType: UnmarshalJSON on bad input: %s", data)
	}
	return nil
}

// State The security level of a page or resource.
type State int

// State as enums.
const (
	StateNotSet State = iota
	StateUnknown
	StateNeutral
	StateInsecure
	StateWarning
	StateSecure
	StateInfo
)

// Valid returns true if enum is set.
func (e State) Valid() bool {
	return e >= 1 && e <= 6
}

func (e State) String() string {
	switch e {
	case 0:
		return "StateNotSet"
	case 1:
		return "unknown"
	case 2:
		return "neutral"
	case 3:
		return "insecure"
	case 4:
		return "warning"
	case 5:
		return "secure"
	case 6:
		return "info"
	}
	return fmt.Sprintf("State(%d)", e)
}

// MarshalJSON encodes enum into a string or null when not set.
func (e State) MarshalJSON() ([]byte, error) {
	if e == 0 {
		return []byte("null"), nil
	}
	if !e.Valid() {
		return nil, errors.New("security.State: MarshalJSON on bad enum value: " + e.String())
	}
	return json.Marshal(e.String())
}

// UnmarshalJSON decodes a string value into a enum.
func (e *State) UnmarshalJSON(data []byte) error {
	switch string(data) {
	case "null":
		*e = 0
	case "\"unknown\"":
		*e = 1
	case "\"neutral\"":
		*e = 2
	case "\"insecure\"":
		*e = 3
	case "\"warning\"":
		*e = 4
	case "\"secure\"":
		*e = 5
	case "\"info\"":
		*e = 6
	default:
		return fmt.Errorf("security.State: UnmarshalJSON on bad input: %s", data)
	}
	return nil
}

// StateExplanation An explanation of an factor contributing to the security state.
type StateExplanation struct {
	SecurityState    State            `json:"securityState"`    // Security state representing the severity of the factor being explained.
	Summary          string           `json:"summary"`          // Short phrase describing the type of factor.
	Description      string           `json:"description"`      // Full text explanation of the factor.
	MixedContentType MixedContentType `json:"mixedContentType"` // The type of mixed content described by the explanation.
	Certificate      []string         `json:"certificate"`      // Page certificate.
}

// InsecureContentStatus Information about insecure content on the page.
type InsecureContentStatus struct {
	RanMixedContent                bool  `json:"ranMixedContent"`                // True if the page was loaded over HTTPS and ran mixed (HTTP) content such as scripts.
	DisplayedMixedContent          bool  `json:"displayedMixedContent"`          // True if the page was loaded over HTTPS and displayed mixed (HTTP) content such as images.
	ContainedMixedForm             bool  `json:"containedMixedForm"`             // True if the page was loaded over HTTPS and contained a form targeting an insecure url.
	RanContentWithCertErrors       bool  `json:"ranContentWithCertErrors"`       // True if the page was loaded over HTTPS without certificate errors, and ran content such as scripts that were loaded with certificate errors.
	DisplayedContentWithCertErrors bool  `json:"displayedContentWithCertErrors"` // True if the page was loaded over HTTPS without certificate errors, and displayed content such as images that were loaded with certificate errors.
	RanInsecureContentStyle        State `json:"ranInsecureContentStyle"`        // Security state representing a page that ran insecure content.
	DisplayedInsecureContentStyle  State `json:"displayedInsecureContentStyle"`  // Security state representing a page that displayed insecure content.
}

// CertificateErrorAction The action to take when a certificate error occurs. continue will continue processing the request and cancel will cancel the request.
type CertificateErrorAction int

// CertificateErrorAction as enums.
const (
	CertificateErrorActionNotSet CertificateErrorAction = iota
	CertificateErrorActionContinue
	CertificateErrorActionCancel
)

// Valid returns true if enum is set.
func (e CertificateErrorAction) Valid() bool {
	return e >= 1 && e <= 2
}

func (e CertificateErrorAction) String() string {
	switch e {
	case 0:
		return "CertificateErrorActionNotSet"
	case 1:
		return "continue"
	case 2:
		return "cancel"
	}
	return fmt.Sprintf("CertificateErrorAction(%d)", e)
}

// MarshalJSON encodes enum into a string or null when not set.
func (e CertificateErrorAction) MarshalJSON() ([]byte, error) {
	if e == 0 {
		return []byte("null"), nil
	}
	if !e.Valid() {
		return nil, errors.New("security.CertificateErrorAction: MarshalJSON on bad enum value: " + e.String())
	}
	return json.Marshal(e.String())
}

// UnmarshalJSON decodes a string value into a enum.
func (e *CertificateErrorAction) UnmarshalJSON(data []byte) error {
	switch string(data) {
	case "null":
		*e = 0
	case "\"continue\"":
		*e = 1
	case "\"cancel\"":
		*e = 2
	default:
		return fmt.Errorf("security.CertificateErrorAction: UnmarshalJSON on bad input: %s", data)
	}
	return nil
}
