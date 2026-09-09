from django.db import models

#======================================================================
# 
# Encapsulates data for model AdAccount
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AdAccount Declaration
#======================================================================
class AdAccount (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	accountCode = models.CharField(max_length=200, null=True)
	defaultCurrency = models.CharField(max_length=200, null=True)
	defaultTimezone = models.CharField(max_length=200, null=True)
	advertiser = models.ForeignKey('Advertiser', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	users = models.ManyToManyField('User',  blank=True, related_name='+')
	campaigns = models.ManyToManyField('Campaign',  blank=True, related_name='+')
	billingProfile = models.ForeignKey('BillingProfile', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	dsp = models.ForeignKey('DSP', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	performanceMetrics = models.ManyToManyField('PerformanceMetric',  blank=True, related_name='+')

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.accountCode
		str = str + self.defaultCurrency
		str = str + self.defaultTimezone
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "AdAccount";
    
	def objectType(self):
		return "AdAccount";
