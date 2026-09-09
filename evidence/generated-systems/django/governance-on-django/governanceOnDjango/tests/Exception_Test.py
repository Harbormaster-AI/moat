import datetime

from django.test import TestCase
from django.utils import timezone
from governanceOnDjango.models.Exception_ import Exception_
from governanceOnDjango.delegates.Exception_Delegate import Exception_Delegate

 #======================================================================
# 
# Encapsulates data for model Exception_
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Exception_Test Declaration
#======================================================================
class Exception_Test (TestCase) :
	def test_crud(self) :
		exception_ = Exception_()
		exception_.title = "default title field value"
		exception_.justification = "default justification field value"
		exception_.startDate = datetime.datetime.now()
		exception_.endDate = datetime.datetime.now()
		exception_.exceptionType = "default exceptionType field value"
		exception_.status = "default status field value"
		
		delegate = Exception_Delegate()
		responseObj = delegate.create(exception_)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


