from django.db import models

#======================================================================
# 
# Encapsulates data for model HealthSystem
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class HealthSystem Declaration
#======================================================================
class HealthSystem (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	legalName = models.CharField(max_length=200, null=True)
	headquartersCountry = models.CharField(max_length=200, null=True)
	website = models.CharField(max_length=200, null=True)
	facilities = models.ManyToManyField('Facility',  blank=True, related_name='+')
	suppliers = models.ManyToManyField('MedicalSupplier',  blank=True, related_name='+')

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
		return "HealthSystem";
    
	def objectType(self):
		return "HealthSystem";
