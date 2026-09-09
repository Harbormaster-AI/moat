import datetime

from django.test import TestCase
from django.utils import timezone
from advertisingOnDjango.models.CreativeAsset import CreativeAsset
from advertisingOnDjango.delegates.CreativeAssetDelegate import CreativeAssetDelegate

 #======================================================================
# 
# Encapsulates data for model CreativeAsset
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CreativeAssetTest Declaration
#======================================================================
class CreativeAssetTest (TestCase) :
	def test_crud(self) :
		creativeAsset = CreativeAsset()
		creativeAsset.name = "default name field value"
		creativeAsset.clickUrl = "default clickUrl field value"
		creativeAsset.landingPage = "default landingPage field value"
		creativeAsset.width = 22
		creativeAsset.height = 22
		creativeAsset.durationSeconds = 22
		creativeAsset.creativeType = "default creativeType field value"
		creativeAsset.adFormat = "default adFormat field value"
		
		delegate = CreativeAssetDelegate()
		responseObj = delegate.create(creativeAsset)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


