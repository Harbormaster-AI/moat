import datetime

from django.test import TestCase
from django.utils import timezone
from fintechOnDjango.models.KYCProfile import KYCProfile
from fintechOnDjango.delegates.KYCProfileDelegate import KYCProfileDelegate

 #======================================================================
# 
# Encapsulates data for model KYCProfile
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class KYCProfileTest Declaration
#======================================================================
class KYCProfileTest (TestCase) :
	def test_crud(self) :
		kYCProfile = KYCProfile()
		kYCProfile.profileId = "default profileId field value"
		kYCProfile.createdAt = "default createdAt field value"
		kYCProfile.status = "default status field value"
		kYCProfile.verificationLevel = "default verificationLevel field value"
		
		delegate = KYCProfileDelegate()
		responseObj = delegate.create(kYCProfile)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


