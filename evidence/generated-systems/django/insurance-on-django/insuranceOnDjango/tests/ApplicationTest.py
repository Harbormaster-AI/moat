import datetime

from django.test import TestCase
from django.utils import timezone
from insuranceOnDjango.models.Application import Application
from insuranceOnDjango.delegates.ApplicationDelegate import ApplicationDelegate

 #======================================================================
# 
# Encapsulates data for model Application
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ApplicationTest Declaration
#======================================================================
class ApplicationTest (TestCase) :
	def test_crud(self) :
		application = Application()
		application.applicationNumber = "default applicationNumber field value"
		application.submissionDate = datetime.datetime.now()
		application.status = "default status field value"
		
		delegate = ApplicationDelegate()
		responseObj = delegate.create(application)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


