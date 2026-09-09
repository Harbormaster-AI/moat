import datetime

from django.test import TestCase
from django.utils import timezone
from analyticsOnDjango.models.Dashboard import Dashboard
from analyticsOnDjango.delegates.DashboardDelegate import DashboardDelegate

 #======================================================================
# 
# Encapsulates data for model Dashboard
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DashboardTest Declaration
#======================================================================
class DashboardTest (TestCase) :
	def test_crud(self) :
		dashboard = Dashboard()
		dashboard.title = "default title field value"
		dashboard.theme = "default theme field value"
		dashboard.status = "default status field value"
		
		delegate = DashboardDelegate()
		responseObj = delegate.create(dashboard)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


