from django.db import models
from fintechOnDjango.models.SecurityType import SecurityType

#======================================================================
# 
# Encapsulates data for model Security
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Security Declaration
#======================================================================
class Security (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	symbol = models.CharField(max_length=200, null=True)
	isin = models.CharField(max_length=200, null=True)
	cusip = models.CharField(max_length=200, null=True)
	currency = models.CharField(max_length=200, null=True)
	positions = models.ManyToManyField('Position',  blank=True, related_name='+')
	trades = models.ManyToManyField('Trade',  blank=True, related_name='+')
	orders = models.ManyToManyField('TradeOrder',  blank=True, related_name='+')
	securityType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in SecurityType])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.symbol
		str = str + self.isin
		str = str + self.cusip
		str = str + self.currency
		str = str + self.securityType
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Security";
    
	def objectType(self):
		return "Security";
