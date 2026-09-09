import datetime

from django.test import TestCase
from django.utils import timezone
from advertisingOnDjango.models.InsertionOrder import InsertionOrder
from advertisingOnDjango.delegates.InsertionOrderDelegate import InsertionOrderDelegate

 #======================================================================
# 
# Encapsulates data for model InsertionOrder
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class InsertionOrderTest Declaration
#======================================================================
class InsertionOrderTest (TestCase) :
	def test_crud(self) :
		insertionOrder = InsertionOrder()
		insertionOrder.ioNumber = "default ioNumber field value"
		insertionOrder.agreedBudget = "default agreedBudget field value"
		insertionOrder.flight = "default flight field value"
		insertionOrder.status = "default status field value"
		
		delegate = InsertionOrderDelegate()
		responseObj = delegate.create(insertionOrder)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


