import datetime

from django.test import TestCase
from django.utils import timezone
from analyticsOnDjango.models.InferenceEndpoint import InferenceEndpoint
from analyticsOnDjango.delegates.InferenceEndpointDelegate import InferenceEndpointDelegate

 #======================================================================
# 
# Encapsulates data for model InferenceEndpoint
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class InferenceEndpointTest Declaration
#======================================================================
class InferenceEndpointTest (TestCase) :
	def test_crud(self) :
		inferenceEndpoint = InferenceEndpoint()
		inferenceEndpoint.name = "default name field value"
		inferenceEndpoint.endpointUrl = "default endpointUrl field value"
		inferenceEndpoint.trafficShare = "default trafficShare field value"
		inferenceEndpoint.mode = "default mode field value"
		
		delegate = InferenceEndpointDelegate()
		responseObj = delegate.create(inferenceEndpoint)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


