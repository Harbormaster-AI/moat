import datetime

from django.test import TestCase
from django.utils import timezone
from healthcareOnDjango.models.Authorization import Authorization
from healthcareOnDjango.delegates.AuthorizationDelegate import AuthorizationDelegate

 #======================================================================
# 
# Encapsulates data for model Authorization
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AuthorizationTest Declaration
#======================================================================
class AuthorizationTest (TestCase) :
	def test_crud(self) :
		authorization = Authorization()
		authorization.authNumber = "default authNumber field value"
		authorization.requestedService = "default requestedService field value"
		authorization.status = "default status field value"
		
		delegate = AuthorizationDelegate()
		responseObj = delegate.create(authorization)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


