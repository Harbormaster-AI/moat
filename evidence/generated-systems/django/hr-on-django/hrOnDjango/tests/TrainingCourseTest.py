import datetime

from django.test import TestCase
from django.utils import timezone
from hrOnDjango.models.TrainingCourse import TrainingCourse
from hrOnDjango.delegates.TrainingCourseDelegate import TrainingCourseDelegate

 #======================================================================
# 
# Encapsulates data for model TrainingCourse
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class TrainingCourseTest Declaration
#======================================================================
class TrainingCourseTest (TestCase) :
	def test_crud(self) :
		trainingCourse = TrainingCourse()
		trainingCourse.code = "default code field value"
		trainingCourse.title = "default title field value"
		trainingCourse.durationHours = "default durationHours field value"
		trainingCourse.deliveryMethod = "default deliveryMethod field value"
		
		delegate = TrainingCourseDelegate()
		responseObj = delegate.create(trainingCourse)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


