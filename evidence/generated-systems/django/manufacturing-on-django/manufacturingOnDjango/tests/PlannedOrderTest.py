import datetime

from django.test import TestCase
from django.utils import timezone
from manufacturingOnDjango.models.PlannedOrder import PlannedOrder
from manufacturingOnDjango.delegates.PlannedOrderDelegate import PlannedOrderDelegate

 #======================================================================
# 
# Encapsulates data for model PlannedOrder
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PlannedOrderTest Declaration
#======================================================================
class PlannedOrderTest (TestCase) :
	def test_crud(self) :
		plannedOrder = PlannedOrder()
		plannedOrder.plannedOrderNumber = "default plannedOrderNumber field value"
		plannedOrder.quantity = "default quantity field value"
		plannedOrder.dueDate = datetime.datetime.now()
		plannedOrder.orderType = "default orderType field value"
		plannedOrder.status = "default status field value"
		
		delegate = PlannedOrderDelegate()
		responseObj = delegate.create(plannedOrder)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


