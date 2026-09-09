from django.db import models
from inventoryOnDjango.models.Disposition import Disposition

#======================================================================
# 
# Encapsulates data for model Quarantine
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Quarantine Declaration
#======================================================================
class Quarantine (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	reason = models.CharField(max_length=200, null=True)
	startedAt = models.DateField(null=True)
	releasedAt = models.DateField(null=True)
	warehouse = models.ForeignKey('Warehouse', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	items = models.ManyToManyField('InventoryItem',  blank=True, related_name='+')
	lot = models.ForeignKey('Lot', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	serialNumbers = models.ManyToManyField('SerialNumber',  blank=True, related_name='+')
	disposition = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in Disposition])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.reason
		str = str + self.startedAt
		str = str + self.releasedAt
		str = str + self.disposition
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Quarantine";
    
	def objectType(self):
		return "Quarantine";
