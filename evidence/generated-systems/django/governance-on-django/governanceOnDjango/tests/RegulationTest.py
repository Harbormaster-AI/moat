import datetime

from django.test import TestCase
from django.utils import timezone
from governanceOnDjango.models.Regulation import Regulation
from governanceOnDjango.delegates.RegulationDelegate import RegulationDelegate

 #======================================================================
# 
# Encapsulates data for model Regulation
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class RegulationTest Declaration
#======================================================================
class RegulationTest (TestCase) :
	def test_crud(self) :
		regulation = Regulation()
		regulation.name = "default name field value"
		regulation.citation = "default citation field value"
		regulation.jurisdiction = "default jurisdiction field value"
		regulation.publicationUrl = "default publicationUrl field value"
		
		delegate = RegulationDelegate()
		responseObj = delegate.create(regulation)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


