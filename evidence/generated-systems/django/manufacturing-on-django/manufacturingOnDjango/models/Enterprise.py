from django.db import models

#======================================================================
# 
# Encapsulates data for model Enterprise
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Enterprise Declaration
#======================================================================
class Enterprise (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	legalName = models.CharField(max_length=200, null=True)
	registrationCountry = models.CharField(max_length=200, null=True)
	website = models.CharField(max_length=200, null=True)
	taxId = models.CharField(max_length=200, null=True)
	businessUnits = models.ManyToManyField('BusinessUnit',  blank=True, related_name='+')
	plants = models.ManyToManyField('Plant',  blank=True, related_name='+')
	suppliers = models.ManyToManyField('Supplier',  blank=True, related_name='+')
	customers = models.ManyToManyField('Customer',  blank=True, related_name='+')

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.legalName
		str = str + self.registrationCountry
		str = str + self.website
		str = str + self.taxId
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Enterprise";
    
	def objectType(self):
		return "Enterprise";
