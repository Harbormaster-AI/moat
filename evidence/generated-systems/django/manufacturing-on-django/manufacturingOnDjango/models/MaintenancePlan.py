from django.db import models
from manufacturingOnDjango.models.MaintenanceStrategy import MaintenanceStrategy

#======================================================================
# 
# Encapsulates data for model MaintenancePlan
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class MaintenancePlan Declaration
#======================================================================
class MaintenancePlan (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	planNumber = models.CharField(max_length=200, null=True)
	interval = TimeDuration
	lastServiceDate = models.DateField(null=True)
	asset = models.ForeignKey('Asset', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	maintenanceOrders = models.ManyToManyField('MaintenanceOrder',  blank=True, related_name='+')
	strategy = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in MaintenanceStrategy])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.planNumber
		str = str + self.interval
		str = str + self.lastServiceDate
		str = str + self.strategy
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "MaintenancePlan";
    
	def objectType(self):
		return "MaintenancePlan";
