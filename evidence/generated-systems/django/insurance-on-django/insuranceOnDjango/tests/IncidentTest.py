import datetime

from django.test import TestCase
from django.utils import timezone
from insuranceOnDjango.models.Incident import Incident
from insuranceOnDjango.delegates.IncidentDelegate import IncidentDelegate

 #======================================================================
# 
# Encapsulates data for model Incident
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class IncidentTest Declaration
#======================================================================
class IncidentTest (TestCase) :
	def test_crud(self) :
		incident = Incident()
		incident.location = "default location field value"
		incident.description = "default description field value"
		incident.incidentType = "default incidentType field value"
		
		delegate = IncidentDelegate()
		responseObj = delegate.create(incident)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


