from django.db import models
from analyticsOnDjango.models.AggregationType import AggregationType

#======================================================================
# 
# Encapsulates data for model Measure
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Measure Declaration
#======================================================================
class Measure (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	format = models.CharField(max_length=200, null=True)
	semanticModel = models.ForeignKey('SemanticModel', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	datasets = models.ManyToManyField('DataSet',  blank=True, related_name='+')
	glossaryTerms = models.ManyToManyField('BusinessGlossaryTerm',  blank=True, related_name='+')
	aggregation = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in AggregationType])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.format
		str = str + self.aggregation
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Measure";
    
	def objectType(self):
		return "Measure";
