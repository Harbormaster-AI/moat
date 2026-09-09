import datetime

from django.test import TestCase
from django.utils import timezone
from analyticsOnDjango.models.SemanticModel import SemanticModel
from analyticsOnDjango.delegates.SemanticModelDelegate import SemanticModelDelegate

 #======================================================================
# 
# Encapsulates data for model SemanticModel
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class SemanticModelTest Declaration
#======================================================================
class SemanticModelTest (TestCase) :
	def test_crud(self) :
		semanticModel = SemanticModel()
		semanticModel.name = "default name field value"
		semanticModel.version = "default version field value"
		semanticModel.grain = "default grain field value"
		
		delegate = SemanticModelDelegate()
		responseObj = delegate.create(semanticModel)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


