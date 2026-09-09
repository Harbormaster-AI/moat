from django.db import models
from healthcareOnDjango.models.SupplierTier import SupplierTier

#======================================================================
# 
# Encapsulates data for model MedicalSupplier
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class MedicalSupplier Declaration
#======================================================================
class MedicalSupplier (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	website = models.CharField(max_length=200, null=True)
	facilities = models.ManyToManyField('Facility',  blank=True, related_name='+')
	inventoryItems = models.ManyToManyField('InventoryItem',  blank=True, related_name='+')
	supplierTier = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in SupplierTier])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.website
		str = str + self.supplierTier
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "MedicalSupplier";
    
	def objectType(self):
		return "MedicalSupplier";
