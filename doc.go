package azuretexttospeech

/*
package azuretexttospeech provides a client for Azure's Cognitive Services (speech services) Text To Speech API. Users of the client
can specify the locale (lanaguage), text in which to speak/digitize as well as the gender in which the gender should be rendered.

For Azure pricing see https://azure.microsoft.com/en-us/pricing/details/cognitive-services/speech-services/ . Note
there is a limited use *FREE* tier available.

Documentation for the TTS Cognitive Services API can be found at https://docs.microsoft.com/en-us/azure/cognitive-services/speech-service/rest-apis#text-to-speech-api

An API key is required to access the Azure API.

USAGE

	func main() {
		// create a key for "Cognitive Services" (kind=SpeechServices). Once the key is available
		// in the Azure portal, push it into an environment variable (export AZUREKEY=SYS64738).
		// By default the free tier keys are served out of West US2
		var apiKey string
		if apiKey = os.Getenv("AZUREKEY"); apiKey == "" {
			exit(fmt.Errorf("Please set your AZUREKEY environment variable"))
		}
		az, err := tts.New(apiKey, tts.RegionEastUS)
		if err != nil {
			exit(fmt.Errorf("failed to create new client, received %v", err))
		}
		defer close(az.TokenRefreshDoneCh)

		// Digitize a text string using the enUS locale, female voice and specify the
		// audio format of a 16Khz, 32kbit mp3 file.
		ctx := context.Background()
		b, err := az.SynthesizeWithContext(
			ctx,
			"64 BASIC BYTES FREE. READY.",
			tts.LocaleEnUS,
			tts.GenderFemale,
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
*/
