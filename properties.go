package azuretexttospeech

// AudioOutput types represent the supported audio encoding formats for the text-to-speech endpoint.
// This type is required when requesting to azuretexttospeech.Synthesize text-to-speed request.
// Each incorporates a bitrate and encoding type. The Speech service supports 24 kHz, 16 kHz, and 8 kHz audio outputs.
// See: https://docs.microsoft.com/en-us/azure/cognitive-services/speech-service/rest-text-to-speech#audio-outputs
type AudioOutput string

// Streaming audio output types
const (
	AudioOutput_amr_wb_16000hz                     AudioOutput = "amr-wb-16000hz"
	AudioOutput_audio_16khz_16bit_32kbps_mono_opus AudioOutput = "audio-16khz-16bit-32kbps-mono-opus"
	AudioOutput_audio_16khz_32kbitrate_mono_mp3    AudioOutput = "audio-16khz-32kbitrate-mono-mp3"
	AudioOutput_audio_16khz_64kbitrate_mono_mp3    AudioOutput = "audio-16khz-64kbitrate-mono-mp3"
	AudioOutput_audio_16khz_128kbitrate_mono_mp3   AudioOutput = "audio-16khz-128kbitrate-mono-mp3"
	AudioOutput_audio_24khz_16bit_24kbps_mono_opus AudioOutput = "audio-24khz-16bit-24kbps-mono-opus"
	AudioOutput_audio_24khz_16bit_48kbps_mono_opus AudioOutput = "audio-24khz-16bit-48kbps-mono-opus"
	AudioOutput_audio_24khz_48kbitrate_mono_mp3    AudioOutput = "audio-24khz-48kbitrate-mono-mp3"
	AudioOutput_audio_24khz_96kbitrate_mono_mp3    AudioOutput = "audio-24khz-96kbitrate-mono-mp3"
	AudioOutput_audio_24khz_160kbitrate_mono_mp3   AudioOutput = "audio-24khz-160kbitrate-mono-mp3"
	AudioOutput_audio_48khz_96kbitrate_mono_mp3    AudioOutput = "audio-48khz-96kbitrate-mono-mp3"
	AudioOutput_audio_48khz_192kbitrate_mono_mp3   AudioOutput = "audio-48khz-192kbitrate-mono-mp3"
	AudioOutput_ogg_16khz_16bit_mono_opus          AudioOutput = "ogg-16khz-16bit-mono-opus"
	AudioOutput_ogg_24khz_16bit_mono_opus          AudioOutput = "ogg-24khz-16bit-mono-opus"
	AudioOutput_ogg_48khz_16bit_mono_opus          AudioOutput = "ogg-48khz-16bit-mono-opus"
	AudioOutput_raw_8khz_8bit_mono_alaw            AudioOutput = "raw-8khz-8bit-mono-alaw"
	AudioOutput_raw_8khz_8bit_mono_mulaw           AudioOutput = "raw-8khz-8bit-mono-mulaw"
	AudioOutput_raw_8khz_16bit_mono_pcm            AudioOutput = "raw-8khz-16bit-mono-pcm"
	AudioOutput_raw_16khz_16bit_mono_pcm           AudioOutput = "raw-16khz-16bit-mono-pcm"
	AudioOutput_raw_16khz_16bit_mono_truesilk      AudioOutput = "raw-16khz-16bit-mono-truesilk"
	AudioOutput_raw_22050hz_16bit_mono_pcm         AudioOutput = "raw-22050hz-16bit-mono-pcm"
	AudioOutput_raw_24khz_16bit_mono_pcm           AudioOutput = "raw-24khz-16bit-mono-pcm"
	AudioOutput_raw_24khz_16bit_mono_truesilk      AudioOutput = "raw-24khz-16bit-mono-truesilk"
	AudioOutput_raw_44100hz_16bit_mono_pcm         AudioOutput = "raw-44100hz-16bit-mono-pcm"
	AudioOutput_raw_48khz_16bit_mono_pcm           AudioOutput = "raw-48khz-16bit-mono-pcm"
	AudioOutput_webm_16khz_16bit_mono_opus         AudioOutput = "webm-16khz-16bit-mono-opus"
	AudioOutput_webm_24khz_16bit_24kbps_mono_opus  AudioOutput = "webm-24khz-16bit-24kbps-mono-opus"
	AudioOutput_webm_24khz_16bit_mono_opus         AudioOutput = "webm-24khz-16bit-mono-opus"
)

// NonStreaming audio output types
const (
	AudioOutput_riff_8khz_8bit_mono_alaw    AudioOutput = "riff-8khz-8bit-mono-alaw"
	AudioOutput_riff_8khz_8bit_mono_mulaw   AudioOutput = "riff-8khz-8bit-mono-mulaw"
	AudioOutput_riff_8khz_16bit_mono_pcm    AudioOutput = "riff-8khz-16bit-mono-pcm"
	AudioOutput_riff_22050hz_16bit_mono_pcm AudioOutput = "riff-22050hz-16bit-mono-pcm"
	AudioOutput_riff_24khz_16bit_mono_pcm   AudioOutput = "riff-24khz-16bit-mono-pcm"
	AudioOutput_riff_44100hz_16bit_mono_pcm AudioOutput = "riff-44100hz-16bit-mono-pcm"
	AudioOutput_riff_48khz_16bit_mono_pcm   AudioOutput = "riff-48khz-16bit-mono-pcm"
)

// Gender type for the digitized language
//
//go:generate enumer -type=Gender -linecomment -json
type Gender int

const (
	// GenderMale , GenderFemale are the static Gender constants for digitized voices.
	// See Gender in https://docs.microsoft.com/en-us/azure/cognitive-services/speech-service/language-support#standard-voices for breakdown
	GenderMale   Gender = iota // Male
	GenderFemale               // Female
)

// Locale references the language or locale for text-to-speech.
// See "locale" in https://docs.microsoft.com/en-us/azure/cognitive-services/speech-service/language-support#standard-voices
type Locale string

const (
	LocaleArEG Locale = "ar-EG"
	LocaleArSA Locale = "ar-SA"
	LocaleBgBG Locale = "bg-BG"
	LocaleCaES Locale = "ca-ES"
	LocaleCsCZ Locale = "cs-CZ"
	LocaleDaDK Locale = "da-DK"
	LocaleDeAT Locale = "de-AT"
	LocaleDeCH Locale = "de-CH"
	LocaleDeDE Locale = "de-DE"
	LocaleElGR Locale = "el-GR"
	LocaleEnAU Locale = "en-AU"
	LocaleEnCA Locale = "en-CA"
	LocaleEnGB Locale = "en-GB"
	LocaleEnIE Locale = "en-IE"
	LocaleEnIN Locale = "en-IN"
	LocaleEnUS Locale = "en-US"
	LocaleEsES Locale = "es-ES"
	LocaleEsMX Locale = "es-MX"
	LocaleEtEE Locale = "et-EE"
	LocaleFiFI Locale = "fi-FI"
	LocaleFrCA Locale = "fr-CA"
	LocaleFrCH Locale = "fr-CH"
	LocaleFrFR Locale = "fr-FR"
	LocaleGaIE Locale = "ga-IE"
	LocaleHeIL Locale = "he-IL"
	LocaleHiIN Locale = "hi-IN"
	LocaleHrHR Locale = "hr-HR"
	LocaleHuHU Locale = "hu-HU"
	LocaleIdID Locale = "id-ID"
	LocaleItIT Locale = "it-IT"
	LocaleJaJP Locale = "ja-JP"
	LocaleKoKR Locale = "ko-KR"
	LocaleLtLT Locale = "lt-LT"
	LocaleLvLV Locale = "lv-LV"
	LocaleMtMT Locale = "mt-MT"
	LocaleMrIN Locale = "mr-IN"
	LocaleMsMY Locale = "ms-MY"
	LocaleNbNO Locale = "nb-NO"
	LocaleNlNL Locale = "nl-NL"
	LocalePlPL Locale = "pl-PL"
	LocalePtBR Locale = "pt-BR"
	LocalePtPT Locale = "pt-PT"
	LocaleRoRO Locale = "ro-RO"
	LocaleRuRU Locale = "ru-RU"
	LocaleSkSK Locale = "sk-SK"
	LocaleSlSI Locale = "sl-SI"
	LocaleSvSE Locale = "sv-SE"
	LocaleTaIN Locale = "ta-IN"
	LocaleTeIN Locale = "te-IN"
	LocaleThTH Locale = "th-TH"
	LocaleTrTR Locale = "tr-TR"
	LocaleViVN Locale = "vi-VN"
	LocaleZhCN Locale = "zh-CN"
	LocaleZhHK Locale = "zh-HK"
	LocaleZhTW Locale = "zh-TW"
)

// Region references the locations of the availability of standard voices.
// See https://docs.microsoft.com/en-us/azure/cognitive-services/speech-service/regions#standard-voices
type Region string

const (
	// Azure regions and their endpoints that support the Text To Speech service.
	RegionAustraliaEast  Region = "australiaeast"
	RegionBrazilSouth    Region = "brazilsouth"
	RegionCanadaCentral  Region = "canadacentral"
	RegionCentralUS      Region = "centralus"
	RegionEastAsia       Region = "eastasia"
	RegionEastUS         Region = "eastus"
	RegionEastUS2        Region = "eastus2"
	RegionFranceCentral  Region = "francecentral"
	RegionIndiaCentral   Region = "indiacentral"
	RegionJapanEast      Region = "japaneast"
	RegionJapanWest      Region = "japanwest"
	RegionKoreaCentral   Region = "koreacentral"
	RegionNorthCentralUS Region = "northcentralus"
	RegionNorthEurope    Region = "northeurope"
	RegionSouthCentralUS Region = "southcentralus"
	RegionSoutheastAsia  Region = "southeastasia"
	RegionUKSouth        Region = "uksouth"
	RegionWestEurope     Region = "westeurope"
	RegionWestUS         Region = "westus"
	RegionWestUS2        Region = "westus2"
)
