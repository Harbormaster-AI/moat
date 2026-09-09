import datetime

from django.test import TestCase
from django.utils import timezone
from aerospaceOnDjango.models.APU import APU
from aerospaceOnDjango.delegates.APUDelegate import APUDelegate

 #======================================================================
# 
# Encapsulates data for model APU
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class APUTest Declaration
#======================================================================
class APUTest (TestCase) :
	def test_crud(self) :
		aPU = APU()
		aPU.model = "default model field value"
		
		delegate = APUDelegate()
		responseObj = delegate.create(aPU)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


