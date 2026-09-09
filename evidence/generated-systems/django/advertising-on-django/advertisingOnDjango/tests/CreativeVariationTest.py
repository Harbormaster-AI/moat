import datetime

from django.test import TestCase
from django.utils import timezone
from advertisingOnDjango.models.CreativeVariation import CreativeVariation
from advertisingOnDjango.delegates.CreativeVariationDelegate import CreativeVariationDelegate

 #======================================================================
# 
# Encapsulates data for model CreativeVariation
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CreativeVariationTest Declaration
#======================================================================
class CreativeVariationTest (TestCase) :
	def test_crud(self) :
		creativeVariation = CreativeVariation()
		creativeVariation.name = "default name field value"
		creativeVariation.language = "default language field value"
		creativeVariation.headline = "default headline field value"
		creativeVariation.bodyText = "default bodyText field value"
		creativeVariation.callToAction = "default callToAction field value"
		
		delegate = CreativeVariationDelegate()
		responseObj = delegate.create(creativeVariation)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


