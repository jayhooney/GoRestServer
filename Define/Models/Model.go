package models

type Topic struct {
	DOCID int
	TERMS string
}

type Emotion struct {
	Seq     int
	CmtID   string
	Emotion string
}
