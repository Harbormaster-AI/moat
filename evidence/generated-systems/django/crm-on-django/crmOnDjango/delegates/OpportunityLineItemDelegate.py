from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from crmOnDjango.models.OpportunityLineItem import OpportunityLineItem
from crmOnDjango.models.Opportunity import Opportunity
from crmOnDjango.models.Product import Product
from crmOnDjango.models.PriceBookEntry import PriceBookEntry
from crmOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model OpportunityLineItem
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class OpportunityLineItemDelegate Declaration
#======================================================================
class OpportunityLineItemDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, opportunityLineItemId ):
		try:	
			opportunityLineItem = OpportunityLineItem.objects.filter(id=opportunityLineItemId)
			return opportunityLineItem.first();
		except OpportunityLineItem.DoesNotExist:
			raise ProcessingError("OpportunityLineItem with id " + str(opportunityLineItemId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, opportunityLineItem):
		for model in serializers.deserialize("json", opportunityLineItem):
			model.save()
			return model;

	def create(self, opportunityLineItem):
		opportunityLineItem.save()
		return opportunityLineItem;

	def saveFromJson(self, opportunityLineItem):
		for model in serializers.deserialize("json", opportunityLineItem):
			model.save()
			return opportunityLineItem;
	
	def save(self, opportunityLineItem):
		opportunityLineItem.save()
		return opportunityLineItem;
	
	def delete(self, opportunityLineItemId ):
		errMsg = "Failed to delete OpportunityLineItem from db using id " + str(opportunityLineItemId)
		
		try:
			opportunityLineItem = OpportunityLineItem.objects.get(id=opportunityLineItemId)
			opportunityLineItem.delete()
			return True
		except OpportunityLineItem.DoesNotExist:
			raise ProcessingError("OpportunityLineItem with id " + str(opportunityLineItemId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = OpportunityLineItem.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all OpportunityLineItem from db")
		except Exception:
			return None;
		
	def assignOpportunity( self, opportunityLineItemId, opportunityId ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.OpportunityDelegate import OpportunityDelegate

		errMsg = "Failed to assign element " + str(opportunityId) + " for Opportunity on OpportunityLineItem"

		try:
			# get the OpportunityLineItem from db
			opportunityLineItem = self.get( opportunityLineItemId ).first()	
			
			# get the Opportunity from db
			opportunity = OpportunityDelegate().get(opportunityId).first();
			
			# assign the Opportunity		
			opportunityLineItem.opportunity = opportunity
			
			#save it
			opportunityLineItem.save()

			# reload and return the appropriate version					
			return self.get( opportunityLineItemId );
		except OpportunityLineItem.DoesNotExist:
			raise ProcessingError(errMsg + " : OpportunityLineItem with id " + str(opportunityLineItemId) + " does not exist.")
		except Opportunity.DoesNotExist:
			raise ProcessingError(errMsg + " : Opportunity with id " + str(opportunityId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignOpportunity( self, opportunityLineItemId ):
		errMsg = "Failed to unassign element " + str(opportunityId) + " for Opportunity on OpportunityLineItem"

		try:
			# get the OpportunityLineItem from db
			opportunityLineItem = self.get( opportunityLineItemId ).first()	
			
			# assign to None for unassignment
			opportunityLineItem.opportunity = None			

			#save it
			opportunityLineItem.save()

			# reload and return the appropriate version					
			return self.get( opportunityLineItemId );
		except OpportunityLineItem.DoesNotExist:
			raise ProcessingError(errMsg + " : OpportunityLineItem with id " + str(opportunityLineItemId) + " does not exist.")
		except Exception:
			return None;
		
	def assignProduct( self, opportunityLineItemId, productId ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.ProductDelegate import ProductDelegate

		errMsg = "Failed to assign element " + str(productId) + " for Product on OpportunityLineItem"

		try:
			# get the OpportunityLineItem from db
			opportunityLineItem = self.get( opportunityLineItemId ).first()	
			
			# get the Product from db
			product = ProductDelegate().get(productId).first();
			
			# assign the Product		
			opportunityLineItem.product = product
			
			#save it
			opportunityLineItem.save()

			# reload and return the appropriate version					
			return self.get( opportunityLineItemId );
		except OpportunityLineItem.DoesNotExist:
			raise ProcessingError(errMsg + " : OpportunityLineItem with id " + str(opportunityLineItemId) + " does not exist.")
		except Product.DoesNotExist:
			raise ProcessingError(errMsg + " : Product with id " + str(productId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignProduct( self, opportunityLineItemId ):
		errMsg = "Failed to unassign element " + str(productId) + " for Product on OpportunityLineItem"

		try:
			# get the OpportunityLineItem from db
			opportunityLineItem = self.get( opportunityLineItemId ).first()	
			
			# assign to None for unassignment
			opportunityLineItem.product = None			

			#save it
			opportunityLineItem.save()

			# reload and return the appropriate version					
			return self.get( opportunityLineItemId );
		except OpportunityLineItem.DoesNotExist:
			raise ProcessingError(errMsg + " : OpportunityLineItem with id " + str(opportunityLineItemId) + " does not exist.")
		except Exception:
			return None;
		
	def assignPriceBookEntry( self, opportunityLineItemId, priceBookEntryId ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.PriceBookEntryDelegate import PriceBookEntryDelegate

		errMsg = "Failed to assign element " + str(priceBookEntryId) + " for PriceBookEntry on OpportunityLineItem"

		try:
			# get the OpportunityLineItem from db
			opportunityLineItem = self.get( opportunityLineItemId ).first()	
			
			# get the PriceBookEntry from db
			priceBookEntry = PriceBookEntryDelegate().get(priceBookEntryId).first();
			
			# assign the PriceBookEntry		
			opportunityLineItem.priceBookEntry = priceBookEntry
			
			#save it
			opportunityLineItem.save()

			# reload and return the appropriate version					
			return self.get( opportunityLineItemId );
		except OpportunityLineItem.DoesNotExist:
			raise ProcessingError(errMsg + " : OpportunityLineItem with id " + str(opportunityLineItemId) + " does not exist.")
		except PriceBookEntry.DoesNotExist:
			raise ProcessingError(errMsg + " : PriceBookEntry with id " + str(priceBookEntryId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignPriceBookEntry( self, opportunityLineItemId ):
		errMsg = "Failed to unassign element " + str(priceBookEntryId) + " for PriceBookEntry on OpportunityLineItem"

		try:
			# get the OpportunityLineItem from db
			opportunityLineItem = self.get( opportunityLineItemId ).first()	
			
			# assign to None for unassignment
			opportunityLineItem.priceBookEntry = None			

			#save it
			opportunityLineItem.save()

			# reload and return the appropriate version					
			return self.get( opportunityLineItemId );
		except OpportunityLineItem.DoesNotExist:
			raise ProcessingError(errMsg + " : OpportunityLineItem with id " + str(opportunityLineItemId) + " does not exist.")
		except Exception:
			return None;
		
