from django.db import models
from manufacturingOnDjango.models.ScheduleStatus import ScheduleStatus

#======================================================================
# 
# Encapsulates data for model ProductionSchedule
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ProductionSchedule Declaration
#======================================================================
class ProductionSchedule (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	scheduleNumber = models.CharField(max_length=200, null=True)
	horizonStart = models.DateField(null=True)
	horizonEnd = models.DateField(null=True)
	plant = models.ForeignKey('Plant', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	workOrders = models.ManyToManyField('WorkOrder',  blank=True, related_name='+')
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in ScheduleStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.scheduleNumber
		str = str + self.horizonStart
		str = str + self.horizonEnd
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "ProductionSchedule";
    
	def objectType(self):
		return "ProductionSchedule";
