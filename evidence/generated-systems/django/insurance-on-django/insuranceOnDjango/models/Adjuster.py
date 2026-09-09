from django.db import models
from insuranceOnDjango.models.AdjusterType import AdjusterType

#======================================================================
# 
# Encapsulates data for model Adjuster
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Adjuster Declaration
#======================================================================
class Adjuster (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	firstName = models.CharField(max_length=200, null=True)
	lastName = models.CharField(max_length=200, null=True)
	licenseNumber = models.CharField(max_length=200, null=True)
	claims = models.ManyToManyField('Claim',  blank=True, related_name='+')
	serviceProviders = models.ManyToManyField('ServiceProvider',  blank=True, related_name='+')
	adjusterType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in AdjusterType])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.firstName
		str = str + self.lastName
		str = str + self.licenseNumber
		str = str + self.adjusterType
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Adjuster";
    
	def objectType(self):
		return "Adjuster";
