import datetime

from django.test import TestCase
from django.utils import timezone
from analyticsOnDjango.models.Feature import Feature
from analyticsOnDjango.delegates.FeatureDelegate import FeatureDelegate

 #======================================================================
# 
# Encapsulates data for model Feature
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class FeatureTest Declaration
#======================================================================
class FeatureTest (TestCase) :
	def test_crud(self) :
		feature = Feature()
		feature.name = "default name field value"
		feature.description = "default description field value"
		feature.dataType = "default dataType field value"
		
		delegate = FeatureDelegate()
		responseObj = delegate.create(feature)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


