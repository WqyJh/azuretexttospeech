package main

import (
	"context"
	"fmt"
	"os"

	tts "github.com/WqyJh/azuretexttospeech"
)

func exit(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %+v\n", err)
		os.Exit(1)
	}
	os.Exit(0)
}
func main() {
	// create a key for "Cognitive Services" (kind=SpeechServices). Once the key is available
	// in the Azure portal, push it into an environment variable (export AZUREKEY=SYS64738).
	// By default the free tier keys are served out of West US2
	var apiKey, region string
	if apiKey = os.Getenv("AZURE_KEY"); apiKey == "" {
		exit(fmt.Errorf("please set your AZURE_KEY environment variable"))
	}
	if region = os.Getenv("AZURE_REGION"); region == "" {
		exit(fmt.Errorf("please set your AZURE_REGION environment variable"))
	}
	az, err := tts.New(apiKey, tts.Region(region))
	if err != nil {
		exit(fmt.Errorf("failed to create new client, received %v", err))
	}
	defer close(az.TokenRefreshDoneCh)

	// Digitize a text string using the enUS locale, female voice and specify the
	// audio format of a 16Khz, 32kbit mp3 file.
	ctx := context.Background()
	b, err := az.SynthesizeWithContext(
		ctx,
		tts.VoiceParam{
			SpeechText: "64 BASIC BYTES FREE. READY.",
			Voice:      "en-US-AvaMultilingualNeural",
			Locale:     tts.LocaleEnUS,
			Gender:     tts.GenderFemale,
		},
		tts.AudioOutput_audio_16khz_32kbitrate_mono_mp3)

	if err != nil {
		exit(fmt.Errorf("unable to synthesize, received: %v", err))
	}

	// send results to disk.
	err = os.WriteFile("audio.mp3", b, 0644)
	if err != nil {
		exit(fmt.Errorf("unable to write file, received %v", err))
	}
}
