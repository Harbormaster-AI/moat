from django.db import models
from ecommerceOnDjango.models.OrderLineStatus import OrderLineStatus

#======================================================================
# 
# Encapsulates data for model OrderLine
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class OrderLine Declaration
#======================================================================
class OrderLine (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	quantity = models.IntegerField(null=True)
	unitPrice = Money
	totalPrice = Money
	taxRate = Percentage
	order = models.ForeignKey('Order', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	variant = models.ForeignKey('ProductVariant', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	appliedPromotions = models.ManyToManyField('Promotion',  blank=True, related_name='+')
	lineStatus = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in OrderLineStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.quantity
		str = str + self.unitPrice
		str = str + self.totalPrice
		str = str + self.taxRate
		str = str + self.lineStatus
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "OrderLine";
    
	def objectType(self):
		return "OrderLine";
