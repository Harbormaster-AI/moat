from django.db import models

#======================================================================
# 
# Encapsulates data for model Warehouse
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Warehouse Declaration
#======================================================================
class Warehouse (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	inventoryItems = models.ManyToManyField('InventoryItem',  blank=True, related_name='+')

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Warehouse";
    
	def objectType(self):
		return "Warehouse";
