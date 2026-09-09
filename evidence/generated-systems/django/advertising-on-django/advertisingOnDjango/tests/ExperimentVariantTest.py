import datetime

from django.test import TestCase
from django.utils import timezone
from advertisingOnDjango.models.ExperimentVariant import ExperimentVariant
from advertisingOnDjango.delegates.ExperimentVariantDelegate import ExperimentVariantDelegate

 #======================================================================
# 
# Encapsulates data for model ExperimentVariant
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ExperimentVariantTest Declaration
#======================================================================
class ExperimentVariantTest (TestCase) :
	def test_crud(self) :
		experimentVariant = ExperimentVariant()
		experimentVariant.name = "default name field value"
		experimentVariant.allocation = "default allocation field value"
		
		delegate = ExperimentVariantDelegate()
		responseObj = delegate.create(experimentVariant)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


