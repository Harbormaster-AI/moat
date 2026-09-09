import datetime

from django.test import TestCase
from django.utils import timezone
from governanceOnDjango.models.ComplianceProgram import ComplianceProgram
from governanceOnDjango.delegates.ComplianceProgramDelegate import ComplianceProgramDelegate

 #======================================================================
# 
# Encapsulates data for model ComplianceProgram
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ComplianceProgramTest Declaration
#======================================================================
class ComplianceProgramTest (TestCase) :
	def test_crud(self) :
		complianceProgram = ComplianceProgram()
		complianceProgram.name = "default name field value"
		complianceProgram.framework = "default framework field value"
		complianceProgram.status = "default status field value"
		
		delegate = ComplianceProgramDelegate()
		responseObj = delegate.create(complianceProgram)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


