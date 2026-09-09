from django.db import models

#======================================================================
# 
# Encapsulates data for model Policy
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Policy Declaration
#======================================================================
class Policy (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	policyNumber = models.CharField(max_length=200, null=True)
	name = models.CharField(max_length=200, null=True)
	effectiveDate = models.DateField(null=True)
	description = models.CharField(max_length=200, null=True)
	organization = models.ForeignKey('Organization', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	acknowledgements = models.ManyToManyField('PolicyAcknowledgement',  blank=True, related_name='+')

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.policyNumber
		str = str + self.name
		str = str + self.effectiveDate
		str = str + self.description
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Policy";
    
	def objectType(self):
		return "Policy";
