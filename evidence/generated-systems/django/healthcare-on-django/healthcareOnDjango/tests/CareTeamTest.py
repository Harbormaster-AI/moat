import datetime

from django.test import TestCase
from django.utils import timezone
from healthcareOnDjango.models.CareTeam import CareTeam
from healthcareOnDjango.delegates.CareTeamDelegate import CareTeamDelegate

 #======================================================================
# 
# Encapsulates data for model CareTeam
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CareTeamTest Declaration
#======================================================================
class CareTeamTest (TestCase) :
	def test_crud(self) :
		careTeam = CareTeam()
		careTeam.name = "default name field value"
		careTeam.careSetting = "default careSetting field value"
		
		delegate = CareTeamDelegate()
		responseObj = delegate.create(careTeam)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


