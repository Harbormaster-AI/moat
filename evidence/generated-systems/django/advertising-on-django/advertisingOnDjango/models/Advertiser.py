from django.db import models

#======================================================================
# 
# Encapsulates data for model Advertiser
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Advertiser Declaration
#======================================================================
class Advertiser (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	legalName = models.CharField(max_length=200, null=True)
	industry = models.CharField(max_length=200, null=True)
	website = models.CharField(max_length=200, null=True)
	agency = models.ForeignKey('Agency', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	adAccounts = models.ManyToManyField('AdAccount',  blank=True, related_name='+')
	billingProfiles = models.ManyToManyField('BillingProfile',  blank=True, related_name='+')
	campaigns = models.ManyToManyField('Campaign',  blank=True, related_name='+')
	trackingPixels = models.ManyToManyField('TrackingPixel',  blank=True, related_name='+')

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.legalName
		str = str + self.industry
		str = str + self.website
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Advertiser";
    
	def objectType(self):
		return "Advertiser";
