from django.db import models
from inventoryOnDjango.models.RotationMethod import RotationMethod

#======================================================================
# 
# Encapsulates data for model ExpirationPolicy
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ExpirationPolicy Declaration
#======================================================================
class ExpirationPolicy (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	rejectIfDaysToExpireLessThan = models.IntegerField(null=True)
	autoQuarantineDaysToExpire = models.IntegerField(null=True)
	sku = models.ForeignKey('StockKeepingUnit', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	warehouse = models.ForeignKey('Warehouse', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	rotationMethod = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in RotationMethod])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.rejectIfDaysToExpireLessThan
		str = str + self.autoQuarantineDaysToExpire
		str = str + self.rotationMethod
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "ExpirationPolicy";
    
	def objectType(self):
		return "ExpirationPolicy";
