import datetime

from django.test import TestCase
from django.utils import timezone
from analyticsOnDjango.models.Model import Model
from analyticsOnDjango.delegates.ModelDelegate import ModelDelegate

 #======================================================================
# 
# Encapsulates data for model Model
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ModelTest Declaration
#======================================================================
class ModelTest (TestCase) :
	def test_crud(self) :
		model = Model()
		model.name = "default name field value"
		model.taskDescription = "default taskDescription field value"
		model.modelType = "default modelType field value"
		
		delegate = ModelDelegate()
		responseObj = delegate.create(model)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


