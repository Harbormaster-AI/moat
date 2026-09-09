import datetime

from django.test import TestCase
from django.utils import timezone
from analyticsOnDjango.models.RunParameter import RunParameter
from analyticsOnDjango.delegates.RunParameterDelegate import RunParameterDelegate

 #======================================================================
# 
# Encapsulates data for model RunParameter
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class RunParameterTest Declaration
#======================================================================
class RunParameterTest (TestCase) :
	def test_crud(self) :
		runParameter = RunParameter()
		runParameter.name = "default name field value"
		runParameter.value = "default value field value"
		
		delegate = RunParameterDelegate()
		responseObj = delegate.create(runParameter)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


