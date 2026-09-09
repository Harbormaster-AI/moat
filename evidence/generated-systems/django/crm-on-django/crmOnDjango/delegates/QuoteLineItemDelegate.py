from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from crmOnDjango.models.QuoteLineItem import QuoteLineItem
from crmOnDjango.models.Quote import Quote
from crmOnDjango.models.Product import Product
from crmOnDjango.models.PriceBookEntry import PriceBookEntry
from crmOnDjango.models.OpportunityLineItem import OpportunityLineItem
from crmOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model QuoteLineItem
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class QuoteLineItemDelegate Declaration
#======================================================================
class QuoteLineItemDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, quoteLineItemId ):
		try:	
			quoteLineItem = QuoteLineItem.objects.filter(id=quoteLineItemId)
			return quoteLineItem.first();
		except QuoteLineItem.DoesNotExist:
			raise ProcessingError("QuoteLineItem with id " + str(quoteLineItemId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, quoteLineItem):
		for model in serializers.deserialize("json", quoteLineItem):
			model.save()
			return model;

	def create(self, quoteLineItem):
		quoteLineItem.save()
		return quoteLineItem;

	def saveFromJson(self, quoteLineItem):
		for model in serializers.deserialize("json", quoteLineItem):
			model.save()
			return quoteLineItem;
	
	def save(self, quoteLineItem):
		quoteLineItem.save()
		return quoteLineItem;
	
	def delete(self, quoteLineItemId ):
		errMsg = "Failed to delete QuoteLineItem from db using id " + str(quoteLineItemId)
		
		try:
			quoteLineItem = QuoteLineItem.objects.get(id=quoteLineItemId)
			quoteLineItem.delete()
			return True
		except QuoteLineItem.DoesNotExist:
			raise ProcessingError("QuoteLineItem with id " + str(quoteLineItemId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = QuoteLineItem.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all QuoteLineItem from db")
		except Exception:
			return None;
		
	def assignQuote( self, quoteLineItemId, quoteId ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.QuoteDelegate import QuoteDelegate

		errMsg = "Failed to assign element " + str(quoteId) + " for Quote on QuoteLineItem"

		try:
			# get the QuoteLineItem from db
			quoteLineItem = self.get( quoteLineItemId ).first()	
			
			# get the Quote from db
			quote = QuoteDelegate().get(quoteId).first();
			
			# assign the Quote		
			quoteLineItem.quote = quote
			
			#save it
			quoteLineItem.save()

			# reload and return the appropriate version					
			return self.get( quoteLineItemId );
		except QuoteLineItem.DoesNotExist:
			raise ProcessingError(errMsg + " : QuoteLineItem with id " + str(quoteLineItemId) + " does not exist.")
		except Quote.DoesNotExist:
			raise ProcessingError(errMsg + " : Quote with id " + str(quoteId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignQuote( self, quoteLineItemId ):
		errMsg = "Failed to unassign element " + str(quoteId) + " for Quote on QuoteLineItem"

		try:
			# get the QuoteLineItem from db
			quoteLineItem = self.get( quoteLineItemId ).first()	
			
			# assign to None for unassignment
			quoteLineItem.quote = None			

			#save it
			quoteLineItem.save()

			# reload and return the appropriate version					
			return self.get( quoteLineItemId );
		except QuoteLineItem.DoesNotExist:
			raise ProcessingError(errMsg + " : QuoteLineItem with id " + str(quoteLineItemId) + " does not exist.")
		except Exception:
			return None;
		
	def assignProduct( self, quoteLineItemId, productId ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.ProductDelegate import ProductDelegate

		errMsg = "Failed to assign element " + str(productId) + " for Product on QuoteLineItem"

		try:
			# get the QuoteLineItem from db
			quoteLineItem = self.get( quoteLineItemId ).first()	
			
			# get the Product from db
			product = ProductDelegate().get(productId).first();
			
			# assign the Product		
			quoteLineItem.product = product
			
			#save it
			quoteLineItem.save()

			# reload and return the appropriate version					
			return self.get( quoteLineItemId );
		except QuoteLineItem.DoesNotExist:
			raise ProcessingError(errMsg + " : QuoteLineItem with id " + str(quoteLineItemId) + " does not exist.")
		except Product.DoesNotExist:
			raise ProcessingError(errMsg + " : Product with id " + str(productId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignProduct( self, quoteLineItemId ):
		errMsg = "Failed to unassign element " + str(productId) + " for Product on QuoteLineItem"

		try:
			# get the QuoteLineItem from db
			quoteLineItem = self.get( quoteLineItemId ).first()	
			
			# assign to None for unassignment
			quoteLineItem.product = None			

			#save it
			quoteLineItem.save()

			# reload and return the appropriate version					
			return self.get( quoteLineItemId );
		except QuoteLineItem.DoesNotExist:
			raise ProcessingError(errMsg + " : QuoteLineItem with id " + str(quoteLineItemId) + " does not exist.")
		except Exception:
			return None;
		
	def assignPriceBookEntry( self, quoteLineItemId, priceBookEntryId ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.PriceBookEntryDelegate import PriceBookEntryDelegate

		errMsg = "Failed to assign element " + str(priceBookEntryId) + " for PriceBookEntry on QuoteLineItem"

		try:
			# get the QuoteLineItem from db
			quoteLineItem = self.get( quoteLineItemId ).first()	
			
			# get the PriceBookEntry from db
			priceBookEntry = PriceBookEntryDelegate().get(priceBookEntryId).first();
			
			# assign the PriceBookEntry		
			quoteLineItem.priceBookEntry = priceBookEntry
			
			#save it
			quoteLineItem.save()

			# reload and return the appropriate version					
			return self.get( quoteLineItemId );
		except QuoteLineItem.DoesNotExist:
			raise ProcessingError(errMsg + " : QuoteLineItem with id " + str(quoteLineItemId) + " does not exist.")
		except PriceBookEntry.DoesNotExist:
			raise ProcessingError(errMsg + " : PriceBookEntry with id " + str(priceBookEntryId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignPriceBookEntry( self, quoteLineItemId ):
		errMsg = "Failed to unassign element " + str(priceBookEntryId) + " for PriceBookEntry on QuoteLineItem"

		try:
			# get the QuoteLineItem from db
			quoteLineItem = self.get( quoteLineItemId ).first()	
			
			# assign to None for unassignment
			quoteLineItem.priceBookEntry = None			

			#save it
			quoteLineItem.save()

			# reload and return the appropriate version					
			return self.get( quoteLineItemId );
		except QuoteLineItem.DoesNotExist:
			raise ProcessingError(errMsg + " : QuoteLineItem with id " + str(quoteLineItemId) + " does not exist.")
		except Exception:
			return None;
		
	def assignOpportunityLineItem( self, quoteLineItemId, opportunityLineItemId ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.OpportunityLineItemDelegate import OpportunityLineItemDelegate

		errMsg = "Failed to assign element " + str(opportunityLineItemId) + " for OpportunityLineItem on QuoteLineItem"

		try:
			# get the QuoteLineItem from db
			quoteLineItem = self.get( quoteLineItemId ).first()	
			
			# get the OpportunityLineItem from db
			opportunityLineItem = OpportunityLineItemDelegate().get(opportunityLineItemId).first();
			
			# assign the OpportunityLineItem		
			quoteLineItem.opportunityLineItem = opportunityLineItem
			
			#save it
			quoteLineItem.save()

			# reload and return the appropriate version					
			return self.get( quoteLineItemId );
		except QuoteLineItem.DoesNotExist:
			raise ProcessingError(errMsg + " : QuoteLineItem with id " + str(quoteLineItemId) + " does not exist.")
		except OpportunityLineItem.DoesNotExist:
			raise ProcessingError(errMsg + " : OpportunityLineItem with id " + str(opportunityLineItemId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignOpportunityLineItem( self, quoteLineItemId ):
		errMsg = "Failed to unassign element " + str(opportunityLineItemId) + " for OpportunityLineItem on QuoteLineItem"

		try:
			# get the QuoteLineItem from db
			quoteLineItem = self.get( quoteLineItemId ).first()	
			
			# assign to None for unassignment
			quoteLineItem.opportunityLineItem = None			

			#save it
			quoteLineItem.save()

			# reload and return the appropriate version					
			return self.get( quoteLineItemId );
		except QuoteLineItem.DoesNotExist:
			raise ProcessingError(errMsg + " : QuoteLineItem with id " + str(quoteLineItemId) + " does not exist.")
		except Exception:
			return None;
		
