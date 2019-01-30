package graphql_endpoint

// these are elastic json models
type PersonKeyword struct {
	Uri   string `json:"uri"`
	Label string `json:"label"`
}

type PersonImage struct {
	Main      string `json:"main"`
	Thumbnail string `json:"thumbnail"`
}

type PersonName struct {
	FirstName  string  `json:"firstName"`
	LastName   string  `json:"lastName"`
	MiddleName *string `json:"middleName"`
}

type PersonType struct {
	Code  string `json:"code"`
	Label string `json:"label"`
}

type OverviewType struct {
	Code  string `json:"code"`
	Label string `json:"label"`
}

type PersonOverview struct {
	Label string       `json:"overview"`
	Type  OverviewType `json:"type"`
}

type Extension struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Person struct {
	Id             string           `json:"id"`
	Uri            string           `json:"uri"`
	SourceId       string           `json:"sourceId"`
	PrimaryTitle   string           `json:"primaryTitle"`
	Name           PersonName       `json:"name" elastic:"type:object"`
	Image          PersonImage      `json:"image" elastic:"type:object"`
	Type           PersonType       `json:"type" elastic:"type:object"`
	OverviewList   []PersonOverview `json:"overviewList" elastic:"type:nested"`
	KeywordList    []PersonKeyword  `json:"keywordList" elastic:"type:nested"`
	Extensions     []Extension      `json:"extensions" elastic:"type:nested"`
	EducationList  []Education      `json:"educationList" elastic:"type:nested"`
	AffliationList []Affiliation    `json:"affiliationList" elastic:"type:nested"`
}

type DateResolution struct {
	DateTime   string `json:"dateTime"`
	Resolution string `json:"resolution"`
}

type Organization struct {
	Id    string `json:"id"`
	Uri   string `json:"uri"`
	Label string `json:"label"`
}

type Institution struct {
	Id    string `json:"id"`
	Uri   string `json:"uri"`
	Label string `json:"label"`
}

type Affiliation struct {
	Id           string         `json:"id"`
	Uri          string         `json:"uri"`
	PersonId     string         `json:"personId"`
	Label        string         `json:"label"`
	StartDate    DateResolution `json:"startDate"`
	Organization Organization   `json:"organization"`
}

type Education struct {
	Id          string      `json:"id"`
	Uri         string      `json:"Uri"`
	Label       string      `json:"label"`
	PersonId    string      `json:"personId"`
	Institution Institution `json:"org" elastic:"type:object"`
}

type FundingRole struct {
	Id       string `json:"id"`
	Uri      string `json:"uri"`
	GrantId  string `json:"grantId"`
	PersonId string `json:"personId"`
	Label    string `json:"label"`
}

type Grant struct {
	Id        string         `json:"id"`
	Uri       string         `json:"uri"`
	Label     string         `json:"label"`
	StartDate DateResolution `json:"startDate"`
	EndDate   DateResolution `json:"endDate"`
}

type Authorship struct {
	Id            string `json:"id"`
	Uri           string `json:"uri"`
	PublicationId string `json:"publicationId"`
	PersonId      string `json:"personId"`
	Label         string `json:"label"`
}

type PublicationVenue struct {
	Uri   string `json:"uri"`
	Label string `json:"label"`
}

type Publication struct {
	Id    string `json:"id"`
	Uri   string `json:"uri"`
	Label string `json:"label"`
	// NOTE: this is supposed to be an array
	AuthorList string           `json:"authorList"`
	Doi        string           `json:"doi"`
	Venue      PublicationVenue `json:"venue"`
}

// these are graphql only json models
type PageInfo struct {
	PerPage     int `json:"perPage"`
	CurrentPage int `json:"page"`
	TotalPages  int `json":totalPages"`
	Count       int `json:"count"`
}

type PersonList struct {
	Results  []Person `json:"data"`
	PageInfo PageInfo `json:"pageInfo"`
}

type PublicationList struct {
	Results  []Publication `json:"data"`
	PageInfo PageInfo      `json:"pageInfo"`
}

type GrantList struct {
	Results  []Grant  `json:"data"`
	PageInfo PageInfo `json:"pageInfo"`
}

