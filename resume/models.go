package resume

// Resume is the root struct
type Resume struct {
	Basics       Basics        `json:"basics"      hcl:"basics,block"`
	Work         []Work        `json:"work"        hcl:"work,block"`
	Education    []Education   `json:"education"   hcl:"education,block"`
	Skills       []Skill       `json:"skills"      hcl:"skill,block"`
	Projects     []Project     `json:"projects,omitempty"    hcl:"project,block"`
	Certificates []Certificate `json:"certificates,omitempty" hcl:"certificate,block"`
	// Add more as needed...
}

type Basics struct {
	Name     string   `json:"name"      hcl:"name"`
	Label    string   `json:"label,omitempty" hcl:"label"`
	Email    string   `json:"email,omitempty" hcl:"email"`
	Phone    string   `json:"phone,omitempty" hcl:"phone"`
	URL      string   `json:"url,omitempty"   hcl:"url,optional"`
	Summary  string   `json:"summary,omitempty" hcl:"summary"`
	Location Location `json:"location,omitempty" hcl:"location,block"`
	Image    string   `json:"image,omitempty" hcl:"image,optional"`
}

type Location struct {
	Address     string `json:"address,omitempty" hcl:"address"`
	PostalCode  string `json:"postalCode,omitempty" hcl:"postal_code"`
	City        string `json:"city,omitempty" hcl:"city"`
	CountryCode string `json:"countryCode,omitempty" hcl:"country_code"`
	Region      string `json:"region,omitempty" hcl:"region"`
}

type Work struct {
	Name       string   `json:"name" hcl:",label"` // Block label (unique name/id)
	Company    string   `json:"company" hcl:"company"`
	Position   string   `json:"position" hcl:"position"`
	Website    string   `json:"website,omitempty" hcl:"website,optional"`
	StartDate  string   `json:"startDate" hcl:"start_date"`
	EndDate    string   `json:"endDate,omitempty" hcl:"end_date"`
	Summary    string   `json:"summary,omitempty" hcl:"summary,optional"`
	Highlights []string `json:"highlights,omitempty" hcl:"highlights,optional"`
	Location   string   `json:"location,omitempty" hcl:"location"`
}

type Education struct {
	Name        string   `json:"name" hcl:",label"`
	Institution string   `json:"institution" hcl:"institution"`
	Area        string   `json:"area,omitempty" hcl:"area"`
	StudyType   string   `json:"studyType,omitempty" hcl:"study_type"`
	StartDate   string   `json:"startDate,omitempty" hcl:"start_date,optional"`
	EndDate     string   `json:"endDate,omitempty" hcl:"end_date,optional"`
	Score       string   `json:"score,omitempty" hcl:"score,optional"`
	Courses     []string `json:"courses,omitempty" hcl:"courses,optional"`
}

type Skill struct {
	Name     string   `json:"name" hcl:"name"`
	Level    string   `json:"level,omitempty" hcl:"level,optional"`
	Keywords []string `json:"keywords,omitempty" hcl:"keywords"`
}

type Project struct {
	Name        string   `json:"name" hcl:",label"`
	Description string   `json:"description,omitempty" hcl:"description"`
	Highlights  []string `json:"highlights,omitempty" hcl:"highlights"`
	Keywords    []string `json:"keywords,omitempty" hcl:"keywords"`
	StartDate   string   `json:"startDate,omitempty" hcl:"start_date"`
	EndDate     string   `json:"endDate,omitempty" hcl:"end_date"`
	Url         string   `json:"url,omitempty" hcl:"url"`
	Roles       []string `json:"roles,omitempty" hcl:"roles"`
	Entity      string   `json:"entity,omitempty" hcl:"entity"`
	Type        string   `json:"type,omitempty" hcl:"type"`
}

type Certificate struct {
	Name   string `json:"name" hcl:"name"`
	Date   string `json:"date,omitempty" hcl:"date,optional"`
	Issuer string `json:"issuer,omitempty" hcl:"issuer,optional"`
	Url    string `json:"url,omitempty" hcl:"url,optional"`
}
