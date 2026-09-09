import datetime

from django.test import TestCase
from django.utils import timezone
from inventoryOnDjango.models.CycleCountEntry import CycleCountEntry
from inventoryOnDjango.delegates.CycleCountEntryDelegate import CycleCountEntryDelegate

 #======================================================================
# 
# Encapsulates data for model CycleCountEntry
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CycleCountEntryTest Declaration
#======================================================================
class CycleCountEntryTest (TestCase) :
	def test_crud(self) :
		cycleCountEntry = CycleCountEntry()
		cycleCountEntry.lineNumber = 22
		cycleCountEntry.systemQuantity = "default systemQuantity field value"
		cycleCountEntry.countedQuantity = "default countedQuantity field value"
		cycleCountEntry.varianceQuantity = "default varianceQuantity field value"
		cycleCountEntry.recountRequired = False
		cycleCountEntry.stockStatus = "default stockStatus field value"
		
		delegate = CycleCountEntryDelegate()
		responseObj = delegate.create(cycleCountEntry)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


