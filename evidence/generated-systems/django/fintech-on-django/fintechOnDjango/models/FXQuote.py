from django.db import models
from fintechOnDjango.models.FXPriceType import FXPriceType

#======================================================================
# 
# Encapsulates data for model FXQuote
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class FXQuote Declaration
#======================================================================
class FXQuote (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	baseCurrency = models.CharField(max_length=200, null=True)
	quoteCurrency = models.CharField(max_length=200, null=True)
	rate = models.CharField(max_length=64, null=True)
	quotedAt = DateTime
	expiresAt = DateTime
	requestedBy = models.ForeignKey('Customer', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	priceType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in FXPriceType])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.baseCurrency
		str = str + self.quoteCurrency
		str = str + self.rate
		str = str + self.quotedAt
		str = str + self.expiresAt
		str = str + self.priceType
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "FXQuote";
    
	def objectType(self):
		return "FXQuote";
