from django.db import models
from inventoryOnDjango.models.SerialStatus import SerialStatus

#======================================================================
# 
# Encapsulates data for model SerialNumber
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class SerialNumber Declaration
#======================================================================
class SerialNumber (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	serial = SerialCode
	activationDate = models.DateField(null=True)
	sku = models.ForeignKey('StockKeepingUnit', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	currentInventoryItem = models.ForeignKey('InventoryItem', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	lot = models.ForeignKey('Lot', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in SerialStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.serial
		str = str + self.activationDate
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "SerialNumber";
    
	def objectType(self):
		return "SerialNumber";
