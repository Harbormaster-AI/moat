import datetime

from django.test import TestCase
from django.utils import timezone
from governanceOnDjango.models.DataSubjectRequest import DataSubjectRequest
from governanceOnDjango.delegates.DataSubjectRequestDelegate import DataSubjectRequestDelegate

 #======================================================================
# 
# Encapsulates data for model DataSubjectRequest
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DataSubjectRequestTest Declaration
#======================================================================
class DataSubjectRequestTest (TestCase) :
	def test_crud(self) :
		dataSubjectRequest = DataSubjectRequest()
		dataSubjectRequest.receivedDate = datetime.datetime.now()
		dataSubjectRequest.dueDate = datetime.datetime.now()
		dataSubjectRequest.requesterCountry = "default requesterCountry field value"
		dataSubjectRequest.requestType = "default requestType field value"
		dataSubjectRequest.status = "default status field value"
		
		delegate = DataSubjectRequestDelegate()
		responseObj = delegate.create(dataSubjectRequest)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


