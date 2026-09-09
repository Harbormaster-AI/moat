import datetime

from django.test import TestCase
from django.utils import timezone
from aerospaceOnDjango.models.ProductionOrder import ProductionOrder
from aerospaceOnDjango.delegates.ProductionOrderDelegate import ProductionOrderDelegate

 #======================================================================
# 
# Encapsulates data for model ProductionOrder
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ProductionOrderTest Declaration
#======================================================================
class ProductionOrderTest (TestCase) :
	def test_crud(self) :
		productionOrder = ProductionOrder()
		productionOrder.orderNumber = "default orderNumber field value"
		productionOrder.status = "default status field value"
		
		delegate = ProductionOrderDelegate()
		responseObj = delegate.create(productionOrder)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


