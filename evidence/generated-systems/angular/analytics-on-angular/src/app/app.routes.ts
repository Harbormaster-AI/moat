import { Routes } from '@angular/router';

import { CreateAnalyticsWorkspaceComponent } from './components/AnalyticsWorkspace/create/create.component';
import { EditAnalyticsWorkspaceComponent } from './components/AnalyticsWorkspace/edit/edit.component';
import { IndexAnalyticsWorkspaceComponent } from './components/AnalyticsWorkspace/index/index.component';
import { CreateDataSourceComponent } from './components/DataSource/create/create.component';
import { EditDataSourceComponent } from './components/DataSource/edit/edit.component';
import { IndexDataSourceComponent } from './components/DataSource/index/index.component';
import { CreateDataSetComponent } from './components/DataSet/create/create.component';
import { EditDataSetComponent } from './components/DataSet/edit/edit.component';
import { IndexDataSetComponent } from './components/DataSet/index/index.component';
import { CreateDataPipelineComponent } from './components/DataPipeline/create/create.component';
import { EditDataPipelineComponent } from './components/DataPipeline/edit/edit.component';
import { IndexDataPipelineComponent } from './components/DataPipeline/index/index.component';
import { CreateDataTaskComponent } from './components/DataTask/create/create.component';
import { EditDataTaskComponent } from './components/DataTask/edit/edit.component';
import { IndexDataTaskComponent } from './components/DataTask/index/index.component';
import { CreateSemanticModelComponent } from './components/SemanticModel/create/create.component';
import { EditSemanticModelComponent } from './components/SemanticModel/edit/edit.component';
import { IndexSemanticModelComponent } from './components/SemanticModel/index/index.component';
import { CreateDimensionComponent } from './components/Dimension/create/create.component';
import { EditDimensionComponent } from './components/Dimension/edit/edit.component';
import { IndexDimensionComponent } from './components/Dimension/index/index.component';
import { CreateMeasureComponent } from './components/Measure/create/create.component';
import { EditMeasureComponent } from './components/Measure/edit/edit.component';
import { IndexMeasureComponent } from './components/Measure/index/index.component';
import { CreateMetricComponent } from './components/Metric/create/create.component';
import { EditMetricComponent } from './components/Metric/edit/edit.component';
import { IndexMetricComponent } from './components/Metric/index/index.component';
import { CreateReportComponent } from './components/Report/create/create.component';
import { EditReportComponent } from './components/Report/edit/edit.component';
import { IndexReportComponent } from './components/Report/index/index.component';
import { CreateDashboardComponent } from './components/Dashboard/create/create.component';
import { EditDashboardComponent } from './components/Dashboard/edit/edit.component';
import { IndexDashboardComponent } from './components/Dashboard/index/index.component';
import { CreateVisualizationComponent } from './components/Visualization/create/create.component';
import { EditVisualizationComponent } from './components/Visualization/edit/edit.component';
import { IndexVisualizationComponent } from './components/Visualization/index/index.component';
import { CreateNotebookComponent } from './components/Notebook/create/create.component';
import { EditNotebookComponent } from './components/Notebook/edit/edit.component';
import { IndexNotebookComponent } from './components/Notebook/index/index.component';
import { CreateBIQueryComponent } from './components/BIQuery/create/create.component';
import { EditBIQueryComponent } from './components/BIQuery/edit/edit.component';
import { IndexBIQueryComponent } from './components/BIQuery/index/index.component';
import { CreateExperimentComponent } from './components/Experiment/create/create.component';
import { EditExperimentComponent } from './components/Experiment/edit/edit.component';
import { IndexExperimentComponent } from './components/Experiment/index/index.component';
import { CreateTrainingRunComponent } from './components/TrainingRun/create/create.component';
import { EditTrainingRunComponent } from './components/TrainingRun/edit/edit.component';
import { IndexTrainingRunComponent } from './components/TrainingRun/index/index.component';
import { CreateRunMetricComponent } from './components/RunMetric/create/create.component';
import { EditRunMetricComponent } from './components/RunMetric/edit/edit.component';
import { IndexRunMetricComponent } from './components/RunMetric/index/index.component';
import { CreateRunParameterComponent } from './components/RunParameter/create/create.component';
import { EditRunParameterComponent } from './components/RunParameter/edit/edit.component';
import { IndexRunParameterComponent } from './components/RunParameter/index/index.component';
import { CreateModel_Component } from './components/Model_/create/create.component';
import { EditModel_Component } from './components/Model_/edit/edit.component';
import { IndexModel_Component } from './components/Model_/index/index.component';
import { CreateModelVersionComponent } from './components/ModelVersion/create/create.component';
import { EditModelVersionComponent } from './components/ModelVersion/edit/edit.component';
import { IndexModelVersionComponent } from './components/ModelVersion/index/index.component';
import { CreateEvaluationMetricComponent } from './components/EvaluationMetric/create/create.component';
import { EditEvaluationMetricComponent } from './components/EvaluationMetric/edit/edit.component';
import { IndexEvaluationMetricComponent } from './components/EvaluationMetric/index/index.component';
import { CreateFeatureSetComponent } from './components/FeatureSet/create/create.component';
import { EditFeatureSetComponent } from './components/FeatureSet/edit/edit.component';
import { IndexFeatureSetComponent } from './components/FeatureSet/index/index.component';
import { CreateFeatureComponent } from './components/Feature/create/create.component';
import { EditFeatureComponent } from './components/Feature/edit/edit.component';
import { IndexFeatureComponent } from './components/Feature/index/index.component';
import { CreateInferenceEndpointComponent } from './components/InferenceEndpoint/create/create.component';
import { EditInferenceEndpointComponent } from './components/InferenceEndpoint/edit/edit.component';
import { IndexInferenceEndpointComponent } from './components/InferenceEndpoint/index/index.component';
import { CreatePredictionComponent } from './components/Prediction/create/create.component';
import { EditPredictionComponent } from './components/Prediction/edit/edit.component';
import { IndexPredictionComponent } from './components/Prediction/index/index.component';
import { CreateForecastComponent } from './components/Forecast/create/create.component';
import { EditForecastComponent } from './components/Forecast/edit/edit.component';
import { IndexForecastComponent } from './components/Forecast/index/index.component';
import { CreateTimeSeriesComponent } from './components/TimeSeries/create/create.component';
import { EditTimeSeriesComponent } from './components/TimeSeries/edit/edit.component';
import { IndexTimeSeriesComponent } from './components/TimeSeries/index/index.component';
import { CreateAnomalyComponent } from './components/Anomaly/create/create.component';
import { EditAnomalyComponent } from './components/Anomaly/edit/edit.component';
import { IndexAnomalyComponent } from './components/Anomaly/index/index.component';
import { CreateQualityRuleComponent } from './components/QualityRule/create/create.component';
import { EditQualityRuleComponent } from './components/QualityRule/edit/edit.component';
import { IndexQualityRuleComponent } from './components/QualityRule/index/index.component';
import { CreateQualityCheckComponent } from './components/QualityCheck/create/create.component';
import { EditQualityCheckComponent } from './components/QualityCheck/edit/edit.component';
import { IndexQualityCheckComponent } from './components/QualityCheck/index/index.component';
import { CreateLineageNodeComponent } from './components/LineageNode/create/create.component';
import { EditLineageNodeComponent } from './components/LineageNode/edit/edit.component';
import { IndexLineageNodeComponent } from './components/LineageNode/index/index.component';
import { CreateTagComponent } from './components/Tag/create/create.component';
import { EditTagComponent } from './components/Tag/edit/edit.component';
import { IndexTagComponent } from './components/Tag/index/index.component';
import { CreateAccessPolicyComponent } from './components/AccessPolicy/create/create.component';
import { EditAccessPolicyComponent } from './components/AccessPolicy/edit/edit.component';
import { IndexAccessPolicyComponent } from './components/AccessPolicy/index/index.component';
import { CreateAlertComponent } from './components/Alert/create/create.component';
import { EditAlertComponent } from './components/Alert/edit/edit.component';
import { IndexAlertComponent } from './components/Alert/index/index.component';
import { CreateSubscriberComponent } from './components/Subscriber/create/create.component';
import { EditSubscriberComponent } from './components/Subscriber/edit/edit.component';
import { IndexSubscriberComponent } from './components/Subscriber/index/index.component';
import { CreateBusinessGlossaryTermComponent } from './components/BusinessGlossaryTerm/create/create.component';
import { EditBusinessGlossaryTermComponent } from './components/BusinessGlossaryTerm/edit/edit.component';
import { IndexBusinessGlossaryTermComponent } from './components/BusinessGlossaryTerm/index/index.component';
import { CreateRecommendationScenarioComponent } from './components/RecommendationScenario/create/create.component';
import { EditRecommendationScenarioComponent } from './components/RecommendationScenario/edit/edit.component';
import { IndexRecommendationScenarioComponent } from './components/RecommendationScenario/index/index.component';
import { CreateFraudScenarioComponent } from './components/FraudScenario/create/create.component';
import { EditFraudScenarioComponent } from './components/FraudScenario/edit/edit.component';
import { IndexFraudScenarioComponent } from './components/FraudScenario/index/index.component';
import { CreateFraudSignalComponent } from './components/FraudSignal/create/create.component';
import { EditFraudSignalComponent } from './components/FraudSignal/edit/edit.component';
import { IndexFraudSignalComponent } from './components/FraudSignal/index/index.component';

export const routes: Routes = [

        {
        path: 'createAnalyticsWorkspace',
        component: CreateAnalyticsWorkspaceComponent
},
{
    path: 'editAnalyticsWorkspace/:id',
        component: EditAnalyticsWorkspaceComponent
},
{
    path: 'indexAnalyticsWorkspace',
        component: IndexAnalyticsWorkspaceComponent
},
    {
        path: 'createDataSource',
        component: CreateDataSourceComponent
},
{
    path: 'editDataSource/:id',
        component: EditDataSourceComponent
},
{
    path: 'indexDataSource',
        component: IndexDataSourceComponent
},
    {
        path: 'createDataSet',
        component: CreateDataSetComponent
},
{
    path: 'editDataSet/:id',
        component: EditDataSetComponent
},
{
    path: 'indexDataSet',
        component: IndexDataSetComponent
},
    {
        path: 'createDataPipeline',
        component: CreateDataPipelineComponent
},
{
    path: 'editDataPipeline/:id',
        component: EditDataPipelineComponent
},
{
    path: 'indexDataPipeline',
        component: IndexDataPipelineComponent
},
    {
        path: 'createDataTask',
        component: CreateDataTaskComponent
},
{
    path: 'editDataTask/:id',
        component: EditDataTaskComponent
},
{
    path: 'indexDataTask',
        component: IndexDataTaskComponent
},
    {
        path: 'createSemanticModel',
        component: CreateSemanticModelComponent
},
{
    path: 'editSemanticModel/:id',
        component: EditSemanticModelComponent
},
{
    path: 'indexSemanticModel',
        component: IndexSemanticModelComponent
},
    {
        path: 'createDimension',
        component: CreateDimensionComponent
},
{
    path: 'editDimension/:id',
        component: EditDimensionComponent
},
{
    path: 'indexDimension',
        component: IndexDimensionComponent
},
    {
        path: 'createMeasure',
        component: CreateMeasureComponent
},
{
    path: 'editMeasure/:id',
        component: EditMeasureComponent
},
{
    path: 'indexMeasure',
        component: IndexMeasureComponent
},
    {
        path: 'createMetric',
        component: CreateMetricComponent
},
{
    path: 'editMetric/:id',
        component: EditMetricComponent
},
{
    path: 'indexMetric',
        component: IndexMetricComponent
},
    {
        path: 'createReport',
        component: CreateReportComponent
},
{
    path: 'editReport/:id',
        component: EditReportComponent
},
{
    path: 'indexReport',
        component: IndexReportComponent
},
    {
        path: 'createDashboard',
        component: CreateDashboardComponent
},
{
    path: 'editDashboard/:id',
        component: EditDashboardComponent
},
{
    path: 'indexDashboard',
        component: IndexDashboardComponent
},
    {
        path: 'createVisualization',
        component: CreateVisualizationComponent
},
{
    path: 'editVisualization/:id',
        component: EditVisualizationComponent
},
{
    path: 'indexVisualization',
        component: IndexVisualizationComponent
},
    {
        path: 'createNotebook',
        component: CreateNotebookComponent
},
{
    path: 'editNotebook/:id',
        component: EditNotebookComponent
},
{
    path: 'indexNotebook',
        component: IndexNotebookComponent
},
    {
        path: 'createBIQuery',
        component: CreateBIQueryComponent
},
{
    path: 'editBIQuery/:id',
        component: EditBIQueryComponent
},
{
    path: 'indexBIQuery',
        component: IndexBIQueryComponent
},
    {
        path: 'createExperiment',
        component: CreateExperimentComponent
},
{
    path: 'editExperiment/:id',
        component: EditExperimentComponent
},
{
    path: 'indexExperiment',
        component: IndexExperimentComponent
},
    {
        path: 'createTrainingRun',
        component: CreateTrainingRunComponent
},
{
    path: 'editTrainingRun/:id',
        component: EditTrainingRunComponent
},
{
    path: 'indexTrainingRun',
        component: IndexTrainingRunComponent
},
    {
        path: 'createRunMetric',
        component: CreateRunMetricComponent
},
{
    path: 'editRunMetric/:id',
        component: EditRunMetricComponent
},
{
    path: 'indexRunMetric',
        component: IndexRunMetricComponent
},
    {
        path: 'createRunParameter',
        component: CreateRunParameterComponent
},
{
    path: 'editRunParameter/:id',
        component: EditRunParameterComponent
},
{
    path: 'indexRunParameter',
        component: IndexRunParameterComponent
},
    {
        path: 'createModel_',
        component: CreateModel_Component
},
{
    path: 'editModel_/:id',
        component: EditModel_Component
},
{
    path: 'indexModel_',
        component: IndexModel_Component
},
    {
        path: 'createModelVersion',
        component: CreateModelVersionComponent
},
{
    path: 'editModelVersion/:id',
        component: EditModelVersionComponent
},
{
    path: 'indexModelVersion',
        component: IndexModelVersionComponent
},
    {
        path: 'createEvaluationMetric',
        component: CreateEvaluationMetricComponent
},
{
    path: 'editEvaluationMetric/:id',
        component: EditEvaluationMetricComponent
},
{
    path: 'indexEvaluationMetric',
        component: IndexEvaluationMetricComponent
},
    {
        path: 'createFeatureSet',
        component: CreateFeatureSetComponent
},
{
    path: 'editFeatureSet/:id',
        component: EditFeatureSetComponent
},
{
    path: 'indexFeatureSet',
        component: IndexFeatureSetComponent
},
    {
        path: 'createFeature',
        component: CreateFeatureComponent
},
{
    path: 'editFeature/:id',
        component: EditFeatureComponent
},
{
    path: 'indexFeature',
        component: IndexFeatureComponent
},
    {
        path: 'createInferenceEndpoint',
        component: CreateInferenceEndpointComponent
},
{
    path: 'editInferenceEndpoint/:id',
        component: EditInferenceEndpointComponent
},
{
    path: 'indexInferenceEndpoint',
        component: IndexInferenceEndpointComponent
},
    {
        path: 'createPrediction',
        component: CreatePredictionComponent
},
{
    path: 'editPrediction/:id',
        component: EditPredictionComponent
},
{
    path: 'indexPrediction',
        component: IndexPredictionComponent
},
    {
        path: 'createForecast',
        component: CreateForecastComponent
},
{
    path: 'editForecast/:id',
        component: EditForecastComponent
},
{
    path: 'indexForecast',
        component: IndexForecastComponent
},
    {
        path: 'createTimeSeries',
        component: CreateTimeSeriesComponent
},
{
    path: 'editTimeSeries/:id',
        component: EditTimeSeriesComponent
},
{
    path: 'indexTimeSeries',
        component: IndexTimeSeriesComponent
},
    {
        path: 'createAnomaly',
        component: CreateAnomalyComponent
},
{
    path: 'editAnomaly/:id',
        component: EditAnomalyComponent
},
{
    path: 'indexAnomaly',
        component: IndexAnomalyComponent
},
    {
        path: 'createQualityRule',
        component: CreateQualityRuleComponent
},
{
    path: 'editQualityRule/:id',
        component: EditQualityRuleComponent
},
{
    path: 'indexQualityRule',
        component: IndexQualityRuleComponent
},
    {
        path: 'createQualityCheck',
        component: CreateQualityCheckComponent
},
{
    path: 'editQualityCheck/:id',
        component: EditQualityCheckComponent
},
{
    path: 'indexQualityCheck',
        component: IndexQualityCheckComponent
},
    {
        path: 'createLineageNode',
        component: CreateLineageNodeComponent
},
{
    path: 'editLineageNode/:id',
        component: EditLineageNodeComponent
},
{
    path: 'indexLineageNode',
        component: IndexLineageNodeComponent
},
    {
        path: 'createTag',
        component: CreateTagComponent
},
{
    path: 'editTag/:id',
        component: EditTagComponent
},
{
    path: 'indexTag',
        component: IndexTagComponent
},
    {
        path: 'createAccessPolicy',
        component: CreateAccessPolicyComponent
},
{
    path: 'editAccessPolicy/:id',
        component: EditAccessPolicyComponent
},
{
    path: 'indexAccessPolicy',
        component: IndexAccessPolicyComponent
},
    {
        path: 'createAlert',
        component: CreateAlertComponent
},
{
    path: 'editAlert/:id',
        component: EditAlertComponent
},
{
    path: 'indexAlert',
        component: IndexAlertComponent
},
    {
        path: 'createSubscriber',
        component: CreateSubscriberComponent
},
{
    path: 'editSubscriber/:id',
        component: EditSubscriberComponent
},
{
    path: 'indexSubscriber',
        component: IndexSubscriberComponent
},
    {
        path: 'createBusinessGlossaryTerm',
        component: CreateBusinessGlossaryTermComponent
},
{
    path: 'editBusinessGlossaryTerm/:id',
        component: EditBusinessGlossaryTermComponent
},
{
    path: 'indexBusinessGlossaryTerm',
        component: IndexBusinessGlossaryTermComponent
},
    {
        path: 'createRecommendationScenario',
        component: CreateRecommendationScenarioComponent
},
{
    path: 'editRecommendationScenario/:id',
        component: EditRecommendationScenarioComponent
},
{
    path: 'indexRecommendationScenario',
        component: IndexRecommendationScenarioComponent
},
    {
        path: 'createFraudScenario',
        component: CreateFraudScenarioComponent
},
{
    path: 'editFraudScenario/:id',
        component: EditFraudScenarioComponent
},
{
    path: 'indexFraudScenario',
        component: IndexFraudScenarioComponent
},
    {
        path: 'createFraudSignal',
        component: CreateFraudSignalComponent
},
{
    path: 'editFraudSignal/:id',
        component: EditFraudSignalComponent
},
{
    path: 'indexFraudSignal',
        component: IndexFraudSignalComponent
}
];