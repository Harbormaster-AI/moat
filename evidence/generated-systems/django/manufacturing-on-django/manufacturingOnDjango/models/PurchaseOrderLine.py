from django.db import models

#======================================================================
# 
# Encapsulates data for model PurchaseOrderLine
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PurchaseOrderLine Declaration
#======================================================================
class PurchaseOrderLine (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	lineNumber = models.IntegerField(null=True)
	quantity = Quantity
	unitPrice = Money
	dueDate = models.DateField(null=True)
	purchaseOrder = models.ForeignKey('PurchaseOrder', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	item = models.ForeignKey('Item', on_delete=models.CASCADE, null=True, blank=True, related_name='+')

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.lineNumber
		str = str + self.quantity
		str = str + self.unitPrice
		str = str + self.dueDate
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "PurchaseOrderLine";
    
	def objectType(self):
		return "PurchaseOrderLine";
