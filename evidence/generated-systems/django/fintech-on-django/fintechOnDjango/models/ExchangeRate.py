from django.db import models

#======================================================================
# 
# Encapsulates data for model ExchangeRate
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ExchangeRate Declaration
#======================================================================
class ExchangeRate (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	baseCurrency = models.CharField(max_length=200, null=True)
	quoteCurrency = models.CharField(max_length=200, null=True)
	rate = models.CharField(max_length=64, null=True)
	asOf = DateTime
	source = models.CharField(max_length=200, null=True)
	usedByQuotes = models.ManyToManyField('FXQuote',  blank=True, related_name='+')

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.baseCurrency
		str = str + self.quoteCurrency
		str = str + self.rate
		str = str + self.asOf
		str = str + self.source
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "ExchangeRate";
    
	def objectType(self):
		return "ExchangeRate";
