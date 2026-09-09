import datetime

from django.test import TestCase
from django.utils import timezone
from aerospaceOnDjango.models.MaintenanceWorkOrder import MaintenanceWorkOrder
from aerospaceOnDjango.delegates.MaintenanceWorkOrderDelegate import MaintenanceWorkOrderDelegate

 #======================================================================
# 
# Encapsulates data for model MaintenanceWorkOrder
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class MaintenanceWorkOrderTest Declaration
#======================================================================
class MaintenanceWorkOrderTest (TestCase) :
	def test_crud(self) :
		maintenanceWorkOrder = MaintenanceWorkOrder()
		maintenanceWorkOrder.workOrderNumber = "default workOrderNumber field value"
		maintenanceWorkOrder.status = "default status field value"
		
		delegate = MaintenanceWorkOrderDelegate()
		responseObj = delegate.create(maintenanceWorkOrder)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


