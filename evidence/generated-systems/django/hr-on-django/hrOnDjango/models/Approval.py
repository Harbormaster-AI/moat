from django.db import models
from hrOnDjango.models.ApprovalStatus import ApprovalStatus

#======================================================================
# 
# Encapsulates data for model Approval
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Approval Declaration
#======================================================================
class Approval (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	approverComment = models.CharField(max_length=200, null=True)
	actionDate = models.DateField(null=True)
	approver = models.ForeignKey('Employee', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	timesheet = models.ForeignKey('Timesheet', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	leaveRequest = models.ForeignKey('LeaveRequest', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in ApprovalStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.approverComment
		str = str + self.actionDate
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Approval";
    
	def objectType(self):
		return "Approval";
