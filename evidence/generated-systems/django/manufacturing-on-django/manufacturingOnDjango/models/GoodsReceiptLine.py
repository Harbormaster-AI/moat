from django.db import models

#======================================================================
# 
# Encapsulates data for model GoodsReceiptLine
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class GoodsReceiptLine Declaration
#======================================================================
class GoodsReceiptLine (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	lineNumber = models.IntegerField(null=True)
	receivedQuantity = Quantity
	acceptedQuantity = Quantity
	rejectedQuantity = Quantity
	lot = LotId
	goodsReceipt = models.ForeignKey('GoodsReceipt', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	item = models.ForeignKey('Item', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	inventoryTransaction = models.OneToOneField('InventoryTransaction', on_delete=models.CASCADE, null=True, blank=True, related_name='+')

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.lineNumber
		str = str + self.receivedQuantity
		str = str + self.acceptedQuantity
		str = str + self.rejectedQuantity
		str = str + self.lot
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "GoodsReceiptLine";
    
	def objectType(self):
		return "GoodsReceiptLine";
