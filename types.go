package main

import "encoding/xml"

type NML struct {
	XMLName xml.Name `xml:"NML"`
	Text    string   `xml:",chardata"`
	VERSION string   `xml:"VERSION,attr"`
	HEAD    struct {
		Text    string `xml:",chardata"`
		COMPANY string `xml:"COMPANY,attr"`
		PROGRAM string `xml:"PROGRAM,attr"`
	} `xml:"HEAD"`
	COLLECTION struct {
		Text    string `xml:",chardata"`
		ENTRIES string `xml:"ENTRIES,attr"`
		ENTRY   []struct {
			Text                 string `xml:",chardata"`
			MODIFIEDDATE         string `xml:"MODIFIED_DATE,attr"`
			MODIFIEDTIME         string `xml:"MODIFIED_TIME,attr"`
			LOCK                 string `xml:"LOCK,attr"`
			LOCKMODIFICATIONTIME string `xml:"LOCK_MODIFICATION_TIME,attr"`
			AUDIOID              string `xml:"AUDIO_ID,attr"`
			TITLE                string `xml:"TITLE,attr"`
			ARTIST               string `xml:"ARTIST,attr"`
			LOCATION             struct {
				Text     string `xml:",chardata"`
				DIR      string `xml:"DIR,attr"`
				FILE     string `xml:"FILE,attr"`
				VOLUME   string `xml:"VOLUME,attr"`
				VOLUMEID string `xml:"VOLUMEID,attr"`
			} `xml:"LOCATION"`
			ALBUM struct {
				Text  string `xml:",chardata"`
				TITLE string `xml:"TITLE,attr"`
			} `xml:"ALBUM"`
			MODIFICATIONINFO struct {
				Text       string `xml:",chardata"`
				AUTHORTYPE string `xml:"AUTHOR_TYPE,attr"`
			} `xml:"MODIFICATION_INFO"`
			INFO struct {
				Text          string `xml:",chardata"`
				BITRATE       string `xml:"BITRATE,attr"`
				GENRE         string `xml:"GENRE,attr"`
				COMMENT       string `xml:"COMMENT,attr"`
				COVERARTID    string `xml:"COVERARTID,attr"`
				KEY           string `xml:"KEY,attr"`
				PLAYCOUNT     string `xml:"PLAYCOUNT,attr"`
				PLAYTIME      string `xml:"PLAYTIME,attr"`
				PLAYTIMEFLOAT string `xml:"PLAYTIME_FLOAT,attr"`
				RANKING       string `xml:"RANKING,attr"`
				IMPORTDATE    string `xml:"IMPORT_DATE,attr"`
				LASTPLAYED    string `xml:"LAST_PLAYED,attr"`
				FLAGS         string `xml:"FLAGS,attr"`
				FILESIZE      string `xml:"FILESIZE,attr"`
				COLOR         string `xml:"COLOR,attr"`
				RELEASEDATE   string `xml:"RELEASE_DATE,attr"`
			} `xml:"INFO"`
			TEMPO struct {
				Text       string `xml:",chardata"`
				BPM        string `xml:"BPM,attr"`
				BPMQUALITY string `xml:"BPM_QUALITY,attr"`
			} `xml:"TEMPO"`
			LOUDNESS struct {
				Text        string `xml:",chardata"`
				PEAKDB      string `xml:"PEAK_DB,attr"`
				PERCEIVEDDB string `xml:"PERCEIVED_DB,attr"`
				ANALYZEDDB  string `xml:"ANALYZED_DB,attr"`
			} `xml:"LOUDNESS"`
			MUSICALKEY struct {
				Text  string `xml:",chardata"`
				VALUE string `xml:"VALUE,attr"`
			} `xml:"MUSICAL_KEY"`
			CUEV2 []struct {
				Text       string `xml:",chardata"`
				NAME       string `xml:"NAME,attr"`
				DISPLORDER string `xml:"DISPL_ORDER,attr"`
				TYPE       string `xml:"TYPE,attr"`
				START      string `xml:"START,attr"`
				LEN        string `xml:"LEN,attr"`
				REPEATS    string `xml:"REPEATS,attr"`
				HOTCUE     string `xml:"HOTCUE,attr"`
				COLOR      string `xml:"COLOR,attr"`
				GRID       struct {
					Text string `xml:",chardata"`
					BPM  string `xml:"BPM,attr"`
				} `xml:"GRID"`
			} `xml:"CUE_V2"`
		} `xml:"ENTRY"`
	} `xml:"COLLECTION"`
	SETS struct {
		Text    string `xml:",chardata"`
		ENTRIES string `xml:"ENTRIES,attr"`
	} `xml:"SETS"`
	PLAYLISTS struct {
		Text string `xml:",chardata"`
		NODE struct {
			Text     string `xml:",chardata"`
			TYPE     string `xml:"TYPE,attr"`
			NAME     string `xml:"NAME,attr"`
			SUBNODES struct {
				Text  string `xml:",chardata"`
				COUNT string `xml:"COUNT,attr"`
				NODE  struct {
					Text     string `xml:",chardata"`
					TYPE     string `xml:"TYPE,attr"`
					NAME     string `xml:"NAME,attr"`
					PLAYLIST struct {
						Text    string `xml:",chardata"`
						ENTRIES string `xml:"ENTRIES,attr"`
						TYPE    string `xml:"TYPE,attr"`
						UUID    string `xml:"UUID,attr"`
						ENTRY   []struct {
							Text       string `xml:",chardata"`
							PRIMARYKEY struct {
								Text string `xml:",chardata"`
								TYPE string `xml:"TYPE,attr"`
								KEY  string `xml:"KEY,attr"`
							} `xml:"PRIMARYKEY"`
						} `xml:"ENTRY"`
					} `xml:"PLAYLIST"`
				} `xml:"NODE"`
			} `xml:"SUBNODES"`
		} `xml:"NODE"`
	} `xml:"PLAYLISTS"`
	INDEXING string `xml:"INDEXING"`
}
