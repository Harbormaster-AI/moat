from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from crmOnDjango.models.PriceBookEntry import PriceBookEntry
from crmOnDjango.models.PriceBook import PriceBook
from crmOnDjango.models.Product import Product
from crmOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model PriceBookEntry
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PriceBookEntryDelegate Declaration
#======================================================================
class PriceBookEntryDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, priceBookEntryId ):
		try:	
			priceBookEntry = PriceBookEntry.objects.filter(id=priceBookEntryId)
			return priceBookEntry.first();
		except PriceBookEntry.DoesNotExist:
			raise ProcessingError("PriceBookEntry with id " + str(priceBookEntryId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, priceBookEntry):
		for model in serializers.deserialize("json", priceBookEntry):
			model.save()
			return model;

	def create(self, priceBookEntry):
		priceBookEntry.save()
		return priceBookEntry;

	def saveFromJson(self, priceBookEntry):
		for model in serializers.deserialize("json", priceBookEntry):
			model.save()
			return priceBookEntry;
	
	def save(self, priceBookEntry):
		priceBookEntry.save()
		return priceBookEntry;
	
	def delete(self, priceBookEntryId ):
		errMsg = "Failed to delete PriceBookEntry from db using id " + str(priceBookEntryId)
		
		try:
			priceBookEntry = PriceBookEntry.objects.get(id=priceBookEntryId)
			priceBookEntry.delete()
			return True
		except PriceBookEntry.DoesNotExist:
			raise ProcessingError("PriceBookEntry with id " + str(priceBookEntryId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = PriceBookEntry.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all PriceBookEntry from db")
		except Exception:
			return None;
		
	def assignPriceBook( self, priceBookEntryId, priceBookId ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.PriceBookDelegate import PriceBookDelegate

		errMsg = "Failed to assign element " + str(priceBookId) + " for PriceBook on PriceBookEntry"

		try:
			# get the PriceBookEntry from db
			priceBookEntry = self.get( priceBookEntryId ).first()	
			
			# get the PriceBook from db
			priceBook = PriceBookDelegate().get(priceBookId).first();
			
			# assign the PriceBook		
			priceBookEntry.priceBook = priceBook
			
			#save it
			priceBookEntry.save()

			# reload and return the appropriate version					
			return self.get( priceBookEntryId );
		except PriceBookEntry.DoesNotExist:
			raise ProcessingError(errMsg + " : PriceBookEntry with id " + str(priceBookEntryId) + " does not exist.")
		except PriceBook.DoesNotExist:
			raise ProcessingError(errMsg + " : PriceBook with id " + str(priceBookId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignPriceBook( self, priceBookEntryId ):
		errMsg = "Failed to unassign element " + str(priceBookId) + " for PriceBook on PriceBookEntry"

		try:
			# get the PriceBookEntry from db
			priceBookEntry = self.get( priceBookEntryId ).first()	
			
			# assign to None for unassignment
			priceBookEntry.priceBook = None			

			#save it
			priceBookEntry.save()

			# reload and return the appropriate version					
			return self.get( priceBookEntryId );
		except PriceBookEntry.DoesNotExist:
			raise ProcessingError(errMsg + " : PriceBookEntry with id " + str(priceBookEntryId) + " does not exist.")
		except Exception:
			return None;
		
	def assignProduct( self, priceBookEntryId, productId ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.ProductDelegate import ProductDelegate

		errMsg = "Failed to assign element " + str(productId) + " for Product on PriceBookEntry"

		try:
			# get the PriceBookEntry from db
			priceBookEntry = self.get( priceBookEntryId ).first()	
			
			# get the Product from db
			product = ProductDelegate().get(productId).first();
			
			# assign the Product		
			priceBookEntry.product = product
			
			#save it
			priceBookEntry.save()

			# reload and return the appropriate version					
			return self.get( priceBookEntryId );
		except PriceBookEntry.DoesNotExist:
			raise ProcessingError(errMsg + " : PriceBookEntry with id " + str(priceBookEntryId) + " does not exist.")
		except Product.DoesNotExist:
			raise ProcessingError(errMsg + " : Product with id " + str(productId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignProduct( self, priceBookEntryId ):
		errMsg = "Failed to unassign element " + str(productId) + " for Product on PriceBookEntry"

		try:
			# get the PriceBookEntry from db
			priceBookEntry = self.get( priceBookEntryId ).first()	
			
			# assign to None for unassignment
			priceBookEntry.product = None			

			#save it
			priceBookEntry.save()

			# reload and return the appropriate version					
			return self.get( priceBookEntryId );
		except PriceBookEntry.DoesNotExist:
			raise ProcessingError(errMsg + " : PriceBookEntry with id " + str(priceBookEntryId) + " does not exist.")
		except Exception:
			return None;
		
