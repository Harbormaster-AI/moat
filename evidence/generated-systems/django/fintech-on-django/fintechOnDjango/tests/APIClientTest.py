import datetime

from django.test import TestCase
from django.utils import timezone
from fintechOnDjango.models.APIClient import APIClient
from fintechOnDjango.delegates.APIClientDelegate import APIClientDelegate

 #======================================================================
# 
# Encapsulates data for model APIClient
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class APIClientTest Declaration
#======================================================================
class APIClientTest (TestCase) :
	def test_crud(self) :
		aPIClient = APIClient()
		aPIClient.name = "default name field value"
		aPIClient.clientId = "default clientId field value"
		aPIClient.redirectUri = "default redirectUri field value"
		aPIClient.clientType = "default clientType field value"
		
		delegate = APIClientDelegate()
		responseObj = delegate.create(aPIClient)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


