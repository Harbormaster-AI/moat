from django.db import models
from analyticsOnDjango.models.TrainingStatus import TrainingStatus

#======================================================================
# 
# Encapsulates data for model TrainingRun
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class TrainingRun Declaration
#======================================================================
class TrainingRun (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	runLabel = models.CharField(max_length=200, null=True)
	startedAt = models.DateField(null=True)
	completedAt = models.DateField(null=True)
	experiment = models.ForeignKey('Experiment', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	modelVersion = models.ForeignKey('ModelVersion', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	inputDatasets = models.ManyToManyField('DataSet',  blank=True, related_name='+')
	features = models.ManyToManyField('Feature',  blank=True, related_name='+')
	runMetrics = models.ManyToManyField('RunMetric',  blank=True, related_name='+')
	runParameters = models.ManyToManyField('RunParameter',  blank=True, related_name='+')
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in TrainingStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.runLabel
		str = str + self.startedAt
		str = str + self.completedAt
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "TrainingRun";
    
	def objectType(self):
		return "TrainingRun";
