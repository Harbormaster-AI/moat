import datetime

from django.test import TestCase
from django.utils import timezone
from analyticsOnDjango.models.ModelVersion import ModelVersion
from analyticsOnDjango.delegates.ModelVersionDelegate import ModelVersionDelegate

 #======================================================================
# 
# Encapsulates data for model ModelVersion
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ModelVersionTest Declaration
#======================================================================
class ModelVersionTest (TestCase) :
	def test_crud(self) :
		modelVersion = ModelVersion()
		modelVersion.version = "default version field value"
		modelVersion.lifecycle = "default lifecycle field value"
		modelVersion.trainingStatus = "default trainingStatus field value"
		
		delegate = ModelVersionDelegate()
		responseObj = delegate.create(modelVersion)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


