from django.db import models

#======================================================================
# 
# Encapsulates data for model AerospaceManufacturer
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AerospaceManufacturer Declaration
#======================================================================
class AerospaceManufacturer (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	legalName = models.CharField(max_length=200, null=True)
	headquartersCountry = models.CharField(max_length=200, null=True)
	website = models.CharField(max_length=200, null=True)
	programs = models.ManyToManyField('AircraftProgram',  blank=True, related_name='+')
	plants = models.ManyToManyField('Plant',  blank=True, related_name='+')
	suppliers = models.ManyToManyField('Supplier',  blank=True, related_name='+')
	productionCertificates = models.ManyToManyField('ProductionCertificate',  blank=True, related_name='+')

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
		return "AerospaceManufacturer";
    
	def objectType(self):
		return "AerospaceManufacturer";
