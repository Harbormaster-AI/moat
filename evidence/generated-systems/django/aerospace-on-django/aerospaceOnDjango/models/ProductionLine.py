from django.db import models
from aerospaceOnDjango.models.ProductionLineType import ProductionLineType

#======================================================================
# 
# Encapsulates data for model ProductionLine
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ProductionLine Declaration
#======================================================================
class ProductionLine (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	plant = models.ForeignKey('Plant', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	workCenters = models.ManyToManyField('WorkCenter',  blank=True, related_name='+')
	lineType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in ProductionLineType])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.lineType
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "ProductionLine";
    
	def objectType(self):
		return "ProductionLine";
