from django.db import models
from hrOnDjango.models.LeaveStatus import LeaveStatus

#======================================================================
# 
# Encapsulates data for model LeaveRequest
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class LeaveRequest Declaration
#======================================================================
class LeaveRequest (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	requestNumber = models.CharField(max_length=200, null=True)
	startDate = models.DateField(null=True)
	endDate = models.DateField(null=True)
	reason = models.CharField(max_length=200, null=True)
	hours = models.CharField(max_length=64, null=True)
	employee = models.ForeignKey('Employee', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	leavePolicy = models.ForeignKey('LeavePolicy', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	approvals = models.ManyToManyField('Approval',  blank=True, related_name='+')
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in LeaveStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.requestNumber
		str = str + self.startDate
		str = str + self.endDate
		str = str + self.reason
		str = str + self.hours
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "LeaveRequest";
    
	def objectType(self):
		return "LeaveRequest";
