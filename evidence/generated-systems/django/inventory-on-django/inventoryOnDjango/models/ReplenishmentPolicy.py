from django.db import models
from inventoryOnDjango.models.ReplenishmentPolicyType import ReplenishmentPolicyType

#======================================================================
# 
# Encapsulates data for model ReplenishmentPolicy
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ReplenishmentPolicy Declaration
#======================================================================
class ReplenishmentPolicy (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	minLevel = models.CharField(max_length=64, null=True)
	maxLevel = models.CharField(max_length=64, null=True)
	reorderPoint = models.CharField(max_length=64, null=True)
	reorderQuantity = models.CharField(max_length=64, null=True)
	leadTimeDays = models.IntegerField(null=True)
	reviewPeriodDays = models.IntegerField(null=True)
	sku = models.ForeignKey('StockKeepingUnit', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	warehouse = models.ForeignKey('Warehouse', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	location = models.ForeignKey('StorageLocation', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	policyType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in ReplenishmentPolicyType])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.minLevel
		str = str + self.maxLevel
		str = str + self.reorderPoint
		str = str + self.reorderQuantity
		str = str + self.leadTimeDays
		str = str + self.reviewPeriodDays
		str = str + self.policyType
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "ReplenishmentPolicy";
    
	def objectType(self):
		return "ReplenishmentPolicy";
