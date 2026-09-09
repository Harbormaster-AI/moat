import datetime

from django.test import TestCase
from django.utils import timezone
from hrOnDjango.models.WorkAuthorization import WorkAuthorization
from hrOnDjango.delegates.WorkAuthorizationDelegate import WorkAuthorizationDelegate

 #======================================================================
# 
# Encapsulates data for model WorkAuthorization
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class WorkAuthorizationTest Declaration
#======================================================================
class WorkAuthorizationTest (TestCase) :
	def test_crud(self) :
		workAuthorization = WorkAuthorization()
		workAuthorization.country = "default country field value"
		workAuthorization.expirationDate = datetime.datetime.now()
		workAuthorization.status = "default status field value"
		
		delegate = WorkAuthorizationDelegate()
		responseObj = delegate.create(workAuthorization)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


