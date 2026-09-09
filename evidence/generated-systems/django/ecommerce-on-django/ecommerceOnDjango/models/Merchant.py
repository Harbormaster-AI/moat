from django.db import models

#======================================================================
# 
# Encapsulates data for model Merchant
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Merchant Declaration
#======================================================================
class Merchant (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	legalName = models.CharField(max_length=200, null=True)
	website = models.CharField(max_length=200, null=True)
	defaultCurrency = models.CharField(max_length=200, null=True)
	defaultLocale = models.CharField(max_length=200, null=True)
	supportEmail = models.CharField(max_length=200, null=True)
	channels = models.ManyToManyField('Channel',  blank=True, related_name='+')
	brands = models.ManyToManyField('Brand',  blank=True, related_name='+')
	fulfillmentCenters = models.ManyToManyField('FulfillmentCenter',  blank=True, related_name='+')
	taxRules = models.ManyToManyField('TaxRule',  blank=True, related_name='+')
	paymentProviders = models.ManyToManyField('PaymentProvider',  blank=True, related_name='+')
	sellers = models.ManyToManyField('Seller',  blank=True, related_name='+')
	promotions = models.ManyToManyField('Promotion',  blank=True, related_name='+')

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.legalName
		str = str + self.website
		str = str + self.defaultCurrency
		str = str + self.defaultLocale
		str = str + self.supportEmail
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Merchant";
    
	def objectType(self):
		return "Merchant";
