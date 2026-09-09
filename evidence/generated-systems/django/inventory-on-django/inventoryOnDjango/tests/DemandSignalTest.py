import datetime

from django.test import TestCase
from django.utils import timezone
from inventoryOnDjango.models.DemandSignal import DemandSignal
from inventoryOnDjango.delegates.DemandSignalDelegate import DemandSignalDelegate

 #======================================================================
# 
# Encapsulates data for model DemandSignal
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DemandSignalTest Declaration
#======================================================================
class DemandSignalTest (TestCase) :
	def test_crud(self) :
		demandSignal = DemandSignal()
		demandSignal.externalReference = "default externalReference field value"
		demandSignal.requestedDate = datetime.datetime.now()
		demandSignal.quantity = "default quantity field value"
		demandSignal.demandType = "default demandType field value"
		
		delegate = DemandSignalDelegate()
		responseObj = delegate.create(demandSignal)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


