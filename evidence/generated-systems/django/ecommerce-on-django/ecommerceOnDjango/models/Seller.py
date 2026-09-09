from django.db import models
from ecommerceOnDjango.models.SellerStatus import SellerStatus

#======================================================================
# 
# Encapsulates data for model Seller
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Seller Declaration
#======================================================================
class Seller (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	sellerCode = models.CharField(max_length=200, null=True)
	contactEmail = models.CharField(max_length=200, null=True)
	merchant = models.ForeignKey('Merchant', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	products = models.ManyToManyField('Product',  blank=True, related_name='+')
	payouts = models.ManyToManyField('Payout',  blank=True, related_name='+')
	orders = models.ManyToManyField('Order',  blank=True, related_name='+')
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in SellerStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.sellerCode
		str = str + self.contactEmail
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Seller";
    
	def objectType(self):
		return "Seller";
