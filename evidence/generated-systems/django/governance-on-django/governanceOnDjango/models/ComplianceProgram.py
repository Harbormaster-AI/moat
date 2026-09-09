from django.db import models
from governanceOnDjango.models.ComplianceStatus import ComplianceStatus

#======================================================================
# 
# Encapsulates data for model ComplianceProgram
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ComplianceProgram Declaration
#======================================================================
class ComplianceProgram (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	framework = models.CharField(max_length=200, null=True)
	organization = models.ForeignKey('Organization', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	requirements = models.ManyToManyField('ComplianceRequirement',  blank=True, related_name='+')
	controls = models.ManyToManyField('Control',  blank=True, related_name='+')
	attestations = models.ManyToManyField('Attestation',  blank=True, related_name='+')
	regulations = models.ManyToManyField('Regulation',  blank=True, related_name='+')
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in ComplianceStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.framework
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "ComplianceProgram";
    
	def objectType(self):
		return "ComplianceProgram";
