from django.db import models
from manufacturingOnDjango.models.AssetStatus import AssetStatus

#======================================================================
# 
# Encapsulates data for model Asset
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Asset Declaration
#======================================================================
class Asset (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	assetTag = models.CharField(max_length=200, null=True)
	assetName = models.CharField(max_length=200, null=True)
	commissioningDate = models.DateField(null=True)
	plant = models.ForeignKey('Plant', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	workCenter = models.ForeignKey('WorkCenter', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	maintenanceOrders = models.ManyToManyField('MaintenanceOrder',  blank=True, related_name='+')
	maintenancePlans = models.ManyToManyField('MaintenancePlan',  blank=True, related_name='+')
	assetStatus = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in AssetStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.assetTag
		str = str + self.assetName
		str = str + self.commissioningDate
		str = str + self.assetStatus
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Asset";
    
	def objectType(self):
		return "Asset";
