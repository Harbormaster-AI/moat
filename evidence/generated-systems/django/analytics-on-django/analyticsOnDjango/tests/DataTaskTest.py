import datetime

from django.test import TestCase
from django.utils import timezone
from analyticsOnDjango.models.DataTask import DataTask
from analyticsOnDjango.delegates.DataTaskDelegate import DataTaskDelegate

 #======================================================================
# 
# Encapsulates data for model DataTask
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DataTaskTest Declaration
#======================================================================
class DataTaskTest (TestCase) :
	def test_crud(self) :
		dataTask = DataTask()
		dataTask.name = "default name field value"
		dataTask.command = "default command field value"
		dataTask.retries = 22
		dataTask.taskType = "default taskType field value"
		
		delegate = DataTaskDelegate()
		responseObj = delegate.create(dataTask)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


