import datetime

from django.test import TestCase
from django.utils import timezone
from manufacturingOnDjango.models.BOM import BOM
from manufacturingOnDjango.delegates.BOMDelegate import BOMDelegate

 #======================================================================
# 
# Encapsulates data for model BOM
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class BOMTest Declaration
#======================================================================
class BOMTest (TestCase) :
	def test_crud(self) :
		bOM = BOM()
		bOM.bomNumber = "default bomNumber field value"
		bOM.revision = "default revision field value"
		bOM.effectivityStart = datetime.datetime.now()
		bOM.effectivityEnd = datetime.datetime.now()
		bOM.status = "default status field value"
		
		delegate = BOMDelegate()
		responseObj = delegate.create(bOM)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


