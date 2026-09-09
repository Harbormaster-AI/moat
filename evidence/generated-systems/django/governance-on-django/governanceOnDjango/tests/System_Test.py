import datetime

from django.test import TestCase
from django.utils import timezone
from governanceOnDjango.models.System_ import System_
from governanceOnDjango.delegates.System_Delegate import System_Delegate

 #======================================================================
# 
# Encapsulates data for model System_
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class System_Test Declaration
#======================================================================
class System_Test (TestCase) :
	def test_crud(self) :
		system_ = System_()
		system_.name = "default name field value"
		system_.ownerDepartment = "default ownerDepartment field value"
		system_.systemType = "default systemType field value"
		
		delegate = System_Delegate()
		responseObj = delegate.create(system_)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


