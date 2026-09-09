import datetime

from django.test import TestCase
from django.utils import timezone
from advertisingOnDjango.models.Team import Team
from advertisingOnDjango.delegates.TeamDelegate import TeamDelegate

 #======================================================================
# 
# Encapsulates data for model Team
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class TeamTest Declaration
#======================================================================
class TeamTest (TestCase) :
	def test_crud(self) :
		team = Team()
		team.name = "default name field value"
		
		delegate = TeamDelegate()
		responseObj = delegate.create(team)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


