from django.db import models
from healthcareOnDjango.models.FacilityType import FacilityType

#======================================================================
# 
# Encapsulates data for model Facility
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Facility Declaration
#======================================================================
class Facility (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	facilityCode = models.CharField(max_length=200, null=True)
	address = Address
	healthSystem = models.ForeignKey('HealthSystem', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	departments = models.ManyToManyField('Department',  blank=True, related_name='+')
	careTeams = models.ManyToManyField('CareTeam',  blank=True, related_name='+')
	laboratories = models.ManyToManyField('Laboratory',  blank=True, related_name='+')
	imagingCenters = models.ManyToManyField('ImagingCenter',  blank=True, related_name='+')
	pharmacies = models.ManyToManyField('Pharmacy',  blank=True, related_name='+')
	inventoryItems = models.ManyToManyField('InventoryItem',  blank=True, related_name='+')
	facilityType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in FacilityType])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.facilityCode
		str = str + self.address
		str = str + self.facilityType
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Facility";
    
	def objectType(self):
		return "Facility";
