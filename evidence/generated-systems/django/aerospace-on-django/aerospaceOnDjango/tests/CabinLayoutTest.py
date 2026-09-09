import datetime

from django.test import TestCase
from django.utils import timezone
from aerospaceOnDjango.models.CabinLayout import CabinLayout
from aerospaceOnDjango.delegates.CabinLayoutDelegate import CabinLayoutDelegate

 #======================================================================
# 
# Encapsulates data for model CabinLayout
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CabinLayoutTest Declaration
#======================================================================
class CabinLayoutTest (TestCase) :
	def test_crud(self) :
		cabinLayout = CabinLayout()
		cabinLayout.layoutCode = "default layoutCode field value"
		cabinLayout.totalSeats = 22
		cabinLayout.classLayout = "default classLayout field value"
		
		delegate = CabinLayoutDelegate()
		responseObj = delegate.create(cabinLayout)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


