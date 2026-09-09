import datetime

from django.test import TestCase
from django.utils import timezone
from governanceOnDjango.models.DataBreach import DataBreach
from governanceOnDjango.delegates.DataBreachDelegate import DataBreachDelegate

 #======================================================================
# 
# Encapsulates data for model DataBreach
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DataBreachTest Declaration
#======================================================================
class DataBreachTest (TestCase) :
	def test_crud(self) :
		dataBreach = DataBreach()
		dataBreach.incidentDate = datetime.datetime.now()
		dataBreach.description = "default description field value"
		dataBreach.recordsAffected = 22
		dataBreach.notificationRequired = False
		dataBreach.severity = "default severity field value"
		dataBreach.status = "default status field value"
		
		delegate = DataBreachDelegate()
		responseObj = delegate.create(dataBreach)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


