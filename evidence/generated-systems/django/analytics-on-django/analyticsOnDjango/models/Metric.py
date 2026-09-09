from django.db import models
from analyticsOnDjango.models.MetricType import MetricType

#======================================================================
# 
# Encapsulates data for model Metric
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Metric Declaration
#======================================================================
class Metric (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	expression = models.CharField(max_length=200, null=True)
	unit = models.CharField(max_length=200, null=True)
	semanticModel = models.ForeignKey('SemanticModel', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	datasets = models.ManyToManyField('DataSet',  blank=True, related_name='+')
	glossaryTerms = models.ManyToManyField('BusinessGlossaryTerm',  blank=True, related_name='+')
	alerts = models.ManyToManyField('Alert',  blank=True, related_name='+')
	visualizations = models.ManyToManyField('Visualization',  blank=True, related_name='+')
	metricType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in MetricType])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.expression
		str = str + self.unit
		str = str + self.metricType
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Metric";
    
	def objectType(self):
		return "Metric";
