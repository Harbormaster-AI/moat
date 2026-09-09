from django.db import models
from hrOnDjango.models.ScheduleType import ScheduleType

#======================================================================
# 
# Encapsulates data for model WorkSchedule
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class WorkSchedule Declaration
#======================================================================
class WorkSchedule (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	standardHoursPerWeek = models.CharField(max_length=64, null=True)
	contracts = models.ManyToManyField('EmploymentContract',  blank=True, related_name='+')
	shifts = models.ManyToManyField('WorkShift',  blank=True, related_name='+')
	exceptions = models.ManyToManyField('ScheduleException',  blank=True, related_name='+')
	scheduleType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in ScheduleType])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.standardHoursPerWeek
		str = str + self.scheduleType
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "WorkSchedule";
    
	def objectType(self):
		return "WorkSchedule";
