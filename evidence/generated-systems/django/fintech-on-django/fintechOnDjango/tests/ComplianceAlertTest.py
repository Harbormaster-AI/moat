import datetime

from django.test import TestCase
from django.utils import timezone
from fintechOnDjango.models.ComplianceAlert import ComplianceAlert
from fintechOnDjango.delegates.ComplianceAlertDelegate import ComplianceAlertDelegate

 #======================================================================
# 
# Encapsulates data for model ComplianceAlert
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ComplianceAlertTest Declaration
#======================================================================
class ComplianceAlertTest (TestCase) :
	def test_crud(self) :
		complianceAlert = ComplianceAlert()
		complianceAlert.alertCode = "default alertCode field value"
		complianceAlert.raisedAt = "default raisedAt field value"
		complianceAlert.notes = "default notes field value"
		complianceAlert.severity = "default severity field value"
		complianceAlert.status = "default status field value"
		
		delegate = ComplianceAlertDelegate()
		responseObj = delegate.create(complianceAlert)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


