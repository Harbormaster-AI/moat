import datetime

from django.test import TestCase
from django.utils import timezone
from inventoryOnDjango.models.CycleCount import CycleCount
from inventoryOnDjango.delegates.CycleCountDelegate import CycleCountDelegate

 #======================================================================
# 
# Encapsulates data for model CycleCount
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CycleCountTest Declaration
#======================================================================
class CycleCountTest (TestCase) :
	def test_crud(self) :
		cycleCount = CycleCount()
		cycleCount.countNumber = "default countNumber field value"
		cycleCount.scheduledDate = datetime.datetime.now()
		cycleCount.performedDate = datetime.datetime.now()
		cycleCount.approvedBy = "default approvedBy field value"
		cycleCount.status = "default status field value"
		
		delegate = CycleCountDelegate()
		responseObj = delegate.create(cycleCount)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


