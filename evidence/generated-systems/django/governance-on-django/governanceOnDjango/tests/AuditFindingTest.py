import datetime

from django.test import TestCase
from django.utils import timezone
from governanceOnDjango.models.AuditFinding import AuditFinding
from governanceOnDjango.delegates.AuditFindingDelegate import AuditFindingDelegate

 #======================================================================
# 
# Encapsulates data for model AuditFinding
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AuditFindingTest Declaration
#======================================================================
class AuditFindingTest (TestCase) :
	def test_crud(self) :
		auditFinding = AuditFinding()
		auditFinding.title = "default title field value"
		auditFinding.description = "default description field value"
		auditFinding.dueDate = datetime.datetime.now()
		auditFinding.severity = "default severity field value"
		auditFinding.status = "default status field value"
		
		delegate = AuditFindingDelegate()
		responseObj = delegate.create(auditFinding)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


