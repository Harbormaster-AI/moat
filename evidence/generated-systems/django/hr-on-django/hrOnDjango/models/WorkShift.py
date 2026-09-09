from django.db import models
from hrOnDjango.models.DayOfWeek import DayOfWeek

#======================================================================
# 
# Encapsulates data for model WorkShift
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class WorkShift Declaration
#======================================================================
class WorkShift (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	startTime = LocalTime
	endTime = LocalTime
	breakMinutes = models.IntegerField(null=True)
	workSchedule = models.ForeignKey('WorkSchedule', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	dayOfWeek = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in DayOfWeek])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.startTime
		str = str + self.endTime
		str = str + self.breakMinutes
		str = str + self.dayOfWeek
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "WorkShift";
    
	def objectType(self):
		return "WorkShift";
