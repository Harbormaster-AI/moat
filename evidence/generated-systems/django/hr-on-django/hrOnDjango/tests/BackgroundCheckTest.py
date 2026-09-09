import datetime

from django.test import TestCase
from django.utils import timezone
from hrOnDjango.models.BackgroundCheck import BackgroundCheck
from hrOnDjango.delegates.BackgroundCheckDelegate import BackgroundCheckDelegate

 #======================================================================
# 
# Encapsulates data for model BackgroundCheck
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class BackgroundCheckTest Declaration
#======================================================================
class BackgroundCheckTest (TestCase) :
	def test_crud(self) :
		backgroundCheck = BackgroundCheck()
		backgroundCheck.checkNumber = "default checkNumber field value"
		backgroundCheck.provider = "default provider field value"
		backgroundCheck.completedDate = datetime.datetime.now()
		backgroundCheck.status = "default status field value"
		
		delegate = BackgroundCheckDelegate()
		responseObj = delegate.create(backgroundCheck)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


