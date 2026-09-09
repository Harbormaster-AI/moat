from django.db import models
from manufacturingOnDjango.models.WorkCenterType import WorkCenterType

#======================================================================
# 
# Encapsulates data for model WorkCenter
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class WorkCenter Declaration
#======================================================================
class WorkCenter (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	code = models.CharField(max_length=200, null=True)
	capacityPerHour = models.IntegerField(null=True)
	oeeTarget = Percentage
	productionLine = models.ForeignKey('ProductionLine', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	assets = models.ManyToManyField('Asset',  blank=True, related_name='+')
	maintenanceOrders = models.ManyToManyField('MaintenanceOrder',  blank=True, related_name='+')
	workCenterType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in WorkCenterType])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.code
		str = str + self.capacityPerHour
		str = str + self.oeeTarget
		str = str + self.workCenterType
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "WorkCenter";
    
	def objectType(self):
		return "WorkCenter";
