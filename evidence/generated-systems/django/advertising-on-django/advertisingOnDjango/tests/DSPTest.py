import datetime

from django.test import TestCase
from django.utils import timezone
from advertisingOnDjango.models.DSP import DSP
from advertisingOnDjango.delegates.DSPDelegate import DSPDelegate

 #======================================================================
# 
# Encapsulates data for model DSP
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DSPTest Declaration
#======================================================================
class DSPTest (TestCase) :
	def test_crud(self) :
		dSP = DSP()
		dSP.name = "default name field value"
		dSP.website = "default website field value"
		dSP.region = "default region field value"
		
		delegate = DSPDelegate()
		responseObj = delegate.create(dSP)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


