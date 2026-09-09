import datetime

from django.test import TestCase
from django.utils import timezone
from manufacturingOnDjango.models.MaintenancePlan import MaintenancePlan
from manufacturingOnDjango.delegates.MaintenancePlanDelegate import MaintenancePlanDelegate

 #======================================================================
# 
# Encapsulates data for model MaintenancePlan
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class MaintenancePlanTest Declaration
#======================================================================
class MaintenancePlanTest (TestCase) :
	def test_crud(self) :
		maintenancePlan = MaintenancePlan()
		maintenancePlan.planNumber = "default planNumber field value"
		maintenancePlan.interval = "default interval field value"
		maintenancePlan.lastServiceDate = datetime.datetime.now()
		maintenancePlan.strategy = "default strategy field value"
		
		delegate = MaintenancePlanDelegate()
		responseObj = delegate.create(maintenancePlan)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


