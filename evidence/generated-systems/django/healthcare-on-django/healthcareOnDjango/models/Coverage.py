from django.db import models
from healthcareOnDjango.models.CoverageType import CoverageType

#======================================================================
# 
# Encapsulates data for model Coverage
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Coverage Declaration
#======================================================================
class Coverage (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	memberId = models.CharField(max_length=200, null=True)
	groupNumber = models.CharField(max_length=200, null=True)
	effectiveDate = models.DateField(null=True)
	endDate = models.DateField(null=True)
	patient = models.ForeignKey('Patient', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	plan = models.ForeignKey('InsurancePlan', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	claims = models.ManyToManyField('Claim',  blank=True, related_name='+')
	authorizations = models.ManyToManyField('Authorization',  blank=True, related_name='+')
	coverageType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in CoverageType])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.memberId
		str = str + self.groupNumber
		str = str + self.effectiveDate
		str = str + self.endDate
		str = str + self.coverageType
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Coverage";
    
	def objectType(self):
		return "Coverage";
