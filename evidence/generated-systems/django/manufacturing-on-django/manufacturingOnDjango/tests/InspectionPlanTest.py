import datetime

from django.test import TestCase
from django.utils import timezone
from manufacturingOnDjango.models.InspectionPlan import InspectionPlan
from manufacturingOnDjango.delegates.InspectionPlanDelegate import InspectionPlanDelegate

 #======================================================================
# 
# Encapsulates data for model InspectionPlan
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class InspectionPlanTest Declaration
#======================================================================
class InspectionPlanTest (TestCase) :
	def test_crud(self) :
		inspectionPlan = InspectionPlan()
		inspectionPlan.planNumber = "default planNumber field value"
		inspectionPlan.revision = "default revision field value"
		inspectionPlan.samplingPlan = "default samplingPlan field value"
		inspectionPlan.status = "default status field value"
		
		delegate = InspectionPlanDelegate()
		responseObj = delegate.create(inspectionPlan)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


