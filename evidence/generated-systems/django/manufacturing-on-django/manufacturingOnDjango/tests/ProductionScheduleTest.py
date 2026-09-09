import datetime

from django.test import TestCase
from django.utils import timezone
from manufacturingOnDjango.models.ProductionSchedule import ProductionSchedule
from manufacturingOnDjango.delegates.ProductionScheduleDelegate import ProductionScheduleDelegate

 #======================================================================
# 
# Encapsulates data for model ProductionSchedule
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ProductionScheduleTest Declaration
#======================================================================
class ProductionScheduleTest (TestCase) :
	def test_crud(self) :
		productionSchedule = ProductionSchedule()
		productionSchedule.scheduleNumber = "default scheduleNumber field value"
		productionSchedule.horizonStart = datetime.datetime.now()
		productionSchedule.horizonEnd = datetime.datetime.now()
		productionSchedule.status = "default status field value"
		
		delegate = ProductionScheduleDelegate()
		responseObj = delegate.create(productionSchedule)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


