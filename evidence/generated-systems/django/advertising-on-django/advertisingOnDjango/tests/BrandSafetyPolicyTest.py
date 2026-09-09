import datetime

from django.test import TestCase
from django.utils import timezone
from advertisingOnDjango.models.BrandSafetyPolicy import BrandSafetyPolicy
from advertisingOnDjango.delegates.BrandSafetyPolicyDelegate import BrandSafetyPolicyDelegate

 #======================================================================
# 
# Encapsulates data for model BrandSafetyPolicy
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class BrandSafetyPolicyTest Declaration
#======================================================================
class BrandSafetyPolicyTest (TestCase) :
	def test_crud(self) :
		brandSafetyPolicy = BrandSafetyPolicy()
		brandSafetyPolicy.level = "default level field value"
		brandSafetyPolicy.contentRatingThreshold = "default contentRatingThreshold field value"
		
		delegate = BrandSafetyPolicyDelegate()
		responseObj = delegate.create(brandSafetyPolicy)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


