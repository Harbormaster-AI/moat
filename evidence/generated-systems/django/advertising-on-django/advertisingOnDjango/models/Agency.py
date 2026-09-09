from django.db import models

#======================================================================
# 
# Encapsulates data for model Agency
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Agency Declaration
#======================================================================
class Agency (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	legalName = models.CharField(max_length=200, null=True)
	headquartersCountry = models.CharField(max_length=200, null=True)
	website = models.CharField(max_length=200, null=True)
	advertisers = models.ManyToManyField('Advertiser',  blank=True, related_name='+')
	teams = models.ManyToManyField('Team',  blank=True, related_name='+')
	users = models.ManyToManyField('User',  blank=True, related_name='+')
	insertionOrders = models.ManyToManyField('InsertionOrder',  blank=True, related_name='+')

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.legalName
		str = str + self.headquartersCountry
		str = str + self.website
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Agency";
    
	def objectType(self):
		return "Agency";
