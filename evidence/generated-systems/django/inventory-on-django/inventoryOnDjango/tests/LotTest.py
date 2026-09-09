import datetime

from django.test import TestCase
from django.utils import timezone
from inventoryOnDjango.models.Lot import Lot
from inventoryOnDjango.delegates.LotDelegate import LotDelegate

 #======================================================================
# 
# Encapsulates data for model Lot
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class LotTest Declaration
#======================================================================
class LotTest (TestCase) :
	def test_crud(self) :
		lot = Lot()
		lot.batchNumber = "default batchNumber field value"
		lot.manufactureDate = datetime.datetime.now()
		lot.expirationDate = datetime.datetime.now()
		lot.lotStatus = "default lotStatus field value"
		
		delegate = LotDelegate()
		responseObj = delegate.create(lot)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


