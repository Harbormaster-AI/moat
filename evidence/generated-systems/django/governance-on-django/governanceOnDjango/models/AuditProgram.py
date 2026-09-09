from django.db import models
from governanceOnDjango.models.AuditCycle import AuditCycle
from governanceOnDjango.models.AuditStatus import AuditStatus

#======================================================================
# 
# Encapsulates data for model AuditProgram
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AuditProgram Declaration
#======================================================================
class AuditProgram (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	scope = models.CharField(max_length=200, null=True)
	organization = models.ForeignKey('Organization', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	engagements = models.ManyToManyField('AuditEngagement',  blank=True, related_name='+')
	cycle = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in AuditCycle])
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in AuditStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.scope
		str = str + self.cycle
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "AuditProgram";
    
	def objectType(self):
		return "AuditProgram";
