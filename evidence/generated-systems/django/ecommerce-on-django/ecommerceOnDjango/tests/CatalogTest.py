import datetime

from django.test import TestCase
from django.utils import timezone
from ecommerceOnDjango.models.Catalog import Catalog
from ecommerceOnDjango.delegates.CatalogDelegate import CatalogDelegate

 #======================================================================
# 
# Encapsulates data for model Catalog
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CatalogTest Declaration
#======================================================================
class CatalogTest (TestCase) :
	def test_crud(self) :
		catalog = Catalog()
		catalog.name = "default name field value"
		catalog.catalogCode = "default catalogCode field value"
		catalog.asActive = False
		
		delegate = CatalogDelegate()
		responseObj = delegate.create(catalog)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


