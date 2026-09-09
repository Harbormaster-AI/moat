import React from 'react';
import './App.css';
import {BrowserRouter as Router, Route, Switch} from 'react-router-dom'
import HomePageComponent from './components/HomePageComponent';
import HeaderComponent from './components/HeaderComponent';
import FooterComponent from './components/FooterComponent';
import ListAnalyticsWorkspaceComponent from './components/ListAnalyticsWorkspaceComponent';
import CreateAnalyticsWorkspaceComponent from './components/CreateAnalyticsWorkspaceComponent';
import ViewAnalyticsWorkspaceComponent from './components/ViewAnalyticsWorkspaceComponent';
import ListDataSourceComponent from './components/ListDataSourceComponent';
import CreateDataSourceComponent from './components/CreateDataSourceComponent';
import ViewDataSourceComponent from './components/ViewDataSourceComponent';
import ListDataSetComponent from './components/ListDataSetComponent';
import CreateDataSetComponent from './components/CreateDataSetComponent';
import ViewDataSetComponent from './components/ViewDataSetComponent';
import ListDataPipelineComponent from './components/ListDataPipelineComponent';
import CreateDataPipelineComponent from './components/CreateDataPipelineComponent';
import ViewDataPipelineComponent from './components/ViewDataPipelineComponent';
import ListDataTaskComponent from './components/ListDataTaskComponent';
import CreateDataTaskComponent from './components/CreateDataTaskComponent';
import ViewDataTaskComponent from './components/ViewDataTaskComponent';
import ListSemanticModelComponent from './components/ListSemanticModelComponent';
import CreateSemanticModelComponent from './components/CreateSemanticModelComponent';
import ViewSemanticModelComponent from './components/ViewSemanticModelComponent';
import ListDimensionComponent from './components/ListDimensionComponent';
import CreateDimensionComponent from './components/CreateDimensionComponent';
import ViewDimensionComponent from './components/ViewDimensionComponent';
import ListMeasureComponent from './components/ListMeasureComponent';
import CreateMeasureComponent from './components/CreateMeasureComponent';
import ViewMeasureComponent from './components/ViewMeasureComponent';
import ListMetricComponent from './components/ListMetricComponent';
import CreateMetricComponent from './components/CreateMetricComponent';
import ViewMetricComponent from './components/ViewMetricComponent';
import ListReportComponent from './components/ListReportComponent';
import CreateReportComponent from './components/CreateReportComponent';
import ViewReportComponent from './components/ViewReportComponent';
import ListDashboardComponent from './components/ListDashboardComponent';
import CreateDashboardComponent from './components/CreateDashboardComponent';
import ViewDashboardComponent from './components/ViewDashboardComponent';
import ListVisualizationComponent from './components/ListVisualizationComponent';
import CreateVisualizationComponent from './components/CreateVisualizationComponent';
import ViewVisualizationComponent from './components/ViewVisualizationComponent';
import ListNotebookComponent from './components/ListNotebookComponent';
import CreateNotebookComponent from './components/CreateNotebookComponent';
import ViewNotebookComponent from './components/ViewNotebookComponent';
import ListBIQueryComponent from './components/ListBIQueryComponent';
import CreateBIQueryComponent from './components/CreateBIQueryComponent';
import ViewBIQueryComponent from './components/ViewBIQueryComponent';
import ListExperimentComponent from './components/ListExperimentComponent';
import CreateExperimentComponent from './components/CreateExperimentComponent';
import ViewExperimentComponent from './components/ViewExperimentComponent';
import ListTrainingRunComponent from './components/ListTrainingRunComponent';
import CreateTrainingRunComponent from './components/CreateTrainingRunComponent';
import ViewTrainingRunComponent from './components/ViewTrainingRunComponent';
import ListRunMetricComponent from './components/ListRunMetricComponent';
import CreateRunMetricComponent from './components/CreateRunMetricComponent';
import ViewRunMetricComponent from './components/ViewRunMetricComponent';
import ListRunParameterComponent from './components/ListRunParameterComponent';
import CreateRunParameterComponent from './components/CreateRunParameterComponent';
import ViewRunParameterComponent from './components/ViewRunParameterComponent';
import ListModelComponent from './components/ListModelComponent';
import CreateModelComponent from './components/CreateModelComponent';
import ViewModelComponent from './components/ViewModelComponent';
import ListModelVersionComponent from './components/ListModelVersionComponent';
import CreateModelVersionComponent from './components/CreateModelVersionComponent';
import ViewModelVersionComponent from './components/ViewModelVersionComponent';
import ListEvaluationMetricComponent from './components/ListEvaluationMetricComponent';
import CreateEvaluationMetricComponent from './components/CreateEvaluationMetricComponent';
import ViewEvaluationMetricComponent from './components/ViewEvaluationMetricComponent';
import ListFeatureSetComponent from './components/ListFeatureSetComponent';
import CreateFeatureSetComponent from './components/CreateFeatureSetComponent';
import ViewFeatureSetComponent from './components/ViewFeatureSetComponent';
import ListFeatureComponent from './components/ListFeatureComponent';
import CreateFeatureComponent from './components/CreateFeatureComponent';
import ViewFeatureComponent from './components/ViewFeatureComponent';
import ListInferenceEndpointComponent from './components/ListInferenceEndpointComponent';
import CreateInferenceEndpointComponent from './components/CreateInferenceEndpointComponent';
import ViewInferenceEndpointComponent from './components/ViewInferenceEndpointComponent';
import ListPredictionComponent from './components/ListPredictionComponent';
import CreatePredictionComponent from './components/CreatePredictionComponent';
import ViewPredictionComponent from './components/ViewPredictionComponent';
import ListForecastComponent from './components/ListForecastComponent';
import CreateForecastComponent from './components/CreateForecastComponent';
import ViewForecastComponent from './components/ViewForecastComponent';
import ListTimeSeriesComponent from './components/ListTimeSeriesComponent';
import CreateTimeSeriesComponent from './components/CreateTimeSeriesComponent';
import ViewTimeSeriesComponent from './components/ViewTimeSeriesComponent';
import ListAnomalyComponent from './components/ListAnomalyComponent';
import CreateAnomalyComponent from './components/CreateAnomalyComponent';
import ViewAnomalyComponent from './components/ViewAnomalyComponent';
import ListQualityRuleComponent from './components/ListQualityRuleComponent';
import CreateQualityRuleComponent from './components/CreateQualityRuleComponent';
import ViewQualityRuleComponent from './components/ViewQualityRuleComponent';
import ListQualityCheckComponent from './components/ListQualityCheckComponent';
import CreateQualityCheckComponent from './components/CreateQualityCheckComponent';
import ViewQualityCheckComponent from './components/ViewQualityCheckComponent';
import ListLineageNodeComponent from './components/ListLineageNodeComponent';
import CreateLineageNodeComponent from './components/CreateLineageNodeComponent';
import ViewLineageNodeComponent from './components/ViewLineageNodeComponent';
import ListTagComponent from './components/ListTagComponent';
import CreateTagComponent from './components/CreateTagComponent';
import ViewTagComponent from './components/ViewTagComponent';
import ListAccessPolicyComponent from './components/ListAccessPolicyComponent';
import CreateAccessPolicyComponent from './components/CreateAccessPolicyComponent';
import ViewAccessPolicyComponent from './components/ViewAccessPolicyComponent';
import ListAlertComponent from './components/ListAlertComponent';
import CreateAlertComponent from './components/CreateAlertComponent';
import ViewAlertComponent from './components/ViewAlertComponent';
import ListSubscriberComponent from './components/ListSubscriberComponent';
import CreateSubscriberComponent from './components/CreateSubscriberComponent';
import ViewSubscriberComponent from './components/ViewSubscriberComponent';
import ListBusinessGlossaryTermComponent from './components/ListBusinessGlossaryTermComponent';
import CreateBusinessGlossaryTermComponent from './components/CreateBusinessGlossaryTermComponent';
import ViewBusinessGlossaryTermComponent from './components/ViewBusinessGlossaryTermComponent';
import ListRecommendationScenarioComponent from './components/ListRecommendationScenarioComponent';
import CreateRecommendationScenarioComponent from './components/CreateRecommendationScenarioComponent';
import ViewRecommendationScenarioComponent from './components/ViewRecommendationScenarioComponent';
import ListFraudScenarioComponent from './components/ListFraudScenarioComponent';
import CreateFraudScenarioComponent from './components/CreateFraudScenarioComponent';
import ViewFraudScenarioComponent from './components/ViewFraudScenarioComponent';
import ListFraudSignalComponent from './components/ListFraudSignalComponent';
import CreateFraudSignalComponent from './components/CreateFraudSignalComponent';
import ViewFraudSignalComponent from './components/ViewFraudSignalComponent';
function App() {
  return (
    <div>
        <Router>
                <HeaderComponent className="header"/>
                <div className="container">
                    <Switch>
                          <Route path = "/" exact component = {HomePageComponent}></Route>
                            <Route path = "/analyticsWorkspaces" component = {ListAnalyticsWorkspaceComponent}></Route>
                            <Route path = "/add-analyticsWorkspace/:id" component = {CreateAnalyticsWorkspaceComponent}></Route>
                            <Route path = "/view-analyticsWorkspace/:id" component = {ViewAnalyticsWorkspaceComponent}></Route>
                          {/* <Route path = "/update-analyticsWorkspace/:id" component = {UpdateAnalyticsWorkspaceComponent}></Route> */}
                            <Route path = "/dataSources" component = {ListDataSourceComponent}></Route>
                            <Route path = "/add-dataSource/:id" component = {CreateDataSourceComponent}></Route>
                            <Route path = "/view-dataSource/:id" component = {ViewDataSourceComponent}></Route>
                          {/* <Route path = "/update-dataSource/:id" component = {UpdateDataSourceComponent}></Route> */}
                            <Route path = "/dataSets" component = {ListDataSetComponent}></Route>
                            <Route path = "/add-dataSet/:id" component = {CreateDataSetComponent}></Route>
                            <Route path = "/view-dataSet/:id" component = {ViewDataSetComponent}></Route>
                          {/* <Route path = "/update-dataSet/:id" component = {UpdateDataSetComponent}></Route> */}
                            <Route path = "/dataPipelines" component = {ListDataPipelineComponent}></Route>
                            <Route path = "/add-dataPipeline/:id" component = {CreateDataPipelineComponent}></Route>
                            <Route path = "/view-dataPipeline/:id" component = {ViewDataPipelineComponent}></Route>
                          {/* <Route path = "/update-dataPipeline/:id" component = {UpdateDataPipelineComponent}></Route> */}
                            <Route path = "/dataTasks" component = {ListDataTaskComponent}></Route>
                            <Route path = "/add-dataTask/:id" component = {CreateDataTaskComponent}></Route>
                            <Route path = "/view-dataTask/:id" component = {ViewDataTaskComponent}></Route>
                          {/* <Route path = "/update-dataTask/:id" component = {UpdateDataTaskComponent}></Route> */}
                            <Route path = "/semanticModels" component = {ListSemanticModelComponent}></Route>
                            <Route path = "/add-semanticModel/:id" component = {CreateSemanticModelComponent}></Route>
                            <Route path = "/view-semanticModel/:id" component = {ViewSemanticModelComponent}></Route>
                          {/* <Route path = "/update-semanticModel/:id" component = {UpdateSemanticModelComponent}></Route> */}
                            <Route path = "/dimensions" component = {ListDimensionComponent}></Route>
                            <Route path = "/add-dimension/:id" component = {CreateDimensionComponent}></Route>
                            <Route path = "/view-dimension/:id" component = {ViewDimensionComponent}></Route>
                          {/* <Route path = "/update-dimension/:id" component = {UpdateDimensionComponent}></Route> */}
                            <Route path = "/measures" component = {ListMeasureComponent}></Route>
                            <Route path = "/add-measure/:id" component = {CreateMeasureComponent}></Route>
                            <Route path = "/view-measure/:id" component = {ViewMeasureComponent}></Route>
                          {/* <Route path = "/update-measure/:id" component = {UpdateMeasureComponent}></Route> */}
                            <Route path = "/metrics" component = {ListMetricComponent}></Route>
                            <Route path = "/add-metric/:id" component = {CreateMetricComponent}></Route>
                            <Route path = "/view-metric/:id" component = {ViewMetricComponent}></Route>
                          {/* <Route path = "/update-metric/:id" component = {UpdateMetricComponent}></Route> */}
                            <Route path = "/reports" component = {ListReportComponent}></Route>
                            <Route path = "/add-report/:id" component = {CreateReportComponent}></Route>
                            <Route path = "/view-report/:id" component = {ViewReportComponent}></Route>
                          {/* <Route path = "/update-report/:id" component = {UpdateReportComponent}></Route> */}
                            <Route path = "/dashboards" component = {ListDashboardComponent}></Route>
                            <Route path = "/add-dashboard/:id" component = {CreateDashboardComponent}></Route>
                            <Route path = "/view-dashboard/:id" component = {ViewDashboardComponent}></Route>
                          {/* <Route path = "/update-dashboard/:id" component = {UpdateDashboardComponent}></Route> */}
                            <Route path = "/visualizations" component = {ListVisualizationComponent}></Route>
                            <Route path = "/add-visualization/:id" component = {CreateVisualizationComponent}></Route>
                            <Route path = "/view-visualization/:id" component = {ViewVisualizationComponent}></Route>
                          {/* <Route path = "/update-visualization/:id" component = {UpdateVisualizationComponent}></Route> */}
                            <Route path = "/notebooks" component = {ListNotebookComponent}></Route>
                            <Route path = "/add-notebook/:id" component = {CreateNotebookComponent}></Route>
                            <Route path = "/view-notebook/:id" component = {ViewNotebookComponent}></Route>
                          {/* <Route path = "/update-notebook/:id" component = {UpdateNotebookComponent}></Route> */}
                            <Route path = "/bIQuerys" component = {ListBIQueryComponent}></Route>
                            <Route path = "/add-bIQuery/:id" component = {CreateBIQueryComponent}></Route>
                            <Route path = "/view-bIQuery/:id" component = {ViewBIQueryComponent}></Route>
                          {/* <Route path = "/update-bIQuery/:id" component = {UpdateBIQueryComponent}></Route> */}
                            <Route path = "/experiments" component = {ListExperimentComponent}></Route>
                            <Route path = "/add-experiment/:id" component = {CreateExperimentComponent}></Route>
                            <Route path = "/view-experiment/:id" component = {ViewExperimentComponent}></Route>
                          {/* <Route path = "/update-experiment/:id" component = {UpdateExperimentComponent}></Route> */}
                            <Route path = "/trainingRuns" component = {ListTrainingRunComponent}></Route>
                            <Route path = "/add-trainingRun/:id" component = {CreateTrainingRunComponent}></Route>
                            <Route path = "/view-trainingRun/:id" component = {ViewTrainingRunComponent}></Route>
                          {/* <Route path = "/update-trainingRun/:id" component = {UpdateTrainingRunComponent}></Route> */}
                            <Route path = "/runMetrics" component = {ListRunMetricComponent}></Route>
                            <Route path = "/add-runMetric/:id" component = {CreateRunMetricComponent}></Route>
                            <Route path = "/view-runMetric/:id" component = {ViewRunMetricComponent}></Route>
                          {/* <Route path = "/update-runMetric/:id" component = {UpdateRunMetricComponent}></Route> */}
                            <Route path = "/runParameters" component = {ListRunParameterComponent}></Route>
                            <Route path = "/add-runParameter/:id" component = {CreateRunParameterComponent}></Route>
                            <Route path = "/view-runParameter/:id" component = {ViewRunParameterComponent}></Route>
                          {/* <Route path = "/update-runParameter/:id" component = {UpdateRunParameterComponent}></Route> */}
                            <Route path = "/models" component = {ListModelComponent}></Route>
                            <Route path = "/add-model/:id" component = {CreateModelComponent}></Route>
                            <Route path = "/view-model/:id" component = {ViewModelComponent}></Route>
                          {/* <Route path = "/update-model/:id" component = {UpdateModelComponent}></Route> */}
                            <Route path = "/modelVersions" component = {ListModelVersionComponent}></Route>
                            <Route path = "/add-modelVersion/:id" component = {CreateModelVersionComponent}></Route>
                            <Route path = "/view-modelVersion/:id" component = {ViewModelVersionComponent}></Route>
                          {/* <Route path = "/update-modelVersion/:id" component = {UpdateModelVersionComponent}></Route> */}
                            <Route path = "/evaluationMetrics" component = {ListEvaluationMetricComponent}></Route>
                            <Route path = "/add-evaluationMetric/:id" component = {CreateEvaluationMetricComponent}></Route>
                            <Route path = "/view-evaluationMetric/:id" component = {ViewEvaluationMetricComponent}></Route>
                          {/* <Route path = "/update-evaluationMetric/:id" component = {UpdateEvaluationMetricComponent}></Route> */}
                            <Route path = "/featureSets" component = {ListFeatureSetComponent}></Route>
                            <Route path = "/add-featureSet/:id" component = {CreateFeatureSetComponent}></Route>
                            <Route path = "/view-featureSet/:id" component = {ViewFeatureSetComponent}></Route>
                          {/* <Route path = "/update-featureSet/:id" component = {UpdateFeatureSetComponent}></Route> */}
                            <Route path = "/features" component = {ListFeatureComponent}></Route>
                            <Route path = "/add-feature/:id" component = {CreateFeatureComponent}></Route>
                            <Route path = "/view-feature/:id" component = {ViewFeatureComponent}></Route>
                          {/* <Route path = "/update-feature/:id" component = {UpdateFeatureComponent}></Route> */}
                            <Route path = "/inferenceEndpoints" component = {ListInferenceEndpointComponent}></Route>
                            <Route path = "/add-inferenceEndpoint/:id" component = {CreateInferenceEndpointComponent}></Route>
                            <Route path = "/view-inferenceEndpoint/:id" component = {ViewInferenceEndpointComponent}></Route>
                          {/* <Route path = "/update-inferenceEndpoint/:id" component = {UpdateInferenceEndpointComponent}></Route> */}
                            <Route path = "/predictions" component = {ListPredictionComponent}></Route>
                            <Route path = "/add-prediction/:id" component = {CreatePredictionComponent}></Route>
                            <Route path = "/view-prediction/:id" component = {ViewPredictionComponent}></Route>
                          {/* <Route path = "/update-prediction/:id" component = {UpdatePredictionComponent}></Route> */}
                            <Route path = "/forecasts" component = {ListForecastComponent}></Route>
                            <Route path = "/add-forecast/:id" component = {CreateForecastComponent}></Route>
                            <Route path = "/view-forecast/:id" component = {ViewForecastComponent}></Route>
                          {/* <Route path = "/update-forecast/:id" component = {UpdateForecastComponent}></Route> */}
                            <Route path = "/timeSeriess" component = {ListTimeSeriesComponent}></Route>
                            <Route path = "/add-timeSeries/:id" component = {CreateTimeSeriesComponent}></Route>
                            <Route path = "/view-timeSeries/:id" component = {ViewTimeSeriesComponent}></Route>
                          {/* <Route path = "/update-timeSeries/:id" component = {UpdateTimeSeriesComponent}></Route> */}
                            <Route path = "/anomalys" component = {ListAnomalyComponent}></Route>
                            <Route path = "/add-anomaly/:id" component = {CreateAnomalyComponent}></Route>
                            <Route path = "/view-anomaly/:id" component = {ViewAnomalyComponent}></Route>
                          {/* <Route path = "/update-anomaly/:id" component = {UpdateAnomalyComponent}></Route> */}
                            <Route path = "/qualityRules" component = {ListQualityRuleComponent}></Route>
                            <Route path = "/add-qualityRule/:id" component = {CreateQualityRuleComponent}></Route>
                            <Route path = "/view-qualityRule/:id" component = {ViewQualityRuleComponent}></Route>
                          {/* <Route path = "/update-qualityRule/:id" component = {UpdateQualityRuleComponent}></Route> */}
                            <Route path = "/qualityChecks" component = {ListQualityCheckComponent}></Route>
                            <Route path = "/add-qualityCheck/:id" component = {CreateQualityCheckComponent}></Route>
                            <Route path = "/view-qualityCheck/:id" component = {ViewQualityCheckComponent}></Route>
                          {/* <Route path = "/update-qualityCheck/:id" component = {UpdateQualityCheckComponent}></Route> */}
                            <Route path = "/lineageNodes" component = {ListLineageNodeComponent}></Route>
                            <Route path = "/add-lineageNode/:id" component = {CreateLineageNodeComponent}></Route>
                            <Route path = "/view-lineageNode/:id" component = {ViewLineageNodeComponent}></Route>
                          {/* <Route path = "/update-lineageNode/:id" component = {UpdateLineageNodeComponent}></Route> */}
                            <Route path = "/tags" component = {ListTagComponent}></Route>
                            <Route path = "/add-tag/:id" component = {CreateTagComponent}></Route>
                            <Route path = "/view-tag/:id" component = {ViewTagComponent}></Route>
                          {/* <Route path = "/update-tag/:id" component = {UpdateTagComponent}></Route> */}
                            <Route path = "/accessPolicys" component = {ListAccessPolicyComponent}></Route>
                            <Route path = "/add-accessPolicy/:id" component = {CreateAccessPolicyComponent}></Route>
                            <Route path = "/view-accessPolicy/:id" component = {ViewAccessPolicyComponent}></Route>
                          {/* <Route path = "/update-accessPolicy/:id" component = {UpdateAccessPolicyComponent}></Route> */}
                            <Route path = "/alerts" component = {ListAlertComponent}></Route>
                            <Route path = "/add-alert/:id" component = {CreateAlertComponent}></Route>
                            <Route path = "/view-alert/:id" component = {ViewAlertComponent}></Route>
                          {/* <Route path = "/update-alert/:id" component = {UpdateAlertComponent}></Route> */}
                            <Route path = "/subscribers" component = {ListSubscriberComponent}></Route>
                            <Route path = "/add-subscriber/:id" component = {CreateSubscriberComponent}></Route>
                            <Route path = "/view-subscriber/:id" component = {ViewSubscriberComponent}></Route>
                          {/* <Route path = "/update-subscriber/:id" component = {UpdateSubscriberComponent}></Route> */}
                            <Route path = "/businessGlossaryTerms" component = {ListBusinessGlossaryTermComponent}></Route>
                            <Route path = "/add-businessGlossaryTerm/:id" component = {CreateBusinessGlossaryTermComponent}></Route>
                            <Route path = "/view-businessGlossaryTerm/:id" component = {ViewBusinessGlossaryTermComponent}></Route>
                          {/* <Route path = "/update-businessGlossaryTerm/:id" component = {UpdateBusinessGlossaryTermComponent}></Route> */}
                            <Route path = "/recommendationScenarios" component = {ListRecommendationScenarioComponent}></Route>
                            <Route path = "/add-recommendationScenario/:id" component = {CreateRecommendationScenarioComponent}></Route>
                            <Route path = "/view-recommendationScenario/:id" component = {ViewRecommendationScenarioComponent}></Route>
                          {/* <Route path = "/update-recommendationScenario/:id" component = {UpdateRecommendationScenarioComponent}></Route> */}
                            <Route path = "/fraudScenarios" component = {ListFraudScenarioComponent}></Route>
                            <Route path = "/add-fraudScenario/:id" component = {CreateFraudScenarioComponent}></Route>
                            <Route path = "/view-fraudScenario/:id" component = {ViewFraudScenarioComponent}></Route>
                          {/* <Route path = "/update-fraudScenario/:id" component = {UpdateFraudScenarioComponent}></Route> */}
                            <Route path = "/fraudSignals" component = {ListFraudSignalComponent}></Route>
                            <Route path = "/add-fraudSignal/:id" component = {CreateFraudSignalComponent}></Route>
                            <Route path = "/view-fraudSignal/:id" component = {ViewFraudSignalComponent}></Route>
                          {/* <Route path = "/update-fraudSignal/:id" component = {UpdateFraudSignalComponent}></Route> */}
                    </Switch>
                </div>
              <FooterComponent />
        </Router>
    </div>
    
  );
}

export default App;
