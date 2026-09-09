import datetime

from django.test import TestCase
from django.utils import timezone
from governanceOnDjango.models.AuditProgram import AuditProgram
from governanceOnDjango.delegates.AuditProgramDelegate import AuditProgramDelegate

 #======================================================================
# 
# Encapsulates data for model AuditProgram
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AuditProgramTest Declaration
#======================================================================
class AuditProgramTest (TestCase) :
	def test_crud(self) :
		auditProgram = AuditProgram()
		auditProgram.name = "default name field value"
		auditProgram.scope = "default scope field value"
		auditProgram.cycle = "default cycle field value"
		auditProgram.status = "default status field value"
		
		delegate = AuditProgramDelegate()
		responseObj = delegate.create(auditProgram)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


