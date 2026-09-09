from django.db import models
from insuranceOnDjango.models.PerilType import PerilType

#======================================================================
# 
# Encapsulates data for model Incident
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Incident Declaration
#======================================================================
class Incident (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	location = Address
	description = models.CharField(max_length=200, null=True)
	claim = models.OneToOneField('Claim', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	insuredObjects = models.ManyToManyField('InsuredObject',  blank=True, related_name='+')
	incidentType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in PerilType])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.location
		str = str + self.description
		str = str + self.incidentType
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Incident";
    
	def objectType(self):
		return "Incident";
