from django.db import models
from manufacturingOnDjango.models.MaintenanceOrderStatus import MaintenanceOrderStatus

#======================================================================
# 
# Encapsulates data for model MaintenanceOrder
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class MaintenanceOrder Declaration
#======================================================================
class MaintenanceOrder (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	orderNumber = models.CharField(max_length=200, null=True)
	priority = models.IntegerField(null=True)
	requestedDate = models.DateField(null=True)
	completionDate = models.DateField(null=True)
	asset = models.ForeignKey('Asset', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	plan = models.ForeignKey('MaintenancePlan', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	workCenter = models.ForeignKey('WorkCenter', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in MaintenanceOrderStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.orderNumber
		str = str + self.priority
		str = str + self.requestedDate
		str = str + self.completionDate
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "MaintenanceOrder";
    
	def objectType(self):
		return "MaintenanceOrder";
