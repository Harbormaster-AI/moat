from django.db import models

#======================================================================
# 
# Encapsulates data for model Organization
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Organization Declaration
#======================================================================
class Organization (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	defaultCurrency = models.CharField(max_length=200, null=True)
	defaultLocale = _Locale
	website = URL
	users = models.ManyToManyField('User',  blank=True, related_name='+')
	accounts = models.ManyToManyField('Account',  blank=True, related_name='+')
	teams = models.ManyToManyField('Team',  blank=True, related_name='+')
	territories = models.ManyToManyField('Territory',  blank=True, related_name='+')
	products = models.ManyToManyField('Product',  blank=True, related_name='+')
	priceBooks = models.ManyToManyField('PriceBook',  blank=True, related_name='+')
	campaigns = models.ManyToManyField('Campaign',  blank=True, related_name='+')

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.defaultCurrency
		str = str + self.defaultLocale
		str = str + self.website
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Organization";
    
	def objectType(self):
		return "Organization";
