from django.db import models
from ecommerceOnDjango.models.CustomerGroup import CustomerGroup

#======================================================================
# 
# Encapsulates data for model Customer
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Customer Declaration
#======================================================================
class Customer (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	firstName = models.CharField(max_length=200, null=True)
	lastName = models.CharField(max_length=200, null=True)
	email = models.CharField(max_length=200, null=True)
	phone = models.CharField(max_length=200, null=True)
	marketingOptIn = models.BooleanField(null=True)
	addresses = models.ManyToManyField('CustomerAddress',  blank=True, related_name='+')
	carts = models.ManyToManyField('Cart',  blank=True, related_name='+')
	orders = models.ManyToManyField('Order',  blank=True, related_name='+')
	payments = models.ManyToManyField('Payment',  blank=True, related_name='+')
	reviews = models.ManyToManyField('Review',  blank=True, related_name='+')
	wishlists = models.ManyToManyField('Wishlist',  blank=True, related_name='+')
	subscriptions = models.ManyToManyField('Subscription',  blank=True, related_name='+')
	couponRedemptions = models.ManyToManyField('CouponRedemption',  blank=True, related_name='+')
	giftCards = models.ManyToManyField('GiftCard',  blank=True, related_name='+')
	customerGroup = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in CustomerGroup])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.firstName
		str = str + self.lastName
		str = str + self.email
		str = str + self.phone
		str = str + self.marketingOptIn
		str = str + self.customerGroup
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Customer";
    
	def objectType(self):
		return "Customer";
