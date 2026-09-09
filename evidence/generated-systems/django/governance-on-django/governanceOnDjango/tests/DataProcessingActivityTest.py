import datetime

from django.test import TestCase
from django.utils import timezone
from governanceOnDjango.models.DataProcessingActivity import DataProcessingActivity
from governanceOnDjango.delegates.DataProcessingActivityDelegate import DataProcessingActivityDelegate

 #======================================================================
# 
# Encapsulates data for model DataProcessingActivity
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DataProcessingActivityTest Declaration
#======================================================================
class DataProcessingActivityTest (TestCase) :
	def test_crud(self) :
		dataProcessingActivity = DataProcessingActivity()
		dataProcessingActivity.name = "default name field value"
		dataProcessingActivity.purpose = "default purpose field value"
		dataProcessingActivity.startDate = datetime.datetime.now()
		dataProcessingActivity.lawfulBasis = "default lawfulBasis field value"
		
		delegate = DataProcessingActivityDelegate()
		responseObj = delegate.create(dataProcessingActivity)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


