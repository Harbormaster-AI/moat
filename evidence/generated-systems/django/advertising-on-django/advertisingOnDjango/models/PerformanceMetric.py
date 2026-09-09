from django.db import models
from advertisingOnDjango.models.MetricType import MetricType

#======================================================================
# 
# Encapsulates data for model PerformanceMetric
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PerformanceMetric Declaration
#======================================================================
class PerformanceMetric (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	date = models.DateField(null=True)
	value = models.CharField(max_length=64, null=True)
	adAccount = models.ForeignKey('AdAccount', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	campaign = models.ForeignKey('Campaign', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	lineItem = models.ForeignKey('LineItem', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	placement = models.ForeignKey('Placement', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	creativeAsset = models.ForeignKey('CreativeAsset', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	metricType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in MetricType])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.date
		str = str + self.value
		str = str + self.metricType
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "PerformanceMetric";
    
	def objectType(self):
		return "PerformanceMetric";
