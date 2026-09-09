from django.db import models
from healthcareOnDjango.models.InsurancePlanType import InsurancePlanType

#======================================================================
# 
# Encapsulates data for model InsurancePlan
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class InsurancePlan Declaration
#======================================================================
class InsurancePlan (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	planCode = models.CharField(max_length=200, null=True)
	payer = models.ForeignKey('InsurancePayer', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	coverages = models.ManyToManyField('Coverage',  blank=True, related_name='+')
	planType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in InsurancePlanType])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.planCode
		str = str + self.planType
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "InsurancePlan";
    
	def objectType(self):
		return "InsurancePlan";
