import {BrowserModule} from '@angular/platform-browser';
import {BrowserAnimationsModule} from '@angular/platform-browser/animations';
import {MatInputModule} from '@angular/material/input';
import {MatDatepickerModule} from '@angular/material/datepicker';
import {MatCheckboxModule} from '@angular/material/checkbox';
import {MatButtonModule} from '@angular/material/button';
import {MatFormFieldModule} from '@angular/material/form-field';
import {MatSelectModule} from '@angular/material/select';
import {MatMomentDateModule} from "@angular/material-moment-adapter";
import {NgModule} from '@angular/core';
import {NgbModule} from '@ng-bootstrap/ng-bootstrap';
import {RouterModule} from '@angular/router';
import {HttpClientModule} from '@angular/common/http';
import {FormsModule} from '@angular/forms';
import {ReactiveFormsModule} from '@angular/forms';
import {AppComponent} from './app.component';
import {MatMenuModule} from '@angular/material/menu';
import {MatToolbarModule} from '@angular/material/toolbar';
import {MatSidenavModule} from '@angular/material/sidenav'

import {IndexAnalyticsWorkspaceComponent} from './components/AnalyticsWorkspace/index/index.component';
import {CreateAnalyticsWorkspaceComponent} from './components/AnalyticsWorkspace/create/create.component';
import {EditAnalyticsWorkspaceComponent} from './components/AnalyticsWorkspace/edit/edit.component';
import {IndexDataSourceComponent} from './components/DataSource/index/index.component';
import {CreateDataSourceComponent} from './components/DataSource/create/create.component';
import {EditDataSourceComponent} from './components/DataSource/edit/edit.component';
import {IndexDataSetComponent} from './components/DataSet/index/index.component';
import {CreateDataSetComponent} from './components/DataSet/create/create.component';
import {EditDataSetComponent} from './components/DataSet/edit/edit.component';
import {IndexDataPipelineComponent} from './components/DataPipeline/index/index.component';
import {CreateDataPipelineComponent} from './components/DataPipeline/create/create.component';
import {EditDataPipelineComponent} from './components/DataPipeline/edit/edit.component';
import {IndexDataTaskComponent} from './components/DataTask/index/index.component';
import {CreateDataTaskComponent} from './components/DataTask/create/create.component';
import {EditDataTaskComponent} from './components/DataTask/edit/edit.component';
import {IndexSemanticModelComponent} from './components/SemanticModel/index/index.component';
import {CreateSemanticModelComponent} from './components/SemanticModel/create/create.component';
import {EditSemanticModelComponent} from './components/SemanticModel/edit/edit.component';
import {IndexDimensionComponent} from './components/Dimension/index/index.component';
import {CreateDimensionComponent} from './components/Dimension/create/create.component';
import {EditDimensionComponent} from './components/Dimension/edit/edit.component';
import {IndexMeasureComponent} from './components/Measure/index/index.component';
import {CreateMeasureComponent} from './components/Measure/create/create.component';
import {EditMeasureComponent} from './components/Measure/edit/edit.component';
import {IndexMetricComponent} from './components/Metric/index/index.component';
import {CreateMetricComponent} from './components/Metric/create/create.component';
import {EditMetricComponent} from './components/Metric/edit/edit.component';
import {IndexReportComponent} from './components/Report/index/index.component';
import {CreateReportComponent} from './components/Report/create/create.component';
import {EditReportComponent} from './components/Report/edit/edit.component';
import {IndexDashboardComponent} from './components/Dashboard/index/index.component';
import {CreateDashboardComponent} from './components/Dashboard/create/create.component';
import {EditDashboardComponent} from './components/Dashboard/edit/edit.component';
import {IndexVisualizationComponent} from './components/Visualization/index/index.component';
import {CreateVisualizationComponent} from './components/Visualization/create/create.component';
import {EditVisualizationComponent} from './components/Visualization/edit/edit.component';
import {IndexNotebookComponent} from './components/Notebook/index/index.component';
import {CreateNotebookComponent} from './components/Notebook/create/create.component';
import {EditNotebookComponent} from './components/Notebook/edit/edit.component';
import {IndexBIQueryComponent} from './components/BIQuery/index/index.component';
import {CreateBIQueryComponent} from './components/BIQuery/create/create.component';
import {EditBIQueryComponent} from './components/BIQuery/edit/edit.component';
import {IndexExperimentComponent} from './components/Experiment/index/index.component';
import {CreateExperimentComponent} from './components/Experiment/create/create.component';
import {EditExperimentComponent} from './components/Experiment/edit/edit.component';
import {IndexTrainingRunComponent} from './components/TrainingRun/index/index.component';
import {CreateTrainingRunComponent} from './components/TrainingRun/create/create.component';
import {EditTrainingRunComponent} from './components/TrainingRun/edit/edit.component';
import {IndexRunMetricComponent} from './components/RunMetric/index/index.component';
import {CreateRunMetricComponent} from './components/RunMetric/create/create.component';
import {EditRunMetricComponent} from './components/RunMetric/edit/edit.component';
import {IndexRunParameterComponent} from './components/RunParameter/index/index.component';
import {CreateRunParameterComponent} from './components/RunParameter/create/create.component';
import {EditRunParameterComponent} from './components/RunParameter/edit/edit.component';
import {IndexModel_Component} from './components/Model_/index/index.component';
import {CreateModel_Component} from './components/Model_/create/create.component';
import {EditModel_Component} from './components/Model_/edit/edit.component';
import {IndexModelVersionComponent} from './components/ModelVersion/index/index.component';
import {CreateModelVersionComponent} from './components/ModelVersion/create/create.component';
import {EditModelVersionComponent} from './components/ModelVersion/edit/edit.component';
import {IndexEvaluationMetricComponent} from './components/EvaluationMetric/index/index.component';
import {CreateEvaluationMetricComponent} from './components/EvaluationMetric/create/create.component';
import {EditEvaluationMetricComponent} from './components/EvaluationMetric/edit/edit.component';
import {IndexFeatureSetComponent} from './components/FeatureSet/index/index.component';
import {CreateFeatureSetComponent} from './components/FeatureSet/create/create.component';
import {EditFeatureSetComponent} from './components/FeatureSet/edit/edit.component';
import {IndexFeatureComponent} from './components/Feature/index/index.component';
import {CreateFeatureComponent} from './components/Feature/create/create.component';
import {EditFeatureComponent} from './components/Feature/edit/edit.component';
import {IndexInferenceEndpointComponent} from './components/InferenceEndpoint/index/index.component';
import {CreateInferenceEndpointComponent} from './components/InferenceEndpoint/create/create.component';
import {EditInferenceEndpointComponent} from './components/InferenceEndpoint/edit/edit.component';
import {IndexPredictionComponent} from './components/Prediction/index/index.component';
import {CreatePredictionComponent} from './components/Prediction/create/create.component';
import {EditPredictionComponent} from './components/Prediction/edit/edit.component';
import {IndexForecastComponent} from './components/Forecast/index/index.component';
import {CreateForecastComponent} from './components/Forecast/create/create.component';
import {EditForecastComponent} from './components/Forecast/edit/edit.component';
import {IndexTimeSeriesComponent} from './components/TimeSeries/index/index.component';
import {CreateTimeSeriesComponent} from './components/TimeSeries/create/create.component';
import {EditTimeSeriesComponent} from './components/TimeSeries/edit/edit.component';
import {IndexAnomalyComponent} from './components/Anomaly/index/index.component';
import {CreateAnomalyComponent} from './components/Anomaly/create/create.component';
import {EditAnomalyComponent} from './components/Anomaly/edit/edit.component';
import {IndexQualityRuleComponent} from './components/QualityRule/index/index.component';
import {CreateQualityRuleComponent} from './components/QualityRule/create/create.component';
import {EditQualityRuleComponent} from './components/QualityRule/edit/edit.component';
import {IndexQualityCheckComponent} from './components/QualityCheck/index/index.component';
import {CreateQualityCheckComponent} from './components/QualityCheck/create/create.component';
import {EditQualityCheckComponent} from './components/QualityCheck/edit/edit.component';
import {IndexLineageNodeComponent} from './components/LineageNode/index/index.component';
import {CreateLineageNodeComponent} from './components/LineageNode/create/create.component';
import {EditLineageNodeComponent} from './components/LineageNode/edit/edit.component';
import {IndexTagComponent} from './components/Tag/index/index.component';
import {CreateTagComponent} from './components/Tag/create/create.component';
import {EditTagComponent} from './components/Tag/edit/edit.component';
import {IndexAccessPolicyComponent} from './components/AccessPolicy/index/index.component';
import {CreateAccessPolicyComponent} from './components/AccessPolicy/create/create.component';
import {EditAccessPolicyComponent} from './components/AccessPolicy/edit/edit.component';
import {IndexAlertComponent} from './components/Alert/index/index.component';
import {CreateAlertComponent} from './components/Alert/create/create.component';
import {EditAlertComponent} from './components/Alert/edit/edit.component';
import {IndexSubscriberComponent} from './components/Subscriber/index/index.component';
import {CreateSubscriberComponent} from './components/Subscriber/create/create.component';
import {EditSubscriberComponent} from './components/Subscriber/edit/edit.component';
import {IndexBusinessGlossaryTermComponent} from './components/BusinessGlossaryTerm/index/index.component';
import {CreateBusinessGlossaryTermComponent} from './components/BusinessGlossaryTerm/create/create.component';
import {EditBusinessGlossaryTermComponent} from './components/BusinessGlossaryTerm/edit/edit.component';
import {IndexRecommendationScenarioComponent} from './components/RecommendationScenario/index/index.component';
import {CreateRecommendationScenarioComponent} from './components/RecommendationScenario/create/create.component';
import {EditRecommendationScenarioComponent} from './components/RecommendationScenario/edit/edit.component';
import {IndexFraudScenarioComponent} from './components/FraudScenario/index/index.component';
import {CreateFraudScenarioComponent} from './components/FraudScenario/create/create.component';
import {EditFraudScenarioComponent} from './components/FraudScenario/edit/edit.component';
import {IndexFraudSignalComponent} from './components/FraudSignal/index/index.component';
import {CreateFraudSignalComponent} from './components/FraudSignal/create/create.component';
import {EditFraudSignalComponent} from './components/FraudSignal/edit/edit.component';

import * as appRoutes from './routerConfig';

import {AnalyticsWorkspaceService} from './services/AnalyticsWorkspace.service';
import {DataSourceService} from './services/DataSource.service';
import {DataSetService} from './services/DataSet.service';
import {DataPipelineService} from './services/DataPipeline.service';
import {DataTaskService} from './services/DataTask.service';
import {SemanticModelService} from './services/SemanticModel.service';
import {DimensionService} from './services/Dimension.service';
import {MeasureService} from './services/Measure.service';
import {MetricService} from './services/Metric.service';
import {ReportService} from './services/Report.service';
import {DashboardService} from './services/Dashboard.service';
import {VisualizationService} from './services/Visualization.service';
import {NotebookService} from './services/Notebook.service';
import {BIQueryService} from './services/BIQuery.service';
import {ExperimentService} from './services/Experiment.service';
import {TrainingRunService} from './services/TrainingRun.service';
import {RunMetricService} from './services/RunMetric.service';
import {RunParameterService} from './services/RunParameter.service';
import {Model_Service} from './services/Model_.service';
import {ModelVersionService} from './services/ModelVersion.service';
import {EvaluationMetricService} from './services/EvaluationMetric.service';
import {FeatureSetService} from './services/FeatureSet.service';
import {FeatureService} from './services/Feature.service';
import {InferenceEndpointService} from './services/InferenceEndpoint.service';
import {PredictionService} from './services/Prediction.service';
import {ForecastService} from './services/Forecast.service';
import {TimeSeriesService} from './services/TimeSeries.service';
import {AnomalyService} from './services/Anomaly.service';
import {QualityRuleService} from './services/QualityRule.service';
import {QualityCheckService} from './services/QualityCheck.service';
import {LineageNodeService} from './services/LineageNode.service';
import {TagService} from './services/Tag.service';
import {AccessPolicyService} from './services/AccessPolicy.service';
import {AlertService} from './services/Alert.service';
import {SubscriberService} from './services/Subscriber.service';
import {BusinessGlossaryTermService} from './services/BusinessGlossaryTerm.service';
import {RecommendationScenarioService} from './services/RecommendationScenario.service';
import {FraudScenarioService} from './services/FraudScenario.service';
import {FraudSignalService} from './services/FraudSignal.service';

@NgModule({
  declarations: [
    IndexAnalyticsWorkspaceComponent,
    CreateAnalyticsWorkspaceComponent,
    EditAnalyticsWorkspaceComponent,
    IndexDataSourceComponent,
    CreateDataSourceComponent,
    EditDataSourceComponent,
    IndexDataSetComponent,
    CreateDataSetComponent,
    EditDataSetComponent,
    IndexDataPipelineComponent,
    CreateDataPipelineComponent,
    EditDataPipelineComponent,
    IndexDataTaskComponent,
    CreateDataTaskComponent,
    EditDataTaskComponent,
    IndexSemanticModelComponent,
    CreateSemanticModelComponent,
    EditSemanticModelComponent,
    IndexDimensionComponent,
    CreateDimensionComponent,
    EditDimensionComponent,
    IndexMeasureComponent,
    CreateMeasureComponent,
    EditMeasureComponent,
    IndexMetricComponent,
    CreateMetricComponent,
    EditMetricComponent,
    IndexReportComponent,
    CreateReportComponent,
    EditReportComponent,
    IndexDashboardComponent,
    CreateDashboardComponent,
    EditDashboardComponent,
    IndexVisualizationComponent,
    CreateVisualizationComponent,
    EditVisualizationComponent,
    IndexNotebookComponent,
    CreateNotebookComponent,
    EditNotebookComponent,
    IndexBIQueryComponent,
    CreateBIQueryComponent,
    EditBIQueryComponent,
    IndexExperimentComponent,
    CreateExperimentComponent,
    EditExperimentComponent,
    IndexTrainingRunComponent,
    CreateTrainingRunComponent,
    EditTrainingRunComponent,
    IndexRunMetricComponent,
    CreateRunMetricComponent,
    EditRunMetricComponent,
    IndexRunParameterComponent,
    CreateRunParameterComponent,
    EditRunParameterComponent,
    IndexModel_Component,
    CreateModel_Component,
    EditModel_Component,
    IndexModelVersionComponent,
    CreateModelVersionComponent,
    EditModelVersionComponent,
    IndexEvaluationMetricComponent,
    CreateEvaluationMetricComponent,
    EditEvaluationMetricComponent,
    IndexFeatureSetComponent,
    CreateFeatureSetComponent,
    EditFeatureSetComponent,
    IndexFeatureComponent,
    CreateFeatureComponent,
    EditFeatureComponent,
    IndexInferenceEndpointComponent,
    CreateInferenceEndpointComponent,
    EditInferenceEndpointComponent,
    IndexPredictionComponent,
    CreatePredictionComponent,
    EditPredictionComponent,
    IndexForecastComponent,
    CreateForecastComponent,
    EditForecastComponent,
    IndexTimeSeriesComponent,
    CreateTimeSeriesComponent,
    EditTimeSeriesComponent,
    IndexAnomalyComponent,
    CreateAnomalyComponent,
    EditAnomalyComponent,
    IndexQualityRuleComponent,
    CreateQualityRuleComponent,
    EditQualityRuleComponent,
    IndexQualityCheckComponent,
    CreateQualityCheckComponent,
    EditQualityCheckComponent,
    IndexLineageNodeComponent,
    CreateLineageNodeComponent,
    EditLineageNodeComponent,
    IndexTagComponent,
    CreateTagComponent,
    EditTagComponent,
    IndexAccessPolicyComponent,
    CreateAccessPolicyComponent,
    EditAccessPolicyComponent,
    IndexAlertComponent,
    CreateAlertComponent,
    EditAlertComponent,
    IndexSubscriberComponent,
    CreateSubscriberComponent,
    EditSubscriberComponent,
    IndexBusinessGlossaryTermComponent,
    CreateBusinessGlossaryTermComponent,
    EditBusinessGlossaryTermComponent,
    IndexRecommendationScenarioComponent,
    CreateRecommendationScenarioComponent,
    EditRecommendationScenarioComponent,
    IndexFraudScenarioComponent,
    CreateFraudScenarioComponent,
    EditFraudScenarioComponent,
    IndexFraudSignalComponent,
    CreateFraudSignalComponent,
    EditFraudSignalComponent,
    AppComponent
  ],
  imports: [

    BrowserModule, 
    NgbModule,
    MatMenuModule,
    MatToolbarModule,
    MatCheckboxModule,
    MatButtonModule,
    MatFormFieldModule,
    MatInputModule,
    MatSelectModule,
    MatDatepickerModule,
	MatMomentDateModule,
    BrowserAnimationsModule,
	HttpClientModule, 
    ReactiveFormsModule,
    FormsModule,
    MatSidenavModule,    
    RouterModule.forRoot(appRoutes.AnalyticsWorkspaceRoutes), 
    RouterModule.forRoot(appRoutes.DataSourceRoutes), 
    RouterModule.forRoot(appRoutes.DataSetRoutes), 
    RouterModule.forRoot(appRoutes.DataPipelineRoutes), 
    RouterModule.forRoot(appRoutes.DataTaskRoutes), 
    RouterModule.forRoot(appRoutes.SemanticModelRoutes), 
    RouterModule.forRoot(appRoutes.DimensionRoutes), 
    RouterModule.forRoot(appRoutes.MeasureRoutes), 
    RouterModule.forRoot(appRoutes.MetricRoutes), 
    RouterModule.forRoot(appRoutes.ReportRoutes), 
    RouterModule.forRoot(appRoutes.DashboardRoutes), 
    RouterModule.forRoot(appRoutes.VisualizationRoutes), 
    RouterModule.forRoot(appRoutes.NotebookRoutes), 
    RouterModule.forRoot(appRoutes.BIQueryRoutes), 
    RouterModule.forRoot(appRoutes.ExperimentRoutes), 
    RouterModule.forRoot(appRoutes.TrainingRunRoutes), 
    RouterModule.forRoot(appRoutes.RunMetricRoutes), 
    RouterModule.forRoot(appRoutes.RunParameterRoutes), 
    RouterModule.forRoot(appRoutes.Model_Routes), 
    RouterModule.forRoot(appRoutes.ModelVersionRoutes), 
    RouterModule.forRoot(appRoutes.EvaluationMetricRoutes), 
    RouterModule.forRoot(appRoutes.FeatureSetRoutes), 
    RouterModule.forRoot(appRoutes.FeatureRoutes), 
    RouterModule.forRoot(appRoutes.InferenceEndpointRoutes), 
    RouterModule.forRoot(appRoutes.PredictionRoutes), 
    RouterModule.forRoot(appRoutes.ForecastRoutes), 
    RouterModule.forRoot(appRoutes.TimeSeriesRoutes), 
    RouterModule.forRoot(appRoutes.AnomalyRoutes), 
    RouterModule.forRoot(appRoutes.QualityRuleRoutes), 
    RouterModule.forRoot(appRoutes.QualityCheckRoutes), 
    RouterModule.forRoot(appRoutes.LineageNodeRoutes), 
    RouterModule.forRoot(appRoutes.TagRoutes), 
    RouterModule.forRoot(appRoutes.AccessPolicyRoutes), 
    RouterModule.forRoot(appRoutes.AlertRoutes), 
    RouterModule.forRoot(appRoutes.SubscriberRoutes), 
    RouterModule.forRoot(appRoutes.BusinessGlossaryTermRoutes), 
    RouterModule.forRoot(appRoutes.RecommendationScenarioRoutes), 
    RouterModule.forRoot(appRoutes.FraudScenarioRoutes), 
    RouterModule.forRoot(appRoutes.FraudSignalRoutes), 
  ],
  providers: [AnalyticsWorkspaceService,DataSourceService,DataSetService,DataPipelineService,DataTaskService,SemanticModelService,DimensionService,MeasureService,MetricService,ReportService,DashboardService,VisualizationService,NotebookService,BIQueryService,ExperimentService,TrainingRunService,RunMetricService,RunParameterService,Model_Service,ModelVersionService,EvaluationMetricService,FeatureSetService,FeatureService,InferenceEndpointService,PredictionService,ForecastService,TimeSeriesService,AnomalyService,QualityRuleService,QualityCheckService,LineageNodeService,TagService,AccessPolicyService,AlertService,SubscriberService,BusinessGlossaryTermService,RecommendationScenarioService,FraudScenarioService,FraudSignalService],
  bootstrap: [AppComponent]
})
export class AppModule { }
