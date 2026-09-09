import datetime

from django.test import TestCase
from django.utils import timezone
from analyticsOnDjango.models.Alert import Alert
from analyticsOnDjango.delegates.AlertDelegate import AlertDelegate

 #======================================================================
# 
# Encapsulates data for model Alert
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AlertTest Declaration
#======================================================================
class AlertTest (TestCase) :
	def test_crud(self) :
		alert = Alert()
		alert.title = "default title field value"
		alert.createdAt = datetime.datetime.now()
		alert.severity = "default severity field value"
		alert.status = "default status field value"
		
		delegate = AlertDelegate()
		responseObj = delegate.create(alert)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


