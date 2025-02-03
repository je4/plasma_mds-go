package grodata

type License struct {
	Name    string `json:"name,omitempty"`
	Uri     string `json:"uri,omitempty"`
	IconUri string `json:"iconUri,omitempty"`
}

type Author struct {
	AuthorName             string `json:"authorName,omitempty"`
	AuthorAffiliation      string `json:"authorAffiliation,omitempty"`
	AuthorIdentifierScheme string `json:"authorIdentifierScheme,omitempty"`
	AuthorIdentifier       string `json:"authorIdentifier,omitempty"`
}

type DatasetContact struct {
	DatasetContactName        string `json:"datasetContactName,omitempty"`
	DatasetContactAffiliation string `json:"datasetContactAffiliation,omitempty"`
	DatasetContactEmail       string `json:"datasetContactEmail,omitempty"`
}

type DsDescription struct {
	DsDescriptionValue string     `json:"dsDescriptionValue,omitempty"`
	DsDescriptionDate  CustomDate `json:"dsDescriptionDate,omitempty"`
}

type Keyword struct {
	KeywordValue         string `json:"keywordValue,omitempty"`
	KeywordVocabulary    string `json:"keywordVocabulary,omitempty"`
	KeywordVocabularyURI string `json:"keywordVocabularyURI,omitempty"`
}

type Publication struct {
	PublicationCitation string `json:"publicationCitation,omitempty"`
	PublicationIDType   string `json:"publicationIDType,omitempty"`
	PublicationIDNumber string `json:"publicationIDNumber,omitempty"`
	PublicationURL      string `json:"publicationURL,omitempty"`
}

type GrantNumber struct {
	GrantNumberAgency string `json:"grantNumberAgency,omitempty"`
	GrantNumberValue  string `json:"grantNumberValue,omitempty"`
}

type MetadataBlock struct {
	DisplayName string `json:"displayName,omitempty"`
	Name        string `json:"name,omitempty"`
	Fields      Fields `json:"fields,omitempty"`
}

type DataFile struct {
	ID                int    `json:"id"`
	PersistentId      string `json:"persistentId,omitempty"`
	PidURL            string `json:"pidURL,omitempty"`
	Filename          string `json:"filename,omitempty"`
	ContentType       string `json:"contentType,omitempty"`
	FriendlyType      string `json:"friendlyType,omitempty"`
	Filesize          int    `json:"filesize,omitempty"`
	Description       string `json:"description,omitempty"`
	StorageIdentifier string `json:"storageIdentifier,omitempty"`
	RootDataFileId    int    `json:"rootDataFileId,omitempty"`
	Md5               string `json:"md5,omitempty"`
	Checksum          struct {
		Type  string `json:"type,omitempty"`
		Value string `json:"value,omitempty"`
	} `json:"checksum,omitempty"`
	TabularData       bool       `json:"tabularData,omitempty"`
	CreationDate      CustomDate `json:"creationDate,omitempty"`
	PublicationDate   CustomDate `json:"publicationDate,omitempty"`
	FileAccessRequest bool       `json:"fileAccessRequest,omitempty"`
}

type File struct {
	Description      string   `json:"description,omitempty"`
	Label            string   `json:"label,omitempty"`
	Restricted       bool     `json:"restricted,omitempty"`
	Version          int      `json:"version,omitempty"`
	DatasetVersionId int      `json:"datasetVersionId,omitempty"`
	DataFile         DataFile `json:"dataFile,omitempty"`
}

type DatasetVersion struct {
	ID                  int        `json:"id"`
	DatasetId           int        `json:"datasetId,omitempty"`
	DatasetPersistentId string     `json:"datasetPersistentId,omitempty"`
	StorageIdentifier   string     `json:"storageIdentifier,omitempty"`
	VersionNumber       int        `json:"versionNumber,omitempty"`
	VersionMinorNumber  int        `json:"versionMinorNumber,omitempty"`
	VersionState        string     `json:"versionState,omitempty"`
	UNF                 string     `json:"UNF,omitempty"`
	LastUpdateTime      CustomDate `json:"lastUpdateTime,omitempty"`
	ReleaseTime         CustomDate `json:"releaseTime,omitempty"`
	CreateTime          CustomDate `json:"createTime,omitempty"`
	PublicationDate     CustomDate `json:"publicationDate,omitempty"`
	CitationDate        CustomDate `json:"citationDate,omitempty"`
	License             License    `json:"license,omitempty"`
	FileAccessRequest   bool       `json:"fileAccessRequest,omitempty"`
	MetadataBlocks      struct {
		Citation   MetadataBlock `json:"citation,omitempty"`
		Geospatial MetadataBlock `json:"geospatial,omitempty"`
		Journal    MetadataBlock `json:"journal,omitempty"`
	} `json:"metadataBlocks,omitempty"`
	Files    []File `json:"files,omitempty"`
	Citation string `json:"citation,omitempty"`
}

type Grodata struct {
	ID                int            `json:"id"`
	Identifier        string         `json:"identifier,omitempty"`
	PersistentUrl     string         `json:"persistentUrl,omitempty"`
	Protocol          string         `json:"protocol,omitempty"`
	Authority         string         `json:"authority,omitempty"`
	Publisher         string         `json:"publisher,omitempty"`
	PublicationDate   CustomDate     `json:"publicationDate,omitempty"`
	StorageIdentifier string         `json:"storageIdentifier,omitempty"`
	DatasetVersion    DatasetVersion `json:"datasetVersion,omitempty"`
}

func (data Grodata) GetTitle() string {
	if t, ok := data.DatasetVersion.MetadataBlocks.Citation.Fields.GetField("title"); ok {
		return t.Value.String()
	}
	return ""
}

func (data Grodata) GetDescription() string {
	var str string
	if dsDesc, ok := data.DatasetVersion.MetadataBlocks.Citation.Fields.GetField("dsDescription"); ok {
		dsDescValue, ok := dsDesc.Value.GetField("dsDescriptionValue")
		if ok {
			str = dsDescValue.Value.String()
			if dsDescDate, ok := dsDescValue.Value.GetField("dsDescriptionDate"); ok {
				str += " (" + dsDescDate.Value.String() + ")"
			}
		}
	}
	return str
}

func (data Grodata) GetAuthors() []Author {
	var authors []Author
	if author, ok := data.DatasetVersion.MetadataBlocks.Citation.Fields.GetField("author"); ok {
		for _, a := range author.Value.Fields {
			var authorName string
			if authorNameStruct, ok := a["authorName"]; ok {
				authorName = authorNameStruct.Value.String()
			}
			var authorAffiliation string
			if authorAffiliationStruct, ok := a["authorAffiliation"]; ok {
				authorAffiliation = authorAffiliationStruct.Value.String()
			}
			var authorIdentifierScheme string
			if authorIdentifierSchemeStruct, ok := a["authorIdentifierScheme"]; ok {
				authorIdentifierScheme = authorIdentifierSchemeStruct.Value.String()
			}
			var authorIdentifier string
			if authorIdentifierStruct, ok := a["authorIdentifier"]; ok {
				authorIdentifier = authorIdentifierStruct.Value.String()
			}
			authors = append(authors, Author{
				AuthorName:             authorName,
				AuthorAffiliation:      authorAffiliation,
				AuthorIdentifierScheme: authorIdentifierScheme,
				AuthorIdentifier:       authorIdentifier,
			})
		}
	}
	return authors
}
