from django.db import models

#======================================================================
# 
# Encapsulates data for model BOMItem
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class BOMItem Declaration
#======================================================================
class BOMItem (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	lineNumber = models.IntegerField(null=True)
	quantity = Quantity
	scrapPercent = Percentage
	bom = models.ForeignKey('BOM', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	component = models.ForeignKey('Item', on_delete=models.CASCADE, null=True, blank=True, related_name='+')

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.lineNumber
		str = str + self.quantity
		str = str + self.scrapPercent
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "BOMItem";
    
	def objectType(self):
		return "BOMItem";
