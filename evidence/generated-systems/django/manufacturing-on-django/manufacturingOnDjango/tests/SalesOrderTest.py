import datetime

from django.test import TestCase
from django.utils import timezone
from manufacturingOnDjango.models.SalesOrder import SalesOrder
from manufacturingOnDjango.delegates.SalesOrderDelegate import SalesOrderDelegate

 #======================================================================
# 
# Encapsulates data for model SalesOrder
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class SalesOrderTest Declaration
#======================================================================
class SalesOrderTest (TestCase) :
	def test_crud(self) :
		salesOrder = SalesOrder()
		salesOrder.orderNumber = "default orderNumber field value"
		salesOrder.orderDate = datetime.datetime.now()
		salesOrder.totalAmount = "default totalAmount field value"
		salesOrder.status = "default status field value"
		
		delegate = SalesOrderDelegate()
		responseObj = delegate.create(salesOrder)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


