from django.db import models

#======================================================================
# 
# Encapsulates data for model BusinessGlossaryTerm
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class BusinessGlossaryTerm Declaration
#======================================================================
class BusinessGlossaryTerm (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	term = models.CharField(max_length=200, null=True)
	definition = models.CharField(max_length=200, null=True)
	steward = models.CharField(max_length=200, null=True)
	relatedTerms = models.ManyToManyField('BusinessGlossaryTerm',  blank=True, related_name='+')
	metrics = models.ManyToManyField('Metric',  blank=True, related_name='+')
	datasets = models.ManyToManyField('DataSet',  blank=True, related_name='+')
	dimensions = models.ManyToManyField('Dimension',  blank=True, related_name='+')
	measures = models.ManyToManyField('Measure',  blank=True, related_name='+')

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.term
		str = str + self.definition
		str = str + self.steward
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "BusinessGlossaryTerm";
    
	def objectType(self):
		return "BusinessGlossaryTerm";
