import datetime

from django.test import TestCase
from django.utils import timezone
from manufacturingOnDjango.models.WorkOrder import WorkOrder
from manufacturingOnDjango.delegates.WorkOrderDelegate import WorkOrderDelegate

 #======================================================================
# 
# Encapsulates data for model WorkOrder
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class WorkOrderTest Declaration
#======================================================================
class WorkOrderTest (TestCase) :
	def test_crud(self) :
		workOrder = WorkOrder()
		workOrder.workOrderNumber = "default workOrderNumber field value"
		workOrder.plannedStart = "default plannedStart field value"
		workOrder.plannedEnd = "default plannedEnd field value"
		workOrder.quantity = "default quantity field value"
		workOrder.priority = 22
		workOrder.status = "default status field value"
		
		delegate = WorkOrderDelegate()
		responseObj = delegate.create(workOrder)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


