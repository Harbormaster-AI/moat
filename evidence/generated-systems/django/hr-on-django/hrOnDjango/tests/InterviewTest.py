import datetime

from django.test import TestCase
from django.utils import timezone
from hrOnDjango.models.Interview import Interview
from hrOnDjango.delegates.InterviewDelegate import InterviewDelegate

 #======================================================================
# 
# Encapsulates data for model Interview
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class InterviewTest Declaration
#======================================================================
class InterviewTest (TestCase) :
	def test_crud(self) :
		interview = Interview()
		interview.interviewDate = datetime.datetime.now()
		interview.feedback = "default feedback field value"
		interview.stage = "default stage field value"
		interview.result = "default result field value"
		
		delegate = InterviewDelegate()
		responseObj = delegate.create(interview)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


