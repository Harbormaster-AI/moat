import datetime

from django.test import TestCase
from django.utils import timezone
from aerospaceOnDjango.models.EngineType import EngineType
from aerospaceOnDjango.delegates.EngineTypeDelegate import EngineTypeDelegate

 #======================================================================
# 
# Encapsulates data for model EngineType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class EngineTypeTest Declaration
#======================================================================
class EngineTypeTest (TestCase) :
	def test_crud(self) :
		engineType = EngineType()
		engineType.engineModelCode = "default engineModelCode field value"
		engineType.maxThrustKn = "default maxThrustKn field value"
		engineType.category = "default category field value"
		
		delegate = EngineTypeDelegate()
		responseObj = delegate.create(engineType)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


