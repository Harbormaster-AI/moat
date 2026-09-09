from django.db import models
from insuranceOnDjango.models.CoverageType import CoverageType

#======================================================================
# 
# Encapsulates data for model CoverageDefinition
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CoverageDefinition Declaration
#======================================================================
class CoverageDefinition (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	defaultLimit = Money
	defaultDeductible = Money
	asMandatory = models.BooleanField(null=True)
	product = models.ForeignKey('InsuranceProduct', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	coverageType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in CoverageType])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.defaultLimit
		str = str + self.defaultDeductible
		str = str + self.asMandatory
		str = str + self.coverageType
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "CoverageDefinition";
    
	def objectType(self):
		return "CoverageDefinition";
