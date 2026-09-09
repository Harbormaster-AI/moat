import datetime

from django.test import TestCase
from django.utils import timezone
from hrOnDjango.models.TrainingEnrollment import TrainingEnrollment
from hrOnDjango.delegates.TrainingEnrollmentDelegate import TrainingEnrollmentDelegate

 #======================================================================
# 
# Encapsulates data for model TrainingEnrollment
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class TrainingEnrollmentTest Declaration
#======================================================================
class TrainingEnrollmentTest (TestCase) :
	def test_crud(self) :
		trainingEnrollment = TrainingEnrollment()
		trainingEnrollment.enrollmentNumber = "default enrollmentNumber field value"
		trainingEnrollment.completionDate = datetime.datetime.now()
		trainingEnrollment.score = "default score field value"
		trainingEnrollment.status = "default status field value"
		
		delegate = TrainingEnrollmentDelegate()
		responseObj = delegate.create(trainingEnrollment)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


