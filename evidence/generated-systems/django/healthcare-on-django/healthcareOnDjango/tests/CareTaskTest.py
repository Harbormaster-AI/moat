import datetime

from django.test import TestCase
from django.utils import timezone
from healthcareOnDjango.models.CareTask import CareTask
from healthcareOnDjango.delegates.CareTaskDelegate import CareTaskDelegate

 #======================================================================
# 
# Encapsulates data for model CareTask
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CareTaskTest Declaration
#======================================================================
class CareTaskTest (TestCase) :
	def test_crud(self) :
		careTask = CareTask()
		careTask.description = "default description field value"
		careTask.dueDate = datetime.datetime.now()
		careTask.status = "default status field value"
		careTask.priority = "default priority field value"
		
		delegate = CareTaskDelegate()
		responseObj = delegate.create(careTask)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


