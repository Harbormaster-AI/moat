import { HttpClient } from '@angular/common/http';
import * as enumTypes from '../models/EnumTypes';

import {AnalyticsWorkspaceService} from '../services/AnalyticsWorkspace.service';
import {DataSourceService} from '../services/DataSource.service';
import {DataSetService} from '../services/DataSet.service';
import {DataPipelineService} from '../services/DataPipeline.service';
import {DataTaskService} from '../services/DataTask.service';
import {SemanticModelService} from '../services/SemanticModel.service';
import {DimensionService} from '../services/Dimension.service';
import {MeasureService} from '../services/Measure.service';
import {MetricService} from '../services/Metric.service';
import {ReportService} from '../services/Report.service';
import {DashboardService} from '../services/Dashboard.service';
import {VisualizationService} from '../services/Visualization.service';
import {NotebookService} from '../services/Notebook.service';
import {BIQueryService} from '../services/BIQuery.service';
import {ExperimentService} from '../services/Experiment.service';
import {TrainingRunService} from '../services/TrainingRun.service';
import {RunMetricService} from '../services/RunMetric.service';
import {RunParameterService} from '../services/RunParameter.service';
import {Model_Service} from '../services/Model_.service';
import {ModelVersionService} from '../services/ModelVersion.service';
import {EvaluationMetricService} from '../services/EvaluationMetric.service';
import {FeatureSetService} from '../services/FeatureSet.service';
import {FeatureService} from '../services/Feature.service';
import {InferenceEndpointService} from '../services/InferenceEndpoint.service';
import {PredictionService} from '../services/Prediction.service';
import {ForecastService} from '../services/Forecast.service';
import {TimeSeriesService} from '../services/TimeSeries.service';
import {AnomalyService} from '../services/Anomaly.service';
import {QualityRuleService} from '../services/QualityRule.service';
import {QualityCheckService} from '../services/QualityCheck.service';
import {LineageNodeService} from '../services/LineageNode.service';
import {TagService} from '../services/Tag.service';
import {AccessPolicyService} from '../services/AccessPolicy.service';
import {AlertService} from '../services/Alert.service';
import {SubscriberService} from '../services/Subscriber.service';
import {BusinessGlossaryTermService} from '../services/BusinessGlossaryTerm.service';
import {RecommendationScenarioService} from '../services/RecommendationScenario.service';
import {FraudScenarioService} from '../services/FraudScenario.service';
import {FraudSignalService} from '../services/FraudSignal.service';

import { Directive } from '@angular/core';

/**
 Base class of all Components.
 For convenience, contains all enums and entity lists
 **/

@Directive()
export class BaseComponent {

    constructor (private http: HttpClient) {}

// enum instances
    DataSourceTypes = Object.keys(enumTypes.DataSourceType);
    DataFormats = Object.keys(enumTypes.DataFormat);
    PipelineTriggerTypes = Object.keys(enumTypes.PipelineTriggerType);
    PipelineStatuss = Object.keys(enumTypes.PipelineStatus);
    DataTaskTypes = Object.keys(enumTypes.DataTaskType);
    DimensionTypes = Object.keys(enumTypes.DimensionType);
    AggregationTypes = Object.keys(enumTypes.AggregationType);
    MetricTypes = Object.keys(enumTypes.MetricType);
    ReportStatuss = Object.keys(enumTypes.ReportStatus);
    DashboardStatuss = Object.keys(enumTypes.DashboardStatus);
    ChartTypes = Object.keys(enumTypes.ChartType);
    NotebookLanguages = Object.keys(enumTypes.NotebookLanguage);
    ExperimentStatuss = Object.keys(enumTypes.ExperimentStatus);
    TrainingStatuss = Object.keys(enumTypes.TrainingStatus);
    ModelTypes = Object.keys(enumTypes.ModelType);
    ModelLifecycles = Object.keys(enumTypes.ModelLifecycle);
    InferenceModes = Object.keys(enumTypes.InferenceMode);
    TimeGranularitys = Object.keys(enumTypes.TimeGranularity);
    QualityDimensions = Object.keys(enumTypes.QualityDimension);
    QualityStatuss = Object.keys(enumTypes.QualityStatus);
    AnomalyTypes = Object.keys(enumTypes.AnomalyType);
    AlertSeveritys = Object.keys(enumTypes.AlertSeverity);
    AlertStatuss = Object.keys(enumTypes.AlertStatus);
    LineageNodeTypes = Object.keys(enumTypes.LineageNodeType);
    SubjectTypes = Object.keys(enumTypes.SubjectType);
    AccessLevels = Object.keys(enumTypes.AccessLevel);
    SQLDialects = Object.keys(enumTypes.SQLDialect);
    RecommendationTypes = Object.keys(enumTypes.RecommendationType);
    FraudDetectionTypes = Object.keys(enumTypes.FraudDetectionType);
    FraudSignalTypes = Object.keys(enumTypes.FraudSignalType);
    FeatureStoreTypes = Object.keys(enumTypes.FeatureStoreType);
    DataTypes = Object.keys(enumTypes.DataType);
    ComparisonOperators = Object.keys(enumTypes.ComparisonOperator);
    NotificationChannels = Object.keys(enumTypes.NotificationChannel);
    TagCategorys = Object.keys(enumTypes.TagCategory);
    GovernanceTiers = Object.keys(enumTypes.GovernanceTier);

// all collection instances
    analyticsWorkspaces : any;
    dataSources : any;
    dataSets : any;
    dataPipelines : any;
    dataTasks : any;
    semanticModels : any;
    dimensions : any;
    measures : any;
    metrics : any;
    reports : any;
    dashboards : any;
    visualizations : any;
    notebooks : any;
    bIQuerys : any;
    experiments : any;
    trainingRuns : any;
    runMetrics : any;
    runParameters : any;
    model_s : any;
    modelVersions : any;
    evaluationMetrics : any;
    featureSets : any;
    features : any;
    inferenceEndpoints : any;
    predictions : any;
    forecasts : any;
    timeSeriess : any;
    anomalys : any;
    qualityRules : any;
    qualityChecks : any;
    lineageNodes : any;
    tags : any;
    accessPolicys : any;
    alerts : any;
    subscribers : any;
    businessGlossaryTerms : any;
    recommendationScenarios : any;
    fraudScenarios : any;
    fraudSignals : any;
  
// initialization  
    ngOnInit() {
    }

    initAnalyticsWorkspaceList() {
        if ( this.analyticsWorkspaces == null ) {
            new AnalyticsWorkspaceService(this.http).getAnalyticsWorkspaces().subscribe(res => {
                this.analyticsWorkspaces = res;
            });
        }
    }
    
    initDataSourceList() {
        if ( this.dataSources == null ) {
            new DataSourceService(this.http).getDataSources().subscribe(res => {
                this.dataSources = res;
            });
        }
    }
    
    initDataSetList() {
        if ( this.dataSets == null ) {
            new DataSetService(this.http).getDataSets().subscribe(res => {
                this.dataSets = res;
            });
        }
    }
    
    initDataPipelineList() {
        if ( this.dataPipelines == null ) {
            new DataPipelineService(this.http).getDataPipelines().subscribe(res => {
                this.dataPipelines = res;
            });
        }
    }
    
    initDataTaskList() {
        if ( this.dataTasks == null ) {
            new DataTaskService(this.http).getDataTasks().subscribe(res => {
                this.dataTasks = res;
            });
        }
    }
    
    initSemanticModelList() {
        if ( this.semanticModels == null ) {
            new SemanticModelService(this.http).getSemanticModels().subscribe(res => {
                this.semanticModels = res;
            });
        }
    }
    
    initDimensionList() {
        if ( this.dimensions == null ) {
            new DimensionService(this.http).getDimensions().subscribe(res => {
                this.dimensions = res;
            });
        }
    }
    
    initMeasureList() {
        if ( this.measures == null ) {
            new MeasureService(this.http).getMeasures().subscribe(res => {
                this.measures = res;
            });
        }
    }
    
    initMetricList() {
        if ( this.metrics == null ) {
            new MetricService(this.http).getMetrics().subscribe(res => {
                this.metrics = res;
            });
        }
    }
    
    initReportList() {
        if ( this.reports == null ) {
            new ReportService(this.http).getReports().subscribe(res => {
                this.reports = res;
            });
        }
    }
    
    initDashboardList() {
        if ( this.dashboards == null ) {
            new DashboardService(this.http).getDashboards().subscribe(res => {
                this.dashboards = res;
            });
        }
    }
    
    initVisualizationList() {
        if ( this.visualizations == null ) {
            new VisualizationService(this.http).getVisualizations().subscribe(res => {
                this.visualizations = res;
            });
        }
    }
    
    initNotebookList() {
        if ( this.notebooks == null ) {
            new NotebookService(this.http).getNotebooks().subscribe(res => {
                this.notebooks = res;
            });
        }
    }
    
    initBIQueryList() {
        if ( this.bIQuerys == null ) {
            new BIQueryService(this.http).getBIQuerys().subscribe(res => {
                this.bIQuerys = res;
            });
        }
    }
    
    initExperimentList() {
        if ( this.experiments == null ) {
            new ExperimentService(this.http).getExperiments().subscribe(res => {
                this.experiments = res;
            });
        }
    }
    
    initTrainingRunList() {
        if ( this.trainingRuns == null ) {
            new TrainingRunService(this.http).getTrainingRuns().subscribe(res => {
                this.trainingRuns = res;
            });
        }
    }
    
    initRunMetricList() {
        if ( this.runMetrics == null ) {
            new RunMetricService(this.http).getRunMetrics().subscribe(res => {
                this.runMetrics = res;
            });
        }
    }
    
    initRunParameterList() {
        if ( this.runParameters == null ) {
            new RunParameterService(this.http).getRunParameters().subscribe(res => {
                this.runParameters = res;
            });
        }
    }
    
    initModel_List() {
        if ( this.model_s == null ) {
            new Model_Service(this.http).getModel_s().subscribe(res => {
                this.model_s = res;
            });
        }
    }
    
    initModelVersionList() {
        if ( this.modelVersions == null ) {
            new ModelVersionService(this.http).getModelVersions().subscribe(res => {
                this.modelVersions = res;
            });
        }
    }
    
    initEvaluationMetricList() {
        if ( this.evaluationMetrics == null ) {
            new EvaluationMetricService(this.http).getEvaluationMetrics().subscribe(res => {
                this.evaluationMetrics = res;
            });
        }
    }
    
    initFeatureSetList() {
        if ( this.featureSets == null ) {
            new FeatureSetService(this.http).getFeatureSets().subscribe(res => {
                this.featureSets = res;
            });
        }
    }
    
    initFeatureList() {
        if ( this.features == null ) {
            new FeatureService(this.http).getFeatures().subscribe(res => {
                this.features = res;
            });
        }
    }
    
    initInferenceEndpointList() {
        if ( this.inferenceEndpoints == null ) {
            new InferenceEndpointService(this.http).getInferenceEndpoints().subscribe(res => {
                this.inferenceEndpoints = res;
            });
        }
    }
    
    initPredictionList() {
        if ( this.predictions == null ) {
            new PredictionService(this.http).getPredictions().subscribe(res => {
                this.predictions = res;
            });
        }
    }
    
    initForecastList() {
        if ( this.forecasts == null ) {
            new ForecastService(this.http).getForecasts().subscribe(res => {
                this.forecasts = res;
            });
        }
    }
    
    initTimeSeriesList() {
        if ( this.timeSeriess == null ) {
            new TimeSeriesService(this.http).getTimeSeriess().subscribe(res => {
                this.timeSeriess = res;
            });
        }
    }
    
    initAnomalyList() {
        if ( this.anomalys == null ) {
            new AnomalyService(this.http).getAnomalys().subscribe(res => {
                this.anomalys = res;
            });
        }
    }
    
    initQualityRuleList() {
        if ( this.qualityRules == null ) {
            new QualityRuleService(this.http).getQualityRules().subscribe(res => {
                this.qualityRules = res;
            });
        }
    }
    
    initQualityCheckList() {
        if ( this.qualityChecks == null ) {
            new QualityCheckService(this.http).getQualityChecks().subscribe(res => {
                this.qualityChecks = res;
            });
        }
    }
    
    initLineageNodeList() {
        if ( this.lineageNodes == null ) {
            new LineageNodeService(this.http).getLineageNodes().subscribe(res => {
                this.lineageNodes = res;
            });
        }
    }
    
    initTagList() {
        if ( this.tags == null ) {
            new TagService(this.http).getTags().subscribe(res => {
                this.tags = res;
            });
        }
    }
    
    initAccessPolicyList() {
        if ( this.accessPolicys == null ) {
            new AccessPolicyService(this.http).getAccessPolicys().subscribe(res => {
                this.accessPolicys = res;
            });
        }
    }
    
    initAlertList() {
        if ( this.alerts == null ) {
            new AlertService(this.http).getAlerts().subscribe(res => {
                this.alerts = res;
            });
        }
    }
    
    initSubscriberList() {
        if ( this.subscribers == null ) {
            new SubscriberService(this.http).getSubscribers().subscribe(res => {
                this.subscribers = res;
            });
        }
    }
    
    initBusinessGlossaryTermList() {
        if ( this.businessGlossaryTerms == null ) {
            new BusinessGlossaryTermService(this.http).getBusinessGlossaryTerms().subscribe(res => {
                this.businessGlossaryTerms = res;
            });
        }
    }
    
    initRecommendationScenarioList() {
        if ( this.recommendationScenarios == null ) {
            new RecommendationScenarioService(this.http).getRecommendationScenarios().subscribe(res => {
                this.recommendationScenarios = res;
            });
        }
    }
    
    initFraudScenarioList() {
        if ( this.fraudScenarios == null ) {
            new FraudScenarioService(this.http).getFraudScenarios().subscribe(res => {
                this.fraudScenarios = res;
            });
        }
    }
    
    initFraudSignalList() {
        if ( this.fraudSignals == null ) {
            new FraudSignalService(this.http).getFraudSignals().subscribe(res => {
                this.fraudSignals = res;
            });
        }
    }
    
    
// comparison function for select controls  
    compareFn(user1: any, user2: any) {
        return user1 == user2
    }    
}
