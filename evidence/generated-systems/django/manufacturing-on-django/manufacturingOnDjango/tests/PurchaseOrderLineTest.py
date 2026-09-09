import datetime

from django.test import TestCase
from django.utils import timezone
from manufacturingOnDjango.models.PurchaseOrderLine import PurchaseOrderLine
from manufacturingOnDjango.delegates.PurchaseOrderLineDelegate import PurchaseOrderLineDelegate

 #======================================================================
# 
# Encapsulates data for model PurchaseOrderLine
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PurchaseOrderLineTest Declaration
#======================================================================
class PurchaseOrderLineTest (TestCase) :
	def test_crud(self) :
		purchaseOrderLine = PurchaseOrderLine()
		purchaseOrderLine.lineNumber = 22
		purchaseOrderLine.quantity = "default quantity field value"
		purchaseOrderLine.unitPrice = "default unitPrice field value"
		purchaseOrderLine.dueDate = datetime.datetime.now()
		
		delegate = PurchaseOrderLineDelegate()
		responseObj = delegate.create(purchaseOrderLine)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


