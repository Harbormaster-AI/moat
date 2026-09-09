from django.db import models
from hrOnDjango.models.TimesheetStatus import TimesheetStatus

#======================================================================
# 
# Encapsulates data for model Timesheet
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Timesheet Declaration
#======================================================================
class Timesheet (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	periodStart = models.DateField(null=True)
	periodEnd = models.DateField(null=True)
	submissionDate = models.DateField(null=True)
	employee = models.ForeignKey('Employee', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	timeEntries = models.ManyToManyField('TimeEntry',  blank=True, related_name='+')
	approvals = models.ManyToManyField('Approval',  blank=True, related_name='+')
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in TimesheetStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.periodStart
		str = str + self.periodEnd
		str = str + self.submissionDate
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Timesheet";
    
	def objectType(self):
		return "Timesheet";
