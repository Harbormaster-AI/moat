
// enum type DataSourceType
export let DataSourceType = {
	Database:"Database",
	File:"File",
	Stream:"Stream",
	API:"API",
	DataWarehouse:"DataWarehouse",
	DataLake:"DataLake",
}

// enum type DataFormat
export let DataFormat = {
	CSV:"CSV",
	JSON:"JSON",
	Parquet:"Parquet",
	Avro:"Avro",
	ORC:"ORC",
	XML:"XML",
}

// enum type PipelineTriggerType
export let PipelineTriggerType = {
	Manual:"Manual",
	Schedule:"Schedule",
	Event:"Event",
}

// enum type PipelineStatus
export let PipelineStatus = {
	Draft:"Draft",
	Active:"Active",
	Paused:"Paused",
	Failed:"Failed",
	Succeeded:"Succeeded",
}

// enum type DataTaskType
export let DataTaskType = {
	Extract:"Extract",
	Transform:"Transform",
	Load:"Load",
	Validate:"Validate",
	Enrich:"Enrich",
}

// enum type DimensionType
export let DimensionType = {
	Categorical:"Categorical",
	Temporal:"Temporal",
	Geospatial:"Geospatial",
	Hierarchical:"Hierarchical",
}

// enum type AggregationType
export let AggregationType = {
	Sum:"Sum",
	Average:"Average",
	Min:"Min",
	Max:"Max",
	Median:"Median",
	Count:"Count",
	DistinctCount:"DistinctCount",
}

// enum type MetricType
export let MetricType = {
	Ratio:"Ratio",
	Rate:"Rate",
	Count:"Count",
	Percentage:"Percentage",
	Index:"Index",
	Score:"Score",
}

// enum type ReportStatus
export let ReportStatus = {
	Draft:"Draft",
	Published:"Published",
	Archived:"Archived",
}

// enum type DashboardStatus
export let DashboardStatus = {
	Draft:"Draft",
	Live:"Live",
	Archived:"Archived",
}

// enum type ChartType
export let ChartType = {
	Table:"Table",
	Bar:"Bar",
	Line:"Line",
	Area:"Area",
	Pie:"Pie",
	Scatter:"Scatter",
	Heatmap:"Heatmap",
	KPI:"KPI",
}

// enum type NotebookLanguage
export let NotebookLanguage = {
	Python:"Python",
	R:"R",
	SQL:"SQL",
	Julia:"Julia",
}

// enum type ExperimentStatus
export let ExperimentStatus = {
	Planned:"Planned",
	Running:"Running",
	Completed:"Completed",
	Failed:"Failed",
	Stopped:"Stopped",
}

// enum type TrainingStatus
export let TrainingStatus = {
	Queued:"Queued",
	Running:"Running",
	Completed:"Completed",
	Failed:"Failed",
}

// enum type ModelType
export let ModelType = {
	Classification:"Classification",
	Regression:"Regression",
	Clustering:"Clustering",
	Forecasting:"Forecasting",
	Ranking:"Ranking",
	NLP:"NLP",
	ComputerVision:"ComputerVision",
}

// enum type ModelLifecycle
export let ModelLifecycle = {
	Draft:"Draft",
	Staging:"Staging",
	Production:"Production",
	Archived:"Archived",
}

// enum type InferenceMode
export let InferenceMode = {
	Batch:"Batch",
	RealTime:"RealTime",
}

// enum type TimeGranularity
export let TimeGranularity = {
	Minute:"Minute",
	Hour:"Hour",
	Day:"Day",
	Week:"Week",
	Month:"Month",
	Quarter:"Quarter",
	Year:"Year",
}

// enum type QualityDimension
export let QualityDimension = {
	Completeness:"Completeness",
	Accuracy:"Accuracy",
	Consistency:"Consistency",
	Timeliness:"Timeliness",
	Uniqueness:"Uniqueness",
	Validity:"Validity",
}

// enum type QualityStatus
export let QualityStatus = {
	Passed:"Passed",
	Failed:"Failed",
	Warning:"Warning",
	Skipped:"Skipped",
}

// enum type AnomalyType
export let AnomalyType = {
	Spike:"Spike",
	Drop:"Drop",
	Drift:"Drift",
	Seasonal:"Seasonal",
	LevelShift:"LevelShift",
}

// enum type AlertSeverity
export let AlertSeverity = {
	Info:"Info",
	Warning:"Warning",
	Critical:"Critical",
}

// enum type AlertStatus
export let AlertStatus = {
	Open:"Open",
	Acknowledged:"Acknowledged",
	Resolved:"Resolved",
	Suppressed:"Suppressed",
}

// enum type LineageNodeType
export let LineageNodeType = {
	Dataset:"Dataset",
	Pipeline:"Pipeline",
	Model_:"Model_",
	Dashboard:"Dashboard",
	Report:"Report",
	FeatureSet:"FeatureSet",
	Notebook:"Notebook",
}

// enum type SubjectType
export let SubjectType = {
	User:"User",
	Group:"Group",
	Service:"Service",
}

// enum type AccessLevel
export let AccessLevel = {
	View:"View",
	Query:"Query",
	Modify:"Modify",
	Admin:"Admin",
}

// enum type SQLDialect
export let SQLDialect = {
	ANSI:"ANSI",
	Postgres:"Postgres",
	MySQL:"MySQL",
	SQLServer:"SQLServer",
	Oracle:"Oracle",
	SparkSQL:"SparkSQL",
	BigQuery:"BigQuery",
}

// enum type RecommendationType
export let RecommendationType = {
	Personalized:"Personalized",
	Trending:"Trending",
	SimilarItems:"SimilarItems",
	FrequentlyBoughtTogether:"FrequentlyBoughtTogether",
	ContentBased:"ContentBased",
}

// enum type FraudDetectionType
export let FraudDetectionType = {
	RuleBased:"RuleBased",
	SupervisedML:"SupervisedML",
	UnsupervisedML:"UnsupervisedML",
	Hybrid:"Hybrid",
}

// enum type FraudSignalType
export let FraudSignalType = {
	Velocity:"Velocity",
	GeolocationMismatch:"GeolocationMismatch",
	AmountOutlier:"AmountOutlier",
	DeviceFingerprint:"DeviceFingerprint",
	BehavioralChange:"BehavioralChange",
}

// enum type FeatureStoreType
export let FeatureStoreType = {
	Online:"Online",
	Offline:"Offline",
	Hybrid:"Hybrid",
}

// enum type DataType
export let DataType = {
	String:"String",
	Integer:"Integer",
	Decimal:"Decimal",
	Boolean:"Boolean",
	Date:"Date",
	DateTime:"DateTime",
}

// enum type ComparisonOperator
export let ComparisonOperator = {
	GreaterThan:"GreaterThan",
	GreaterThanOrEqual:"GreaterThanOrEqual",
	LessThan:"LessThan",
	LessThanOrEqual:"LessThanOrEqual",
	Equal:"Equal",
	NotEqual:"NotEqual",
}

// enum type NotificationChannel
export let NotificationChannel = {
	Email:"Email",
	SMS:"SMS",
	Webhook:"Webhook",
	Chat:"Chat",
}

// enum type TagCategory
export let TagCategory = {
	Domain:"Domain",
	Sensitivity:"Sensitivity",
	Priority:"Priority",
	Lifecycle:"Lifecycle",
}

// enum type GovernanceTier
export let GovernanceTier = {
	Open:"Open",
	Internal:"Internal",
	Restricted:"Restricted",
	Confidential:"Confidential",
}
