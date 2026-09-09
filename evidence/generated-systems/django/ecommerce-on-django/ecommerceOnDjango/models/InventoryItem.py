from django.db import models
from ecommerceOnDjango.models.InventoryStatus import InventoryStatus

#======================================================================
# 
# Encapsulates data for model InventoryItem
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class InventoryItem Declaration
#======================================================================
class InventoryItem (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	quantityOnHand = models.IntegerField(null=True)
	quantityReserved = models.IntegerField(null=True)
	safetyStock = models.IntegerField(null=True)
	variant = models.ForeignKey('ProductVariant', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	fulfillmentCenter = models.ForeignKey('FulfillmentCenter', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in InventoryStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.quantityOnHand
		str = str + self.quantityReserved
		str = str + self.safetyStock
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "InventoryItem";
    
	def objectType(self):
		return "InventoryItem";
