from django.db import models
from hrOnDjango.models.BenefitType import BenefitType

#======================================================================
# 
# Encapsulates data for model BenefitPlan
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class BenefitPlan Declaration
#======================================================================
class BenefitPlan (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	providerName = models.CharField(max_length=200, null=True)
	employeeContributionRate = Percentage
	employerContributionRate = Percentage
	eligibilityRules = models.CharField(max_length=200, null=True)
	organization = models.ForeignKey('Organization', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	enrollments = models.ManyToManyField('BenefitEnrollment',  blank=True, related_name='+')
	benefitType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in BenefitType])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.providerName
		str = str + self.employeeContributionRate
		str = str + self.employerContributionRate
		str = str + self.eligibilityRules
		str = str + self.benefitType
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "BenefitPlan";
    
	def objectType(self):
		return "BenefitPlan";
