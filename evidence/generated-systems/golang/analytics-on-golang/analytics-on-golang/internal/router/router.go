package router

import (

    AnalyticsWorkspaceController "analytics-on-golang/internal/controller"
    DataSourceController "analytics-on-golang/internal/controller"
    DataSetController "analytics-on-golang/internal/controller"
    DataPipelineController "analytics-on-golang/internal/controller"
    DataTaskController "analytics-on-golang/internal/controller"
    SemanticModelController "analytics-on-golang/internal/controller"
    DimensionController "analytics-on-golang/internal/controller"
    MeasureController "analytics-on-golang/internal/controller"
    MetricController "analytics-on-golang/internal/controller"
    ReportController "analytics-on-golang/internal/controller"
    DashboardController "analytics-on-golang/internal/controller"
    VisualizationController "analytics-on-golang/internal/controller"
    NotebookController "analytics-on-golang/internal/controller"
    BIQueryController "analytics-on-golang/internal/controller"
    ExperimentController "analytics-on-golang/internal/controller"
    TrainingRunController "analytics-on-golang/internal/controller"
    RunMetricController "analytics-on-golang/internal/controller"
    RunParameterController "analytics-on-golang/internal/controller"
    Model_Controller "analytics-on-golang/internal/controller"
    ModelVersionController "analytics-on-golang/internal/controller"
    EvaluationMetricController "analytics-on-golang/internal/controller"
    FeatureSetController "analytics-on-golang/internal/controller"
    FeatureController "analytics-on-golang/internal/controller"
    InferenceEndpointController "analytics-on-golang/internal/controller"
    PredictionController "analytics-on-golang/internal/controller"
    ForecastController "analytics-on-golang/internal/controller"
    TimeSeriesController "analytics-on-golang/internal/controller"
    AnomalyController "analytics-on-golang/internal/controller"
    QualityRuleController "analytics-on-golang/internal/controller"
    QualityCheckController "analytics-on-golang/internal/controller"
    LineageNodeController "analytics-on-golang/internal/controller"
    TagController "analytics-on-golang/internal/controller"
    AccessPolicyController "analytics-on-golang/internal/controller"
    AlertController "analytics-on-golang/internal/controller"
    SubscriberController "analytics-on-golang/internal/controller"
    BusinessGlossaryTermController "analytics-on-golang/internal/controller"
    RecommendationScenarioController "analytics-on-golang/internal/controller"
    FraudScenarioController "analytics-on-golang/internal/controller"
    FraudSignalController "analytics-on-golang/internal/controller"
    jsonResponseFormatter "analytics-on-golang/internal/response"
    "github.com/gorilla/mux"

    PulseIndicatorController__ "analytics-on-golang/internal/controller"

)

// Router is exported and used in main.go
func Router() *mux.Router {

    router := mux.NewRouter()

    //----------------------------------------------------------------------------
    // default controllers for health and availability checking
    //----------------------------------------------------------------------------

    router.HandleFunc("/", jsonResponseFormatter.FormatToJSON(PulseIndicatorController__.Default__)).Methods("GET", "OPTIONS")
    router.HandleFunc("/health", jsonResponseFormatter.FormatToJSON(PulseIndicatorController__.Health__)).Methods("GET", "OPTIONS")


    //----------------------------------------------------------------------------
    // AnalyticsWorkspace Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AnalyticsWorkspace/{id}", jsonResponseFormatter.FormatToJSON(AnalyticsWorkspaceController.GetAnalyticsWorkspace)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/AnalyticsWorkspace", jsonResponseFormatter.FormatToJSON(AnalyticsWorkspaceController.GetAllAnalyticsWorkspace)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewAnalyticsWorkspace", jsonResponseFormatter.FormatToJSON(AnalyticsWorkspaceController.CreateAnalyticsWorkspace)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/AnalyticsWorkspace/{id}", jsonResponseFormatter.FormatToJSON(AnalyticsWorkspaceController.UpdateAnalyticsWorkspace)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteAnalyticsWorkspace/{id}", jsonResponseFormatter.FormatToJSON(AnalyticsWorkspaceController.DeleteAnalyticsWorkspace)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddDatasetsToAnalyticsWorkspace/{parentId}/datasetsId", jsonResponseFormatter.FormatToJSON(AnalyticsWorkspaceController.AddDatasetsToAnalyticsWorkspace)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveDatasetsFromAnalyticsWorkspace/{parentId}/datasetsIds", jsonResponseFormatter.FormatToJSON(AnalyticsWorkspaceController.RemoveDatasetsFromAnalyticsWorkspace)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddDataSourcesToAnalyticsWorkspace/{parentId}/dataSourcesId", jsonResponseFormatter.FormatToJSON(AnalyticsWorkspaceController.AddDataSourcesToAnalyticsWorkspace)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveDataSourcesFromAnalyticsWorkspace/{parentId}/dataSourcesIds", jsonResponseFormatter.FormatToJSON(AnalyticsWorkspaceController.RemoveDataSourcesFromAnalyticsWorkspace)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddPipelinesToAnalyticsWorkspace/{parentId}/pipelinesId", jsonResponseFormatter.FormatToJSON(AnalyticsWorkspaceController.AddPipelinesToAnalyticsWorkspace)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemovePipelinesFromAnalyticsWorkspace/{parentId}/pipelinesIds", jsonResponseFormatter.FormatToJSON(AnalyticsWorkspaceController.RemovePipelinesFromAnalyticsWorkspace)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddDashboardsToAnalyticsWorkspace/{parentId}/dashboardsId", jsonResponseFormatter.FormatToJSON(AnalyticsWorkspaceController.AddDashboardsToAnalyticsWorkspace)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveDashboardsFromAnalyticsWorkspace/{parentId}/dashboardsIds", jsonResponseFormatter.FormatToJSON(AnalyticsWorkspaceController.RemoveDashboardsFromAnalyticsWorkspace)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddReportsToAnalyticsWorkspace/{parentId}/reportsId", jsonResponseFormatter.FormatToJSON(AnalyticsWorkspaceController.AddReportsToAnalyticsWorkspace)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveReportsFromAnalyticsWorkspace/{parentId}/reportsIds", jsonResponseFormatter.FormatToJSON(AnalyticsWorkspaceController.RemoveReportsFromAnalyticsWorkspace)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddNotebooksToAnalyticsWorkspace/{parentId}/notebooksId", jsonResponseFormatter.FormatToJSON(AnalyticsWorkspaceController.AddNotebooksToAnalyticsWorkspace)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveNotebooksFromAnalyticsWorkspace/{parentId}/notebooksIds", jsonResponseFormatter.FormatToJSON(AnalyticsWorkspaceController.RemoveNotebooksFromAnalyticsWorkspace)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddModelsToAnalyticsWorkspace/{parentId}/modelsId", jsonResponseFormatter.FormatToJSON(AnalyticsWorkspaceController.AddModelsToAnalyticsWorkspace)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveModelsFromAnalyticsWorkspace/{parentId}/modelsIds", jsonResponseFormatter.FormatToJSON(AnalyticsWorkspaceController.RemoveModelsFromAnalyticsWorkspace)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddFeatureSetsToAnalyticsWorkspace/{parentId}/featureSetsId", jsonResponseFormatter.FormatToJSON(AnalyticsWorkspaceController.AddFeatureSetsToAnalyticsWorkspace)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveFeatureSetsFromAnalyticsWorkspace/{parentId}/featureSetsIds", jsonResponseFormatter.FormatToJSON(AnalyticsWorkspaceController.RemoveFeatureSetsFromAnalyticsWorkspace)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddPoliciesToAnalyticsWorkspace/{parentId}/policiesId", jsonResponseFormatter.FormatToJSON(AnalyticsWorkspaceController.AddPoliciesToAnalyticsWorkspace)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemovePoliciesFromAnalyticsWorkspace/{parentId}/policiesIds", jsonResponseFormatter.FormatToJSON(AnalyticsWorkspaceController.RemovePoliciesFromAnalyticsWorkspace)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddLineageNodesToAnalyticsWorkspace/{parentId}/lineageNodesId", jsonResponseFormatter.FormatToJSON(AnalyticsWorkspaceController.AddLineageNodesToAnalyticsWorkspace)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveLineageNodesFromAnalyticsWorkspace/{parentId}/lineageNodesIds", jsonResponseFormatter.FormatToJSON(AnalyticsWorkspaceController.RemoveLineageNodesFromAnalyticsWorkspace)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // DataSource Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/DataSource/{id}", jsonResponseFormatter.FormatToJSON(DataSourceController.GetDataSource)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/DataSource", jsonResponseFormatter.FormatToJSON(DataSourceController.GetAllDataSource)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewDataSource", jsonResponseFormatter.FormatToJSON(DataSourceController.CreateDataSource)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/DataSource/{id}", jsonResponseFormatter.FormatToJSON(DataSourceController.UpdateDataSource)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteDataSource/{id}", jsonResponseFormatter.FormatToJSON(DataSourceController.DeleteDataSource)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignWorkspaceToDataSource/{parentId}/workspaceId", jsonResponseFormatter.FormatToJSON(DataSourceController.AssignWorkspaceToDataSource)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignWorkspaceFromDataSource/{parentId}", jsonResponseFormatter.FormatToJSON(DataSourceController.UnassignWorkspaceFromDataSource)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddProducedDatasetsToDataSource/{parentId}/producedDatasetsId", jsonResponseFormatter.FormatToJSON(DataSourceController.AddProducedDatasetsToDataSource)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveProducedDatasetsFromDataSource/{parentId}/producedDatasetsIds", jsonResponseFormatter.FormatToJSON(DataSourceController.RemoveProducedDatasetsFromDataSource)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddPipelinesToDataSource/{parentId}/pipelinesId", jsonResponseFormatter.FormatToJSON(DataSourceController.AddPipelinesToDataSource)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemovePipelinesFromDataSource/{parentId}/pipelinesIds", jsonResponseFormatter.FormatToJSON(DataSourceController.RemovePipelinesFromDataSource)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // DataSet Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/DataSet/{id}", jsonResponseFormatter.FormatToJSON(DataSetController.GetDataSet)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/DataSet", jsonResponseFormatter.FormatToJSON(DataSetController.GetAllDataSet)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewDataSet", jsonResponseFormatter.FormatToJSON(DataSetController.CreateDataSet)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/DataSet/{id}", jsonResponseFormatter.FormatToJSON(DataSetController.UpdateDataSet)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteDataSet/{id}", jsonResponseFormatter.FormatToJSON(DataSetController.DeleteDataSet)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignWorkspaceToDataSet/{parentId}/workspaceId", jsonResponseFormatter.FormatToJSON(DataSetController.AssignWorkspaceToDataSet)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignWorkspaceFromDataSet/{parentId}", jsonResponseFormatter.FormatToJSON(DataSetController.UnassignWorkspaceFromDataSet)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignLineageNodeToDataSet/{parentId}/lineageNodeId", jsonResponseFormatter.FormatToJSON(DataSetController.AssignLineageNodeToDataSet)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignLineageNodeFromDataSet/{parentId}", jsonResponseFormatter.FormatToJSON(DataSetController.UnassignLineageNodeFromDataSet)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddSourcesToDataSet/{parentId}/sourcesId", jsonResponseFormatter.FormatToJSON(DataSetController.AddSourcesToDataSet)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveSourcesFromDataSet/{parentId}/sourcesIds", jsonResponseFormatter.FormatToJSON(DataSetController.RemoveSourcesFromDataSet)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddPipelinesToDataSet/{parentId}/pipelinesId", jsonResponseFormatter.FormatToJSON(DataSetController.AddPipelinesToDataSet)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemovePipelinesFromDataSet/{parentId}/pipelinesIds", jsonResponseFormatter.FormatToJSON(DataSetController.RemovePipelinesFromDataSet)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddSemanticModelsToDataSet/{parentId}/semanticModelsId", jsonResponseFormatter.FormatToJSON(DataSetController.AddSemanticModelsToDataSet)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveSemanticModelsFromDataSet/{parentId}/semanticModelsIds", jsonResponseFormatter.FormatToJSON(DataSetController.RemoveSemanticModelsFromDataSet)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddDimensionsToDataSet/{parentId}/dimensionsId", jsonResponseFormatter.FormatToJSON(DataSetController.AddDimensionsToDataSet)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveDimensionsFromDataSet/{parentId}/dimensionsIds", jsonResponseFormatter.FormatToJSON(DataSetController.RemoveDimensionsFromDataSet)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddMeasuresToDataSet/{parentId}/measuresId", jsonResponseFormatter.FormatToJSON(DataSetController.AddMeasuresToDataSet)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveMeasuresFromDataSet/{parentId}/measuresIds", jsonResponseFormatter.FormatToJSON(DataSetController.RemoveMeasuresFromDataSet)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddMetricsToDataSet/{parentId}/metricsId", jsonResponseFormatter.FormatToJSON(DataSetController.AddMetricsToDataSet)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveMetricsFromDataSet/{parentId}/metricsIds", jsonResponseFormatter.FormatToJSON(DataSetController.RemoveMetricsFromDataSet)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddQualityRulesToDataSet/{parentId}/qualityRulesId", jsonResponseFormatter.FormatToJSON(DataSetController.AddQualityRulesToDataSet)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveQualityRulesFromDataSet/{parentId}/qualityRulesIds", jsonResponseFormatter.FormatToJSON(DataSetController.RemoveQualityRulesFromDataSet)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddTagsToDataSet/{parentId}/tagsId", jsonResponseFormatter.FormatToJSON(DataSetController.AddTagsToDataSet)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveTagsFromDataSet/{parentId}/tagsIds", jsonResponseFormatter.FormatToJSON(DataSetController.RemoveTagsFromDataSet)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // DataPipeline Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/DataPipeline/{id}", jsonResponseFormatter.FormatToJSON(DataPipelineController.GetDataPipeline)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/DataPipeline", jsonResponseFormatter.FormatToJSON(DataPipelineController.GetAllDataPipeline)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewDataPipeline", jsonResponseFormatter.FormatToJSON(DataPipelineController.CreateDataPipeline)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/DataPipeline/{id}", jsonResponseFormatter.FormatToJSON(DataPipelineController.UpdateDataPipeline)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteDataPipeline/{id}", jsonResponseFormatter.FormatToJSON(DataPipelineController.DeleteDataPipeline)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignWorkspaceToDataPipeline/{parentId}/workspaceId", jsonResponseFormatter.FormatToJSON(DataPipelineController.AssignWorkspaceToDataPipeline)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignWorkspaceFromDataPipeline/{parentId}", jsonResponseFormatter.FormatToJSON(DataPipelineController.UnassignWorkspaceFromDataPipeline)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignLineageNodeToDataPipeline/{parentId}/lineageNodeId", jsonResponseFormatter.FormatToJSON(DataPipelineController.AssignLineageNodeToDataPipeline)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignLineageNodeFromDataPipeline/{parentId}", jsonResponseFormatter.FormatToJSON(DataPipelineController.UnassignLineageNodeFromDataPipeline)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddTasksToDataPipeline/{parentId}/tasksId", jsonResponseFormatter.FormatToJSON(DataPipelineController.AddTasksToDataPipeline)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveTasksFromDataPipeline/{parentId}/tasksIds", jsonResponseFormatter.FormatToJSON(DataPipelineController.RemoveTasksFromDataPipeline)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddSourcesToDataPipeline/{parentId}/sourcesId", jsonResponseFormatter.FormatToJSON(DataPipelineController.AddSourcesToDataPipeline)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveSourcesFromDataPipeline/{parentId}/sourcesIds", jsonResponseFormatter.FormatToJSON(DataPipelineController.RemoveSourcesFromDataPipeline)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddOutputsToDataPipeline/{parentId}/outputsId", jsonResponseFormatter.FormatToJSON(DataPipelineController.AddOutputsToDataPipeline)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveOutputsFromDataPipeline/{parentId}/outputsIds", jsonResponseFormatter.FormatToJSON(DataPipelineController.RemoveOutputsFromDataPipeline)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // DataTask Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/DataTask/{id}", jsonResponseFormatter.FormatToJSON(DataTaskController.GetDataTask)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/DataTask", jsonResponseFormatter.FormatToJSON(DataTaskController.GetAllDataTask)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewDataTask", jsonResponseFormatter.FormatToJSON(DataTaskController.CreateDataTask)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/DataTask/{id}", jsonResponseFormatter.FormatToJSON(DataTaskController.UpdateDataTask)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteDataTask/{id}", jsonResponseFormatter.FormatToJSON(DataTaskController.DeleteDataTask)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignPipelineToDataTask/{parentId}/pipelineId", jsonResponseFormatter.FormatToJSON(DataTaskController.AssignPipelineToDataTask)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignPipelineFromDataTask/{parentId}", jsonResponseFormatter.FormatToJSON(DataTaskController.UnassignPipelineFromDataTask)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddInputDatasetsToDataTask/{parentId}/inputDatasetsId", jsonResponseFormatter.FormatToJSON(DataTaskController.AddInputDatasetsToDataTask)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveInputDatasetsFromDataTask/{parentId}/inputDatasetsIds", jsonResponseFormatter.FormatToJSON(DataTaskController.RemoveInputDatasetsFromDataTask)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddOutputDatasetsToDataTask/{parentId}/outputDatasetsId", jsonResponseFormatter.FormatToJSON(DataTaskController.AddOutputDatasetsToDataTask)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveOutputDatasetsFromDataTask/{parentId}/outputDatasetsIds", jsonResponseFormatter.FormatToJSON(DataTaskController.RemoveOutputDatasetsFromDataTask)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // SemanticModel Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/SemanticModel/{id}", jsonResponseFormatter.FormatToJSON(SemanticModelController.GetSemanticModel)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/SemanticModel", jsonResponseFormatter.FormatToJSON(SemanticModelController.GetAllSemanticModel)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewSemanticModel", jsonResponseFormatter.FormatToJSON(SemanticModelController.CreateSemanticModel)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/SemanticModel/{id}", jsonResponseFormatter.FormatToJSON(SemanticModelController.UpdateSemanticModel)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteSemanticModel/{id}", jsonResponseFormatter.FormatToJSON(SemanticModelController.DeleteSemanticModel)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddDatasetsToSemanticModel/{parentId}/datasetsId", jsonResponseFormatter.FormatToJSON(SemanticModelController.AddDatasetsToSemanticModel)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveDatasetsFromSemanticModel/{parentId}/datasetsIds", jsonResponseFormatter.FormatToJSON(SemanticModelController.RemoveDatasetsFromSemanticModel)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddMetricsToSemanticModel/{parentId}/metricsId", jsonResponseFormatter.FormatToJSON(SemanticModelController.AddMetricsToSemanticModel)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveMetricsFromSemanticModel/{parentId}/metricsIds", jsonResponseFormatter.FormatToJSON(SemanticModelController.RemoveMetricsFromSemanticModel)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddDimensionsToSemanticModel/{parentId}/dimensionsId", jsonResponseFormatter.FormatToJSON(SemanticModelController.AddDimensionsToSemanticModel)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveDimensionsFromSemanticModel/{parentId}/dimensionsIds", jsonResponseFormatter.FormatToJSON(SemanticModelController.RemoveDimensionsFromSemanticModel)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddMeasuresToSemanticModel/{parentId}/measuresId", jsonResponseFormatter.FormatToJSON(SemanticModelController.AddMeasuresToSemanticModel)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveMeasuresFromSemanticModel/{parentId}/measuresIds", jsonResponseFormatter.FormatToJSON(SemanticModelController.RemoveMeasuresFromSemanticModel)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddGlossaryTermsToSemanticModel/{parentId}/glossaryTermsId", jsonResponseFormatter.FormatToJSON(SemanticModelController.AddGlossaryTermsToSemanticModel)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveGlossaryTermsFromSemanticModel/{parentId}/glossaryTermsIds", jsonResponseFormatter.FormatToJSON(SemanticModelController.RemoveGlossaryTermsFromSemanticModel)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Dimension Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Dimension/{id}", jsonResponseFormatter.FormatToJSON(DimensionController.GetDimension)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Dimension", jsonResponseFormatter.FormatToJSON(DimensionController.GetAllDimension)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewDimension", jsonResponseFormatter.FormatToJSON(DimensionController.CreateDimension)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Dimension/{id}", jsonResponseFormatter.FormatToJSON(DimensionController.UpdateDimension)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteDimension/{id}", jsonResponseFormatter.FormatToJSON(DimensionController.DeleteDimension)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignSemanticModelToDimension/{parentId}/semanticModelId", jsonResponseFormatter.FormatToJSON(DimensionController.AssignSemanticModelToDimension)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignSemanticModelFromDimension/{parentId}", jsonResponseFormatter.FormatToJSON(DimensionController.UnassignSemanticModelFromDimension)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddDatasetsToDimension/{parentId}/datasetsId", jsonResponseFormatter.FormatToJSON(DimensionController.AddDatasetsToDimension)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveDatasetsFromDimension/{parentId}/datasetsIds", jsonResponseFormatter.FormatToJSON(DimensionController.RemoveDatasetsFromDimension)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddGlossaryTermsToDimension/{parentId}/glossaryTermsId", jsonResponseFormatter.FormatToJSON(DimensionController.AddGlossaryTermsToDimension)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveGlossaryTermsFromDimension/{parentId}/glossaryTermsIds", jsonResponseFormatter.FormatToJSON(DimensionController.RemoveGlossaryTermsFromDimension)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Measure Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Measure/{id}", jsonResponseFormatter.FormatToJSON(MeasureController.GetMeasure)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Measure", jsonResponseFormatter.FormatToJSON(MeasureController.GetAllMeasure)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewMeasure", jsonResponseFormatter.FormatToJSON(MeasureController.CreateMeasure)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Measure/{id}", jsonResponseFormatter.FormatToJSON(MeasureController.UpdateMeasure)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteMeasure/{id}", jsonResponseFormatter.FormatToJSON(MeasureController.DeleteMeasure)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignSemanticModelToMeasure/{parentId}/semanticModelId", jsonResponseFormatter.FormatToJSON(MeasureController.AssignSemanticModelToMeasure)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignSemanticModelFromMeasure/{parentId}", jsonResponseFormatter.FormatToJSON(MeasureController.UnassignSemanticModelFromMeasure)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddDatasetsToMeasure/{parentId}/datasetsId", jsonResponseFormatter.FormatToJSON(MeasureController.AddDatasetsToMeasure)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveDatasetsFromMeasure/{parentId}/datasetsIds", jsonResponseFormatter.FormatToJSON(MeasureController.RemoveDatasetsFromMeasure)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddGlossaryTermsToMeasure/{parentId}/glossaryTermsId", jsonResponseFormatter.FormatToJSON(MeasureController.AddGlossaryTermsToMeasure)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveGlossaryTermsFromMeasure/{parentId}/glossaryTermsIds", jsonResponseFormatter.FormatToJSON(MeasureController.RemoveGlossaryTermsFromMeasure)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Metric Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Metric/{id}", jsonResponseFormatter.FormatToJSON(MetricController.GetMetric)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Metric", jsonResponseFormatter.FormatToJSON(MetricController.GetAllMetric)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewMetric", jsonResponseFormatter.FormatToJSON(MetricController.CreateMetric)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Metric/{id}", jsonResponseFormatter.FormatToJSON(MetricController.UpdateMetric)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteMetric/{id}", jsonResponseFormatter.FormatToJSON(MetricController.DeleteMetric)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignSemanticModelToMetric/{parentId}/semanticModelId", jsonResponseFormatter.FormatToJSON(MetricController.AssignSemanticModelToMetric)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignSemanticModelFromMetric/{parentId}", jsonResponseFormatter.FormatToJSON(MetricController.UnassignSemanticModelFromMetric)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddDatasetsToMetric/{parentId}/datasetsId", jsonResponseFormatter.FormatToJSON(MetricController.AddDatasetsToMetric)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveDatasetsFromMetric/{parentId}/datasetsIds", jsonResponseFormatter.FormatToJSON(MetricController.RemoveDatasetsFromMetric)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddGlossaryTermsToMetric/{parentId}/glossaryTermsId", jsonResponseFormatter.FormatToJSON(MetricController.AddGlossaryTermsToMetric)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveGlossaryTermsFromMetric/{parentId}/glossaryTermsIds", jsonResponseFormatter.FormatToJSON(MetricController.RemoveGlossaryTermsFromMetric)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddAlertsToMetric/{parentId}/alertsId", jsonResponseFormatter.FormatToJSON(MetricController.AddAlertsToMetric)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveAlertsFromMetric/{parentId}/alertsIds", jsonResponseFormatter.FormatToJSON(MetricController.RemoveAlertsFromMetric)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddVisualizationsToMetric/{parentId}/visualizationsId", jsonResponseFormatter.FormatToJSON(MetricController.AddVisualizationsToMetric)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveVisualizationsFromMetric/{parentId}/visualizationsIds", jsonResponseFormatter.FormatToJSON(MetricController.RemoveVisualizationsFromMetric)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Report Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Report/{id}", jsonResponseFormatter.FormatToJSON(ReportController.GetReport)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Report", jsonResponseFormatter.FormatToJSON(ReportController.GetAllReport)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewReport", jsonResponseFormatter.FormatToJSON(ReportController.CreateReport)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Report/{id}", jsonResponseFormatter.FormatToJSON(ReportController.UpdateReport)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteReport/{id}", jsonResponseFormatter.FormatToJSON(ReportController.DeleteReport)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignWorkspaceToReport/{parentId}/workspaceId", jsonResponseFormatter.FormatToJSON(ReportController.AssignWorkspaceToReport)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignWorkspaceFromReport/{parentId}", jsonResponseFormatter.FormatToJSON(ReportController.UnassignWorkspaceFromReport)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddVisualizationsToReport/{parentId}/visualizationsId", jsonResponseFormatter.FormatToJSON(ReportController.AddVisualizationsToReport)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveVisualizationsFromReport/{parentId}/visualizationsIds", jsonResponseFormatter.FormatToJSON(ReportController.RemoveVisualizationsFromReport)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddDatasetsToReport/{parentId}/datasetsId", jsonResponseFormatter.FormatToJSON(ReportController.AddDatasetsToReport)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveDatasetsFromReport/{parentId}/datasetsIds", jsonResponseFormatter.FormatToJSON(ReportController.RemoveDatasetsFromReport)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddSemanticModelsToReport/{parentId}/semanticModelsId", jsonResponseFormatter.FormatToJSON(ReportController.AddSemanticModelsToReport)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveSemanticModelsFromReport/{parentId}/semanticModelsIds", jsonResponseFormatter.FormatToJSON(ReportController.RemoveSemanticModelsFromReport)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddQueriesToReport/{parentId}/queriesId", jsonResponseFormatter.FormatToJSON(ReportController.AddQueriesToReport)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveQueriesFromReport/{parentId}/queriesIds", jsonResponseFormatter.FormatToJSON(ReportController.RemoveQueriesFromReport)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddTagsToReport/{parentId}/tagsId", jsonResponseFormatter.FormatToJSON(ReportController.AddTagsToReport)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveTagsFromReport/{parentId}/tagsIds", jsonResponseFormatter.FormatToJSON(ReportController.RemoveTagsFromReport)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Dashboard Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Dashboard/{id}", jsonResponseFormatter.FormatToJSON(DashboardController.GetDashboard)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Dashboard", jsonResponseFormatter.FormatToJSON(DashboardController.GetAllDashboard)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewDashboard", jsonResponseFormatter.FormatToJSON(DashboardController.CreateDashboard)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Dashboard/{id}", jsonResponseFormatter.FormatToJSON(DashboardController.UpdateDashboard)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteDashboard/{id}", jsonResponseFormatter.FormatToJSON(DashboardController.DeleteDashboard)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignWorkspaceToDashboard/{parentId}/workspaceId", jsonResponseFormatter.FormatToJSON(DashboardController.AssignWorkspaceToDashboard)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignWorkspaceFromDashboard/{parentId}", jsonResponseFormatter.FormatToJSON(DashboardController.UnassignWorkspaceFromDashboard)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddVisualizationsToDashboard/{parentId}/visualizationsId", jsonResponseFormatter.FormatToJSON(DashboardController.AddVisualizationsToDashboard)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveVisualizationsFromDashboard/{parentId}/visualizationsIds", jsonResponseFormatter.FormatToJSON(DashboardController.RemoveVisualizationsFromDashboard)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddReportsToDashboard/{parentId}/reportsId", jsonResponseFormatter.FormatToJSON(DashboardController.AddReportsToDashboard)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveReportsFromDashboard/{parentId}/reportsIds", jsonResponseFormatter.FormatToJSON(DashboardController.RemoveReportsFromDashboard)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddDatasetsToDashboard/{parentId}/datasetsId", jsonResponseFormatter.FormatToJSON(DashboardController.AddDatasetsToDashboard)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveDatasetsFromDashboard/{parentId}/datasetsIds", jsonResponseFormatter.FormatToJSON(DashboardController.RemoveDatasetsFromDashboard)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddAlertsToDashboard/{parentId}/alertsId", jsonResponseFormatter.FormatToJSON(DashboardController.AddAlertsToDashboard)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveAlertsFromDashboard/{parentId}/alertsIds", jsonResponseFormatter.FormatToJSON(DashboardController.RemoveAlertsFromDashboard)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddQueriesToDashboard/{parentId}/queriesId", jsonResponseFormatter.FormatToJSON(DashboardController.AddQueriesToDashboard)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveQueriesFromDashboard/{parentId}/queriesIds", jsonResponseFormatter.FormatToJSON(DashboardController.RemoveQueriesFromDashboard)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddTagsToDashboard/{parentId}/tagsId", jsonResponseFormatter.FormatToJSON(DashboardController.AddTagsToDashboard)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveTagsFromDashboard/{parentId}/tagsIds", jsonResponseFormatter.FormatToJSON(DashboardController.RemoveTagsFromDashboard)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Visualization Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Visualization/{id}", jsonResponseFormatter.FormatToJSON(VisualizationController.GetVisualization)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Visualization", jsonResponseFormatter.FormatToJSON(VisualizationController.GetAllVisualization)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewVisualization", jsonResponseFormatter.FormatToJSON(VisualizationController.CreateVisualization)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Visualization/{id}", jsonResponseFormatter.FormatToJSON(VisualizationController.UpdateVisualization)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteVisualization/{id}", jsonResponseFormatter.FormatToJSON(VisualizationController.DeleteVisualization)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignDashboardToVisualization/{parentId}/dashboardId", jsonResponseFormatter.FormatToJSON(VisualizationController.AssignDashboardToVisualization)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignDashboardFromVisualization/{parentId}", jsonResponseFormatter.FormatToJSON(VisualizationController.UnassignDashboardFromVisualization)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignReportToVisualization/{parentId}/reportId", jsonResponseFormatter.FormatToJSON(VisualizationController.AssignReportToVisualization)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignReportFromVisualization/{parentId}", jsonResponseFormatter.FormatToJSON(VisualizationController.UnassignReportFromVisualization)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddMetricsToVisualization/{parentId}/metricsId", jsonResponseFormatter.FormatToJSON(VisualizationController.AddMetricsToVisualization)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveMetricsFromVisualization/{parentId}/metricsIds", jsonResponseFormatter.FormatToJSON(VisualizationController.RemoveMetricsFromVisualization)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddDimensionsToVisualization/{parentId}/dimensionsId", jsonResponseFormatter.FormatToJSON(VisualizationController.AddDimensionsToVisualization)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveDimensionsFromVisualization/{parentId}/dimensionsIds", jsonResponseFormatter.FormatToJSON(VisualizationController.RemoveDimensionsFromVisualization)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddDatasetsToVisualization/{parentId}/datasetsId", jsonResponseFormatter.FormatToJSON(VisualizationController.AddDatasetsToVisualization)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveDatasetsFromVisualization/{parentId}/datasetsIds", jsonResponseFormatter.FormatToJSON(VisualizationController.RemoveDatasetsFromVisualization)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Notebook Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Notebook/{id}", jsonResponseFormatter.FormatToJSON(NotebookController.GetNotebook)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Notebook", jsonResponseFormatter.FormatToJSON(NotebookController.GetAllNotebook)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewNotebook", jsonResponseFormatter.FormatToJSON(NotebookController.CreateNotebook)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Notebook/{id}", jsonResponseFormatter.FormatToJSON(NotebookController.UpdateNotebook)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteNotebook/{id}", jsonResponseFormatter.FormatToJSON(NotebookController.DeleteNotebook)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignWorkspaceToNotebook/{parentId}/workspaceId", jsonResponseFormatter.FormatToJSON(NotebookController.AssignWorkspaceToNotebook)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignWorkspaceFromNotebook/{parentId}", jsonResponseFormatter.FormatToJSON(NotebookController.UnassignWorkspaceFromNotebook)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddDatasetsToNotebook/{parentId}/datasetsId", jsonResponseFormatter.FormatToJSON(NotebookController.AddDatasetsToNotebook)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveDatasetsFromNotebook/{parentId}/datasetsIds", jsonResponseFormatter.FormatToJSON(NotebookController.RemoveDatasetsFromNotebook)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddExperimentsToNotebook/{parentId}/experimentsId", jsonResponseFormatter.FormatToJSON(NotebookController.AddExperimentsToNotebook)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveExperimentsFromNotebook/{parentId}/experimentsIds", jsonResponseFormatter.FormatToJSON(NotebookController.RemoveExperimentsFromNotebook)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddQueriesToNotebook/{parentId}/queriesId", jsonResponseFormatter.FormatToJSON(NotebookController.AddQueriesToNotebook)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveQueriesFromNotebook/{parentId}/queriesIds", jsonResponseFormatter.FormatToJSON(NotebookController.RemoveQueriesFromNotebook)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // BIQuery Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/BIQuery/{id}", jsonResponseFormatter.FormatToJSON(BIQueryController.GetBIQuery)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/BIQuery", jsonResponseFormatter.FormatToJSON(BIQueryController.GetAllBIQuery)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewBIQuery", jsonResponseFormatter.FormatToJSON(BIQueryController.CreateBIQuery)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/BIQuery/{id}", jsonResponseFormatter.FormatToJSON(BIQueryController.UpdateBIQuery)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteBIQuery/{id}", jsonResponseFormatter.FormatToJSON(BIQueryController.DeleteBIQuery)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignWorkspaceToBIQuery/{parentId}/workspaceId", jsonResponseFormatter.FormatToJSON(BIQueryController.AssignWorkspaceToBIQuery)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignWorkspaceFromBIQuery/{parentId}", jsonResponseFormatter.FormatToJSON(BIQueryController.UnassignWorkspaceFromBIQuery)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddDatasetsToBIQuery/{parentId}/datasetsId", jsonResponseFormatter.FormatToJSON(BIQueryController.AddDatasetsToBIQuery)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveDatasetsFromBIQuery/{parentId}/datasetsIds", jsonResponseFormatter.FormatToJSON(BIQueryController.RemoveDatasetsFromBIQuery)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddReportsToBIQuery/{parentId}/reportsId", jsonResponseFormatter.FormatToJSON(BIQueryController.AddReportsToBIQuery)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveReportsFromBIQuery/{parentId}/reportsIds", jsonResponseFormatter.FormatToJSON(BIQueryController.RemoveReportsFromBIQuery)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddDashboardsToBIQuery/{parentId}/dashboardsId", jsonResponseFormatter.FormatToJSON(BIQueryController.AddDashboardsToBIQuery)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveDashboardsFromBIQuery/{parentId}/dashboardsIds", jsonResponseFormatter.FormatToJSON(BIQueryController.RemoveDashboardsFromBIQuery)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddNotebooksToBIQuery/{parentId}/notebooksId", jsonResponseFormatter.FormatToJSON(BIQueryController.AddNotebooksToBIQuery)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveNotebooksFromBIQuery/{parentId}/notebooksIds", jsonResponseFormatter.FormatToJSON(BIQueryController.RemoveNotebooksFromBIQuery)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Experiment Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Experiment/{id}", jsonResponseFormatter.FormatToJSON(ExperimentController.GetExperiment)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Experiment", jsonResponseFormatter.FormatToJSON(ExperimentController.GetAllExperiment)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewExperiment", jsonResponseFormatter.FormatToJSON(ExperimentController.CreateExperiment)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Experiment/{id}", jsonResponseFormatter.FormatToJSON(ExperimentController.UpdateExperiment)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteExperiment/{id}", jsonResponseFormatter.FormatToJSON(ExperimentController.DeleteExperiment)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignWorkspaceToExperiment/{parentId}/workspaceId", jsonResponseFormatter.FormatToJSON(ExperimentController.AssignWorkspaceToExperiment)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignWorkspaceFromExperiment/{parentId}", jsonResponseFormatter.FormatToJSON(ExperimentController.UnassignWorkspaceFromExperiment)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddTrainingRunsToExperiment/{parentId}/trainingRunsId", jsonResponseFormatter.FormatToJSON(ExperimentController.AddTrainingRunsToExperiment)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveTrainingRunsFromExperiment/{parentId}/trainingRunsIds", jsonResponseFormatter.FormatToJSON(ExperimentController.RemoveTrainingRunsFromExperiment)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddModelsToExperiment/{parentId}/modelsId", jsonResponseFormatter.FormatToJSON(ExperimentController.AddModelsToExperiment)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveModelsFromExperiment/{parentId}/modelsIds", jsonResponseFormatter.FormatToJSON(ExperimentController.RemoveModelsFromExperiment)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddNotebooksToExperiment/{parentId}/notebooksId", jsonResponseFormatter.FormatToJSON(ExperimentController.AddNotebooksToExperiment)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveNotebooksFromExperiment/{parentId}/notebooksIds", jsonResponseFormatter.FormatToJSON(ExperimentController.RemoveNotebooksFromExperiment)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // TrainingRun Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/TrainingRun/{id}", jsonResponseFormatter.FormatToJSON(TrainingRunController.GetTrainingRun)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/TrainingRun", jsonResponseFormatter.FormatToJSON(TrainingRunController.GetAllTrainingRun)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewTrainingRun", jsonResponseFormatter.FormatToJSON(TrainingRunController.CreateTrainingRun)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/TrainingRun/{id}", jsonResponseFormatter.FormatToJSON(TrainingRunController.UpdateTrainingRun)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteTrainingRun/{id}", jsonResponseFormatter.FormatToJSON(TrainingRunController.DeleteTrainingRun)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignExperimentToTrainingRun/{parentId}/experimentId", jsonResponseFormatter.FormatToJSON(TrainingRunController.AssignExperimentToTrainingRun)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignExperimentFromTrainingRun/{parentId}", jsonResponseFormatter.FormatToJSON(TrainingRunController.UnassignExperimentFromTrainingRun)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignModelVersionToTrainingRun/{parentId}/modelVersionId", jsonResponseFormatter.FormatToJSON(TrainingRunController.AssignModelVersionToTrainingRun)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignModelVersionFromTrainingRun/{parentId}", jsonResponseFormatter.FormatToJSON(TrainingRunController.UnassignModelVersionFromTrainingRun)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddInputDatasetsToTrainingRun/{parentId}/inputDatasetsId", jsonResponseFormatter.FormatToJSON(TrainingRunController.AddInputDatasetsToTrainingRun)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveInputDatasetsFromTrainingRun/{parentId}/inputDatasetsIds", jsonResponseFormatter.FormatToJSON(TrainingRunController.RemoveInputDatasetsFromTrainingRun)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddFeaturesToTrainingRun/{parentId}/featuresId", jsonResponseFormatter.FormatToJSON(TrainingRunController.AddFeaturesToTrainingRun)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveFeaturesFromTrainingRun/{parentId}/featuresIds", jsonResponseFormatter.FormatToJSON(TrainingRunController.RemoveFeaturesFromTrainingRun)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddRunMetricsToTrainingRun/{parentId}/runMetricsId", jsonResponseFormatter.FormatToJSON(TrainingRunController.AddRunMetricsToTrainingRun)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveRunMetricsFromTrainingRun/{parentId}/runMetricsIds", jsonResponseFormatter.FormatToJSON(TrainingRunController.RemoveRunMetricsFromTrainingRun)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddRunParametersToTrainingRun/{parentId}/runParametersId", jsonResponseFormatter.FormatToJSON(TrainingRunController.AddRunParametersToTrainingRun)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveRunParametersFromTrainingRun/{parentId}/runParametersIds", jsonResponseFormatter.FormatToJSON(TrainingRunController.RemoveRunParametersFromTrainingRun)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // RunMetric Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/RunMetric/{id}", jsonResponseFormatter.FormatToJSON(RunMetricController.GetRunMetric)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/RunMetric", jsonResponseFormatter.FormatToJSON(RunMetricController.GetAllRunMetric)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewRunMetric", jsonResponseFormatter.FormatToJSON(RunMetricController.CreateRunMetric)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/RunMetric/{id}", jsonResponseFormatter.FormatToJSON(RunMetricController.UpdateRunMetric)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteRunMetric/{id}", jsonResponseFormatter.FormatToJSON(RunMetricController.DeleteRunMetric)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignTrainingRunToRunMetric/{parentId}/trainingRunId", jsonResponseFormatter.FormatToJSON(RunMetricController.AssignTrainingRunToRunMetric)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignTrainingRunFromRunMetric/{parentId}", jsonResponseFormatter.FormatToJSON(RunMetricController.UnassignTrainingRunFromRunMetric)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignMetricToRunMetric/{parentId}/metricId", jsonResponseFormatter.FormatToJSON(RunMetricController.AssignMetricToRunMetric)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignMetricFromRunMetric/{parentId}", jsonResponseFormatter.FormatToJSON(RunMetricController.UnassignMetricFromRunMetric)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignDatasetToRunMetric/{parentId}/datasetId", jsonResponseFormatter.FormatToJSON(RunMetricController.AssignDatasetToRunMetric)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignDatasetFromRunMetric/{parentId}", jsonResponseFormatter.FormatToJSON(RunMetricController.UnassignDatasetFromRunMetric)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // RunParameter Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/RunParameter/{id}", jsonResponseFormatter.FormatToJSON(RunParameterController.GetRunParameter)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/RunParameter", jsonResponseFormatter.FormatToJSON(RunParameterController.GetAllRunParameter)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewRunParameter", jsonResponseFormatter.FormatToJSON(RunParameterController.CreateRunParameter)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/RunParameter/{id}", jsonResponseFormatter.FormatToJSON(RunParameterController.UpdateRunParameter)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteRunParameter/{id}", jsonResponseFormatter.FormatToJSON(RunParameterController.DeleteRunParameter)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignTrainingRunToRunParameter/{parentId}/trainingRunId", jsonResponseFormatter.FormatToJSON(RunParameterController.AssignTrainingRunToRunParameter)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignTrainingRunFromRunParameter/{parentId}", jsonResponseFormatter.FormatToJSON(RunParameterController.UnassignTrainingRunFromRunParameter)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Model_ Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Model_/{id}", jsonResponseFormatter.FormatToJSON(Model_Controller.GetModel_)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Model_", jsonResponseFormatter.FormatToJSON(Model_Controller.GetAllModel_)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewModel_", jsonResponseFormatter.FormatToJSON(Model_Controller.CreateModel_)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Model_/{id}", jsonResponseFormatter.FormatToJSON(Model_Controller.UpdateModel_)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteModel_/{id}", jsonResponseFormatter.FormatToJSON(Model_Controller.DeleteModel_)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignWorkspaceToModel_/{parentId}/workspaceId", jsonResponseFormatter.FormatToJSON(Model_Controller.AssignWorkspaceToModel_)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignWorkspaceFromModel_/{parentId}", jsonResponseFormatter.FormatToJSON(Model_Controller.UnassignWorkspaceFromModel_)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddVersionsToModel_/{parentId}/versionsId", jsonResponseFormatter.FormatToJSON(Model_Controller.AddVersionsToModel_)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveVersionsFromModel_/{parentId}/versionsIds", jsonResponseFormatter.FormatToJSON(Model_Controller.RemoveVersionsFromModel_)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddFeatureSetsToModel_/{parentId}/featureSetsId", jsonResponseFormatter.FormatToJSON(Model_Controller.AddFeatureSetsToModel_)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveFeatureSetsFromModel_/{parentId}/featureSetsIds", jsonResponseFormatter.FormatToJSON(Model_Controller.RemoveFeatureSetsFromModel_)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddExperimentsToModel_/{parentId}/experimentsId", jsonResponseFormatter.FormatToJSON(Model_Controller.AddExperimentsToModel_)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveExperimentsFromModel_/{parentId}/experimentsIds", jsonResponseFormatter.FormatToJSON(Model_Controller.RemoveExperimentsFromModel_)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddTagsToModel_/{parentId}/tagsId", jsonResponseFormatter.FormatToJSON(Model_Controller.AddTagsToModel_)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveTagsFromModel_/{parentId}/tagsIds", jsonResponseFormatter.FormatToJSON(Model_Controller.RemoveTagsFromModel_)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // ModelVersion Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/ModelVersion/{id}", jsonResponseFormatter.FormatToJSON(ModelVersionController.GetModelVersion)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/ModelVersion", jsonResponseFormatter.FormatToJSON(ModelVersionController.GetAllModelVersion)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewModelVersion", jsonResponseFormatter.FormatToJSON(ModelVersionController.CreateModelVersion)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/ModelVersion/{id}", jsonResponseFormatter.FormatToJSON(ModelVersionController.UpdateModelVersion)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteModelVersion/{id}", jsonResponseFormatter.FormatToJSON(ModelVersionController.DeleteModelVersion)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignModel_ToModelVersion/{parentId}/model_Id", jsonResponseFormatter.FormatToJSON(ModelVersionController.AssignModel_ToModelVersion)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignModel_FromModelVersion/{parentId}", jsonResponseFormatter.FormatToJSON(ModelVersionController.UnassignModel_FromModelVersion)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignTrainingRunToModelVersion/{parentId}/trainingRunId", jsonResponseFormatter.FormatToJSON(ModelVersionController.AssignTrainingRunToModelVersion)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignTrainingRunFromModelVersion/{parentId}", jsonResponseFormatter.FormatToJSON(ModelVersionController.UnassignTrainingRunFromModelVersion)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddEvaluationMetricsToModelVersion/{parentId}/evaluationMetricsId", jsonResponseFormatter.FormatToJSON(ModelVersionController.AddEvaluationMetricsToModelVersion)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveEvaluationMetricsFromModelVersion/{parentId}/evaluationMetricsIds", jsonResponseFormatter.FormatToJSON(ModelVersionController.RemoveEvaluationMetricsFromModelVersion)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddDeploymentsToModelVersion/{parentId}/deploymentsId", jsonResponseFormatter.FormatToJSON(ModelVersionController.AddDeploymentsToModelVersion)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveDeploymentsFromModelVersion/{parentId}/deploymentsIds", jsonResponseFormatter.FormatToJSON(ModelVersionController.RemoveDeploymentsFromModelVersion)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddFeatureSetsToModelVersion/{parentId}/featureSetsId", jsonResponseFormatter.FormatToJSON(ModelVersionController.AddFeatureSetsToModelVersion)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveFeatureSetsFromModelVersion/{parentId}/featureSetsIds", jsonResponseFormatter.FormatToJSON(ModelVersionController.RemoveFeatureSetsFromModelVersion)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddDatasetsToModelVersion/{parentId}/datasetsId", jsonResponseFormatter.FormatToJSON(ModelVersionController.AddDatasetsToModelVersion)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveDatasetsFromModelVersion/{parentId}/datasetsIds", jsonResponseFormatter.FormatToJSON(ModelVersionController.RemoveDatasetsFromModelVersion)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // EvaluationMetric Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/EvaluationMetric/{id}", jsonResponseFormatter.FormatToJSON(EvaluationMetricController.GetEvaluationMetric)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/EvaluationMetric", jsonResponseFormatter.FormatToJSON(EvaluationMetricController.GetAllEvaluationMetric)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewEvaluationMetric", jsonResponseFormatter.FormatToJSON(EvaluationMetricController.CreateEvaluationMetric)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/EvaluationMetric/{id}", jsonResponseFormatter.FormatToJSON(EvaluationMetricController.UpdateEvaluationMetric)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteEvaluationMetric/{id}", jsonResponseFormatter.FormatToJSON(EvaluationMetricController.DeleteEvaluationMetric)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignModelVersionToEvaluationMetric/{parentId}/modelVersionId", jsonResponseFormatter.FormatToJSON(EvaluationMetricController.AssignModelVersionToEvaluationMetric)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignModelVersionFromEvaluationMetric/{parentId}", jsonResponseFormatter.FormatToJSON(EvaluationMetricController.UnassignModelVersionFromEvaluationMetric)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignMetricToEvaluationMetric/{parentId}/metricId", jsonResponseFormatter.FormatToJSON(EvaluationMetricController.AssignMetricToEvaluationMetric)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignMetricFromEvaluationMetric/{parentId}", jsonResponseFormatter.FormatToJSON(EvaluationMetricController.UnassignMetricFromEvaluationMetric)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignDatasetToEvaluationMetric/{parentId}/datasetId", jsonResponseFormatter.FormatToJSON(EvaluationMetricController.AssignDatasetToEvaluationMetric)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignDatasetFromEvaluationMetric/{parentId}", jsonResponseFormatter.FormatToJSON(EvaluationMetricController.UnassignDatasetFromEvaluationMetric)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // FeatureSet Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/FeatureSet/{id}", jsonResponseFormatter.FormatToJSON(FeatureSetController.GetFeatureSet)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/FeatureSet", jsonResponseFormatter.FormatToJSON(FeatureSetController.GetAllFeatureSet)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewFeatureSet", jsonResponseFormatter.FormatToJSON(FeatureSetController.CreateFeatureSet)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/FeatureSet/{id}", jsonResponseFormatter.FormatToJSON(FeatureSetController.UpdateFeatureSet)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteFeatureSet/{id}", jsonResponseFormatter.FormatToJSON(FeatureSetController.DeleteFeatureSet)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignWorkspaceToFeatureSet/{parentId}/workspaceId", jsonResponseFormatter.FormatToJSON(FeatureSetController.AssignWorkspaceToFeatureSet)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignWorkspaceFromFeatureSet/{parentId}", jsonResponseFormatter.FormatToJSON(FeatureSetController.UnassignWorkspaceFromFeatureSet)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddFeaturesToFeatureSet/{parentId}/featuresId", jsonResponseFormatter.FormatToJSON(FeatureSetController.AddFeaturesToFeatureSet)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveFeaturesFromFeatureSet/{parentId}/featuresIds", jsonResponseFormatter.FormatToJSON(FeatureSetController.RemoveFeaturesFromFeatureSet)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddDatasetsToFeatureSet/{parentId}/datasetsId", jsonResponseFormatter.FormatToJSON(FeatureSetController.AddDatasetsToFeatureSet)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveDatasetsFromFeatureSet/{parentId}/datasetsIds", jsonResponseFormatter.FormatToJSON(FeatureSetController.RemoveDatasetsFromFeatureSet)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddModelsToFeatureSet/{parentId}/modelsId", jsonResponseFormatter.FormatToJSON(FeatureSetController.AddModelsToFeatureSet)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveModelsFromFeatureSet/{parentId}/modelsIds", jsonResponseFormatter.FormatToJSON(FeatureSetController.RemoveModelsFromFeatureSet)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddModelVersionsToFeatureSet/{parentId}/modelVersionsId", jsonResponseFormatter.FormatToJSON(FeatureSetController.AddModelVersionsToFeatureSet)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveModelVersionsFromFeatureSet/{parentId}/modelVersionsIds", jsonResponseFormatter.FormatToJSON(FeatureSetController.RemoveModelVersionsFromFeatureSet)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddTagsToFeatureSet/{parentId}/tagsId", jsonResponseFormatter.FormatToJSON(FeatureSetController.AddTagsToFeatureSet)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveTagsFromFeatureSet/{parentId}/tagsIds", jsonResponseFormatter.FormatToJSON(FeatureSetController.RemoveTagsFromFeatureSet)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Feature Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Feature/{id}", jsonResponseFormatter.FormatToJSON(FeatureController.GetFeature)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Feature", jsonResponseFormatter.FormatToJSON(FeatureController.GetAllFeature)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewFeature", jsonResponseFormatter.FormatToJSON(FeatureController.CreateFeature)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Feature/{id}", jsonResponseFormatter.FormatToJSON(FeatureController.UpdateFeature)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteFeature/{id}", jsonResponseFormatter.FormatToJSON(FeatureController.DeleteFeature)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignFeatureSetToFeature/{parentId}/featureSetId", jsonResponseFormatter.FormatToJSON(FeatureController.AssignFeatureSetToFeature)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignFeatureSetFromFeature/{parentId}", jsonResponseFormatter.FormatToJSON(FeatureController.UnassignFeatureSetFromFeature)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddSourceDatasetsToFeature/{parentId}/sourceDatasetsId", jsonResponseFormatter.FormatToJSON(FeatureController.AddSourceDatasetsToFeature)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveSourceDatasetsFromFeature/{parentId}/sourceDatasetsIds", jsonResponseFormatter.FormatToJSON(FeatureController.RemoveSourceDatasetsFromFeature)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddModelsToFeature/{parentId}/modelsId", jsonResponseFormatter.FormatToJSON(FeatureController.AddModelsToFeature)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveModelsFromFeature/{parentId}/modelsIds", jsonResponseFormatter.FormatToJSON(FeatureController.RemoveModelsFromFeature)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddTrainingRunsToFeature/{parentId}/trainingRunsId", jsonResponseFormatter.FormatToJSON(FeatureController.AddTrainingRunsToFeature)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveTrainingRunsFromFeature/{parentId}/trainingRunsIds", jsonResponseFormatter.FormatToJSON(FeatureController.RemoveTrainingRunsFromFeature)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // InferenceEndpoint Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/InferenceEndpoint/{id}", jsonResponseFormatter.FormatToJSON(InferenceEndpointController.GetInferenceEndpoint)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/InferenceEndpoint", jsonResponseFormatter.FormatToJSON(InferenceEndpointController.GetAllInferenceEndpoint)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewInferenceEndpoint", jsonResponseFormatter.FormatToJSON(InferenceEndpointController.CreateInferenceEndpoint)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/InferenceEndpoint/{id}", jsonResponseFormatter.FormatToJSON(InferenceEndpointController.UpdateInferenceEndpoint)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteInferenceEndpoint/{id}", jsonResponseFormatter.FormatToJSON(InferenceEndpointController.DeleteInferenceEndpoint)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignModelVersionToInferenceEndpoint/{parentId}/modelVersionId", jsonResponseFormatter.FormatToJSON(InferenceEndpointController.AssignModelVersionToInferenceEndpoint)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignModelVersionFromInferenceEndpoint/{parentId}", jsonResponseFormatter.FormatToJSON(InferenceEndpointController.UnassignModelVersionFromInferenceEndpoint)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignWorkspaceToInferenceEndpoint/{parentId}/workspaceId", jsonResponseFormatter.FormatToJSON(InferenceEndpointController.AssignWorkspaceToInferenceEndpoint)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignWorkspaceFromInferenceEndpoint/{parentId}", jsonResponseFormatter.FormatToJSON(InferenceEndpointController.UnassignWorkspaceFromInferenceEndpoint)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddPredictionsToInferenceEndpoint/{parentId}/predictionsId", jsonResponseFormatter.FormatToJSON(InferenceEndpointController.AddPredictionsToInferenceEndpoint)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemovePredictionsFromInferenceEndpoint/{parentId}/predictionsIds", jsonResponseFormatter.FormatToJSON(InferenceEndpointController.RemovePredictionsFromInferenceEndpoint)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Prediction Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Prediction/{id}", jsonResponseFormatter.FormatToJSON(PredictionController.GetPrediction)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Prediction", jsonResponseFormatter.FormatToJSON(PredictionController.GetAllPrediction)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewPrediction", jsonResponseFormatter.FormatToJSON(PredictionController.CreatePrediction)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Prediction/{id}", jsonResponseFormatter.FormatToJSON(PredictionController.UpdatePrediction)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeletePrediction/{id}", jsonResponseFormatter.FormatToJSON(PredictionController.DeletePrediction)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignEndpointToPrediction/{parentId}/endpointId", jsonResponseFormatter.FormatToJSON(PredictionController.AssignEndpointToPrediction)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignEndpointFromPrediction/{parentId}", jsonResponseFormatter.FormatToJSON(PredictionController.UnassignEndpointFromPrediction)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignModelVersionToPrediction/{parentId}/modelVersionId", jsonResponseFormatter.FormatToJSON(PredictionController.AssignModelVersionToPrediction)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignModelVersionFromPrediction/{parentId}", jsonResponseFormatter.FormatToJSON(PredictionController.UnassignModelVersionFromPrediction)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignDatasetToPrediction/{parentId}/datasetId", jsonResponseFormatter.FormatToJSON(PredictionController.AssignDatasetToPrediction)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignDatasetFromPrediction/{parentId}", jsonResponseFormatter.FormatToJSON(PredictionController.UnassignDatasetFromPrediction)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Forecast Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Forecast/{id}", jsonResponseFormatter.FormatToJSON(ForecastController.GetForecast)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Forecast", jsonResponseFormatter.FormatToJSON(ForecastController.GetAllForecast)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewForecast", jsonResponseFormatter.FormatToJSON(ForecastController.CreateForecast)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Forecast/{id}", jsonResponseFormatter.FormatToJSON(ForecastController.UpdateForecast)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteForecast/{id}", jsonResponseFormatter.FormatToJSON(ForecastController.DeleteForecast)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignModelVersionToForecast/{parentId}/modelVersionId", jsonResponseFormatter.FormatToJSON(ForecastController.AssignModelVersionToForecast)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignModelVersionFromForecast/{parentId}", jsonResponseFormatter.FormatToJSON(ForecastController.UnassignModelVersionFromForecast)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignTimeSeriesToForecast/{parentId}/timeSeriesId", jsonResponseFormatter.FormatToJSON(ForecastController.AssignTimeSeriesToForecast)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignTimeSeriesFromForecast/{parentId}", jsonResponseFormatter.FormatToJSON(ForecastController.UnassignTimeSeriesFromForecast)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddDatasetsToForecast/{parentId}/datasetsId", jsonResponseFormatter.FormatToJSON(ForecastController.AddDatasetsToForecast)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveDatasetsFromForecast/{parentId}/datasetsIds", jsonResponseFormatter.FormatToJSON(ForecastController.RemoveDatasetsFromForecast)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // TimeSeries Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/TimeSeries/{id}", jsonResponseFormatter.FormatToJSON(TimeSeriesController.GetTimeSeries)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/TimeSeries", jsonResponseFormatter.FormatToJSON(TimeSeriesController.GetAllTimeSeries)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewTimeSeries", jsonResponseFormatter.FormatToJSON(TimeSeriesController.CreateTimeSeries)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/TimeSeries/{id}", jsonResponseFormatter.FormatToJSON(TimeSeriesController.UpdateTimeSeries)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteTimeSeries/{id}", jsonResponseFormatter.FormatToJSON(TimeSeriesController.DeleteTimeSeries)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddDatasetsToTimeSeries/{parentId}/datasetsId", jsonResponseFormatter.FormatToJSON(TimeSeriesController.AddDatasetsToTimeSeries)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveDatasetsFromTimeSeries/{parentId}/datasetsIds", jsonResponseFormatter.FormatToJSON(TimeSeriesController.RemoveDatasetsFromTimeSeries)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddForecastsToTimeSeries/{parentId}/forecastsId", jsonResponseFormatter.FormatToJSON(TimeSeriesController.AddForecastsToTimeSeries)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveForecastsFromTimeSeries/{parentId}/forecastsIds", jsonResponseFormatter.FormatToJSON(TimeSeriesController.RemoveForecastsFromTimeSeries)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddAnomaliesToTimeSeries/{parentId}/anomaliesId", jsonResponseFormatter.FormatToJSON(TimeSeriesController.AddAnomaliesToTimeSeries)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveAnomaliesFromTimeSeries/{parentId}/anomaliesIds", jsonResponseFormatter.FormatToJSON(TimeSeriesController.RemoveAnomaliesFromTimeSeries)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Anomaly Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Anomaly/{id}", jsonResponseFormatter.FormatToJSON(AnomalyController.GetAnomaly)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Anomaly", jsonResponseFormatter.FormatToJSON(AnomalyController.GetAllAnomaly)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewAnomaly", jsonResponseFormatter.FormatToJSON(AnomalyController.CreateAnomaly)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Anomaly/{id}", jsonResponseFormatter.FormatToJSON(AnomalyController.UpdateAnomaly)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteAnomaly/{id}", jsonResponseFormatter.FormatToJSON(AnomalyController.DeleteAnomaly)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignTimeSeriesToAnomaly/{parentId}/timeSeriesId", jsonResponseFormatter.FormatToJSON(AnomalyController.AssignTimeSeriesToAnomaly)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignTimeSeriesFromAnomaly/{parentId}", jsonResponseFormatter.FormatToJSON(AnomalyController.UnassignTimeSeriesFromAnomaly)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignAlertToAnomaly/{parentId}/alertId", jsonResponseFormatter.FormatToJSON(AnomalyController.AssignAlertToAnomaly)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignAlertFromAnomaly/{parentId}", jsonResponseFormatter.FormatToJSON(AnomalyController.UnassignAlertFromAnomaly)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignDatasetToAnomaly/{parentId}/datasetId", jsonResponseFormatter.FormatToJSON(AnomalyController.AssignDatasetToAnomaly)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignDatasetFromAnomaly/{parentId}", jsonResponseFormatter.FormatToJSON(AnomalyController.UnassignDatasetFromAnomaly)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // QualityRule Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/QualityRule/{id}", jsonResponseFormatter.FormatToJSON(QualityRuleController.GetQualityRule)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/QualityRule", jsonResponseFormatter.FormatToJSON(QualityRuleController.GetAllQualityRule)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewQualityRule", jsonResponseFormatter.FormatToJSON(QualityRuleController.CreateQualityRule)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/QualityRule/{id}", jsonResponseFormatter.FormatToJSON(QualityRuleController.UpdateQualityRule)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteQualityRule/{id}", jsonResponseFormatter.FormatToJSON(QualityRuleController.DeleteQualityRule)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignDatasetToQualityRule/{parentId}/datasetId", jsonResponseFormatter.FormatToJSON(QualityRuleController.AssignDatasetToQualityRule)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignDatasetFromQualityRule/{parentId}", jsonResponseFormatter.FormatToJSON(QualityRuleController.UnassignDatasetFromQualityRule)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddChecksToQualityRule/{parentId}/checksId", jsonResponseFormatter.FormatToJSON(QualityRuleController.AddChecksToQualityRule)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveChecksFromQualityRule/{parentId}/checksIds", jsonResponseFormatter.FormatToJSON(QualityRuleController.RemoveChecksFromQualityRule)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // QualityCheck Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/QualityCheck/{id}", jsonResponseFormatter.FormatToJSON(QualityCheckController.GetQualityCheck)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/QualityCheck", jsonResponseFormatter.FormatToJSON(QualityCheckController.GetAllQualityCheck)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewQualityCheck", jsonResponseFormatter.FormatToJSON(QualityCheckController.CreateQualityCheck)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/QualityCheck/{id}", jsonResponseFormatter.FormatToJSON(QualityCheckController.UpdateQualityCheck)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteQualityCheck/{id}", jsonResponseFormatter.FormatToJSON(QualityCheckController.DeleteQualityCheck)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignRuleToQualityCheck/{parentId}/ruleId", jsonResponseFormatter.FormatToJSON(QualityCheckController.AssignRuleToQualityCheck)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignRuleFromQualityCheck/{parentId}", jsonResponseFormatter.FormatToJSON(QualityCheckController.UnassignRuleFromQualityCheck)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignDatasetToQualityCheck/{parentId}/datasetId", jsonResponseFormatter.FormatToJSON(QualityCheckController.AssignDatasetToQualityCheck)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignDatasetFromQualityCheck/{parentId}", jsonResponseFormatter.FormatToJSON(QualityCheckController.UnassignDatasetFromQualityCheck)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // LineageNode Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/LineageNode/{id}", jsonResponseFormatter.FormatToJSON(LineageNodeController.GetLineageNode)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/LineageNode", jsonResponseFormatter.FormatToJSON(LineageNodeController.GetAllLineageNode)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewLineageNode", jsonResponseFormatter.FormatToJSON(LineageNodeController.CreateLineageNode)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/LineageNode/{id}", jsonResponseFormatter.FormatToJSON(LineageNodeController.UpdateLineageNode)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteLineageNode/{id}", jsonResponseFormatter.FormatToJSON(LineageNodeController.DeleteLineageNode)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignWorkspaceToLineageNode/{parentId}/workspaceId", jsonResponseFormatter.FormatToJSON(LineageNodeController.AssignWorkspaceToLineageNode)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignWorkspaceFromLineageNode/{parentId}", jsonResponseFormatter.FormatToJSON(LineageNodeController.UnassignWorkspaceFromLineageNode)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddInputsToLineageNode/{parentId}/inputsId", jsonResponseFormatter.FormatToJSON(LineageNodeController.AddInputsToLineageNode)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveInputsFromLineageNode/{parentId}/inputsIds", jsonResponseFormatter.FormatToJSON(LineageNodeController.RemoveInputsFromLineageNode)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddOutputsToLineageNode/{parentId}/outputsId", jsonResponseFormatter.FormatToJSON(LineageNodeController.AddOutputsToLineageNode)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveOutputsFromLineageNode/{parentId}/outputsIds", jsonResponseFormatter.FormatToJSON(LineageNodeController.RemoveOutputsFromLineageNode)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddDatasetsToLineageNode/{parentId}/datasetsId", jsonResponseFormatter.FormatToJSON(LineageNodeController.AddDatasetsToLineageNode)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveDatasetsFromLineageNode/{parentId}/datasetsIds", jsonResponseFormatter.FormatToJSON(LineageNodeController.RemoveDatasetsFromLineageNode)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddModelsToLineageNode/{parentId}/modelsId", jsonResponseFormatter.FormatToJSON(LineageNodeController.AddModelsToLineageNode)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveModelsFromLineageNode/{parentId}/modelsIds", jsonResponseFormatter.FormatToJSON(LineageNodeController.RemoveModelsFromLineageNode)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddPipelinesToLineageNode/{parentId}/pipelinesId", jsonResponseFormatter.FormatToJSON(LineageNodeController.AddPipelinesToLineageNode)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemovePipelinesFromLineageNode/{parentId}/pipelinesIds", jsonResponseFormatter.FormatToJSON(LineageNodeController.RemovePipelinesFromLineageNode)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddDashboardsToLineageNode/{parentId}/dashboardsId", jsonResponseFormatter.FormatToJSON(LineageNodeController.AddDashboardsToLineageNode)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveDashboardsFromLineageNode/{parentId}/dashboardsIds", jsonResponseFormatter.FormatToJSON(LineageNodeController.RemoveDashboardsFromLineageNode)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddReportsToLineageNode/{parentId}/reportsId", jsonResponseFormatter.FormatToJSON(LineageNodeController.AddReportsToLineageNode)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveReportsFromLineageNode/{parentId}/reportsIds", jsonResponseFormatter.FormatToJSON(LineageNodeController.RemoveReportsFromLineageNode)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Tag Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Tag/{id}", jsonResponseFormatter.FormatToJSON(TagController.GetTag)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Tag", jsonResponseFormatter.FormatToJSON(TagController.GetAllTag)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewTag", jsonResponseFormatter.FormatToJSON(TagController.CreateTag)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Tag/{id}", jsonResponseFormatter.FormatToJSON(TagController.UpdateTag)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteTag/{id}", jsonResponseFormatter.FormatToJSON(TagController.DeleteTag)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddDatasetsToTag/{parentId}/datasetsId", jsonResponseFormatter.FormatToJSON(TagController.AddDatasetsToTag)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveDatasetsFromTag/{parentId}/datasetsIds", jsonResponseFormatter.FormatToJSON(TagController.RemoveDatasetsFromTag)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddModelsToTag/{parentId}/modelsId", jsonResponseFormatter.FormatToJSON(TagController.AddModelsToTag)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveModelsFromTag/{parentId}/modelsIds", jsonResponseFormatter.FormatToJSON(TagController.RemoveModelsFromTag)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddModelVersionsToTag/{parentId}/modelVersionsId", jsonResponseFormatter.FormatToJSON(TagController.AddModelVersionsToTag)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveModelVersionsFromTag/{parentId}/modelVersionsIds", jsonResponseFormatter.FormatToJSON(TagController.RemoveModelVersionsFromTag)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddDashboardsToTag/{parentId}/dashboardsId", jsonResponseFormatter.FormatToJSON(TagController.AddDashboardsToTag)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveDashboardsFromTag/{parentId}/dashboardsIds", jsonResponseFormatter.FormatToJSON(TagController.RemoveDashboardsFromTag)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddReportsToTag/{parentId}/reportsId", jsonResponseFormatter.FormatToJSON(TagController.AddReportsToTag)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveReportsFromTag/{parentId}/reportsIds", jsonResponseFormatter.FormatToJSON(TagController.RemoveReportsFromTag)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddFeatureSetsToTag/{parentId}/featureSetsId", jsonResponseFormatter.FormatToJSON(TagController.AddFeatureSetsToTag)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveFeatureSetsFromTag/{parentId}/featureSetsIds", jsonResponseFormatter.FormatToJSON(TagController.RemoveFeatureSetsFromTag)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddMetricsToTag/{parentId}/metricsId", jsonResponseFormatter.FormatToJSON(TagController.AddMetricsToTag)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveMetricsFromTag/{parentId}/metricsIds", jsonResponseFormatter.FormatToJSON(TagController.RemoveMetricsFromTag)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // AccessPolicy Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AccessPolicy/{id}", jsonResponseFormatter.FormatToJSON(AccessPolicyController.GetAccessPolicy)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/AccessPolicy", jsonResponseFormatter.FormatToJSON(AccessPolicyController.GetAllAccessPolicy)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewAccessPolicy", jsonResponseFormatter.FormatToJSON(AccessPolicyController.CreateAccessPolicy)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/AccessPolicy/{id}", jsonResponseFormatter.FormatToJSON(AccessPolicyController.UpdateAccessPolicy)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteAccessPolicy/{id}", jsonResponseFormatter.FormatToJSON(AccessPolicyController.DeleteAccessPolicy)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignWorkspaceToAccessPolicy/{parentId}/workspaceId", jsonResponseFormatter.FormatToJSON(AccessPolicyController.AssignWorkspaceToAccessPolicy)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignWorkspaceFromAccessPolicy/{parentId}", jsonResponseFormatter.FormatToJSON(AccessPolicyController.UnassignWorkspaceFromAccessPolicy)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddDatasetsToAccessPolicy/{parentId}/datasetsId", jsonResponseFormatter.FormatToJSON(AccessPolicyController.AddDatasetsToAccessPolicy)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveDatasetsFromAccessPolicy/{parentId}/datasetsIds", jsonResponseFormatter.FormatToJSON(AccessPolicyController.RemoveDatasetsFromAccessPolicy)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddDashboardsToAccessPolicy/{parentId}/dashboardsId", jsonResponseFormatter.FormatToJSON(AccessPolicyController.AddDashboardsToAccessPolicy)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveDashboardsFromAccessPolicy/{parentId}/dashboardsIds", jsonResponseFormatter.FormatToJSON(AccessPolicyController.RemoveDashboardsFromAccessPolicy)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddReportsToAccessPolicy/{parentId}/reportsId", jsonResponseFormatter.FormatToJSON(AccessPolicyController.AddReportsToAccessPolicy)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveReportsFromAccessPolicy/{parentId}/reportsIds", jsonResponseFormatter.FormatToJSON(AccessPolicyController.RemoveReportsFromAccessPolicy)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddModelsToAccessPolicy/{parentId}/modelsId", jsonResponseFormatter.FormatToJSON(AccessPolicyController.AddModelsToAccessPolicy)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveModelsFromAccessPolicy/{parentId}/modelsIds", jsonResponseFormatter.FormatToJSON(AccessPolicyController.RemoveModelsFromAccessPolicy)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddFeatureSetsToAccessPolicy/{parentId}/featureSetsId", jsonResponseFormatter.FormatToJSON(AccessPolicyController.AddFeatureSetsToAccessPolicy)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveFeatureSetsFromAccessPolicy/{parentId}/featureSetsIds", jsonResponseFormatter.FormatToJSON(AccessPolicyController.RemoveFeatureSetsFromAccessPolicy)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Alert Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Alert/{id}", jsonResponseFormatter.FormatToJSON(AlertController.GetAlert)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Alert", jsonResponseFormatter.FormatToJSON(AlertController.GetAllAlert)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewAlert", jsonResponseFormatter.FormatToJSON(AlertController.CreateAlert)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Alert/{id}", jsonResponseFormatter.FormatToJSON(AlertController.UpdateAlert)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteAlert/{id}", jsonResponseFormatter.FormatToJSON(AlertController.DeleteAlert)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignMetricToAlert/{parentId}/metricId", jsonResponseFormatter.FormatToJSON(AlertController.AssignMetricToAlert)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignMetricFromAlert/{parentId}", jsonResponseFormatter.FormatToJSON(AlertController.UnassignMetricFromAlert)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignDashboardToAlert/{parentId}/dashboardId", jsonResponseFormatter.FormatToJSON(AlertController.AssignDashboardToAlert)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignDashboardFromAlert/{parentId}", jsonResponseFormatter.FormatToJSON(AlertController.UnassignDashboardFromAlert)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignDatasetToAlert/{parentId}/datasetId", jsonResponseFormatter.FormatToJSON(AlertController.AssignDatasetToAlert)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignDatasetFromAlert/{parentId}", jsonResponseFormatter.FormatToJSON(AlertController.UnassignDatasetFromAlert)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignRuleToAlert/{parentId}/ruleId", jsonResponseFormatter.FormatToJSON(AlertController.AssignRuleToAlert)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignRuleFromAlert/{parentId}", jsonResponseFormatter.FormatToJSON(AlertController.UnassignRuleFromAlert)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddAnomaliesToAlert/{parentId}/anomaliesId", jsonResponseFormatter.FormatToJSON(AlertController.AddAnomaliesToAlert)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveAnomaliesFromAlert/{parentId}/anomaliesIds", jsonResponseFormatter.FormatToJSON(AlertController.RemoveAnomaliesFromAlert)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddSubscribersToAlert/{parentId}/subscribersId", jsonResponseFormatter.FormatToJSON(AlertController.AddSubscribersToAlert)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveSubscribersFromAlert/{parentId}/subscribersIds", jsonResponseFormatter.FormatToJSON(AlertController.RemoveSubscribersFromAlert)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Subscriber Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Subscriber/{id}", jsonResponseFormatter.FormatToJSON(SubscriberController.GetSubscriber)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Subscriber", jsonResponseFormatter.FormatToJSON(SubscriberController.GetAllSubscriber)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewSubscriber", jsonResponseFormatter.FormatToJSON(SubscriberController.CreateSubscriber)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Subscriber/{id}", jsonResponseFormatter.FormatToJSON(SubscriberController.UpdateSubscriber)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteSubscriber/{id}", jsonResponseFormatter.FormatToJSON(SubscriberController.DeleteSubscriber)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddAlertsToSubscriber/{parentId}/alertsId", jsonResponseFormatter.FormatToJSON(SubscriberController.AddAlertsToSubscriber)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveAlertsFromSubscriber/{parentId}/alertsIds", jsonResponseFormatter.FormatToJSON(SubscriberController.RemoveAlertsFromSubscriber)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // BusinessGlossaryTerm Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/BusinessGlossaryTerm/{id}", jsonResponseFormatter.FormatToJSON(BusinessGlossaryTermController.GetBusinessGlossaryTerm)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/BusinessGlossaryTerm", jsonResponseFormatter.FormatToJSON(BusinessGlossaryTermController.GetAllBusinessGlossaryTerm)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewBusinessGlossaryTerm", jsonResponseFormatter.FormatToJSON(BusinessGlossaryTermController.CreateBusinessGlossaryTerm)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/BusinessGlossaryTerm/{id}", jsonResponseFormatter.FormatToJSON(BusinessGlossaryTermController.UpdateBusinessGlossaryTerm)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteBusinessGlossaryTerm/{id}", jsonResponseFormatter.FormatToJSON(BusinessGlossaryTermController.DeleteBusinessGlossaryTerm)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddRelatedTermsToBusinessGlossaryTerm/{parentId}/relatedTermsId", jsonResponseFormatter.FormatToJSON(BusinessGlossaryTermController.AddRelatedTermsToBusinessGlossaryTerm)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveRelatedTermsFromBusinessGlossaryTerm/{parentId}/relatedTermsIds", jsonResponseFormatter.FormatToJSON(BusinessGlossaryTermController.RemoveRelatedTermsFromBusinessGlossaryTerm)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddMetricsToBusinessGlossaryTerm/{parentId}/metricsId", jsonResponseFormatter.FormatToJSON(BusinessGlossaryTermController.AddMetricsToBusinessGlossaryTerm)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveMetricsFromBusinessGlossaryTerm/{parentId}/metricsIds", jsonResponseFormatter.FormatToJSON(BusinessGlossaryTermController.RemoveMetricsFromBusinessGlossaryTerm)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddDatasetsToBusinessGlossaryTerm/{parentId}/datasetsId", jsonResponseFormatter.FormatToJSON(BusinessGlossaryTermController.AddDatasetsToBusinessGlossaryTerm)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveDatasetsFromBusinessGlossaryTerm/{parentId}/datasetsIds", jsonResponseFormatter.FormatToJSON(BusinessGlossaryTermController.RemoveDatasetsFromBusinessGlossaryTerm)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddDimensionsToBusinessGlossaryTerm/{parentId}/dimensionsId", jsonResponseFormatter.FormatToJSON(BusinessGlossaryTermController.AddDimensionsToBusinessGlossaryTerm)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveDimensionsFromBusinessGlossaryTerm/{parentId}/dimensionsIds", jsonResponseFormatter.FormatToJSON(BusinessGlossaryTermController.RemoveDimensionsFromBusinessGlossaryTerm)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddMeasuresToBusinessGlossaryTerm/{parentId}/measuresId", jsonResponseFormatter.FormatToJSON(BusinessGlossaryTermController.AddMeasuresToBusinessGlossaryTerm)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveMeasuresFromBusinessGlossaryTerm/{parentId}/measuresIds", jsonResponseFormatter.FormatToJSON(BusinessGlossaryTermController.RemoveMeasuresFromBusinessGlossaryTerm)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // RecommendationScenario Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/RecommendationScenario/{id}", jsonResponseFormatter.FormatToJSON(RecommendationScenarioController.GetRecommendationScenario)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/RecommendationScenario", jsonResponseFormatter.FormatToJSON(RecommendationScenarioController.GetAllRecommendationScenario)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewRecommendationScenario", jsonResponseFormatter.FormatToJSON(RecommendationScenarioController.CreateRecommendationScenario)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/RecommendationScenario/{id}", jsonResponseFormatter.FormatToJSON(RecommendationScenarioController.UpdateRecommendationScenario)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteRecommendationScenario/{id}", jsonResponseFormatter.FormatToJSON(RecommendationScenarioController.DeleteRecommendationScenario)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddModelsToRecommendationScenario/{parentId}/modelsId", jsonResponseFormatter.FormatToJSON(RecommendationScenarioController.AddModelsToRecommendationScenario)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveModelsFromRecommendationScenario/{parentId}/modelsIds", jsonResponseFormatter.FormatToJSON(RecommendationScenarioController.RemoveModelsFromRecommendationScenario)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddDatasetsToRecommendationScenario/{parentId}/datasetsId", jsonResponseFormatter.FormatToJSON(RecommendationScenarioController.AddDatasetsToRecommendationScenario)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveDatasetsFromRecommendationScenario/{parentId}/datasetsIds", jsonResponseFormatter.FormatToJSON(RecommendationScenarioController.RemoveDatasetsFromRecommendationScenario)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddExperimentsToRecommendationScenario/{parentId}/experimentsId", jsonResponseFormatter.FormatToJSON(RecommendationScenarioController.AddExperimentsToRecommendationScenario)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveExperimentsFromRecommendationScenario/{parentId}/experimentsIds", jsonResponseFormatter.FormatToJSON(RecommendationScenarioController.RemoveExperimentsFromRecommendationScenario)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddAlertsToRecommendationScenario/{parentId}/alertsId", jsonResponseFormatter.FormatToJSON(RecommendationScenarioController.AddAlertsToRecommendationScenario)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveAlertsFromRecommendationScenario/{parentId}/alertsIds", jsonResponseFormatter.FormatToJSON(RecommendationScenarioController.RemoveAlertsFromRecommendationScenario)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // FraudScenario Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/FraudScenario/{id}", jsonResponseFormatter.FormatToJSON(FraudScenarioController.GetFraudScenario)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/FraudScenario", jsonResponseFormatter.FormatToJSON(FraudScenarioController.GetAllFraudScenario)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewFraudScenario", jsonResponseFormatter.FormatToJSON(FraudScenarioController.CreateFraudScenario)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/FraudScenario/{id}", jsonResponseFormatter.FormatToJSON(FraudScenarioController.UpdateFraudScenario)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteFraudScenario/{id}", jsonResponseFormatter.FormatToJSON(FraudScenarioController.DeleteFraudScenario)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddModelsToFraudScenario/{parentId}/modelsId", jsonResponseFormatter.FormatToJSON(FraudScenarioController.AddModelsToFraudScenario)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveModelsFromFraudScenario/{parentId}/modelsIds", jsonResponseFormatter.FormatToJSON(FraudScenarioController.RemoveModelsFromFraudScenario)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddDatasetsToFraudScenario/{parentId}/datasetsId", jsonResponseFormatter.FormatToJSON(FraudScenarioController.AddDatasetsToFraudScenario)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveDatasetsFromFraudScenario/{parentId}/datasetsIds", jsonResponseFormatter.FormatToJSON(FraudScenarioController.RemoveDatasetsFromFraudScenario)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddAlertsToFraudScenario/{parentId}/alertsId", jsonResponseFormatter.FormatToJSON(FraudScenarioController.AddAlertsToFraudScenario)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveAlertsFromFraudScenario/{parentId}/alertsIds", jsonResponseFormatter.FormatToJSON(FraudScenarioController.RemoveAlertsFromFraudScenario)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddSignalsToFraudScenario/{parentId}/signalsId", jsonResponseFormatter.FormatToJSON(FraudScenarioController.AddSignalsToFraudScenario)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveSignalsFromFraudScenario/{parentId}/signalsIds", jsonResponseFormatter.FormatToJSON(FraudScenarioController.RemoveSignalsFromFraudScenario)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // FraudSignal Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/FraudSignal/{id}", jsonResponseFormatter.FormatToJSON(FraudSignalController.GetFraudSignal)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/FraudSignal", jsonResponseFormatter.FormatToJSON(FraudSignalController.GetAllFraudSignal)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewFraudSignal", jsonResponseFormatter.FormatToJSON(FraudSignalController.CreateFraudSignal)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/FraudSignal/{id}", jsonResponseFormatter.FormatToJSON(FraudSignalController.UpdateFraudSignal)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteFraudSignal/{id}", jsonResponseFormatter.FormatToJSON(FraudSignalController.DeleteFraudSignal)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignScenarioToFraudSignal/{parentId}/scenarioId", jsonResponseFormatter.FormatToJSON(FraudSignalController.AssignScenarioToFraudSignal)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignScenarioFromFraudSignal/{parentId}", jsonResponseFormatter.FormatToJSON(FraudSignalController.UnassignScenarioFromFraudSignal)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignDatasetToFraudSignal/{parentId}/datasetId", jsonResponseFormatter.FormatToJSON(FraudSignalController.AssignDatasetToFraudSignal)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignDatasetFromFraudSignal/{parentId}", jsonResponseFormatter.FormatToJSON(FraudSignalController.UnassignDatasetFromFraudSignal)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignModelVersionToFraudSignal/{parentId}/modelVersionId", jsonResponseFormatter.FormatToJSON(FraudSignalController.AssignModelVersionToFraudSignal)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignModelVersionFromFraudSignal/{parentId}", jsonResponseFormatter.FormatToJSON(FraudSignalController.UnassignModelVersionFromFraudSignal)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    return router
}
