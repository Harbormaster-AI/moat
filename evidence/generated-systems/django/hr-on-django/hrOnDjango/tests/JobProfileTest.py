import datetime

from django.test import TestCase
from django.utils import timezone
from hrOnDjango.models.JobProfile import JobProfile
from hrOnDjango.delegates.JobProfileDelegate import JobProfileDelegate

 #======================================================================
# 
# Encapsulates data for model JobProfile
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class JobProfileTest Declaration
#======================================================================
class JobProfileTest (TestCase) :
	def test_crud(self) :
		jobProfile = JobProfile()
		jobProfile.title = "default title field value"
		jobProfile.jobCode = "default jobCode field value"
		jobProfile.jobLevel = "default jobLevel field value"
		jobProfile.exemptStatus = "default exemptStatus field value"
		
		delegate = JobProfileDelegate()
		responseObj = delegate.create(jobProfile)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


