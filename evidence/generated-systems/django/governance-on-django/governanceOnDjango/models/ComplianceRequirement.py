from django.db import models
from governanceOnDjango.models.Applicability import Applicability
from governanceOnDjango.models.ComplianceStatus import ComplianceStatus

#======================================================================
# 
# Encapsulates data for model ComplianceRequirement
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ComplianceRequirement Declaration
#======================================================================
class ComplianceRequirement (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	source = models.CharField(max_length=200, null=True)
	citation = models.CharField(max_length=200, null=True)
	complianceProgram = models.ForeignKey('ComplianceProgram', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	policies = models.ManyToManyField('Policy',  blank=True, related_name='+')
	controls = models.ManyToManyField('Control',  blank=True, related_name='+')
	obligations = models.ManyToManyField('Obligation',  blank=True, related_name='+')
	applicability = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in Applicability])
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in ComplianceStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.source
		str = str + self.citation
		str = str + self.applicability
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "ComplianceRequirement";
    
	def objectType(self):
		return "ComplianceRequirement";
