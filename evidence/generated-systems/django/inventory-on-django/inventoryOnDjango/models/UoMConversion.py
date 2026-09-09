from django.db import models
from inventoryOnDjango.models.UnitOfMeasure import UnitOfMeasure
from inventoryOnDjango.models.UnitOfMeasure import UnitOfMeasure

#======================================================================
# 
# Encapsulates data for model UoMConversion
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class UoMConversion Declaration
#======================================================================
class UoMConversion (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	factor = models.CharField(max_length=64, null=True)
	precision = models.IntegerField(null=True)
	sku = models.ForeignKey('StockKeepingUnit', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	fromUnit = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in UnitOfMeasure])
	toUnit = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in UnitOfMeasure])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.factor
		str = str + self.precision
		str = str + self.fromUnit
		str = str + self.toUnit
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "UoMConversion";
    
	def objectType(self):
		return "UoMConversion";
