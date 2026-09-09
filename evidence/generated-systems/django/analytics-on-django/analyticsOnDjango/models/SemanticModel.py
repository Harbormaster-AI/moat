from django.db import models

#======================================================================
# 
# Encapsulates data for model SemanticModel
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class SemanticModel Declaration
#======================================================================
class SemanticModel (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	version = models.CharField(max_length=200, null=True)
	grain = models.CharField(max_length=200, null=True)
	datasets = models.ManyToManyField('DataSet',  blank=True, related_name='+')
	metrics = models.ManyToManyField('Metric',  blank=True, related_name='+')
	dimensions = models.ManyToManyField('Dimension',  blank=True, related_name='+')
	measures = models.ManyToManyField('Measure',  blank=True, related_name='+')
	glossaryTerms = models.ManyToManyField('BusinessGlossaryTerm',  blank=True, related_name='+')

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.version
		str = str + self.grain
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "SemanticModel";
    
	def objectType(self):
		return "SemanticModel";
