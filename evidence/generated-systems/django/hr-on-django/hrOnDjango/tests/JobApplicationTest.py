import datetime

from django.test import TestCase
from django.utils import timezone
from hrOnDjango.models.JobApplication import JobApplication
from hrOnDjango.delegates.JobApplicationDelegate import JobApplicationDelegate

 #======================================================================
# 
# Encapsulates data for model JobApplication
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class JobApplicationTest Declaration
#======================================================================
class JobApplicationTest (TestCase) :
	def test_crud(self) :
		jobApplication = JobApplication()
		jobApplication.applicationNumber = "default applicationNumber field value"
		jobApplication.appliedDate = datetime.datetime.now()
		jobApplication.resumeUrl = "default resumeUrl field value"
		jobApplication.status = "default status field value"
		
		delegate = JobApplicationDelegate()
		responseObj = delegate.create(jobApplication)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


