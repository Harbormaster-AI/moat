import datetime

from django.test import TestCase
from django.utils import timezone
from ecommerceOnDjango.models.MediaAsset import MediaAsset
from ecommerceOnDjango.delegates.MediaAssetDelegate import MediaAssetDelegate

 #======================================================================
# 
# Encapsulates data for model MediaAsset
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class MediaAssetTest Declaration
#======================================================================
class MediaAssetTest (TestCase) :
	def test_crud(self) :
		mediaAsset = MediaAsset()
		mediaAsset.url = "default url field value"
		mediaAsset.altText = "default altText field value"
		mediaAsset.position = 22
		mediaAsset.mediaType = "default mediaType field value"
		
		delegate = MediaAssetDelegate()
		responseObj = delegate.create(mediaAsset)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


