from django.db import models
from manufacturingOnDjango.models.InspectionType import InspectionType
from manufacturingOnDjango.models.InspectionStatus import InspectionStatus

#======================================================================
# 
# Encapsulates data for model InspectionLot
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class InspectionLot Declaration
#======================================================================
class InspectionLot (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	lotNumber = models.CharField(max_length=200, null=True)
	quantity = Quantity
	sampleSize = models.IntegerField(null=True)
	createdOn = models.CharField(max_length=64, null=True)
	item = models.ForeignKey('Item', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	workOrder = models.ForeignKey('WorkOrder', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	goodsReceipt = models.ForeignKey('GoodsReceipt', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	results = models.ManyToManyField('InspectionResult',  blank=True, related_name='+')
	inspectionType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in InspectionType])
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in InspectionStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.lotNumber
		str = str + self.quantity
		str = str + self.sampleSize
		str = str + self.createdOn
		str = str + self.inspectionType
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "InspectionLot";
    
	def objectType(self):
		return "InspectionLot";
