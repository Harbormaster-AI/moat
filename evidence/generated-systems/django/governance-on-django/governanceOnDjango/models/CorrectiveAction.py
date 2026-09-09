from django.db import models
from governanceOnDjango.models.ActionStatus import ActionStatus

#======================================================================
# 
# Encapsulates data for model CorrectiveAction
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CorrectiveAction Declaration
#======================================================================
class CorrectiveAction (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	actionTitle = models.CharField(max_length=200, null=True)
	owner = models.CharField(max_length=200, null=True)
	targetDate = models.DateField(null=True)
	finding = models.ForeignKey('AuditFinding', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	issue = models.ForeignKey('Issue', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in ActionStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.actionTitle
		str = str + self.owner
		str = str + self.targetDate
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "CorrectiveAction";
    
	def objectType(self):
		return "CorrectiveAction";
