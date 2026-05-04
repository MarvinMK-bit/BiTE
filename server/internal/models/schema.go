package models

import (
	"time"

	"github.com/lib/pq"
)

var db = Db()

type User struct {
	ID                                string    `gorm:"column:id;type:uuid;primaryKey" json:"id"`
	Name                              string    `gorm:"column:name;not null;index" json:"name"`
	Password                          string    `gorm:"column:password;not null" json:"password"`
	Role                              string    `gorm:"column:role;default:'USER';not null;index" json:"role"`
	Email                             string    `gorm:"column:email;not null;unique;index" json:"email"`
	Gender                            string    `gorm:"column:gender;default:null" json:"gender,omitempty"`
	DateOfBirth                       string    `gorm:"column:dateOfBirth;default:null" json:"dateOfBirth,omitempty"`
	Country                           string    `gorm:"column:country;default:null" json:"country,omitempty"`
	ImageUrl                          string    `gorm:"column:imageUrl;default:null" json:"imageUrl"`
	ImagePath                         string    `gorm:"column:imagePath;default:null" json:"imagePath,omitempty"`
	ProfileBgColor                    string    `gorm:"column:profileBgColor;default:null" json:"profileBgColor"`
	AgreedTermsOfService              bool      `gorm:"column:agreedTermsOfService;default:true" json:"agreedTermsOfService"`
	AgreedNewsLetterOrMarketingEmails bool      `gorm:"column:agreedNewsLetterOrMarketingEmails;default:false" json:"agreedNewsLetterOrMarketingEmails"` // Agreed to receive marketing or newsletter emails
	CreatedAt                         time.Time `gorm:"column:createdAt" json:"createdAt"`
	UpdatedAt                         time.Time `gorm:"column:updatedAt" json:"updatedAt"`

	// Relationships
	OTPs             []OTP                 `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"otps,omitempty"`
	Sessions         []*Session            `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"sessions,omitempty"`
	Locations        []*Location           `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"locations,omitempty"`
	Quizzes          []*Quiz               `gorm:"foreignKey:PostedByUserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"quizzes,omitempty"`
	Questions        []*Question           `gorm:"foreignKey:PostedByUserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"questions,omitempty"`
	Answers          []*Answer             `gorm:"foreignKey:PostedByUserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"answers,omitempty"`
	Attempts         []*Attempt            `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"attempts,omitempty"`
	AttemptDurations []*AttemptDuration    `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"attemptDurations,omitempty"`
	Rankings         *Ranking              `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"rankings,omitempty"`
	QuizUserProgress []*QuizUserProgress   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"quizUserProgress,omitempty"`
	SiteVisits       []*SiteVisit          `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"siteVisits,omitempty"`
	Attachments      []*Attachment         `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"attachments,omitempty"`
	Certificates     []*CertificateAwarded `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"certificates,omitempty"`
}

type OTP struct {
	ID         string    `gorm:"column:id;type:uuid;primaryKey" json:"id"`
	UserID     string    `gorm:"column:userID;not null;index" json:"userID"`
	OTP        string    `gorm:"column:OTP;not null;index" json:"OTP"`
	IsUsed     bool      `gorm:"column:isUsed;default:false" json:"isUsed"`
	IsVerified bool      `gorm:"column:isVerified;default:false" json:"isVerified"`
	ExpiresAt  time.Time `gorm:"column:expiresAt;not null;index" json:"expiresAt"`
	CreatedAt  time.Time `gorm:"column:createdAt;index" json:"createdAt"`
	UpdatedAt  time.Time `gorm:"column:updatedAt;index" json:"updatedAt"`

	// Relationships
	User *User `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"user,omitempty"`
}

type Session struct {
	ID           string    `gorm:"column:id;type:uuid;primaryKey" json:"id"`
	UserID       string    `gorm:"column:userID;not null;index" json:"userID"`
	AccessToken  string    `gorm:"column:accessToken;not null;index" json:"accessToken"`
	RefreshToken string    `gorm:"column:refreshToken;not null;index" json:"refreshToken"`
	GeneratedVia string    `gorm:"column:generatedVia;not null;index" json:"generatedVia"`
	Device       string    `gorm:"column:Device;default:'Unknown Device';index" json:"device"`
	LocationID   string    `gorm:"column:locationID;default:null" json:"locationID"`
	IsRevoked    bool      `gorm:"column:isRevoked;default:false" json:"isRevoked"`
	CreatedAt    time.Time `gorm:"column:createdAt;index" json:"createdAt"`
	UpdatedAt    time.Time `gorm:"column:updatedAt;index" json:"updatedAt"`

	// Relationships
	User     *User     `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"user,omitempty"`
	Location *Location `gorm:"foreignKey:LocationID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"location,omitempty"`
}

type Location struct {
	ID        string    `gorm:"column:id;type:uuid;primaryKey" json:"id"`
	UserID    string    `gorm:"column:userID;not null;index" json:"userID"`
	Info      JSONB     `gorm:"column:info;type:jsonb;not null;" json:"info"`
	CreatedAt time.Time `gorm:"column:createdAt;index" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updatedAt;index" json:"updatedAt"`

	// Relationships
	User *User `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"user,omitempty"`
}

type QuizCategory struct {
	ID        string    `gorm:"column:id;type:uuid;primaryKey" json:"id"`
	Name      string    `gorm:"column:name;unique;not null;index" json:"name"`
	Color     string    `gorm:"column:color;default:null;" json:"color"`
	CreatedAt time.Time `gorm:"column:createdAt;index" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updatedAt;index" json:"updatedAt"`

	// Relationships
	Quizzes     []*Quiz              `gorm:"foreignKey:QuizCategoryID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"quizzes,omitempty"`
	Attachments []*Attachment        `gorm:"foreignKey:QuizCategoryID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"attachments"`
	Certificate *CategoryCertificate `gorm:"foreignKey:QuizCategoryID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"certificate"`
}

type Quiz struct {
	ID                string    `gorm:"column:id;type:uuid;primaryKey" json:"id"`
	Title             string    `gorm:"column:title;not null;index" json:"title"`
	TitleDelta        string    `gorm:"column:titleDelta;default:null;index" json:"titleDelta"`
	TitleHTML         string    `gorm:"column:titleHTML;default:null;index" json:"titleHTML"`
	Instructions      string    `gorm:"column:instructions;default:null" json:"instructions"`
	InstructionsDelta string    `gorm:"column:instructionsDelta;default:null" json:"instructionsDelta"`
	InstructionsHTML  string    `gorm:"column:instructionsHTML;default:null" json:"instructionsHTML"`
	Introduction      string    `gorm:"column:introduction;default:null" json:"introduction"`
	IntroductionDelta string    `gorm:"column:introductionDelta;default:null" json:"introductionDelta"`
	IntroductionHTML  string    `gorm:"column:introductionHTML;default:null" json:"introductionHTML"`
	IsDeltaDefault    bool      `gorm:"column:isDeltaDefault;default:false;index" json:"isDeltaDefault"`
	PostedByUserID    string    `gorm:"column:postedByUserID;not null;index" json:"postedByUserID"`
	QuizCategoryID    string    `gorm:"column:quizCategoryID;default:null;index" json:"quizCategoryID"`
	StartsAt          time.Time `gorm:"column:startsAt;index" json:"startsAt"`
	EndsAt            time.Time `gorm:"column:endsAt;index" json:"endsAt"`
	CanBeAttempted    bool      `gorm:"column:canBeAttempted;default:false;index" json:"canBeAttempted"`
	IsDuplicate       bool      `gorm:"column:isDuplicate;default:false;index" json:"isDuplicate"`
	ShowQuiz          bool      `gorm:"column:showQuiz;default:false;index" json:"showQuiz"`
	CreatedAt         time.Time `gorm:"column:createdAt;index" json:"createdAt"`
	UpdatedAt         time.Time `gorm:"column:updatedAt;index" json:"updatedAt"`

	// Relationships
	PostedByUser     *User              `gorm:"foreignKey:PostedByUserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"postedByUser,omitempty"`
	Questions        []*Question        `gorm:"foreignKey:QuizID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"questions"`
	Attachments      []*Attachment      `gorm:"foreignKey:QuizID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"attachments"`
	Attempts         []*Attempt         `gorm:"foreignKey:QuizID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"attempts,omitempty"`
	AttemptDurations []*AttemptDuration `gorm:"foreignKey:QuizID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"attemptDurations,omitempty"`
	QuizCategory     *QuizCategory      `gorm:"foreignKey:QuizCategoryID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"quizCategory,omitempty"`
}

type QuizDuplicate struct {
	ID                   string    `gorm:"column:id;type:uuid;primaryKey" json:"id"`
	DuplicatedByUserID   string    `gorm:"column:duplicatedByUserID;not null;index" json:"duplicatedByUserID"`
	DuplicatedFromQuizID string    `gorm:"column:duplicatedFromQuizID;not null;index" json:"duplicatedFromQuizID"`
	DuplicatedQuizID     string    `gorm:"column:duplicatedQuizID;not null;index" json:"duplicatedQuizID"`
	CreatedAt            time.Time `gorm:"column:createdAt;index" json:"createdAt"`
	UpdatedAt            time.Time `gorm:"column:updatedAt;index" json:"updatedAt"`

	// Relationships
	DuplicatedByUser   *User `gorm:"foreignKey:DuplicatedByUserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"duplicatedByUser,omitempty"`
	DuplicatedFromQuiz *Quiz `gorm:"foreignKey:DuplicatedFromQuizID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"duplicatedFromQuiz,omitempty"`
	DuplicatedQuiz     *Quiz `gorm:"foreignKey:DuplicatedQuizID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"duplicatedQuiz,omitempty"`
}

type Question struct {
	ID                        string    `gorm:"column:id;type:uuid;primaryKey" json:"id"`
	Title                     string    `gorm:"column:title;not null;index" json:"title"`
	TitleDelta                string    `gorm:"column:titleDelta;default:null;index" json:"titleDelta"`
	TitleHTML                 string    `gorm:"column:titleHTML;default:null;index" json:"titleHTML"`
	Introduction              string    `gorm:"column:introduction;default:null" json:"introduction"`
	IntroductionDelta         string    `gorm:"column:introductionDelta;default:null" json:"introductionDelta"`
	IntroductionHTML          string    `gorm:"column:introductionHTML;default:null" json:"introductionHTML"`
	IsDeltaDefault            bool      `gorm:"column:isDeltaDefault;default:false;index" json:"isDeltaDefault"`
	PostedByUserID            string    `gorm:"column:postedByUserID;not null;index" json:"postedByUserID"`
	QuizID                    string    `gorm:"column:quizID;not null;index" json:"quizID"`
	SequenceNumber            int64     `gorm:"column:sequenceNumber;not null;index" json:"sequenceNumber"`
	HasMultipleCorrectAnswers bool      `gorm:"column:hasMultipleCorrectAnswers;default:false;index" json:"hasMultipleCorrectAnswers"`
	RequiresNumericalAnswer   bool      `gorm:"column:requiresNumericalAnswer;default:false;index" json:"requiresNumericalAnswer"`
	ShowAIPreview             bool      `gorm:"column:showAIPreview;default:false;index" json:"showAIPreview"`
	CreatedAt                 time.Time `gorm:"column:createdAt;index" json:"createdAt"`
	UpdatedAt                 time.Time `gorm:"column:updatedAt;index" json:"updatedAt"`

	// Relationships
	PostedByUser    *User            `gorm:"foreignKey:PostedByUserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"postedByUser,omitempty"`
	Quiz            *Quiz            `gorm:"foreignKey:QuizID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"quiz,omitempty"`
	Answers         []*Answer        `gorm:"foreignKey:QuestionID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"answers"`
	Attachments     []*Attachment    `gorm:"foreignKey:QuestionID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"attachments"`
	Attempts        []*Attempt       `gorm:"foreignKey:QuestionID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"attempts,omitempty"`
	AttemptStatuses []*AttemptStatus `gorm:"foreignKey:QuestionID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"attemptStatuses,omitempty"`
	AIPreviews      []*AIPreview     `gorm:"foreignKey:QuestionID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"aiPreviews,omitempty"`
}

type Answer struct {
	ID             string    `gorm:"column:id;type:uuid;primaryKey" json:"id"`
	Title          string    `gorm:"column:title;not null;index" json:"title"`
	TitleDelta     string    `gorm:"column:titleDelta;default:null;index" json:"titleDelta"`
	TitleHTML      string    `gorm:"column:titleHTML;default:null;index" json:"titleHTML"`
	IsDeltaDefault bool      `gorm:"column:isDeltaDefault;default:false;index" json:"isDeltaDefault"`
	PostedByUserID string    `gorm:"column:postedByUserID;not null;index" json:"postedByUserID"`
	QuestionID     string    `gorm:"column:questionID;not null;index" json:"questionID"`
	SequenceNumber int64     `gorm:"column:sequenceNumber;not null;index" json:"sequenceNumber"`
	IsCorrect      bool      `gorm:"column:isCorrect;default:false;index" json:"isCorrect"`
	CreatedAt      time.Time `gorm:"column:createdAt;index" json:"createdAt"`
	UpdatedAt      time.Time `gorm:"column:updatedAt;index" json:"updatedAt"`

	// Relationships
	PostedByUser *User         `gorm:"foreignKey:PostedByUserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"postedByUser,omitempty"`
	Question     *Question     `gorm:"foreignKey:QuestionID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"question,omitempty"`
	Attachments  []*Attachment `gorm:"foreignKey:AnswerID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"attachments"`
	Attempts     []*Attempt    `gorm:"foreignKey:AnswerID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"attempts,omitempty"`
}

type Attempt struct {
	ID          string    `gorm:"column:id;type:uuid;primaryKey" json:"id"`
	UserID      string    `gorm:"column:userID;not null;index" json:"userID"`
	QuizID      string    `gorm:"column:quizID;not null;index" json:"quizID"`
	QuestionID  string    `gorm:"column:questionID;not null;index" json:"questionID"`
	AnswerID    string    `gorm:"column:answerID;not null;index" json:"answerID"`
	AnswerInput string    `gorm:"column:answerInput;default:null" json:"answerInput"`
	CreatedAt   time.Time `gorm:"column:createdAt;index" json:"createdAt"`
	UpdatedAt   time.Time `gorm:"column:updatedAt;index" json:"updatedAt"`

	// Relationships
	User          *User          `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"user,omitempty"`
	Quiz          *Quiz          `gorm:"foreignKey:QuizID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"quiz,omitempty"`
	Question      *Question      `gorm:"foreignKey:QuestionID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"question,omitempty"`
	Answer        *Answer        `gorm:"foreignKey:AnswerID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"answer,omitempty"`
	AttemptStatus *AttemptStatus `gorm:"foreignKey:AttemptID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"attemptStatus,omitempty"`
}

type AttemptStatus struct {
	ID         string    `gorm:"column:id;type:uuid;primaryKey" json:"id"`
	UserID     string    `gorm:"column:userID;not null;index" json:"userID"`
	AttemptID  string    `gorm:"column:attemptID;not null;index" json:"attemptID"`
	QuestionID string    `gorm:"column:questionID;not null;index" json:"questionID"`
	IsCorrect  bool      `gorm:"column:IsCorrect;not null;default:false;index" json:"IsCorrect"`
	CreatedAt  time.Time `gorm:"column:createdAt;index" json:"createdAt"`
	UpdatedAt  time.Time `gorm:"column:updatedAt;index" json:"updatedAt"`

	// Relationships
	User     *User     `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"user,omitempty"`
	Attempt  *Attempt  `gorm:"foreignKey:AttemptID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"attempt,omitempty"`
	Question *Question `gorm:"foreignKey:QuestionID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"question,omitempty"`
}

type AttemptDuration struct {
	ID        string    `gorm:"column:id;type:uuid;primaryKey" json:"id"`
	UserID    string    `gorm:"column:userID;not null;index" json:"userID"`
	QuizID    string    `gorm:"column:quizID;not null;index" json:"quizID"`
	Duration  int64     `gorm:"column:duration;not null;index" json:"duration"` // Duration in seconds
	CreatedAt time.Time `gorm:"column:createdAt;index" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updatedAt;index" json:"updatedAt"`

	// Relationships
	User *User `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"user,omitempty"`
	Quiz *Quiz `gorm:"foreignKey:QuizID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"quiz,omitempty"`
}

type AIPreview struct {
	ID         string    `gorm:"column:id;type:uuid;primaryKey" json:"id"`
	QuestionID string    `gorm:"column:questionID;not null;index" json:"questionID"`
	Summary    string    `gorm:"column:summary;not null;index" json:"summary"`
	Prompt     string    `gorm:"column:prompt;not null;index" json:"prompt"`
	IsDefault  bool      `gorm:"column:isDefault;not null;default:false;index" json:"isDefault"`
	CreatedAt  time.Time `gorm:"column:createdAt;index" json:"createdAt"`
	UpdatedAt  time.Time `gorm:"column:updatedAt;index" json:"updatedAt"`

	// Relationships
	Question *Question `gorm:"foreignKey:QuestionID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"question,omitempty"`
}

type Ranking struct {
	ID                   string    `gorm:"column:id;type:uuid;primaryKey" json:"id"`
	UserID               string    `gorm:"column:userID;not null;index" json:"userID"`
	TotalDuration        int64     `gorm:"column:totalDuration;not null;index" json:"totalDuration"` // Duration in seconds
	TotalAttempts        int64     `gorm:"column:totalAttempts;not null;index" json:"totalAttempts"`
	TotalCorrectAttempts int64     `gorm:"column:totalCorrectAttempts;not null;index" json:"totalCorrectAttempts"`
	Rank                 int64     `gorm:"column:rank;not null;index" json:"rank"`
	CreatedAt            time.Time `gorm:"column:createdAt;index" json:"createdAt"`
	UpdatedAt            time.Time `gorm:"column:updatedAt;index" json:"updatedAt"`

	// Relationships
	User *User `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"user,omitempty"`
}
type QuizUserProgress struct {
	ID                      string    `gorm:"column:id;type:uuid;primaryKey" json:"id"`
	UserID                  string    `gorm:"column:userID;not null;index" json:"userID"`
	QuizID                  string    `gorm:"column:quizID;not null;index" json:"quizID"`
	QuizCategoryID          string    `gorm:"column:quizCategoryID;default:null;index" json:"quizCategoryID"`
	TotalQuestions          int64     `gorm:"column:totalQuestions;not null;index" json:"totalQuestions"`
	TotalAttemptedQuestions int64     `gorm:"column:totalAttemptedQuestions;not null" json:"totalAttemptedQuestions"`
	Status                  string    `gorm:"column:status;not null;index" json:"status"` // COMPLETED, IN_PROGRESS
	CreatedAt               time.Time `gorm:"column:createdAt;index" json:"createdAt"`
	UpdatedAt               time.Time `gorm:"column:updatedAt;index" json:"updatedAt"`

	// Relationships
	User         *User         `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"user,omitempty"`
	Quiz         *Quiz         `gorm:"foreignKey:QuizID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"quiz,omitempty"`
	QuizCategory *QuizCategory `gorm:"foreignKey:QuizCategoryID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"quizCategory,omitempty"`
}

type SiteVisit struct {
	ID         string    `gorm:"column:id;type:uuid;primaryKey" json:"id"`
	UserID     string    `gorm:"column:userID;default:null;index" json:"userID"`
	Device     string    `gorm:"column:device;default:UNKNOWN" json:"device"`
	Page       string    `gorm:"column:page;not null;index" json:"page"`
	Path       string    `gorm:"column:path;not null;index" json:"path"`
	LocationID string    `gorm:"column:locationID;not null;index" json:"locationID"`
	CapturedAt time.Time `gorm:"column:capturedAt;not null;index" json:"capturedAt"`
	CreatedAt  time.Time `gorm:"column:createdAt;index" json:"createdAt"`
	UpdatedAt  time.Time `gorm:"column:updatedAt;index" json:"updatedAt"`

	// Relationships
	User     *User     `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"user,omitempty"`
	Location *Location `gorm:"foreignKey:LocationID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"location,omitempty"`
}

type Attachment struct {
	ID                   string    `gorm:"column:id;type:uuid;primaryKey" json:"id"`
	Type                 string    `gorm:"column:type;not null;index" json:"type"` //USER, QUIZ,QUIZCATEGORY, QUESTION, ANSWER, CERTIFICATE_PDF, CERTIFICATE_PNG
	UserID               string    `gorm:"column:userID;default:null;index" json:"userID,omitempty"`
	QuizID               string    `gorm:"column:quizID;default:null;index" json:"quizID,omitempty"`
	QuizCategoryID       string    `gorm:"column:quizCategoryID;default:null;index" json:"quizCategoryID,omitempty"`
	QuestionID           string    `gorm:"column:questionID;default:null;index" json:"questionID,omitempty"`
	AnswerID             string    `gorm:"column:answerID;default:null;index" json:"answerID,omitempty"`
	CertificateAwardedID string    `gorm:"column:certificateAwardedID;default:null;index" json:"certificateAwardedID,omitempty"`
	Filename             string    `gorm:"column:filename;not null" json:"filename"`
	Url                  string    `gorm:"column:url;not null" json:"url"`
	Size                 int64     `gorm:"column:size;not null" json:"size"`
	ContentType          string    `gorm:"column:contentType;not null" json:"contentType"`
	CreatedAt            time.Time `gorm:"column:createdAt;index" json:"createdAt"`
	UpdatedAt            time.Time `gorm:"column:updatedAt;index" json:"updatedAt"`

	// Relationships
	User         *User               `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"user,omitempty"`
	Quiz         *Quiz               `gorm:"foreignKey:QuizID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"quiz,omitempty"`
	QuizCategory *QuizCategory       `gorm:"foreignKey:QuizCategoryID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"quizCategory,omitempty"`
	Question     *Question           `gorm:"foreignKey:QuestionID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"question,omitempty"`
	Answer       *Answer             `gorm:"foreignKey:AnswerID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"answer,omitempty"`
	Certificate  *CertificateAwarded `gorm:"foreignKey:CertificateAwardedID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"certificate,omitempty"`
}

type SatsReward struct {
	ID                  string    `gorm:"column:id;type:uuid;primaryKey" json:"id"`
	UserID              string    `gorm:"column:userID;not null;index" json:"userID"`
	QuizID              string    `gorm:"column:quizID;not null;index" json:"quizID"`
	SatsRewardAddressID string    `gorm:"column:satsRewardAddressID;not null;index" json:"satsRewardAddressID"`
	Status              string    `gorm:"column:status;not null;index" json:"status"` // PENDING, COMPLETED, FAILED
	CreatedAt           time.Time `gorm:"column:createdAt;index" json:"createdAt"`
	UpdatedAt           time.Time `gorm:"column:updatedAt;index" json:"updatedAt"`

	// Relationships
	User                  *User                    `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"user,omitempty"`
	Quiz                  *Quiz                    `gorm:"foreignKey:QuizID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"quiz,omitempty"`
	SatsRewardAddress     *SatsRewardAddress       `gorm:"foreignKey:SatsRewardAddressID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"satsRewardAddress,omitempty"`
	SatsRewardTransaction []*SatsRewardTransaction `gorm:"foreignKey:SatsRewardID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"satsRewardTransactions,omitempty"`
	SatsRewardOperation   []*SatsRewardOperation   `gorm:"foreignKey:SatsRewardID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"satsRewardOperations,omitempty"`
}

type SatsRewardAddress struct {
	ID         string    `gorm:"column:id;type:uuid;primaryKey" json:"id"`
	UserID     string    `gorm:"column:userID;not null;index" json:"userID"`
	Address    string    `gorm:"column:address;not null;unique;index" json:"address"`
	IsVerified bool      `gorm:"column:isVerified;default:false;index" json:"isVerified"`
	IsDefault  bool      `gorm:"column:isDefault;default:false;index" json:"isDefault"`
	Info       JSONB     `gorm:"column:info" json:"info"`
	CreatedAt  time.Time `gorm:"column:createdAt;index" json:"createdAt"`
	UpdatedAt  time.Time `gorm:"column:updatedAt;index" json:"updatedAt"`

	// Relationships
	User       *User         `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"user,omitempty"`
	SatsReward []*SatsReward `gorm:"foreignKey:SatsRewardAddressID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"satsRewards,omitempty"`
}

type SatsRewardTransaction struct {
	ID                  string    `gorm:"column:id;type:uuid;primaryKey" json:"id"`
	SatsRewardID        string    `gorm:"column:satsRewardID;not null;index" json:"satsRewardID"`
	Transaction         JSONB     `gorm:"column:transaction;not null" json:"transaction"`
	RewardedQuestionIDs JSONB     `gorm:"column:rewardedQuestionIDs;not null" json:"rewardedQuestionIDs"`
	CreatedAt           time.Time `gorm:"column:createdAt;index" json:"createdAt"`
	UpdatedAt           time.Time `gorm:"column:updatedAt;index" json:"updatedAt"`

	// Relationships
	SatsReward *SatsReward `gorm:"foreignKey:SatsRewardID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"satsReward,omitempty"`
}

type SatsRewardOperation struct {
	ID           string    `gorm:"column:id;type:uuid;primaryKey" json:"id"`
	SatsRewardID string    `gorm:"column:satsRewardID;not null;index" json:"satsRewardID"`
	Status       string    `gorm:"column:status;not null" json:"status"` //SUCCESS, FAIL
	Info         string    `gorm:"column:info;not null" json:"info"`
	CreatedAt    time.Time `gorm:"column:createdAt;index" json:"createdAt"`
	UpdatedAt    time.Time `gorm:"column:updatedAt;index" json:"updatedAt"`

	// Relationships
	SatsReward *SatsReward `gorm:"foreignKey:SatsRewardID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"satsReward,omitempty"`
}

type CategoryCertificate struct {
	ID             string    `gorm:"column:id;type:uuid;primaryKey" json:"id"`
	QuizCategoryID string    `gorm:"column:quizCategoryID;not null;index" json:"quizCategoryID"`
	CreatedAt      time.Time `gorm:"column:createdAt;index" json:"createdAt"`
	UpdatedAt      time.Time `gorm:"column:updatedAt;index" json:"updatedAt"`

	// Relationships
	QuizCategory               *QuizCategory                 `gorm:"foreignKey:QuizCategoryID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"quizCategory,omitempty"`
	CategoryCertificateQuizzes []*CategoryCertificateQuizzes `gorm:"foreignKey:CategoryCertificateID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"categoryCertificateQuizzes,omitempty"`
}

type CategoryCertificateQuizzes struct {
	ID                    string    `gorm:"column:id;type:uuid;primaryKey" json:"id"`
	CategoryCertificateID string    `gorm:"column:categoryCertificateID;not null;index" json:"categoryCertificateID"`
	QuizID                string    `gorm:"column:quizID;not null;index" json:"quizID"`
	CreatedAt             time.Time `gorm:"column:createdAt;index" json:"createdAt"`
	UpdatedAt             time.Time `gorm:"column:updatedAt;index" json:"updatedAt"`

	// Relationships
	CategoryCertificate *CategoryCertificate `gorm:"foreignKey:CategoryCertificateID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"categoryCertificate,omitempty"`
	Quiz                *Quiz                `gorm:"foreignKey:QuizID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"quiz,omitempty"`
}

type CertificateAwarded struct {
	ID                    string    `gorm:"column:id;type:uuid;primaryKey" json:"id"`
	UserID                string    `gorm:"column:userID;not null;index" json:"userID"`
	CategoryCertificateID string    `gorm:"column:categoryCertificateID;not null;index" json:"categoryCertificateID"`
	AwardedQuizIDs        JSONB     `gorm:"column:awardedQuizIDs;not null" json:"awardedQuizIDs"`
	CreatedAt             time.Time `gorm:"column:createdAt;index" json:"createdAt"`
	UpdatedAt             time.Time `gorm:"column:updatedAt;index" json:"updatedAt"`

	// Relationships
	User                *User                `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"user,omitempty"`
	CategoryCertificate *CategoryCertificate `gorm:"foreignKey:CategoryCertificateID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"categoryCertificate,omitempty"`
	Attachments         []*Attachment        `gorm:"foreignKey:CertificateAwardedID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"attachments,omitempty"`
}

type ChessPuzzle struct {
	ID              string         `gorm:"column:id;primaryKey"`
	FEN             string         `gorm:"column:fen;not null"`
	Moves           string         `gorm:"column:moves;not null"`
	Rating          int            `gorm:"column:rating;not null;default:1500"`
	RatingDeviation int            `gorm:"column:ratingDeviation;not null;default:80"`
	Volatility      float64        `gorm:"column:volatility;not null;default:0.09"`
	Popularity      int16          `gorm:"column:popularity;not null;default:0"`
	NbPlays         int            `gorm:"column:nbPlays;not null;default:0"`
	Themes          pq.StringArray `gorm:"column:themes;type:text[];not null;default:'{}'"`
	GameUrl         string         `gorm:"column:gameUrl"`
	OpeningTags     pq.StringArray `gorm:"column:openingTags;type:text[];not null;default:'{}'"`
	Color           string         `gorm:"column:color;not null;default:'w'"`
	CreatedAt       time.Time      `gorm:"column:createdAt;not null;autoCreateTime"`
	UpdatedAt       time.Time      `gorm:"column:updatedAt;not null;autoUpdateTime"`
}

func (ChessPuzzle) TableName() string {
	return "chess_puzzles"
}

type ChessPuzzleRound struct {
	ID                 string    `gorm:"column:id;primaryKey"`
	ChessPuzzleID      string    `gorm:"column:chessPuzzleId;not null"`
	UserID             string    `gorm:"column:userId;not null"`
	Angle              string    `gorm:"column:angle;not null;default:'mix'"`
	Win                bool      `gorm:"column:win;not null"`
	TimeMs             int       `gorm:"column:timeMs"`
	SatsEarned         int       `gorm:"column:satsEarned;not null;default:0"`
	PuzzleRatingBefore int       `gorm:"column:puzzleRatingBefore;not null"`
	PuzzleRatingAfter  int       `gorm:"column:puzzleRatingAfter;not null"`
	UserRatingBefore   int       `gorm:"column:userRatingBefore;not null"`
	UserRatingAfter    int       `gorm:"column:userRatingAfter;not null"`
	CreatedAt          time.Time `gorm:"column:createdAt;not null;autoCreateTime"`
	UpdatedAt          time.Time `gorm:"column:updatedAt;not null;autoUpdateTime"`
}

func (ChessPuzzleRound) TableName() string {
	return "chess_puzzle_rounds"
}

type ChessUserPuzzleRating struct {
	ID              string     `gorm:"column:id;primaryKey"`
	UserID          string     `gorm:"column:userId;not null;uniqueIndex"`
	Rating          int        `gorm:"column:rating;not null;default:1500"`
	RatingDeviation int        `gorm:"column:ratingDeviation;not null;default:500"`
	Volatility      float64    `gorm:"column:volatility;not null;default:0.09"`
	GamesPlayed     int        `gorm:"column:gamesPlayed;not null;default:0"`
	LastPlayedAt    *time.Time `gorm:"column:lastPlayedAt"`
	CreatedAt       time.Time  `gorm:"column:createdAt;not null;autoCreateTime"`
	UpdatedAt       time.Time  `gorm:"column:updatedAt;not null;autoUpdateTime"`
}

func (ChessUserPuzzleRating) TableName() string {
	return "chess_user_puzzle_ratings"
}
