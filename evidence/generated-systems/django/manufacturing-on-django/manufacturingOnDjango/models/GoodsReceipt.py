from django.db import models
from manufacturingOnDjango.models.ReceiptStatus import ReceiptStatus

#======================================================================
# 
# Encapsulates data for model GoodsReceipt
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class GoodsReceipt Declaration
#======================================================================
class GoodsReceipt (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	receiptNumber = models.CharField(max_length=200, null=True)
	receiptDate = models.DateField(null=True)
	purchaseOrder = models.ForeignKey('PurchaseOrder', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	warehouse = models.ForeignKey('Warehouse', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	lines = models.ManyToManyField('GoodsReceiptLine',  blank=True, related_name='+')
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in ReceiptStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.receiptNumber
		str = str + self.receiptDate
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "GoodsReceipt";
    
	def objectType(self):
		return "GoodsReceipt";
