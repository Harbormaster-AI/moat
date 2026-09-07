package model


//==============================================================
// DataSourceType Declaration
//==============================================================
type DataSourceType int
const (
    DataSourceTypeDatabase DataSourceType = iota
	DataSourceTypeFile
	DataSourceTypeStream
	DataSourceTypeAPI
	DataSourceTypeDataWarehouse
	DataSourceTypeDataLake
)


//==============================================================
// DataFormat Declaration
//==============================================================
type DataFormat int
const (
    DataFormatCSV DataFormat = iota
	DataFormatJSON
	DataFormatParquet
	DataFormatAvro
	DataFormatORC
	DataFormatXML
)


//==============================================================
// PipelineTriggerType Declaration
//==============================================================
type PipelineTriggerType int
const (
    PipelineTriggerTypeManual PipelineTriggerType = iota
	PipelineTriggerTypeSchedule
	PipelineTriggerTypeEvent
)


//==============================================================
// PipelineStatus Declaration
//==============================================================
type PipelineStatus int
const (
    PipelineStatusDraft PipelineStatus = iota
	PipelineStatusActive
	PipelineStatusPaused
	PipelineStatusFailed
	PipelineStatusSucceeded
)


//==============================================================
// DataTaskType Declaration
//==============================================================
type DataTaskType int
const (
    DataTaskTypeExtract DataTaskType = iota
	DataTaskTypeTransform
	DataTaskTypeLoad
	DataTaskTypeValidate
	DataTaskTypeEnrich
)


//==============================================================
// DimensionType Declaration
//==============================================================
type DimensionType int
const (
    DimensionTypeCategorical DimensionType = iota
	DimensionTypeTemporal
	DimensionTypeGeospatial
	DimensionTypeHierarchical
)


//==============================================================
// AggregationType Declaration
//==============================================================
type AggregationType int
const (
    AggregationTypeSum AggregationType = iota
	AggregationTypeAverage
	AggregationTypeMin
	AggregationTypeMax
	AggregationTypeMedian
	AggregationTypeCount
	AggregationTypeDistinctCount
)


//==============================================================
// MetricType Declaration
//==============================================================
type MetricType int
const (
    MetricTypeRatio MetricType = iota
	MetricTypeRate
	MetricTypeCount
	MetricTypePercentage
	MetricTypeIndex
	MetricTypeScore
)


//==============================================================
// ReportStatus Declaration
//==============================================================
type ReportStatus int
const (
    ReportStatusDraft ReportStatus = iota
	ReportStatusPublished
	ReportStatusArchived
)


//==============================================================
// DashboardStatus Declaration
//==============================================================
type DashboardStatus int
const (
    DashboardStatusDraft DashboardStatus = iota
	DashboardStatusLive
	DashboardStatusArchived
)


//==============================================================
// ChartType Declaration
//==============================================================
type ChartType int
const (
    ChartTypeTable ChartType = iota
	ChartTypeBar
	ChartTypeLine
	ChartTypeArea
	ChartTypePie
	ChartTypeScatter
	ChartTypeHeatmap
	ChartTypeKPI
)


//==============================================================
// NotebookLanguage Declaration
//==============================================================
type NotebookLanguage int
const (
    NotebookLanguagePython NotebookLanguage = iota
	NotebookLanguageR
	NotebookLanguageSQL
	NotebookLanguageJulia
)


//==============================================================
// ExperimentStatus Declaration
//==============================================================
type ExperimentStatus int
const (
    ExperimentStatusPlanned ExperimentStatus = iota
	ExperimentStatusRunning
	ExperimentStatusCompleted
	ExperimentStatusFailed
	ExperimentStatusStopped
)


//==============================================================
// TrainingStatus Declaration
//==============================================================
type TrainingStatus int
const (
    TrainingStatusQueued TrainingStatus = iota
	TrainingStatusRunning
	TrainingStatusCompleted
	TrainingStatusFailed
)


//==============================================================
// ModelType Declaration
//==============================================================
type ModelType int
const (
    ModelTypeClassification ModelType = iota
	ModelTypeRegression
	ModelTypeClustering
	ModelTypeForecasting
	ModelTypeRanking
	ModelTypeNLP
	ModelTypeComputerVision
)


//==============================================================
// ModelLifecycle Declaration
//==============================================================
type ModelLifecycle int
const (
    ModelLifecycleDraft ModelLifecycle = iota
	ModelLifecycleStaging
	ModelLifecycleProduction
	ModelLifecycleArchived
)


//==============================================================
// InferenceMode Declaration
//==============================================================
type InferenceMode int
const (
    InferenceModeBatch InferenceMode = iota
	InferenceModeRealTime
)


//==============================================================
// TimeGranularity Declaration
//==============================================================
type TimeGranularity int
const (
    TimeGranularityMinute TimeGranularity = iota
	TimeGranularityHour
	TimeGranularityDay
	TimeGranularityWeek
	TimeGranularityMonth
	TimeGranularityQuarter
	TimeGranularityYear
)


//==============================================================
// QualityDimension Declaration
//==============================================================
type QualityDimension int
const (
    QualityDimensionCompleteness QualityDimension = iota
	QualityDimensionAccuracy
	QualityDimensionConsistency
	QualityDimensionTimeliness
	QualityDimensionUniqueness
	QualityDimensionValidity
)


//==============================================================
// QualityStatus Declaration
//==============================================================
type QualityStatus int
const (
    QualityStatusPassed QualityStatus = iota
	QualityStatusFailed
	QualityStatusWarning
	QualityStatusSkipped
)


//==============================================================
// AnomalyType Declaration
//==============================================================
type AnomalyType int
const (
    AnomalyTypeSpike AnomalyType = iota
	AnomalyTypeDrop
	AnomalyTypeDrift
	AnomalyTypeSeasonal
	AnomalyTypeLevelShift
)


//==============================================================
// AlertSeverity Declaration
//==============================================================
type AlertSeverity int
const (
    AlertSeverityInfo AlertSeverity = iota
	AlertSeverityWarning
	AlertSeverityCritical
)


//==============================================================
// AlertStatus Declaration
//==============================================================
type AlertStatus int
const (
    AlertStatusOpen AlertStatus = iota
	AlertStatusAcknowledged
	AlertStatusResolved
	AlertStatusSuppressed
)


//==============================================================
// LineageNodeType Declaration
//==============================================================
type LineageNodeType int
const (
    LineageNodeTypeDataset LineageNodeType = iota
	LineageNodeTypePipeline
	LineageNodeTypeModel_
	LineageNodeTypeDashboard
	LineageNodeTypeReport
	LineageNodeTypeFeatureSet
	LineageNodeTypeNotebook
)


//==============================================================
// SubjectType Declaration
//==============================================================
type SubjectType int
const (
    SubjectTypeUser SubjectType = iota
	SubjectTypeGroup
	SubjectTypeService
)


//==============================================================
// AccessLevel Declaration
//==============================================================
type AccessLevel int
const (
    AccessLevelView AccessLevel = iota
	AccessLevelQuery
	AccessLevelModify
	AccessLevelAdmin
)


//==============================================================
// SQLDialect Declaration
//==============================================================
type SQLDialect int
const (
    SQLDialectANSI SQLDialect = iota
	SQLDialectPostgres
	SQLDialectMySQL
	SQLDialectSQLServer
	SQLDialectOracle
	SQLDialectSparkSQL
	SQLDialectBigQuery
)


//==============================================================
// RecommendationType Declaration
//==============================================================
type RecommendationType int
const (
    RecommendationTypePersonalized RecommendationType = iota
	RecommendationTypeTrending
	RecommendationTypeSimilarItems
	RecommendationTypeFrequentlyBoughtTogether
	RecommendationTypeContentBased
)


//==============================================================
// FraudDetectionType Declaration
//==============================================================
type FraudDetectionType int
const (
    FraudDetectionTypeRuleBased FraudDetectionType = iota
	FraudDetectionTypeSupervisedML
	FraudDetectionTypeUnsupervisedML
	FraudDetectionTypeHybrid
)


//==============================================================
// FraudSignalType Declaration
//==============================================================
type FraudSignalType int
const (
    FraudSignalTypeVelocity FraudSignalType = iota
	FraudSignalTypeGeolocationMismatch
	FraudSignalTypeAmountOutlier
	FraudSignalTypeDeviceFingerprint
	FraudSignalTypeBehavioralChange
)


//==============================================================
// FeatureStoreType Declaration
//==============================================================
type FeatureStoreType int
const (
    FeatureStoreTypeOnline FeatureStoreType = iota
	FeatureStoreTypeOffline
	FeatureStoreTypeHybrid
)


//==============================================================
// DataType Declaration
//==============================================================
type DataType int
const (
    DataTypeString DataType = iota
	DataTypeInteger
	DataTypeDecimal
	DataTypeBoolean
	DataTypeDate
	DataTypeDateTime
)


//==============================================================
// ComparisonOperator Declaration
//==============================================================
type ComparisonOperator int
const (
    ComparisonOperatorGreaterThan ComparisonOperator = iota
	ComparisonOperatorGreaterThanOrEqual
	ComparisonOperatorLessThan
	ComparisonOperatorLessThanOrEqual
	ComparisonOperatorEqual
	ComparisonOperatorNotEqual
)


//==============================================================
// NotificationChannel Declaration
//==============================================================
type NotificationChannel int
const (
    NotificationChannelEmail NotificationChannel = iota
	NotificationChannelSMS
	NotificationChannelWebhook
	NotificationChannelChat
)


//==============================================================
// TagCategory Declaration
//==============================================================
type TagCategory int
const (
    TagCategoryDomain TagCategory = iota
	TagCategorySensitivity
	TagCategoryPriority
	TagCategoryLifecycle
)


//==============================================================
// GovernanceTier Declaration
//==============================================================
type GovernanceTier int
const (
    GovernanceTierOpen GovernanceTier = iota
	GovernanceTierInternal
	GovernanceTierRestricted
	GovernanceTierConfidential
)

