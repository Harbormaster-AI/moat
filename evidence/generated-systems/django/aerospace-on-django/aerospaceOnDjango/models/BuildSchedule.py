from django.db import models
from aerospaceOnDjango.models.ScheduleStatus import ScheduleStatus

#======================================================================
# 
# Encapsulates data for model BuildSchedule
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class BuildSchedule Declaration
#======================================================================
class BuildSchedule (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	scheduleNumber = models.CharField(max_length=200, null=True)
	productionOrders = models.ManyToManyField('ProductionOrder',  blank=True, related_name='+')
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in ScheduleStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.scheduleNumber
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "BuildSchedule";
    
	def objectType(self):
		return "BuildSchedule";
