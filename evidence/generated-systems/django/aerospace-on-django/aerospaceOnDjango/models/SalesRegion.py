from django.db import models

#======================================================================
# 
# Encapsulates data for model SalesRegion
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class SalesRegion Declaration
#======================================================================
class SalesRegion (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	regionCode = models.CharField(max_length=200, null=True)
	operators = models.ManyToManyField('Operator',  blank=True, related_name='+')
	salesCampaigns = models.ManyToManyField('SalesCampaign',  blank=True, related_name='+')

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.regionCode
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "SalesRegion";
    
	def objectType(self):
		return "SalesRegion";
