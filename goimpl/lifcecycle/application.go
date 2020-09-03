package lifecycle

// Application is an application life cycle state
type Application byte

const (
	// ApplicationInstalled is the INSTALLED application lifecycle state
	ApplicationInstalled Application = 0x03
	// ApplicationSelectable is the SELECTABLE application lifecycle state
	ApplicationSelectable Application = 0x07
	// ApplicationLocked is the LOCKED application lifecycle state
	ApplicationLocked Application = 0x83
)

// IsValidCustomApplicationState indicates whether a give byte represents a valid application-specific application life cycle state
func IsValidCustomApplicationState(in Application) bool {
	// Must be SELECTABLE + any b4-b7 set + 0 in b8
	// This does mean that SELECTABLE by itself counts as a valid "custom" state, which is consistent with the spec
	return in&ApplicationSelectable == ApplicationSelectable && in&b8 == 0
}
