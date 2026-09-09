from django.db import models
from insuranceOnDjango.models.CoverageType import CoverageType

#======================================================================
# 
# Encapsulates data for model PolicyCoverage
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PolicyCoverage Declaration
#======================================================================
class PolicyCoverage (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	limit = Money
	deductible = Money
	premium = Money
	policy = models.ForeignKey('Policy', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	insuredObjects = models.ManyToManyField('InsuredObject',  blank=True, related_name='+')
	coverageType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in CoverageType])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.limit
		str = str + self.deductible
		str = str + self.premium
		str = str + self.coverageType
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "PolicyCoverage";
    
	def objectType(self):
		return "PolicyCoverage";
