import datetime

from django.test import TestCase
from django.utils import timezone
from hrOnDjango.models.Goal import Goal
from hrOnDjango.delegates.GoalDelegate import GoalDelegate

 #======================================================================
# 
# Encapsulates data for model Goal
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class GoalTest Declaration
#======================================================================
class GoalTest (TestCase) :
	def test_crud(self) :
		goal = Goal()
		goal.title = "default title field value"
		goal.description = "default description field value"
		goal.targetDate = datetime.datetime.now()
		goal.weight = "default weight field value"
		goal.status = "default status field value"
		
		delegate = GoalDelegate()
		responseObj = delegate.create(goal)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


