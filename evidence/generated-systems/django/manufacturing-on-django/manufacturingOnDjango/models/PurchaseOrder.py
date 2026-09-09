from django.db import models
from manufacturingOnDjango.models.PurchaseOrderStatus import PurchaseOrderStatus

#======================================================================
# 
# Encapsulates data for model PurchaseOrder
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PurchaseOrder Declaration
#======================================================================
class PurchaseOrder (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	poNumber = models.CharField(max_length=200, null=True)
	orderDate = models.DateField(null=True)
	totalAmount = Money
	supplier = models.ForeignKey('Supplier', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	plant = models.ForeignKey('Plant', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	lines = models.ManyToManyField('PurchaseOrderLine',  blank=True, related_name='+')
	goodsReceipts = models.ManyToManyField('GoodsReceipt',  blank=True, related_name='+')
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in PurchaseOrderStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.poNumber
		str = str + self.orderDate
		str = str + self.totalAmount
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "PurchaseOrder";
    
	def objectType(self):
		return "PurchaseOrder";
