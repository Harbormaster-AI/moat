from django.db import models
from aerospaceOnDjango.models.SalesCampaignStatus import SalesCampaignStatus

#======================================================================
# 
# Encapsulates data for model SalesCampaign
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class SalesCampaign Declaration
#======================================================================
class SalesCampaign (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	campaignCode = models.CharField(max_length=200, null=True)
	region = models.ForeignKey('SalesRegion', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	operator = models.ForeignKey('Operator', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	quotes = models.ManyToManyField('Quote',  blank=True, related_name='+')
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in SalesCampaignStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.campaignCode
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "SalesCampaign";
    
	def objectType(self):
		return "SalesCampaign";
