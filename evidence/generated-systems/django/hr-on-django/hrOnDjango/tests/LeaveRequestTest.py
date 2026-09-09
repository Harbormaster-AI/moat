import datetime

from django.test import TestCase
from django.utils import timezone
from hrOnDjango.models.LeaveRequest import LeaveRequest
from hrOnDjango.delegates.LeaveRequestDelegate import LeaveRequestDelegate

 #======================================================================
# 
# Encapsulates data for model LeaveRequest
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class LeaveRequestTest Declaration
#======================================================================
class LeaveRequestTest (TestCase) :
	def test_crud(self) :
		leaveRequest = LeaveRequest()
		leaveRequest.requestNumber = "default requestNumber field value"
		leaveRequest.startDate = datetime.datetime.now()
		leaveRequest.endDate = datetime.datetime.now()
		leaveRequest.reason = "default reason field value"
		leaveRequest.hours = "default hours field value"
		leaveRequest.status = "default status field value"
		
		delegate = LeaveRequestDelegate()
		responseObj = delegate.create(leaveRequest)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


