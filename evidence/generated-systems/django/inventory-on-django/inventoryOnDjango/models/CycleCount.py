from django.db import models
from inventoryOnDjango.models.CountStatus import CountStatus

#======================================================================
# 
# Encapsulates data for model CycleCount
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CycleCount Declaration
#======================================================================
class CycleCount (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	countNumber = models.CharField(max_length=200, null=True)
	scheduledDate = models.DateField(null=True)
	performedDate = models.DateField(null=True)
	approvedBy = models.CharField(max_length=200, null=True)
	warehouse = models.ForeignKey('Warehouse', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	locations = models.ManyToManyField('StorageLocation',  blank=True, related_name='+')
	entries = models.ManyToManyField('CycleCountEntry',  blank=True, related_name='+')
	transactions = models.ManyToManyField('InventoryTransaction',  blank=True, related_name='+')
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in CountStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.countNumber
		str = str + self.scheduledDate
		str = str + self.performedDate
		str = str + self.approvedBy
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "CycleCount";
    
	def objectType(self):
		return "CycleCount";
