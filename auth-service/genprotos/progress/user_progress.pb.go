// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v3.12.4
// source: user_progress.proto

package progress

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DailyProgress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Date             string `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	XpEarned         int32  `protobuf:"varint,2,opt,name=xp_earned,json=xpEarned,proto3" json:"xp_earned,omitempty"`
	LessonsCompleted int32  `protobuf:"varint,3,opt,name=lessons_completed,json=lessonsCompleted,proto3" json:"lessons_completed,omitempty"`
	NewWordsLearned  int32  `protobuf:"varint,4,opt,name=new_words_learned,json=newWordsLearned,proto3" json:"new_words_learned,omitempty"`
	MinutesPracticed int32  `protobuf:"varint,5,opt,name=minutes_practiced,json=minutesPracticed,proto3" json:"minutes_practiced,omitempty"`
}

func (x *DailyProgress) Reset() {
	*x = DailyProgress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_progress_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DailyProgress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DailyProgress) ProtoMessage() {}

func (x *DailyProgress) ProtoReflect() protoreflect.Message {
	mi := &file_user_progress_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DailyProgress.ProtoReflect.Descriptor instead.
func (*DailyProgress) Descriptor() ([]byte, []int) {
	return file_user_progress_proto_rawDescGZIP(), []int{0}
}

func (x *DailyProgress) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *DailyProgress) GetXpEarned() int32 {
	if x != nil {
		return x.XpEarned
	}
	return 0
}

func (x *DailyProgress) GetLessonsCompleted() int32 {
	if x != nil {
		return x.LessonsCompleted
	}
	return 0
}

func (x *DailyProgress) GetNewWordsLearned() int32 {
	if x != nil {
		return x.NewWordsLearned
	}
	return 0
}

func (x *DailyProgress) GetMinutesPracticed() int32 {
	if x != nil {
		return x.MinutesPracticed
	}
	return 0
}

type Achievement struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title       string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	EarnedAt    string `protobuf:"bytes,4,opt,name=earned_at,json=earnedAt,proto3" json:"earned_at,omitempty"`
}

func (x *Achievement) Reset() {
	*x = Achievement{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_progress_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Achievement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Achievement) ProtoMessage() {}

func (x *Achievement) ProtoReflect() protoreflect.Message {
	mi := &file_user_progress_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Achievement.ProtoReflect.Descriptor instead.
func (*Achievement) Descriptor() ([]byte, []int) {
	return file_user_progress_proto_rawDescGZIP(), []int{1}
}

func (x *Achievement) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Achievement) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Achievement) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Achievement) GetEarnedAt() string {
	if x != nil {
		return x.EarnedAt
	}
	return ""
}

type Goal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	GoalType     string `protobuf:"bytes,2,opt,name=goal_type,json=goalType,proto3" json:"goal_type,omitempty"`
	TargetValue  int32  `protobuf:"varint,3,opt,name=target_value,json=targetValue,proto3" json:"target_value,omitempty"`
	CurrentValue int32  `protobuf:"varint,4,opt,name=current_value,json=currentValue,proto3" json:"current_value,omitempty"`
	StartDate    string `protobuf:"bytes,5,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	EndDate      string `protobuf:"bytes,6,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
}

func (x *Goal) Reset() {
	*x = Goal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_progress_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Goal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Goal) ProtoMessage() {}

func (x *Goal) ProtoReflect() protoreflect.Message {
	mi := &file_user_progress_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Goal.ProtoReflect.Descriptor instead.
func (*Goal) Descriptor() ([]byte, []int) {
	return file_user_progress_proto_rawDescGZIP(), []int{2}
}

func (x *Goal) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Goal) GetGoalType() string {
	if x != nil {
		return x.GoalType
	}
	return ""
}

func (x *Goal) GetTargetValue() int32 {
	if x != nil {
		return x.TargetValue
	}
	return 0
}

func (x *Goal) GetCurrentValue() int32 {
	if x != nil {
		return x.CurrentValue
	}
	return 0
}

func (x *Goal) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *Goal) GetEndDate() string {
	if x != nil {
		return x.EndDate
	}
	return ""
}

type UserProgress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                     string           `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId                 string           `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	LanguageCode           string           `protobuf:"bytes,3,opt,name=language_code,json=languageCode,proto3" json:"language_code,omitempty"`
	Xp                     int32            `protobuf:"varint,4,opt,name=xp,proto3" json:"xp,omitempty"`
	Level                  string           `protobuf:"bytes,5,opt,name=level,proto3" json:"level,omitempty"`
	LessonsCompleted       int32            `protobuf:"varint,6,opt,name=lessons_completed,json=lessonsCompleted,proto3" json:"lessons_completed,omitempty"`
	VocabularyLearned      int32            `protobuf:"varint,7,opt,name=vocabulary_learned,json=vocabularyLearned,proto3" json:"vocabulary_learned,omitempty"`
	GrammarRulesMastered   int32            `protobuf:"varint,8,opt,name=grammar_rules_mastered,json=grammarRulesMastered,proto3" json:"grammar_rules_mastered,omitempty"`
	ListeningComprehension int32            `protobuf:"varint,9,opt,name=listening_comprehension,json=listeningComprehension,proto3" json:"listening_comprehension,omitempty"`
	SpeakingFluency        int32            `protobuf:"varint,10,opt,name=speaking_fluency,json=speakingFluency,proto3" json:"speaking_fluency,omitempty"`
	DailyProgress          []*DailyProgress `protobuf:"bytes,11,rep,name=daily_progress,json=dailyProgress,proto3" json:"daily_progress,omitempty"`
	Achievements           []*Achievement   `protobuf:"bytes,12,rep,name=achievements,proto3" json:"achievements,omitempty"`
	Goals                  []*Goal          `protobuf:"bytes,13,rep,name=goals,proto3" json:"goals,omitempty"`
	LastUpdated            string           `protobuf:"bytes,14,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty"`
}

func (x *UserProgress) Reset() {
	*x = UserProgress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_progress_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserProgress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserProgress) ProtoMessage() {}

func (x *UserProgress) ProtoReflect() protoreflect.Message {
	mi := &file_user_progress_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserProgress.ProtoReflect.Descriptor instead.
func (*UserProgress) Descriptor() ([]byte, []int) {
	return file_user_progress_proto_rawDescGZIP(), []int{3}
}

func (x *UserProgress) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UserProgress) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UserProgress) GetLanguageCode() string {
	if x != nil {
		return x.LanguageCode
	}
	return ""
}

func (x *UserProgress) GetXp() int32 {
	if x != nil {
		return x.Xp
	}
	return 0
}

func (x *UserProgress) GetLevel() string {
	if x != nil {
		return x.Level
	}
	return ""
}

func (x *UserProgress) GetLessonsCompleted() int32 {
	if x != nil {
		return x.LessonsCompleted
	}
	return 0
}

func (x *UserProgress) GetVocabularyLearned() int32 {
	if x != nil {
		return x.VocabularyLearned
	}
	return 0
}

func (x *UserProgress) GetGrammarRulesMastered() int32 {
	if x != nil {
		return x.GrammarRulesMastered
	}
	return 0
}

func (x *UserProgress) GetListeningComprehension() int32 {
	if x != nil {
		return x.ListeningComprehension
	}
	return 0
}

func (x *UserProgress) GetSpeakingFluency() int32 {
	if x != nil {
		return x.SpeakingFluency
	}
	return 0
}

func (x *UserProgress) GetDailyProgress() []*DailyProgress {
	if x != nil {
		return x.DailyProgress
	}
	return nil
}

func (x *UserProgress) GetAchievements() []*Achievement {
	if x != nil {
		return x.Achievements
	}
	return nil
}

func (x *UserProgress) GetGoals() []*Goal {
	if x != nil {
		return x.Goals
	}
	return nil
}

func (x *UserProgress) GetLastUpdated() string {
	if x != nil {
		return x.LastUpdated
	}
	return ""
}

type UserProgressReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId                 string           `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	LanguageCode           string           `protobuf:"bytes,2,opt,name=language_code,json=languageCode,proto3" json:"language_code,omitempty"`
	Xp                     int32            `protobuf:"varint,3,opt,name=xp,proto3" json:"xp,omitempty"`
	Level                  string           `protobuf:"bytes,4,opt,name=level,proto3" json:"level,omitempty"`
	LessonsCompleted       int32            `protobuf:"varint,5,opt,name=lessons_completed,json=lessonsCompleted,proto3" json:"lessons_completed,omitempty"`
	VocabularyLearned      int32            `protobuf:"varint,6,opt,name=vocabulary_learned,json=vocabularyLearned,proto3" json:"vocabulary_learned,omitempty"`
	GrammarRulesMastered   int32            `protobuf:"varint,7,opt,name=grammar_rules_mastered,json=grammarRulesMastered,proto3" json:"grammar_rules_mastered,omitempty"`
	ListeningComprehension int32            `protobuf:"varint,8,opt,name=listening_comprehension,json=listeningComprehension,proto3" json:"listening_comprehension,omitempty"`
	SpeakingFluency        int32            `protobuf:"varint,9,opt,name=speaking_fluency,json=speakingFluency,proto3" json:"speaking_fluency,omitempty"`
	DailyProgress          []*DailyProgress `protobuf:"bytes,10,rep,name=daily_progress,json=dailyProgress,proto3" json:"daily_progress,omitempty"`
	Achievements           []*Achievement   `protobuf:"bytes,11,rep,name=achievements,proto3" json:"achievements,omitempty"`
	Goals                  []*Goal          `protobuf:"bytes,12,rep,name=goals,proto3" json:"goals,omitempty"`
	LastUpdated            string           `protobuf:"bytes,13,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty"`
}

func (x *UserProgressReq) Reset() {
	*x = UserProgressReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_progress_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserProgressReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserProgressReq) ProtoMessage() {}

func (x *UserProgressReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_progress_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserProgressReq.ProtoReflect.Descriptor instead.
func (*UserProgressReq) Descriptor() ([]byte, []int) {
	return file_user_progress_proto_rawDescGZIP(), []int{4}
}

func (x *UserProgressReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UserProgressReq) GetLanguageCode() string {
	if x != nil {
		return x.LanguageCode
	}
	return ""
}

func (x *UserProgressReq) GetXp() int32 {
	if x != nil {
		return x.Xp
	}
	return 0
}

func (x *UserProgressReq) GetLevel() string {
	if x != nil {
		return x.Level
	}
	return ""
}

func (x *UserProgressReq) GetLessonsCompleted() int32 {
	if x != nil {
		return x.LessonsCompleted
	}
	return 0
}

func (x *UserProgressReq) GetVocabularyLearned() int32 {
	if x != nil {
		return x.VocabularyLearned
	}
	return 0
}

func (x *UserProgressReq) GetGrammarRulesMastered() int32 {
	if x != nil {
		return x.GrammarRulesMastered
	}
	return 0
}

func (x *UserProgressReq) GetListeningComprehension() int32 {
	if x != nil {
		return x.ListeningComprehension
	}
	return 0
}

func (x *UserProgressReq) GetSpeakingFluency() int32 {
	if x != nil {
		return x.SpeakingFluency
	}
	return 0
}

func (x *UserProgressReq) GetDailyProgress() []*DailyProgress {
	if x != nil {
		return x.DailyProgress
	}
	return nil
}

func (x *UserProgressReq) GetAchievements() []*Achievement {
	if x != nil {
		return x.Achievements
	}
	return nil
}

func (x *UserProgressReq) GetGoals() []*Goal {
	if x != nil {
		return x.Goals
	}
	return nil
}

func (x *UserProgressReq) GetLastUpdated() string {
	if x != nil {
		return x.LastUpdated
	}
	return ""
}

type UserProgressFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId       string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	LanguageCode string `protobuf:"bytes,2,opt,name=language_code,json=languageCode,proto3" json:"language_code,omitempty"`
	Limit        int32  `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset       int32  `protobuf:"varint,4,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *UserProgressFilter) Reset() {
	*x = UserProgressFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_progress_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserProgressFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserProgressFilter) ProtoMessage() {}

func (x *UserProgressFilter) ProtoReflect() protoreflect.Message {
	mi := &file_user_progress_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserProgressFilter.ProtoReflect.Descriptor instead.
func (*UserProgressFilter) Descriptor() ([]byte, []int) {
	return file_user_progress_proto_rawDescGZIP(), []int{5}
}

func (x *UserProgressFilter) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UserProgressFilter) GetLanguageCode() string {
	if x != nil {
		return x.LanguageCode
	}
	return ""
}

func (x *UserProgressFilter) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *UserProgressFilter) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type AllUserProgress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserProgress []*UserProgress `protobuf:"bytes,1,rep,name=user_progress,json=userProgress,proto3" json:"user_progress,omitempty"`
	Count        int32           `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *AllUserProgress) Reset() {
	*x = AllUserProgress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_progress_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllUserProgress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllUserProgress) ProtoMessage() {}

func (x *AllUserProgress) ProtoReflect() protoreflect.Message {
	mi := &file_user_progress_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllUserProgress.ProtoReflect.Descriptor instead.
func (*AllUserProgress) Descriptor() ([]byte, []int) {
	return file_user_progress_proto_rawDescGZIP(), []int{6}
}

func (x *AllUserProgress) GetUserProgress() []*UserProgress {
	if x != nil {
		return x.UserProgress
	}
	return nil
}

func (x *AllUserProgress) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_user_progress_proto protoreflect.FileDescriptor

var file_user_progress_proto_rawDesc = []byte{
	0x0a, 0x13, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x67,
	0x72, 0x65, 0x73, 0x73, 0x1a, 0x0f, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc6, 0x01, 0x0a, 0x0d, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x50,
	0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x78,
	0x70, 0x5f, 0x65, 0x61, 0x72, 0x6e, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x78, 0x70, 0x45, 0x61, 0x72, 0x6e, 0x65, 0x64, 0x12, 0x2b, 0x0a, 0x11, 0x6c, 0x65, 0x73, 0x73,
	0x6f, 0x6e, 0x73, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x10, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x73, 0x43, 0x6f, 0x6d, 0x70,
	0x6c, 0x65, 0x74, 0x65, 0x64, 0x12, 0x2a, 0x0a, 0x11, 0x6e, 0x65, 0x77, 0x5f, 0x77, 0x6f, 0x72,
	0x64, 0x73, 0x5f, 0x6c, 0x65, 0x61, 0x72, 0x6e, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0f, 0x6e, 0x65, 0x77, 0x57, 0x6f, 0x72, 0x64, 0x73, 0x4c, 0x65, 0x61, 0x72, 0x6e, 0x65,
	0x64, 0x12, 0x2b, 0x0a, 0x11, 0x6d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x73, 0x5f, 0x70, 0x72, 0x61,
	0x63, 0x74, 0x69, 0x63, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x6d, 0x69,
	0x6e, 0x75, 0x74, 0x65, 0x73, 0x50, 0x72, 0x61, 0x63, 0x74, 0x69, 0x63, 0x65, 0x64, 0x22, 0x72,
	0x0a, 0x0b, 0x41, 0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x61, 0x72, 0x6e, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x61, 0x72, 0x6e, 0x65, 0x64,
	0x41, 0x74, 0x22, 0xb5, 0x01, 0x0a, 0x04, 0x47, 0x6f, 0x61, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x67,
	0x6f, 0x61, 0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x67, 0x6f, 0x61, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b,
	0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12,
	0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x22, 0xcb, 0x04, 0x0a, 0x0c, 0x55,
	0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65,
	0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6c, 0x61, 0x6e,
	0x67, 0x75, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x78, 0x70, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x78, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76,
	0x65, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12,
	0x2b, 0x0a, 0x11, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x73, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6c,
	0x65, 0x74, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x6c, 0x65, 0x73, 0x73,
	0x6f, 0x6e, 0x73, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x12, 0x2d, 0x0a, 0x12,
	0x76, 0x6f, 0x63, 0x61, 0x62, 0x75, 0x6c, 0x61, 0x72, 0x79, 0x5f, 0x6c, 0x65, 0x61, 0x72, 0x6e,
	0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x76, 0x6f, 0x63, 0x61, 0x62, 0x75,
	0x6c, 0x61, 0x72, 0x79, 0x4c, 0x65, 0x61, 0x72, 0x6e, 0x65, 0x64, 0x12, 0x34, 0x0a, 0x16, 0x67,
	0x72, 0x61, 0x6d, 0x6d, 0x61, 0x72, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x5f, 0x6d, 0x61, 0x73,
	0x74, 0x65, 0x72, 0x65, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x14, 0x67, 0x72, 0x61,
	0x6d, 0x6d, 0x61, 0x72, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x65,
	0x64, 0x12, 0x37, 0x0a, 0x17, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x63,
	0x6f, 0x6d, 0x70, 0x72, 0x65, 0x68, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x16, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6d,
	0x70, 0x72, 0x65, 0x68, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x29, 0x0a, 0x10, 0x73, 0x70,
	0x65, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x66, 0x6c, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x73, 0x70, 0x65, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x46, 0x6c,
	0x75, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x43, 0x0a, 0x0e, 0x64, 0x61, 0x69, 0x6c, 0x79, 0x5f, 0x70,
	0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x44, 0x61,
	0x69, 0x6c, 0x79, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x0d, 0x64, 0x61, 0x69,
	0x6c, 0x79, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x3e, 0x0a, 0x0c, 0x61, 0x63,
	0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73,
	0x2e, 0x41, 0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0c, 0x61, 0x63,
	0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x29, 0x0a, 0x05, 0x67, 0x6f,
	0x61, 0x6c, 0x73, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x47, 0x6f, 0x61, 0x6c, 0x52, 0x05,
	0x67, 0x6f, 0x61, 0x6c, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6c, 0x61, 0x73,
	0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x22, 0xbe, 0x04, 0x0a, 0x0f, 0x55, 0x73, 0x65,
	0x72, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67,
	0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6c, 0x61,
	0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x78, 0x70,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x78, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65,
	0x76, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c,
	0x12, 0x2b, 0x0a, 0x11, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x73, 0x5f, 0x63, 0x6f, 0x6d, 0x70,
	0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x6c, 0x65, 0x73,
	0x73, 0x6f, 0x6e, 0x73, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x12, 0x2d, 0x0a,
	0x12, 0x76, 0x6f, 0x63, 0x61, 0x62, 0x75, 0x6c, 0x61, 0x72, 0x79, 0x5f, 0x6c, 0x65, 0x61, 0x72,
	0x6e, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x76, 0x6f, 0x63, 0x61, 0x62,
	0x75, 0x6c, 0x61, 0x72, 0x79, 0x4c, 0x65, 0x61, 0x72, 0x6e, 0x65, 0x64, 0x12, 0x34, 0x0a, 0x16,
	0x67, 0x72, 0x61, 0x6d, 0x6d, 0x61, 0x72, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x5f, 0x6d, 0x61,
	0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x14, 0x67, 0x72,
	0x61, 0x6d, 0x6d, 0x61, 0x72, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72,
	0x65, 0x64, 0x12, 0x37, 0x0a, 0x17, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x69, 0x6e, 0x67, 0x5f,
	0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x68, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x16, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x69, 0x6e, 0x67, 0x43, 0x6f,
	0x6d, 0x70, 0x72, 0x65, 0x68, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x29, 0x0a, 0x10, 0x73,
	0x70, 0x65, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x66, 0x6c, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x73, 0x70, 0x65, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x46,
	0x6c, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x43, 0x0a, 0x0e, 0x64, 0x61, 0x69, 0x6c, 0x79, 0x5f,
	0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x44,
	0x61, 0x69, 0x6c, 0x79, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x0d, 0x64, 0x61,
	0x69, 0x6c, 0x79, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x3e, 0x0a, 0x0c, 0x61,
	0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73,
	0x73, 0x2e, 0x41, 0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0c, 0x61,
	0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x29, 0x0a, 0x05, 0x67,
	0x6f, 0x61, 0x6c, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x47, 0x6f, 0x61, 0x6c, 0x52,
	0x05, 0x67, 0x6f, 0x61, 0x6c, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6c, 0x61,
	0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x22, 0x80, 0x01, 0x0a, 0x12, 0x55, 0x73,
	0x65, 0x72, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x6c, 0x61, 0x6e,
	0x67, 0x75, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x69, 0x0a, 0x0f,
	0x41, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x40, 0x0a, 0x0d, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x70, 0x72,
	0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x67, 0x72,
	0x65, 0x73, 0x73, 0x52, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0xc4, 0x02, 0x0a, 0x13, 0x55, 0x73, 0x65, 0x72,
	0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x3b, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1e, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72,
	0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x0f, 0x2e, 0x65, 0x78, 0x65, 0x72,
	0x63, 0x69, 0x73, 0x65, 0x73, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x06,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x70, 0x72,
	0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x67, 0x72,
	0x65, 0x73, 0x73, 0x1a, 0x0f, 0x2e, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x73, 0x2e,
	0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x12, 0x2c, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x12, 0x0f, 0x2e, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x73, 0x2e, 0x42, 0x79, 0x49,
	0x64, 0x1a, 0x0f, 0x2e, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x73, 0x2e, 0x56, 0x6f,
	0x69, 0x64, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12,
	0x0f, 0x2e, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x73, 0x2e, 0x42, 0x79, 0x49, 0x64,
	0x1a, 0x1b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x22, 0x00, 0x12,
	0x4d, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x21, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72,
	0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x1a, 0x1e, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x41, 0x6c, 0x6c,
	0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x22, 0x00, 0x42, 0x14,
	0x5a, 0x12, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x67,
	0x72, 0x65, 0x73, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_progress_proto_rawDescOnce sync.Once
	file_user_progress_proto_rawDescData = file_user_progress_proto_rawDesc
)

func file_user_progress_proto_rawDescGZIP() []byte {
	file_user_progress_proto_rawDescOnce.Do(func() {
		file_user_progress_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_progress_proto_rawDescData)
	})
	return file_user_progress_proto_rawDescData
}

var file_user_progress_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_user_progress_proto_goTypes = []interface{}{
	(*DailyProgress)(nil),      // 0: user_progress.DailyProgress
	(*Achievement)(nil),        // 1: user_progress.Achievement
	(*Goal)(nil),               // 2: user_progress.Goal
	(*UserProgress)(nil),       // 3: user_progress.UserProgress
	(*UserProgressReq)(nil),    // 4: user_progress.UserProgressReq
	(*UserProgressFilter)(nil), // 5: user_progress.UserProgressFilter
	(*AllUserProgress)(nil),    // 6: user_progress.AllUserProgress
	(*ById)(nil),               // 7: exercises.ById
	(*Void)(nil),               // 8: exercises.Void
}
var file_user_progress_proto_depIdxs = []int32{
	0,  // 0: user_progress.UserProgress.daily_progress:type_name -> user_progress.DailyProgress
	1,  // 1: user_progress.UserProgress.achievements:type_name -> user_progress.Achievement
	2,  // 2: user_progress.UserProgress.goals:type_name -> user_progress.Goal
	0,  // 3: user_progress.UserProgressReq.daily_progress:type_name -> user_progress.DailyProgress
	1,  // 4: user_progress.UserProgressReq.achievements:type_name -> user_progress.Achievement
	2,  // 5: user_progress.UserProgressReq.goals:type_name -> user_progress.Goal
	3,  // 6: user_progress.AllUserProgress.user_progress:type_name -> user_progress.UserProgress
	4,  // 7: user_progress.UserProgressService.Create:input_type -> user_progress.UserProgressReq
	3,  // 8: user_progress.UserProgressService.Update:input_type -> user_progress.UserProgress
	7,  // 9: user_progress.UserProgressService.Delete:input_type -> exercises.ById
	7,  // 10: user_progress.UserProgressService.GetById:input_type -> exercises.ById
	5,  // 11: user_progress.UserProgressService.GetAll:input_type -> user_progress.UserProgressFilter
	8,  // 12: user_progress.UserProgressService.Create:output_type -> exercises.Void
	8,  // 13: user_progress.UserProgressService.Update:output_type -> exercises.Void
	8,  // 14: user_progress.UserProgressService.Delete:output_type -> exercises.Void
	3,  // 15: user_progress.UserProgressService.GetById:output_type -> user_progress.UserProgress
	6,  // 16: user_progress.UserProgressService.GetAll:output_type -> user_progress.AllUserProgress
	12, // [12:17] is the sub-list for method output_type
	7,  // [7:12] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_user_progress_proto_init() }
func file_user_progress_proto_init() {
	if File_user_progress_proto != nil {
		return
	}
	file_exercises_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_user_progress_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DailyProgress); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_progress_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Achievement); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_progress_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Goal); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_progress_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserProgress); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_progress_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserProgressReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_progress_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserProgressFilter); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_progress_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllUserProgress); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_user_progress_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_progress_proto_goTypes,
		DependencyIndexes: file_user_progress_proto_depIdxs,
		MessageInfos:      file_user_progress_proto_msgTypes,
	}.Build()
	File_user_progress_proto = out.File
	file_user_progress_proto_rawDesc = nil
	file_user_progress_proto_goTypes = nil
	file_user_progress_proto_depIdxs = nil
}
