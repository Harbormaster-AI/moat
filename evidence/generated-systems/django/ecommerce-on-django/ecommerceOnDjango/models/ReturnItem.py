from django.db import models
from ecommerceOnDjango.models.ReturnReason import ReturnReason
from ecommerceOnDjango.models.ReturnItemCondition import ReturnItemCondition

#======================================================================
# 
# Encapsulates data for model ReturnItem
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ReturnItem Declaration
#======================================================================
class ReturnItem (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	quantity = models.IntegerField(null=True)
	returnRequest = models.ForeignKey('ReturnRequest', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	orderLine = models.ForeignKey('OrderLine', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	reason = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in ReturnReason])
	condition = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in ReturnItemCondition])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.quantity
		str = str + self.reason
		str = str + self.condition
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "ReturnItem";
    
	def objectType(self):
		return "ReturnItem";
