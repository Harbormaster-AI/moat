import datetime

from django.test import TestCase
from django.utils import timezone
from advertisingOnDjango.models.Advertiser import Advertiser
from advertisingOnDjango.delegates.AdvertiserDelegate import AdvertiserDelegate

 #======================================================================
# 
# Encapsulates data for model Advertiser
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AdvertiserTest Declaration
#======================================================================
class AdvertiserTest (TestCase) :
	def test_crud(self) :
		advertiser = Advertiser()
		advertiser.name = "default name field value"
		advertiser.legalName = "default legalName field value"
		advertiser.industry = "default industry field value"
		advertiser.website = "default website field value"
		
		delegate = AdvertiserDelegate()
		responseObj = delegate.create(advertiser)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


