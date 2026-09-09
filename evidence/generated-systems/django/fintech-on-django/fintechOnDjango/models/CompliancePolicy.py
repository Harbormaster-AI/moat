from django.db import models
from fintechOnDjango.models.PolicyStatus import PolicyStatus

#======================================================================
# 
# Encapsulates data for model CompliancePolicy
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CompliancePolicy Declaration
#======================================================================
class CompliancePolicy (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	policyCode = models.CharField(max_length=200, null=True)
	description = models.CharField(max_length=200, null=True)
	institution = models.ForeignKey('FinancialInstitution', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in PolicyStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.policyCode
		str = str + self.description
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "CompliancePolicy";
    
	def objectType(self):
		return "CompliancePolicy";
