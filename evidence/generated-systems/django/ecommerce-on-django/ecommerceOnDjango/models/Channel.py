from django.db import models
from ecommerceOnDjango.models.ChannelType import ChannelType

#======================================================================
# 
# Encapsulates data for model Channel
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Channel Declaration
#======================================================================
class Channel (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	channelCode = models.CharField(max_length=200, null=True)
	locale = models.CharField(max_length=200, null=True)
	domain = models.CharField(max_length=200, null=True)
	asActive = models.BooleanField(null=True)
	defaultCurrency = models.CharField(max_length=200, null=True)
	merchant = models.ForeignKey('Merchant', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	catalogs = models.ManyToManyField('Catalog',  blank=True, related_name='+')
	promotions = models.ManyToManyField('Promotion',  blank=True, related_name='+')
	shippingMethods = models.ManyToManyField('ShippingMethod',  blank=True, related_name='+')
	paymentProviders = models.ManyToManyField('PaymentProvider',  blank=True, related_name='+')
	channelType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in ChannelType])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.channelCode
		str = str + self.locale
		str = str + self.domain
		str = str + self.asActive
		str = str + self.defaultCurrency
		str = str + self.channelType
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Channel";
    
	def objectType(self):
		return "Channel";
