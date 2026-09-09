from django.db import models
from inventoryOnDjango.models.InventoryAlertType import InventoryAlertType
from inventoryOnDjango.models.AlertStatus import AlertStatus

#======================================================================
# 
# Encapsulates data for model InventoryThresholdAlert
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class InventoryThresholdAlert Declaration
#======================================================================
class InventoryThresholdAlert (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	alertNumber = models.CharField(max_length=200, null=True)
	detectedAt = models.DateField(null=True)
	message = models.CharField(max_length=200, null=True)
	sku = models.ForeignKey('StockKeepingUnit', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	warehouse = models.ForeignKey('Warehouse', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	location = models.ForeignKey('StorageLocation', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	relatedPolicy = models.ForeignKey('ReplenishmentPolicy', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	alertType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in InventoryAlertType])
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in AlertStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.alertNumber
		str = str + self.detectedAt
		str = str + self.message
		str = str + self.alertType
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "InventoryThresholdAlert";
    
	def objectType(self):
		return "InventoryThresholdAlert";
