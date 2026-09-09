import datetime

from django.test import TestCase
from django.utils import timezone
from manufacturingOnDjango.models.Asset import Asset
from manufacturingOnDjango.delegates.AssetDelegate import AssetDelegate

 #======================================================================
# 
# Encapsulates data for model Asset
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AssetTest Declaration
#======================================================================
class AssetTest (TestCase) :
	def test_crud(self) :
		asset = Asset()
		asset.assetTag = "default assetTag field value"
		asset.assetName = "default assetName field value"
		asset.commissioningDate = datetime.datetime.now()
		asset.assetStatus = "default assetStatus field value"
		
		delegate = AssetDelegate()
		responseObj = delegate.create(asset)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


