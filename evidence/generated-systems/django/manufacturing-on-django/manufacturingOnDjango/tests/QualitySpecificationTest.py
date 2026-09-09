import datetime

from django.test import TestCase
from django.utils import timezone
from manufacturingOnDjango.models.QualitySpecification import QualitySpecification
from manufacturingOnDjango.delegates.QualitySpecificationDelegate import QualitySpecificationDelegate

 #======================================================================
# 
# Encapsulates data for model QualitySpecification
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class QualitySpecificationTest Declaration
#======================================================================
class QualitySpecificationTest (TestCase) :
	def test_crud(self) :
		qualitySpecification = QualitySpecification()
		qualitySpecification.specCode = "default specCode field value"
		qualitySpecification.name = "default name field value"
		qualitySpecification.version = "default version field value"
		
		delegate = QualitySpecificationDelegate()
		responseObj = delegate.create(qualitySpecification)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


