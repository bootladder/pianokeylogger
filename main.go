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
    //listen only key stroke event, EV_KEY, type 1
    //also, key press has value 1
    if i.Type == keylogger.EV_KEY && i.Value == 1 {
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
  wavpath := pianoMap[s]
	cmd := exec.Command("aplay", wavpath)
	cmd.Start()
  return nil
}

var pianoMap = map[string]string{
  "A":   "/opt/util/sound/a1.wav",
  "B":   "/opt/util/sound/b1.wav",
  "C":   "/opt/util/sound/c1.wav",
  "D":   "/opt/util/sound/d1.wav",
  "E":   "/opt/util/sound/e1.wav",
  "F":   "/opt/util/sound/f1.wav",
  "G":   "/opt/util/sound/g1.wav",
  "Q":   "/opt/util/sound/a1s.wav",
  "V":   "/opt/util/sound/c1s.wav",
  "S":   "/opt/util/sound/d1s.wav",
  "R":   "/opt/util/sound/f1s.wav",
  "T":   "/opt/util/sound/g1s.wav",
}
