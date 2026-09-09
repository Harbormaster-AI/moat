from django.db import models
from hrOnDjango.models.TimeEntryType import TimeEntryType

#======================================================================
# 
# Encapsulates data for model TimeEntry
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class TimeEntry Declaration
#======================================================================
class TimeEntry (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	entryDate = models.DateField(null=True)
	hoursWorked = models.CharField(max_length=64, null=True)
	timesheet = models.ForeignKey('Timesheet', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	employee = models.ForeignKey('Employee', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	costCenter = models.ForeignKey('CostCenter', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	entryType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in TimeEntryType])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.entryDate
		str = str + self.hoursWorked
		str = str + self.entryType
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "TimeEntry";
    
	def objectType(self):
		return "TimeEntry";
