package main

import (
  "fmt"
  "time"
  "github.com/jetpacktuxedo/portmidi"
)

func print_info() {
  device_count := portmidi.CountDevices()
  default_out := portmidi.GetDefaultOutputDeviceId()
  default_in := portmidi.GetDefaultInputDeviceId()
  fmt.Println("Total Deivces:", device_count)
  fmt.Println("Default Output Device is:", default_out)
  fmt.Println("Default Input Device is: ", default_in)

  for i := 0; i < device_count; i++ {
    device_info := portmidi.GetDeviceInfo(portmidi.DeviceId(i))
    fmt.Println("Device", i)
    fmt.Println("\tInterface:\t", device_info.Interface)
    fmt.Println("\tName:\t", device_info.Name)
    fmt.Println("\tInput:\t", device_info.IsInputAvailable)
    fmt.Println("\tOutput:\t", device_info.IsOutputAvailable)
    fmt.Println("\tOpen:\t", device_info.IsOpened)
  }
}

func test_audio() {
  default_out := portmidi.GetDefaultOutputDeviceId()
  out, err := portmidi.NewOutputStream(default_out, 1024, 0)
  if err != nil {
    panic(err)
  }

  out.WriteShort(0x90, 60, 100)
  out.WriteShort(0x90, 64, 100)
  out.WriteShort(0x90, 67, 100)

  time.Sleep(2 * time.Second)

  out.WriteShort(0x80, 60, 100)
  out.WriteShort(0x80, 64, 100)
  out.WriteShort(0x80, 67, 100)

  out.Close()
}

func main() {
  portmidi.Initialize()
  print_info()
  test_audio()
  portmidi.Terminate()
}
