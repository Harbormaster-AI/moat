from django.db import models
from advertisingOnDjango.models.MetricType import MetricType

#======================================================================
# 
# Encapsulates data for model KPI
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class KPI Declaration
#======================================================================
class KPI (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	targetValue = models.CharField(max_length=64, null=True)
	campaign = models.ForeignKey('Campaign', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	metricType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in MetricType])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.targetValue
		str = str + self.metricType
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "KPI";
    
	def objectType(self):
		return "KPI";
