import datetime

from django.test import TestCase
from django.utils import timezone
from advertisingOnDjango.models.AdSlot import AdSlot
from advertisingOnDjango.delegates.AdSlotDelegate import AdSlotDelegate

 #======================================================================
# 
# Encapsulates data for model AdSlot
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AdSlotTest Declaration
#======================================================================
class AdSlotTest (TestCase) :
	def test_crud(self) :
		adSlot = AdSlot()
		adSlot.slotCode = "default slotCode field value"
		adSlot.width = 22
		adSlot.height = 22
		adSlot.floorPrice = "default floorPrice field value"
		adSlot.format = "default format field value"
		
		delegate = AdSlotDelegate()
		responseObj = delegate.create(adSlot)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


