import datetime

from django.test import TestCase
from django.utils import timezone
from analyticsOnDjango.models.AnalyticsWorkspace import AnalyticsWorkspace
from analyticsOnDjango.delegates.AnalyticsWorkspaceDelegate import AnalyticsWorkspaceDelegate

 #======================================================================
# 
# Encapsulates data for model AnalyticsWorkspace
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AnalyticsWorkspaceTest Declaration
#======================================================================
class AnalyticsWorkspaceTest (TestCase) :
	def test_crud(self) :
		analyticsWorkspace = AnalyticsWorkspace()
		analyticsWorkspace.name = "default name field value"
		analyticsWorkspace.businessDomain = "default businessDomain field value"
		analyticsWorkspace.ownerTeam = "default ownerTeam field value"
		analyticsWorkspace.governanceTier = "default governanceTier field value"
		
		delegate = AnalyticsWorkspaceDelegate()
		responseObj = delegate.create(analyticsWorkspace)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


