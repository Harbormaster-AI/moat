import datetime

from django.test import TestCase
from django.utils import timezone
from analyticsOnDjango.models.FeatureSet import FeatureSet
from analyticsOnDjango.delegates.FeatureSetDelegate import FeatureSetDelegate

 #======================================================================
# 
# Encapsulates data for model FeatureSet
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class FeatureSetTest Declaration
#======================================================================
class FeatureSetTest (TestCase) :
	def test_crud(self) :
		featureSet = FeatureSet()
		featureSet.name = "default name field value"
		featureSet.refreshSchedule = "default refreshSchedule field value"
		featureSet.storeType = "default storeType field value"
		
		delegate = FeatureSetDelegate()
		responseObj = delegate.create(featureSet)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


