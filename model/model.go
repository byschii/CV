package model

type CV struct {
	Title       string         `json:"title"`
	TextSuggest string         `json:"text_suggestion"`
	Description string         `json:"description"`
	WE          WorkExperience `json:"we"`
	SP          []SideProject  `json:"sp"`
	ED          Education      `json:"ed"`
	ME          AboutMe        `json:"me"`
	SK          Skills         `json:"skills"`
}

type WorkExperience struct {
	Head    string                  `json:"head"`
	History []WorkExperienceHistory `json:"history"`
}

type WorkExperienceHistory struct { // work
	FromPeriod  string `json:"from_period"`
	ToPeriod    string `json:"to_period"`
	Company     string `json:"company"`
	Description string `json:"description"`
}

type SideProject struct {
	Name        string `json:"name"`
	Link        string `json:"link"`
	Description string `json:"description"`
	Hide        bool   `json:"hide"`
}

type Education struct {
	Corso string `json:"corso"`
	Tesi  string `json:"tesi"`
}

type AboutMe struct {
	Head string `json:"head"`
}

type Skills struct {
	Head     string         `json:"head"`
	Sections []SkillSection `json:"sections"`
}

type SkillSection struct {
	SectName string         `json:"section_name"`
	SectVals []SectionValue `json:"section_values"`
}

type SectionValue struct {
	Name  string `json:"name"`
	Level string `json:"level"`
}
