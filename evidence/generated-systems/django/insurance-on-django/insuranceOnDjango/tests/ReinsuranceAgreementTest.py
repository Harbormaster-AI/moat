import datetime

from django.test import TestCase
from django.utils import timezone
from insuranceOnDjango.models.ReinsuranceAgreement import ReinsuranceAgreement
from insuranceOnDjango.delegates.ReinsuranceAgreementDelegate import ReinsuranceAgreementDelegate

 #======================================================================
# 
# Encapsulates data for model ReinsuranceAgreement
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ReinsuranceAgreementTest Declaration
#======================================================================
class ReinsuranceAgreementTest (TestCase) :
	def test_crud(self) :
		reinsuranceAgreement = ReinsuranceAgreement()
		reinsuranceAgreement.agreementNumber = "default agreementNumber field value"
		reinsuranceAgreement.effectivePeriod = "default effectivePeriod field value"
		reinsuranceAgreement.retention = "default retention field value"
		reinsuranceAgreement.limit = "default limit field value"
		reinsuranceAgreement.cessionPercentage = "default cessionPercentage field value"
		reinsuranceAgreement.reinsuranceType = "default reinsuranceType field value"
		reinsuranceAgreement.treatyType = "default treatyType field value"
		
		delegate = ReinsuranceAgreementDelegate()
		responseObj = delegate.create(reinsuranceAgreement)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


