from django.db import models
from aerospaceOnDjango.models.SupplierType import SupplierType
from aerospaceOnDjango.models.SupplierApprovalStatus import SupplierApprovalStatus

#======================================================================
# 
# Encapsulates data for model Supplier
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Supplier Declaration
#======================================================================
class Supplier (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	manufacturers = models.ManyToManyField('AerospaceManufacturer',  blank=True, related_name='+')
	components = models.ManyToManyField('Component_',  blank=True, related_name='+')
	engineTypes = models.ManyToManyField('EngineType',  blank=True, related_name='+')
	avionicsSuites = models.ManyToManyField('AvionicsSuite',  blank=True, related_name='+')
	apus = models.ManyToManyField('APU',  blank=True, related_name='+')
	landingGears = models.ManyToManyField('LandingGear',  blank=True, related_name='+')
	supplierType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in SupplierType])
	approvalStatus = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in SupplierApprovalStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.supplierType
		str = str + self.approvalStatus
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Supplier";
    
	def objectType(self):
		return "Supplier";
