from django.db import models
from inventoryOnDjango.models.StockStatus import StockStatus

#======================================================================
# 
# Encapsulates data for model CycleCountEntry
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CycleCountEntry Declaration
#======================================================================
class CycleCountEntry (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	lineNumber = models.IntegerField(null=True)
	systemQuantity = models.CharField(max_length=64, null=True)
	countedQuantity = models.CharField(max_length=64, null=True)
	varianceQuantity = models.CharField(max_length=64, null=True)
	recountRequired = models.BooleanField(null=True)
	cycleCount = models.ForeignKey('CycleCount', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	sku = models.ForeignKey('StockKeepingUnit', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	lot = models.ForeignKey('Lot', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	location = models.ForeignKey('StorageLocation', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	serialNumbers = models.ManyToManyField('SerialNumber',  blank=True, related_name='+')
	stockStatus = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in StockStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.lineNumber
		str = str + self.systemQuantity
		str = str + self.countedQuantity
		str = str + self.varianceQuantity
		str = str + self.recountRequired
		str = str + self.stockStatus
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "CycleCountEntry";
    
	def objectType(self):
		return "CycleCountEntry";
