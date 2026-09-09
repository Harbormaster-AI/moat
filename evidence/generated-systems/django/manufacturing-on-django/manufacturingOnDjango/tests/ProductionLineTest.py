import datetime

from django.test import TestCase
from django.utils import timezone
from manufacturingOnDjango.models.ProductionLine import ProductionLine
from manufacturingOnDjango.delegates.ProductionLineDelegate import ProductionLineDelegate

 #======================================================================
# 
# Encapsulates data for model ProductionLine
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ProductionLineTest Declaration
#======================================================================
class ProductionLineTest (TestCase) :
	def test_crud(self) :
		productionLine = ProductionLine()
		productionLine.name = "default name field value"
		productionLine.lineCode = "default lineCode field value"
		productionLine.lineType = "default lineType field value"
		
		delegate = ProductionLineDelegate()
		responseObj = delegate.create(productionLine)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


