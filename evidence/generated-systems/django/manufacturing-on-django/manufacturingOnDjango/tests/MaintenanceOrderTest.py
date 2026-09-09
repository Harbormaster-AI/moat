import datetime

from django.test import TestCase
from django.utils import timezone
from manufacturingOnDjango.models.MaintenanceOrder import MaintenanceOrder
from manufacturingOnDjango.delegates.MaintenanceOrderDelegate import MaintenanceOrderDelegate

 #======================================================================
# 
# Encapsulates data for model MaintenanceOrder
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class MaintenanceOrderTest Declaration
#======================================================================
class MaintenanceOrderTest (TestCase) :
	def test_crud(self) :
		maintenanceOrder = MaintenanceOrder()
		maintenanceOrder.orderNumber = "default orderNumber field value"
		maintenanceOrder.priority = 22
		maintenanceOrder.requestedDate = datetime.datetime.now()
		maintenanceOrder.completionDate = datetime.datetime.now()
		maintenanceOrder.status = "default status field value"
		
		delegate = MaintenanceOrderDelegate()
		responseObj = delegate.create(maintenanceOrder)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


