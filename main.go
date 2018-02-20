package main
import (
 "fmt"
 "github.com/MarinX/keylogger"
 "os/exec"
)

func main() {
  devs, err := keylogger.NewDevices()
  if err != nil {
    fmt.Println(err)
    return
  }
  for _, val := range devs {
    fmt.Println("Id->", val.Id, "Device->", val.Name)
  }
  //our keyboard..on your system, it will be diffrent
  rd := keylogger.NewKeyLogger(devs[3])
  in, err := rd.Read()
  if err != nil {
    fmt.Println(err)
    return
  }
  for i := range in {
    //listen only key stroke event
    if i.Type == keylogger.EV_KEY {
      fmt.Printf("\n\n\n")
      fmt.Printf("i is: %d\n",i)
      fmt.Printf("i keycode is : %d\n",i.Code)
      fmt.Printf("i keystring is : %s\n", i.KeyString())
      Play(i.KeyString())
    }
  }
}

func Play(s string) error {

	//mystrings := strings.Fields(out)
	cmd := exec.Command("aplay", "/opt/util/sound/a1.wav")
	cmd.Start()
  return nil
}

pianoMap = map[string]string{
  "A":   "/opt/util/sound/a1.wav",
  "B":   "/opt/util/sound/b1.wav",
  "C":   "/opt/util/sound/c1.wav",
  "D":   "/opt/util/sound/d1.wav",
  "E":   "/opt/util/sound/e1.wav",
  "F":   "/opt/util/sound/f1.wav",
  "G":   "/opt/util/sound/g1.wav",
  "Q":   "/opt/util/sound/a1.wav",
  "V":   "/opt/util/sound/a1.wav",
  "S":   "/opt/util/sound/a1.wav",
  "W":   "/opt/util/sound/a1.wav",
  "R":   "/opt/util/sound/a1.wav",
  "T":   "/opt/util/sound/a1.wav",
  "X":   "/opt/util/sound/a1.wav",
}
