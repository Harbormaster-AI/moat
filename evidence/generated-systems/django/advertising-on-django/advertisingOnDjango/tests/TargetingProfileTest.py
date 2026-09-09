import datetime

from django.test import TestCase
from django.utils import timezone
from advertisingOnDjango.models.TargetingProfile import TargetingProfile
from advertisingOnDjango.delegates.TargetingProfileDelegate import TargetingProfileDelegate

 #======================================================================
# 
# Encapsulates data for model TargetingProfile
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class TargetingProfileTest Declaration
#======================================================================
class TargetingProfileTest (TestCase) :
	def test_crud(self) :
		targetingProfile = TargetingProfile()
		targetingProfile.name = "default name field value"
		
		delegate = TargetingProfileDelegate()
		responseObj = delegate.create(targetingProfile)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


