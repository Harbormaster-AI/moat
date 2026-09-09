from django.db import models
from advertisingOnDjango.models.DealType import DealType

#======================================================================
# 
# Encapsulates data for model Deal
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Deal Declaration
#======================================================================
class Deal (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	floorPrice = Money
	publisher = models.ForeignKey('Publisher', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	inventorySources = models.ManyToManyField('InventorySource',  blank=True, related_name='+')
	placements = models.ManyToManyField('Placement',  blank=True, related_name='+')
	dealType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in DealType])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.floorPrice
		str = str + self.dealType
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Deal";
    
	def objectType(self):
		return "Deal";
