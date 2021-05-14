package gotion

import "time"

type Response struct {
	Object  ObjectType `json:"object"`
	Code    string     `json:"code"`
	Message string     `json:"message"`
}

type DatabaseList struct {
	Results    []*DatabaseObject `json:"results,omitempty"`
	NextCursor string            `json:"next_cursor"`
	HasMore    bool              `json:"has_more"`
}

type UUID = string

type ObjectType string

const (
	User     ObjectType = "user"
	Page                = "page"
	Database            = "database"
	Error               = "error"
)

type ParentType string

const (
	DatabaseParent  ParentType = "database_id"
	PageParent                 = "page_id"
	WorkspaceParent            = "workspace"
)

type Parent struct {
	Type       ParentType `json:"type,omitempty"`
	DatabaseId UUID       `json:"database_id,omitempty"`
	PageId     UUID       `json:"page_id,omitempty"`
}

type PageObject struct {
	Object     ObjectType                `json:"object,omitempty"`
	Id         UUID                      `json:"id,omitempty"`
	CreatedAt  *time.Time                `json:"created_time,omitempty"`
	ModifiedAt *time.Time                `json:"last_edited_time,omitempty"`
	Archived   bool                      `json:"archived,omitempty"`
	Properties map[string]*PropertyValue `json:"properties"`
	Parent     *Parent                   `json:"parent,omitempty"`
}

type DatabaseObject struct {
	Object     ObjectType                 `json:"object,omitempty"`
	Id         UUID                       `json:"id,omitempty"`
	CreatedAt  time.Time                  `json:"created_time,omitempty"`
	ModifiedAt time.Time                  `json:"last_edited_time,omitempty"`
	Title      []*RichTextObject          `json:"title"`
	Properties map[string]*PropertyConfig `json:"properties"`
}

type PropertyType string

const (
	Title        PropertyType = "title"
	RichText                  = "rich_text"
	Number                    = "number"
	Select                    = "select"
	MultiSelect               = "multi_select"
	DateProperty              = "date"
	People                    = "people"
	File                      = "file"
	Checkbox                  = "checkbox"
	Url                       = "url"
	Email                     = "email"
	PhoneNumber               = "phone_number"
	Formula                   = "formula"
	Relation                  = "relation"
	Rollup                    = "rollup"
	CreatedAt                 = "created_time"
	CreatedBy                 = "created_by"
	ModifiedAt                = "last_edited_time"
	ModifiedBy                = "last_edited_by"
)

type NumberFormat string

const (
	PlainNumber      NumberFormat = "number"
	NumberWithCommas              = "number_with_commas"
	Percent                       = "percent"
	Dollar                        = "dollar"
	Euro                          = "euro"
	Pound                         = "pound"
	Yen                           = "yen"
	Ruble                         = "ruble"
	Rupee                         = "rupee"
	Won                           = "won"
	Yuan                          = "yuan"
)

type NumberConfig struct {
	Format NumberFormat `json:"format"`
}

type SelectOption struct {
	Name  string `json:"name"`
	Id    UUID   `json:"id"`
	Color Color  `json:"color"`
}

type SelectConfig struct {
	Options []*SelectOption `json:"options"`
}

type MultiSelectConfig struct {
	Options []*SelectOption `json:"options"`
}

type FormulaConfig struct {
	Expression string `json:"expression"`
}

type RelationConfig struct {
	DatabaseId   UUID   `json:"database_id"`
	PropertyName string `json:"synced_property_name,omitempty"`
	PropertyId   string `json:"synced_property_id,omitempty"`
}

type RollupFunction string

const (
	CountAll          RollupFunction = "count_all"
	CountValues                      = "count_values"
	CountUniqueValues                = "count_unique_values"
	CountEmpty                       = "count_empty"
	CountNotEmpty                    = "count_not_empty"
	PercentEmpty                     = "percent_empty"
	PercentNotEmpty                  = "percent_not_empty"
	Sum                              = "sum"
	Average                          = "average"
	Median                           = "median"
	Min                              = "min"
	Max                              = "max"
	Range                            = "range"
)

type RollupConfig struct {
	RelationPropertyName string         `json:"relation_property_name"`
	RelationPropertyId   string         `json:"relation_property_id"`
	RollupPropertyName   string         `json:"rollup_property_name"`
	RollupPropertyId     string         `json:"rollup_property_id"`
	Function             RollupFunction `json:"function"`
}

type PropertyConfig struct {
	Id   string       `json:"id"`
	Type PropertyType `json:"type"`

	// Title       struct{}          `json:"title,omitempty"`
	// Text        struct{}          `json:"rich_text,omitempty"`
	Number      *NumberConfig      `json:"number,omitempty"`
	Select      *SelectConfig      `json:"select,omitempty"`
	MultiSelect *MultiSelectConfig `json:"multi_select,omitempty"`
	// Date        struct{}          `json:"date,omitempty"`
	// People      struct{}          `json:"people,omitempty"`
	// File        struct{}          `json:"file,omitempty"`
	// Checkbox    struct{}          `json:"checkbox,omitempty"`
	// Url         struct{}          `json:"url,omitempty"`
	// Email       struct{}          `json:"email,omitempty"`
	// PhoneNumber struct{}          `json:"phone_number,omitempty"`
	Formula  *FormulaConfig  `json:"formula,omitempty"`
	Relation *RelationConfig `json:"relation,omitempty"`
	Rollup   *RollupConfig   `json:"rollup,omitempty"`
	// CreatedAt   struct{}          `json:"created_time,omitempty"`
	// CreatedBy   struct{}          `json:"created_by,omitempty"`
	// ModifiedAt  struct{}          `json:"last_edited_time,omitempty"`
	// ModifiedBy  struct{}          `json:"last_edited_by,omitempty"`
}

type DateValue struct {
	Start *time.Time `json:"start"`
	End   *time.Time `json:"end,omitempty"`
}

type FormulaResultType string

const (
	StringResult  FormulaResultType = "string"
	NumberResult                    = "number"
	BooleanResult                   = "boolean"
	DateResult                      = "date"
)

type FormulaResult struct {
	Type    FormulaResultType `json:"type"`
	String  string            `json:"string,omitempty"`
	Number  float64           `json:"number,omitempty"`
	Boolean bool              `json:"boolean,omitempty"`
	Date    *DateValue        `json:"date,omitempty"`
}

type RollupType string

const (
	NumberRollup RollupType = "number"
	DateRollup              = "date"
	ArrayRollup             = "array"
)

type FileObject struct {
	Name string `json:"name"`
}

type RollupValue struct {
	Type   RollupType       `json:"type"`
	Number float64          `json:"number,omitempty"`
	Date   *DateValue       `json:"date,omitempty"`
	Array  []*PropertyValue `json:"array,omitempty"`
}

type PropertyValue struct {
	Id   string       `json:"id,omitempty"`
	Type PropertyType `json:"type"`

	Title       []*RichTextObject  `json:"title,omitempty"`
	Text        []*RichTextObject  `json:"rich_text,omitempty"`
	Number      float64            `json:"number,omitempty"`
	Select      *SelectOption      `json:"select,omitempty"`
	MultiSelect []*SelectOption    `json:"multi_select,omitempty"`
	Date        *DateValue         `json:"date,omitempty"`
	Formula     *FormulaResult     `json:"formula,omitempty"`
	Relation    []*ObjectReference `json:"relation,omitempty"`
	Rollup      *RollupValue       `json:"rollup,omitempty"`
	People      []*UserObject      `json:"people,omitempty"`
	File        *FileObject        `json:"file,omitempty"`
	Checkbox    bool               `json:"checkbox,omitempty"`
	Url         string             `json:"url,omitempty"`
	Email       string             `json:"email,omitempty"`
	PhoneNumber string             `json:"phone_number,omitempty"`
	CreatedAt   *time.Time         `json:"created_time,omitempty"`
	CreatedBy   *UserObject        `json:"created_by,omitempty"`
	ModifiedAt  *time.Time         `json:"last_edited_time,omitempty"`
	ModifiedBy  *UserObject        `json:"last_edited_by,omitempty"`
}

type UserType string

const (
	Person UserType = "person"
	Bot             = "bot"
)

type PersonObject struct {
	Email string `json:"email"`
}

type BotObject struct {
}

type UserObject struct {
	Object    ObjectType `json:"object"`
	Id        UUID       `json:"id"`
	Type      UserType   `json:"type,omitempty"`
	AvatarUrl string     `json:"avatar_url,omitempty"`

	Person PersonObject `json:"person,omitempty"`
	Bot    BotObject    `json:"bot,omitemtpy"`
}

type Color string

const (
	Default          Color = "default"
	Gray                   = "gray"
	Brown                  = "brown"
	Orange                 = "orange"
	Yellow                 = "yellow"
	Green                  = "green"
	Blue                   = "blue"
	Purple                 = "purple"
	Pink                   = "pink"
	Red                    = "red"
	GrayBackground         = "gray_background"
	BrownBackground        = "brown_background"
	OrangeBackground       = "orange_background"
	YellowBackground       = "yellow_background"
	GreenBackground        = "green_background"
	BlueBackground         = "blue_background"
	PurpleBackground       = "purple_background"
	PinkBackground         = "pink_background"
	RedBackground          = "red_background"
)

type Annotations struct {
	Bold          bool  `json:"bold"`
	Italic        bool  `json:"italic"`
	Strikethrough bool  `json:"strikethrough"`
	Underline     bool  `json:"underline"`
	Code          bool  `json:"code"`
	Color         Color `json:"color"`
}

type RichTextType string

const (
	Text     RichTextType = "text"
	Mention               = "mention"
	Equation              = "equation"
)

type LinkType string

const (
	LinkUrl LinkType = "url"
)

type LinkObject struct {
	Type LinkType `json:"type"`
	Url  string   `json:"url"`
}

type TextObject struct {
	Content string      `json:"content"`
	Link    *LinkObject `json:"link,omitempty"`
}

type ObjectReference struct {
	Object ObjectType `json:"object,omitempty"`
	Id     UUID       `json:"id"`
}

type MentionObject struct {
	Type ObjectType `json:"type"`

	User     *UserObject      `json:"user,omitempty"`
	Page     *ObjectReference `json:"page,omitempty"`
	Database *ObjectReference `json:"database,omitempty"`
}

type RichTextObject struct {
	PlainText   string        `json:"plain_text"`
	Href        string        `json:"href,omitempty"`
	Annotations *Annotations  `json:"annotations"`
	Type        *RichTextType `json:"type,omitempty"`

	Text     *TextObject    `json:"text,omitempty"`
	Mention  *MentionObject `json:"mention,omitempty"`
	Date     *DateValue     `json:"date,omitempty"`
	Equation *FormulaConfig `json:"equation,omitempty"`
}
