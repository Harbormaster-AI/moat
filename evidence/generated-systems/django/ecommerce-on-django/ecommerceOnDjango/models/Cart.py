from django.db import models
from ecommerceOnDjango.models.CartStatus import CartStatus

#======================================================================
# 
# Encapsulates data for model Cart
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Cart Declaration
#======================================================================
class Cart (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	cartNumber = models.CharField(max_length=200, null=True)
	createdAt = models.DateField(null=True)
	currency = models.CharField(max_length=200, null=True)
	shippingAddress = Address
	billingAddress = Address
	customer = models.ForeignKey('Customer', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	channel = models.ForeignKey('Channel', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	items = models.ManyToManyField('CartItem',  blank=True, related_name='+')
	appliedPromotions = models.ManyToManyField('Promotion',  blank=True, related_name='+')
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in CartStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.cartNumber
		str = str + self.createdAt
		str = str + self.currency
		str = str + self.shippingAddress
		str = str + self.billingAddress
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Cart";
    
	def objectType(self):
		return "Cart";
