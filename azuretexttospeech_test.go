package azuretexttospeech

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVoiceXML(t *testing.T) {
	expect := "<speak version='1.0' xml:lang='en-US'><voice xml:lang='en-US' xml:gender='Female' name='ar-EG-Hoda'>Microsoft Speech Service Text-to-Speech API</voice></speak>"
	xml, err := voiceXMLRender(VoiceParam{
		SpeechText: "Microsoft Speech Service Text-to-Speech API",
		Voice:      "ar-EG-Hoda",
		Locale:     LocaleEnUS,
		Gender:     GenderFemale,
	})
	assert.NoError(t, err)
	assert.Equal(t, expect, xml)
}

func TestSynthesize(t *testing.T) {
	az := &AzureCSTextToSpeech{SubscriptionKey: "SYS64738", accessToken: "SYS49152"}

	// payload should be nil and err should be true, since DeCH + Female is not a valid combination
	payload, err := az.Synthesize(VoiceParam{
		SpeechText: "test-speech",
		Voice:      "SYS2064",
		Locale:     LocaleDeCH,
		Gender:     GenderFemale,
	}, AudioOutput_riff_8khz_8bit_mono_alaw)
	assert.Error(t, err, "should raise an error")
	assert.Nil(t, payload, "payload should be nil")

	ts := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("SYS4096"))
		}),
	)
	defer ts.Close()

	az.textToSpeechURL = ts.URL
	// request should now be successful with a valid locale and gender.
	payload, err = az.Synthesize(VoiceParam{
		SpeechText: "test-speech",
		Voice:      "SYS4096",
		Locale:     LocaleDeCH,
		Gender:     GenderMale,
	}, AudioOutput_riff_8khz_8bit_mono_alaw)
	assert.NoError(t, err)
	assert.Equal(t, payload, []byte("SYS4096"))
}

// TestRefreshToken validates logic for fetching of the refreshToken
func TestRefreshToken(t *testing.T) {
	az := &AzureCSTextToSpeech{SubscriptionKey: "ThisIsMySubscriptionKeyAndToBeToken"}

	ts := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// return the SubscriptionKey as the token, for test case only.
			w.Write([]byte(az.SubscriptionKey))
		}),
	)
	defer ts.Close()
	az.tokenRefreshURL = ts.URL
	err := az.refreshToken()

	assert.NoError(t, err, "should not return an error")
	assert.Equal(t, az.SubscriptionKey, az.accessToken, "values should be equal")
}
