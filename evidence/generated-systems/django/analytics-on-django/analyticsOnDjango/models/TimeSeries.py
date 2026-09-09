from django.db import models
from analyticsOnDjango.models.TimeGranularity import TimeGranularity

#======================================================================
# 
# Encapsulates data for model TimeSeries
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class TimeSeries Declaration
#======================================================================
class TimeSeries (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	timezone = models.CharField(max_length=200, null=True)
	datasets = models.ManyToManyField('DataSet',  blank=True, related_name='+')
	forecasts = models.ManyToManyField('Forecast',  blank=True, related_name='+')
	anomalies = models.ManyToManyField('Anomaly',  blank=True, related_name='+')
	granularity = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in TimeGranularity])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.timezone
		str = str + self.granularity
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "TimeSeries";
    
	def objectType(self):
		return "TimeSeries";
