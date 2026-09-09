import datetime

from django.test import TestCase
from django.utils import timezone
from aerospaceOnDjango.models.AvionicsSuite import AvionicsSuite
from aerospaceOnDjango.delegates.AvionicsSuiteDelegate import AvionicsSuiteDelegate

 #======================================================================
# 
# Encapsulates data for model AvionicsSuite
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AvionicsSuiteTest Declaration
#======================================================================
class AvionicsSuiteTest (TestCase) :
	def test_crud(self) :
		avionicsSuite = AvionicsSuite()
		avionicsSuite.suiteName = "default suiteName field value"
		avionicsSuite.softwareBaseline = "default softwareBaseline field value"
		
		delegate = AvionicsSuiteDelegate()
		responseObj = delegate.create(avionicsSuite)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


