from django.db import models

#======================================================================
# 
# Encapsulates data for model ExperimentVariant
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ExperimentVariant Declaration
#======================================================================
class ExperimentVariant (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	allocation = Percentage
	experiment = models.ForeignKey('Experiment', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	creativeVariation = models.ForeignKey('CreativeVariation', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	lineItem = models.ForeignKey('LineItem', on_delete=models.CASCADE, null=True, blank=True, related_name='+')

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.allocation
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "ExperimentVariant";
    
	def objectType(self):
		return "ExperimentVariant";
