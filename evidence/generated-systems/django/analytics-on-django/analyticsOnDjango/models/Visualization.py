from django.db import models
from analyticsOnDjango.models.ChartType import ChartType

#======================================================================
# 
# Encapsulates data for model Visualization
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Visualization Declaration
#======================================================================
class Visualization (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	title = models.CharField(max_length=200, null=True)
	options = ChartOptions
	dashboard = models.ForeignKey('Dashboard', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	report = models.ForeignKey('Report', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	metrics = models.ManyToManyField('Metric',  blank=True, related_name='+')
	dimensions = models.ManyToManyField('Dimension',  blank=True, related_name='+')
	datasets = models.ManyToManyField('DataSet',  blank=True, related_name='+')
	chartType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in ChartType])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.title
		str = str + self.options
		str = str + self.chartType
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Visualization";
    
	def objectType(self):
		return "Visualization";
