//go:build soc.rp2040 || soc.rp2350

package rp2

// WatchdogConfig holds configuration for the watchdog timer.
type WatchdogConfig struct {
	// The timeout (in milliseconds) before the watchdog fires.
	//
	// If the requested timeout exceeds `MaxTimeout` it will be rounded
	// down.
	TimeoutMillis uint32
}

// watchdog must be implemented by any platform supporting watchdog functionality
type watchdog interface {
	// Configure the watchdog.
	//
	// This method should not be called after the watchdog is started and on
	// some platforms attempting to reconfigure after starting the watchdog
	// is explicitly forbidden / will not work.
	Configure(config WatchdogConfig) error

	// Starts the watchdog.
	Start() error

	// Update the watchdog, indicating that the app is healthy.
	Update()
}

// Ensure required public symbols var exists and meets interface spec
var _ = watchdog(Watchdog)

// Ensure required public constants exist
const _ = WatchdogMaxTimeout
