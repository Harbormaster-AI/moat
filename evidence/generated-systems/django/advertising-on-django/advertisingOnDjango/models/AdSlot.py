from django.db import models
from advertisingOnDjango.models.AdFormat import AdFormat

#======================================================================
# 
# Encapsulates data for model AdSlot
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AdSlot Declaration
#======================================================================
class AdSlot (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	slotCode = models.CharField(max_length=200, null=True)
	width = models.IntegerField(null=True)
	height = models.IntegerField(null=True)
	floorPrice = Money
	inventorySource = models.ForeignKey('InventorySource', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	placements = models.ManyToManyField('Placement',  blank=True, related_name='+')
	rates = models.ManyToManyField('Rate',  blank=True, related_name='+')
	format = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in AdFormat])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.slotCode
		str = str + self.width
		str = str + self.height
		str = str + self.floorPrice
		str = str + self.format
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "AdSlot";
    
	def objectType(self):
		return "AdSlot";
