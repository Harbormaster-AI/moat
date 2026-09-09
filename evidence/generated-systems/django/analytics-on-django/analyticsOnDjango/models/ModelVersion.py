from django.db import models
from analyticsOnDjango.models.ModelLifecycle import ModelLifecycle
from analyticsOnDjango.models.TrainingStatus import TrainingStatus

#======================================================================
# 
# Encapsulates data for model ModelVersion
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ModelVersion Declaration
#======================================================================
class ModelVersion (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	version = models.CharField(max_length=200, null=True)
	model = models.ForeignKey('Model', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	trainingRun = models.ForeignKey('TrainingRun', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	evaluationMetrics = models.ManyToManyField('EvaluationMetric',  blank=True, related_name='+')
	deployments = models.ManyToManyField('InferenceEndpoint',  blank=True, related_name='+')
	featureSets = models.ManyToManyField('FeatureSet',  blank=True, related_name='+')
	datasets = models.ManyToManyField('DataSet',  blank=True, related_name='+')
	lifecycle = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in ModelLifecycle])
	trainingStatus = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in TrainingStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.version
		str = str + self.lifecycle
		str = str + self.trainingStatus
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "ModelVersion";
    
	def objectType(self):
		return "ModelVersion";
