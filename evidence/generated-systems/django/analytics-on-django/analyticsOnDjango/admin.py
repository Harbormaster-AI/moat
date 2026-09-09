from django.contrib import admin

# Register your models here.
from .models.AnalyticsWorkspace import AnalyticsWorkspace
from .models.DataSource import DataSource
from .models.DataSet import DataSet
from .models.DataPipeline import DataPipeline
from .models.DataTask import DataTask
from .models.SemanticModel import SemanticModel
from .models.Dimension import Dimension
from .models.Measure import Measure
from .models.Metric import Metric
from .models.Report import Report
from .models.Dashboard import Dashboard
from .models.Visualization import Visualization
from .models.Notebook import Notebook
from .models.BIQuery import BIQuery
from .models.Experiment import Experiment
from .models.TrainingRun import TrainingRun
from .models.RunMetric import RunMetric
from .models.RunParameter import RunParameter
from .models.Model import Model
from .models.ModelVersion import ModelVersion
from .models.EvaluationMetric import EvaluationMetric
from .models.FeatureSet import FeatureSet
from .models.Feature import Feature
from .models.InferenceEndpoint import InferenceEndpoint
from .models.Prediction import Prediction
from .models.Forecast import Forecast
from .models.TimeSeries import TimeSeries
from .models.Anomaly import Anomaly
from .models.QualityRule import QualityRule
from .models.QualityCheck import QualityCheck
from .models.LineageNode import LineageNode
from .models.Tag import Tag
from .models.AccessPolicy import AccessPolicy
from .models.Alert import Alert
from .models.Subscriber import Subscriber
from .models.BusinessGlossaryTerm import BusinessGlossaryTerm
from .models.RecommendationScenario import RecommendationScenario
from .models.FraudScenario import FraudScenario
from .models.FraudSignal import FraudSignal

# Need to add this for each model that requires managing

admin.site.register(AnalyticsWorkspace)
admin.site.register(DataSource)
admin.site.register(DataSet)
admin.site.register(DataPipeline)
admin.site.register(DataTask)
admin.site.register(SemanticModel)
admin.site.register(Dimension)
admin.site.register(Measure)
admin.site.register(Metric)
admin.site.register(Report)
admin.site.register(Dashboard)
admin.site.register(Visualization)
admin.site.register(Notebook)
admin.site.register(BIQuery)
admin.site.register(Experiment)
admin.site.register(TrainingRun)
admin.site.register(RunMetric)
admin.site.register(RunParameter)
admin.site.register(Model)
admin.site.register(ModelVersion)
admin.site.register(EvaluationMetric)
admin.site.register(FeatureSet)
admin.site.register(Feature)
admin.site.register(InferenceEndpoint)
admin.site.register(Prediction)
admin.site.register(Forecast)
admin.site.register(TimeSeries)
admin.site.register(Anomaly)
admin.site.register(QualityRule)
admin.site.register(QualityCheck)
admin.site.register(LineageNode)
admin.site.register(Tag)
admin.site.register(AccessPolicy)
admin.site.register(Alert)
admin.site.register(Subscriber)
admin.site.register(BusinessGlossaryTerm)
admin.site.register(RecommendationScenario)
admin.site.register(FraudScenario)
admin.site.register(FraudSignal)
