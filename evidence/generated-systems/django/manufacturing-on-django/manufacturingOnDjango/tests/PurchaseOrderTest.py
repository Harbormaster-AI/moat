import datetime

from django.test import TestCase
from django.utils import timezone
from manufacturingOnDjango.models.PurchaseOrder import PurchaseOrder
from manufacturingOnDjango.delegates.PurchaseOrderDelegate import PurchaseOrderDelegate

 #======================================================================
# 
# Encapsulates data for model PurchaseOrder
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PurchaseOrderTest Declaration
#======================================================================
class PurchaseOrderTest (TestCase) :
	def test_crud(self) :
		purchaseOrder = PurchaseOrder()
		purchaseOrder.poNumber = "default poNumber field value"
		purchaseOrder.orderDate = datetime.datetime.now()
		purchaseOrder.totalAmount = "default totalAmount field value"
		purchaseOrder.status = "default status field value"
		
		delegate = PurchaseOrderDelegate()
		responseObj = delegate.create(purchaseOrder)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


