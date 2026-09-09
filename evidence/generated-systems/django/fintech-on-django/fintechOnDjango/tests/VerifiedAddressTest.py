import datetime

from django.test import TestCase
from django.utils import timezone
from fintechOnDjango.models.VerifiedAddress import VerifiedAddress
from fintechOnDjango.delegates.VerifiedAddressDelegate import VerifiedAddressDelegate

 #======================================================================
# 
# Encapsulates data for model VerifiedAddress
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class VerifiedAddressTest Declaration
#======================================================================
class VerifiedAddressTest (TestCase) :
	def test_crud(self) :
		verifiedAddress = VerifiedAddress()
		verifiedAddress.address = "default address field value"
		verifiedAddress.verifiedAt = "default verifiedAt field value"
		verifiedAddress.verificationStatus = "default verificationStatus field value"
		
		delegate = VerifiedAddressDelegate()
		responseObj = delegate.create(verifiedAddress)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


