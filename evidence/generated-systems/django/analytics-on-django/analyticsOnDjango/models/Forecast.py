from django.db import models
from analyticsOnDjango.models.TimeGranularity import TimeGranularity

#======================================================================
# 
# Encapsulates data for model Forecast
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Forecast Declaration
#======================================================================
class Forecast (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	horizon = models.IntegerField(null=True)
	modelVersion = models.ForeignKey('ModelVersion', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	timeSeries = models.ForeignKey('TimeSeries', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	datasets = models.ManyToManyField('DataSet',  blank=True, related_name='+')
	granularity = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in TimeGranularity])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.horizon
		str = str + self.granularity
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Forecast";
    
	def objectType(self):
		return "Forecast";
