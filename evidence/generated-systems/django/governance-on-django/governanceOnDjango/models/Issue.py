from django.db import models
from governanceOnDjango.models.IssueType import IssueType
from governanceOnDjango.models.Priority import Priority
from governanceOnDjango.models.IssueStatus import IssueStatus

#======================================================================
# 
# Encapsulates data for model Issue
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Issue Declaration
#======================================================================
class Issue (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	title = models.CharField(max_length=200, null=True)
	openedDate = models.DateField(null=True)
	closedDate = models.DateField(null=True)
	risk = models.ForeignKey('Risk', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	finding = models.ForeignKey('AuditFinding', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	correctiveActions = models.ManyToManyField('CorrectiveAction',  blank=True, related_name='+')
	control = models.ForeignKey('Control', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	issueType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in IssueType])
	priority = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in Priority])
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in IssueStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.title
		str = str + self.openedDate
		str = str + self.closedDate
		str = str + self.issueType
		str = str + self.priority
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Issue";
    
	def objectType(self):
		return "Issue";
