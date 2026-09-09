from django.db import models
from aerospaceOnDjango.models.WorkOrderStatus import WorkOrderStatus

#======================================================================
# 
# Encapsulates data for model MaintenanceWorkOrder
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class MaintenanceWorkOrder Declaration
#======================================================================
class MaintenanceWorkOrder (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	workOrderNumber = models.CharField(max_length=200, null=True)
	aircraft = models.ForeignKey('Aircraft', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	airworthinessDirective = models.ForeignKey('AirworthinessDirective', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	serviceBulletin = models.ForeignKey('ServiceBulletin', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in WorkOrderStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.workOrderNumber
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "MaintenanceWorkOrder";
    
	def objectType(self):
		return "MaintenanceWorkOrder";
