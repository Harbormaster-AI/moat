from django.db import models
from governanceOnDjango.models.FindingSeverity import FindingSeverity
from governanceOnDjango.models.FindingStatus import FindingStatus

#======================================================================
# 
# Encapsulates data for model AuditFinding
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AuditFinding Declaration
#======================================================================
class AuditFinding (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	title = models.CharField(max_length=200, null=True)
	description = models.CharField(max_length=200, null=True)
	dueDate = models.DateField(null=True)
	engagement = models.ForeignKey('AuditEngagement', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	workpaper = models.ForeignKey('AuditWorkpaper', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	correctiveActions = models.ManyToManyField('CorrectiveAction',  blank=True, related_name='+')
	relatedRisks = models.ManyToManyField('Risk',  blank=True, related_name='+')
	relatedControls = models.ManyToManyField('Control',  blank=True, related_name='+')
	issues = models.ManyToManyField('Issue',  blank=True, related_name='+')
	severity = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in FindingSeverity])
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in FindingStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.title
		str = str + self.description
		str = str + self.dueDate
		str = str + self.severity
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "AuditFinding";
    
	def objectType(self):
		return "AuditFinding";
