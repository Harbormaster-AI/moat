import datetime

from django.test import TestCase
from django.utils import timezone
from aerospaceOnDjango.models.AircraftOrder import AircraftOrder
from aerospaceOnDjango.delegates.AircraftOrderDelegate import AircraftOrderDelegate

 #======================================================================
# 
# Encapsulates data for model AircraftOrder
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AircraftOrderTest Declaration
#======================================================================
class AircraftOrderTest (TestCase) :
	def test_crud(self) :
		aircraftOrder = AircraftOrder()
		aircraftOrder.orderNumber = "default orderNumber field value"
		aircraftOrder.totalAmount = "default totalAmount field value"
		aircraftOrder.status = "default status field value"
		
		delegate = AircraftOrderDelegate()
		responseObj = delegate.create(aircraftOrder)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


