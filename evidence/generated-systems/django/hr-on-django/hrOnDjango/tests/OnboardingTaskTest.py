import datetime

from django.test import TestCase
from django.utils import timezone
from hrOnDjango.models.OnboardingTask import OnboardingTask
from hrOnDjango.delegates.OnboardingTaskDelegate import OnboardingTaskDelegate

 #======================================================================
# 
# Encapsulates data for model OnboardingTask
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class OnboardingTaskTest Declaration
#======================================================================
class OnboardingTaskTest (TestCase) :
	def test_crud(self) :
		onboardingTask = OnboardingTask()
		onboardingTask.taskNumber = "default taskNumber field value"
		onboardingTask.name = "default name field value"
		onboardingTask.dueDate = datetime.datetime.now()
		onboardingTask.status = "default status field value"
		
		delegate = OnboardingTaskDelegate()
		responseObj = delegate.create(onboardingTask)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


