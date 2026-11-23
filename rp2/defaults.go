//go:build soc.rp2040 || soc.rp2350

package rp2

import "runtime"

const xoscFreq = 12

func ticks() uint64 { return Ticks() }

func gosched() { runtime.Gosched() }

// Default peripheral pin assignments (boards may override before configuring).
var (
	UART_TX_PIN = GPIO0
	UART_RX_PIN = GPIO1

	I2C0_SDA_PIN = GPIO4
	I2C0_SCL_PIN = GPIO5
	I2C1_SDA_PIN = GPIO2
	I2C1_SCL_PIN = GPIO3

	SPI0_SCK_PIN = GPIO18
	SPI0_SDO_PIN = GPIO19
	SPI0_SDI_PIN = GPIO16
	SPI1_SCK_PIN = GPIO10
	SPI1_SDO_PIN = GPIO11
	SPI1_SDI_PIN = GPIO12
)
