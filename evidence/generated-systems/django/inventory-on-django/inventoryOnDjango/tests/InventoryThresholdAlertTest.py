import datetime

from django.test import TestCase
from django.utils import timezone
from inventoryOnDjango.models.InventoryThresholdAlert import InventoryThresholdAlert
from inventoryOnDjango.delegates.InventoryThresholdAlertDelegate import InventoryThresholdAlertDelegate

 #======================================================================
# 
# Encapsulates data for model InventoryThresholdAlert
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class InventoryThresholdAlertTest Declaration
#======================================================================
class InventoryThresholdAlertTest (TestCase) :
	def test_crud(self) :
		inventoryThresholdAlert = InventoryThresholdAlert()
		inventoryThresholdAlert.alertNumber = "default alertNumber field value"
		inventoryThresholdAlert.detectedAt = datetime.datetime.now()
		inventoryThresholdAlert.message = "default message field value"
		inventoryThresholdAlert.alertType = "default alertType field value"
		inventoryThresholdAlert.status = "default status field value"
		
		delegate = InventoryThresholdAlertDelegate()
		responseObj = delegate.create(inventoryThresholdAlert)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


