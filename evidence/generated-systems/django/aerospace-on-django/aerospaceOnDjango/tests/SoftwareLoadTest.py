import datetime

from django.test import TestCase
from django.utils import timezone
from aerospaceOnDjango.models.SoftwareLoad import SoftwareLoad
from aerospaceOnDjango.delegates.SoftwareLoadDelegate import SoftwareLoadDelegate

 #======================================================================
# 
# Encapsulates data for model SoftwareLoad
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class SoftwareLoadTest Declaration
#======================================================================
class SoftwareLoadTest (TestCase) :
	def test_crud(self) :
		softwareLoad = SoftwareLoad()
		softwareLoad.version = "default version field value"
		softwareLoad.loadType = "default loadType field value"
		
		delegate = SoftwareLoadDelegate()
		responseObj = delegate.create(softwareLoad)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


