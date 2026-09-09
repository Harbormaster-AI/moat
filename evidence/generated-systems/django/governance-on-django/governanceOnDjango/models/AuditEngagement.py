from django.db import models
from governanceOnDjango.models.AuditStatus import AuditStatus

#======================================================================
# 
# Encapsulates data for model AuditEngagement
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AuditEngagement Declaration
#======================================================================
class AuditEngagement (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	title = models.CharField(max_length=200, null=True)
	startDate = models.DateField(null=True)
	endDate = models.DateField(null=True)
	auditProgram = models.ForeignKey('AuditProgram', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	businessUnits = models.ManyToManyField('BusinessUnit',  blank=True, related_name='+')
	controlTests = models.ManyToManyField('ControlTest_',  blank=True, related_name='+')
	workpapers = models.ManyToManyField('AuditWorkpaper',  blank=True, related_name='+')
	findings = models.ManyToManyField('AuditFinding',  blank=True, related_name='+')
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in AuditStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.title
		str = str + self.startDate
		str = str + self.endDate
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "AuditEngagement";
    
	def objectType(self):
		return "AuditEngagement";
