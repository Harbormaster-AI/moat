from django.db import models
from aerospaceOnDjango.models.WarrantyType import WarrantyType

#======================================================================
# 
# Encapsulates data for model Warranty
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Warranty Declaration
#======================================================================
class Warranty (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	coverageMonths = models.IntegerField(null=True)
	aircraft = models.OneToOneField('Aircraft', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	warrantyType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in WarrantyType])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.coverageMonths
		str = str + self.warrantyType
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Warranty";
    
	def objectType(self):
		return "Warranty";
