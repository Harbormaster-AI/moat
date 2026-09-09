"""mainsite URL Configuration

The `urlpatterns` list routes URLs to views. For more information please see:
    https://docs.djangoproject.com/en/2.1/topics/http/urls/
Examples:
Function views
    1. Add an import:  from my_app import views
    2. Add a URL to urlpatterns:  path('', views.home, name='home')
Class-based views
    1. Add an import:  from other_app.views import Home
    2. Add a URL to urlpatterns:  path('', Home.as_view(), name='home')
Including another URLconf
    1. Import the include() function: from django.urls import include, path
    2. Add a URL to urlpatterns:  path('blog/', include('blog.urls'))
"""
from django.contrib import admin
from django.urls import path, include
urlpatterns = [
    path('AnalyticsWorkspace/', include('analyticsOnDjango.urls.AnalyticsWorkspaceUrls')),
    path('DataSource/', include('analyticsOnDjango.urls.DataSourceUrls')),
    path('DataSet/', include('analyticsOnDjango.urls.DataSetUrls')),
    path('DataPipeline/', include('analyticsOnDjango.urls.DataPipelineUrls')),
    path('DataTask/', include('analyticsOnDjango.urls.DataTaskUrls')),
    path('SemanticModel/', include('analyticsOnDjango.urls.SemanticModelUrls')),
    path('Dimension/', include('analyticsOnDjango.urls.DimensionUrls')),
    path('Measure/', include('analyticsOnDjango.urls.MeasureUrls')),
    path('Metric/', include('analyticsOnDjango.urls.MetricUrls')),
    path('Report/', include('analyticsOnDjango.urls.ReportUrls')),
    path('Dashboard/', include('analyticsOnDjango.urls.DashboardUrls')),
    path('Visualization/', include('analyticsOnDjango.urls.VisualizationUrls')),
    path('Notebook/', include('analyticsOnDjango.urls.NotebookUrls')),
    path('BIQuery/', include('analyticsOnDjango.urls.BIQueryUrls')),
    path('Experiment/', include('analyticsOnDjango.urls.ExperimentUrls')),
    path('TrainingRun/', include('analyticsOnDjango.urls.TrainingRunUrls')),
    path('RunMetric/', include('analyticsOnDjango.urls.RunMetricUrls')),
    path('RunParameter/', include('analyticsOnDjango.urls.RunParameterUrls')),
    path('Model/', include('analyticsOnDjango.urls.ModelUrls')),
    path('ModelVersion/', include('analyticsOnDjango.urls.ModelVersionUrls')),
    path('EvaluationMetric/', include('analyticsOnDjango.urls.EvaluationMetricUrls')),
    path('FeatureSet/', include('analyticsOnDjango.urls.FeatureSetUrls')),
    path('Feature/', include('analyticsOnDjango.urls.FeatureUrls')),
    path('InferenceEndpoint/', include('analyticsOnDjango.urls.InferenceEndpointUrls')),
    path('Prediction/', include('analyticsOnDjango.urls.PredictionUrls')),
    path('Forecast/', include('analyticsOnDjango.urls.ForecastUrls')),
    path('TimeSeries/', include('analyticsOnDjango.urls.TimeSeriesUrls')),
    path('Anomaly/', include('analyticsOnDjango.urls.AnomalyUrls')),
    path('QualityRule/', include('analyticsOnDjango.urls.QualityRuleUrls')),
    path('QualityCheck/', include('analyticsOnDjango.urls.QualityCheckUrls')),
    path('LineageNode/', include('analyticsOnDjango.urls.LineageNodeUrls')),
    path('Tag/', include('analyticsOnDjango.urls.TagUrls')),
    path('AccessPolicy/', include('analyticsOnDjango.urls.AccessPolicyUrls')),
    path('Alert/', include('analyticsOnDjango.urls.AlertUrls')),
    path('Subscriber/', include('analyticsOnDjango.urls.SubscriberUrls')),
    path('BusinessGlossaryTerm/', include('analyticsOnDjango.urls.BusinessGlossaryTermUrls')),
    path('RecommendationScenario/', include('analyticsOnDjango.urls.RecommendationScenarioUrls')),
    path('FraudScenario/', include('analyticsOnDjango.urls.FraudScenarioUrls')),
    path('FraudSignal/', include('analyticsOnDjango.urls.FraudSignalUrls')),
    path('admin/', admin.site.urls),
    path('', admin.site.urls),
]