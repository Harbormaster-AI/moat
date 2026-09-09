import datetime

from django.test import TestCase
from django.utils import timezone
from inventoryOnDjango.models.TransferOrder import TransferOrder
from inventoryOnDjango.delegates.TransferOrderDelegate import TransferOrderDelegate

 #======================================================================
# 
# Encapsulates data for model TransferOrder
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class TransferOrderTest Declaration
#======================================================================
class TransferOrderTest (TestCase) :
	def test_crud(self) :
		transferOrder = TransferOrder()
		transferOrder.orderNumber = "default orderNumber field value"
		transferOrder.requestedShipDate = datetime.datetime.now()
		transferOrder.requestedReceiveDate = datetime.datetime.now()
		transferOrder.shippedDate = datetime.datetime.now()
		transferOrder.receivedDate = datetime.datetime.now()
		transferOrder.status = "default status field value"
		
		delegate = TransferOrderDelegate()
		responseObj = delegate.create(transferOrder)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


