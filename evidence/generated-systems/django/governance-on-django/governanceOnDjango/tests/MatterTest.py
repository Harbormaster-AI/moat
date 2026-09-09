import datetime

from django.test import TestCase
from django.utils import timezone
from governanceOnDjango.models.Matter import Matter
from governanceOnDjango.delegates.MatterDelegate import MatterDelegate

 #======================================================================
# 
# Encapsulates data for model Matter
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class MatterTest Declaration
#======================================================================
class MatterTest (TestCase) :
	def test_crud(self) :
		matter = Matter()
		matter.matterName = "default matterName field value"
		matter.leadCounsel = "default leadCounsel field value"
		matter.matterType = "default matterType field value"
		matter.status = "default status field value"
		
		delegate = MatterDelegate()
		responseObj = delegate.create(matter)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


