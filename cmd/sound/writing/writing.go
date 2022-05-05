//см еще https://github.com/faiface/beep/blob/master/examples/tone-player/main.go

// программа создает звуковой файл
package main

import (
	"fmt"
	"math"
	"os"

	"github.com/zenwerk/go-wave"
)

func main() {
	file, err := os.Create("./cmd/sound/writing/write_test.wav")
	defer file.Close()
	if err != nil {
		fmt.Println("error:", err.Error())
		return
	}
	param := wave.WriterParam{
		Out:           file,
		Channel:       1,
		SampleRate:    44100,
		BitsPerSample: 16,
	}

	w, err := wave.NewWriter(param)
	defer w.Close()
	if err != nil {
		fmt.Println("error:", err.Error())
		return
	}

	amplitude := 0.1                // громкость 20%
	hz := 440.0                     // частота звука
	length := param.SampleRate * 10 // длительность исполнения 10 сек

	for i := 0; i < length; i++ {

		if i%44100 == 0 {
			amplitude += 0.05
			hz += 50.0
		}

		d := amplitude * math.Sin(2.0*math.Pi*hz*float64(i)/float64(param.SampleRate))

		d = (d + 1.0) / 2.0 * 65536.0

		if d > 65535.0 {
			d = 65535.0
		} else if d < 0.0 {
			d = 0.0
		}

		data := int16(d+0.5) - int16(32767)

		_, err = w.WriteSample16([]int16{data})
		if err != nil {
			fmt.Println("error:", err.Error())
			return
		}
	}
}
