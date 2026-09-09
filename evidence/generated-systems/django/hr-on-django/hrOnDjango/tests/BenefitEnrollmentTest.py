import datetime

from django.test import TestCase
from django.utils import timezone
from hrOnDjango.models.BenefitEnrollment import BenefitEnrollment
from hrOnDjango.delegates.BenefitEnrollmentDelegate import BenefitEnrollmentDelegate

 #======================================================================
# 
# Encapsulates data for model BenefitEnrollment
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class BenefitEnrollmentTest Declaration
#======================================================================
class BenefitEnrollmentTest (TestCase) :
	def test_crud(self) :
		benefitEnrollment = BenefitEnrollment()
		benefitEnrollment.enrollmentId = "default enrollmentId field value"
		benefitEnrollment.effectiveFrom = datetime.datetime.now()
		benefitEnrollment.effectiveTo = datetime.datetime.now()
		benefitEnrollment.status = "default status field value"
		benefitEnrollment.coverageLevel = "default coverageLevel field value"
		
		delegate = BenefitEnrollmentDelegate()
		responseObj = delegate.create(benefitEnrollment)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


