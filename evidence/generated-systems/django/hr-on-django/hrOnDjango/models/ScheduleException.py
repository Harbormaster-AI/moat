from django.db import models

#======================================================================
# 
# Encapsulates data for model ScheduleException
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ScheduleException Declaration
#======================================================================
class ScheduleException (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	date = models.DateField(null=True)
	reason = models.CharField(max_length=200, null=True)
	hours = models.CharField(max_length=64, null=True)
	workSchedule = models.ForeignKey('WorkSchedule', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	employee = models.ForeignKey('Employee', on_delete=models.CASCADE, null=True, blank=True, related_name='+')

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.date
		str = str + self.reason
		str = str + self.hours
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "ScheduleException";
    
	def objectType(self):
		return "ScheduleException";
