import datetime

from django.test import TestCase
from django.utils import timezone
from aerospaceOnDjango.models.Registration import Registration
from aerospaceOnDjango.delegates.RegistrationDelegate import RegistrationDelegate

 #======================================================================
# 
# Encapsulates data for model Registration
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class RegistrationTest Declaration
#======================================================================
class RegistrationTest (TestCase) :
	def test_crud(self) :
		registration = Registration()
		registration.tailNumber = "default tailNumber field value"
		registration.registryCountry = "default registryCountry field value"
		
		delegate = RegistrationDelegate()
		responseObj = delegate.create(registration)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


