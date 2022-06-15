// https://socketloop.com/tutorials/golang-record-voice-audio-from-microphone-to-wav-file

package main

import "C"

import (
	"fmt"
	"github.com/gordonklaus/portaudio"
	wave "github.com/zenwerk/go-wave"
	"math/rand"
	"os"
	"time"
)

func errCheck(err error) {
	if err != nil {
		fmt.Println("error:", err.Error())
	}
}

func main() {
	//if len(os.Args) != 2 {
	//	fmt.Printf("Usage : %s <audiofilename.wav>\n", os.Args[0])
	//	os.Exit(0)
	//}
	//
	//audioFileName := os.Args[1]
	//
	fmt.Println("Recording. Press ESC to quit.")

	//if !strings.HasSuffix(audioFileName, ".wav") {
	//	audioFileName += ".wav"
	//}

	audioFileName := "audioFileName.wav"
	waveFile, err := os.Create(audioFileName)
	errCheck(err)

	// www.people.csail.mit.edu/hubert/pyaudio/  - under the Record tab
	inputChannels := 1
	outputChannels := 0
	sampleRate := 44100
	framesPerBuffer := make([]byte, 64)

	// init PortAudio

	portaudio.Initialize()
	defer portaudio.Terminate()

	stream, err := portaudio.OpenDefaultStream(inputChannels, outputChannels, float64(sampleRate), len(framesPerBuffer), framesPerBuffer)
	errCheck(err)
	defer stream.Close()

	// setup Wave file writer

	param := wave.WriterParam{
		Out:           waveFile,
		Channel:       inputChannels,
		SampleRate:    sampleRate,
		BitsPerSample: 8, // if 16, change to WriteSample16()
	}

	waveWriter, err := wave.NewWriter(param)
	errCheck(err)

	//defer waveWriter.Close()

	go func() {
		key := C.getch()
		fmt.Println()
		fmt.Println("Cleaning up ...")
		if key == 27 {
			// better to control
			// how we close then relying on defer
			waveWriter.Close()
			stream.Close()
			portaudio.Terminate()
			fmt.Println("Play", audioFileName, "with a audio player to hear the result.")
			os.Exit(0)
		}
	}()

	// recording in progress ticker. From good old DOS days.
	ticker := []string{
		"-",
		"\\",
		"/",
		"|",
	}
	rand.Seed(time.Now().UnixNano())

	// start reading from microphone
	errCheck(stream.Start())
	defer errCheck(stream.Stop())

	for {
		errCheck(stream.Read())

		fmt.Printf("\rRecording is live now. Say something to your microphone! [%v]", ticker[rand.Intn(len(ticker)-1)])

		// write to wave file
		_, err := waveWriter.Write(framesPerBuffer) // WriteSample16 for 16 bits
		errCheck(err)
	}
}
