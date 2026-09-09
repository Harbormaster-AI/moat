import datetime

from django.test import TestCase
from django.utils import timezone
from analyticsOnDjango.models.DataPipeline import DataPipeline
from analyticsOnDjango.delegates.DataPipelineDelegate import DataPipelineDelegate

 #======================================================================
# 
# Encapsulates data for model DataPipeline
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DataPipelineTest Declaration
#======================================================================
class DataPipelineTest (TestCase) :
	def test_crud(self) :
		dataPipeline = DataPipeline()
		dataPipeline.name = "default name field value"
		dataPipeline.schedule = "default schedule field value"
		dataPipeline.triggerType = "default triggerType field value"
		dataPipeline.status = "default status field value"
		
		delegate = DataPipelineDelegate()
		responseObj = delegate.create(dataPipeline)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


