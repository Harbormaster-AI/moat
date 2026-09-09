import datetime

from django.test import TestCase
from django.utils import timezone
from governanceOnDjango.models.AuditEngagement import AuditEngagement
from governanceOnDjango.delegates.AuditEngagementDelegate import AuditEngagementDelegate

 #======================================================================
# 
# Encapsulates data for model AuditEngagement
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AuditEngagementTest Declaration
#======================================================================
class AuditEngagementTest (TestCase) :
	def test_crud(self) :
		auditEngagement = AuditEngagement()
		auditEngagement.title = "default title field value"
		auditEngagement.startDate = datetime.datetime.now()
		auditEngagement.endDate = datetime.datetime.now()
		auditEngagement.status = "default status field value"
		
		delegate = AuditEngagementDelegate()
		responseObj = delegate.create(auditEngagement)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


