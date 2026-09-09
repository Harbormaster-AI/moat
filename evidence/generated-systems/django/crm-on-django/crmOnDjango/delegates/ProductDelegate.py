from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from crmOnDjango.models.Product import Product
from crmOnDjango.models.Organization import Organization
from crmOnDjango.models.PriceBookEntry import PriceBookEntry
from crmOnDjango.models.OpportunityLineItem import OpportunityLineItem
from crmOnDjango.models.QuoteLineItem import QuoteLineItem
from crmOnDjango.models.OrderItem import OrderItem
from crmOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Product
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ProductDelegate Declaration
#======================================================================
class ProductDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, productId ):
		try:	
			product = Product.objects.filter(id=productId)
			return product.first();
		except Product.DoesNotExist:
			raise ProcessingError("Product with id " + str(productId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, product):
		for model in serializers.deserialize("json", product):
			model.save()
			return model;

	def create(self, product):
		product.save()
		return product;

	def saveFromJson(self, product):
		for model in serializers.deserialize("json", product):
			model.save()
			return product;
	
	def save(self, product):
		product.save()
		return product;
	
	def delete(self, productId ):
		errMsg = "Failed to delete Product from db using id " + str(productId)
		
		try:
			product = Product.objects.get(id=productId)
			product.delete()
			return True
		except Product.DoesNotExist:
			raise ProcessingError("Product with id " + str(productId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Product.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Product from db")
		except Exception:
			return None;
		
	def assignOrganization( self, productId, organizationId ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.OrganizationDelegate import OrganizationDelegate

		errMsg = "Failed to assign element " + str(organizationId) + " for Organization on Product"

		try:
			# get the Product from db
			product = self.get( productId ).first()	
			
			# get the Organization from db
			organization = OrganizationDelegate().get(organizationId).first();
			
			# assign the Organization		
			product.organization = organization
			
			#save it
			product.save()

			# reload and return the appropriate version					
			return self.get( productId );
		except Product.DoesNotExist:
			raise ProcessingError(errMsg + " : Product with id " + str(productId) + " does not exist.")
		except Organization.DoesNotExist:
			raise ProcessingError(errMsg + " : Organization with id " + str(organizationId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignOrganization( self, productId ):
		errMsg = "Failed to unassign element " + str(organizationId) + " for Organization on Product"

		try:
			# get the Product from db
			product = self.get( productId ).first()	
			
			# assign to None for unassignment
			product.organization = None			

			#save it
			product.save()

			# reload and return the appropriate version					
			return self.get( productId );
		except Product.DoesNotExist:
			raise ProcessingError(errMsg + " : Product with id " + str(productId) + " does not exist.")
		except Exception:
			return None;
		
	def addPriceBookEntries( self, productId, priceBookEntriesIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.PriceBookEntryDelegate import PriceBookEntryDelegate

		errMsg = "Failed to add elements " + str(priceBookEntriesIds) + " for PriceBookEntries on Product"

		try:
			# get the Product
			product = self.get( productId ).first()
				
			# split on a comma with no spaces
			idList = priceBookEntriesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the PriceBookEntry		
				priceBookEntry = PriceBookEntryDelegate().get(id).first();	
				# add the PriceBookEntry
				product.priceBookEntries.add(priceBookEntry)
				
			# save it		
			product.save()
			
			# reload and return the appropriate version
			return self.get( productId );
		except Product.DoesNotExist:
			raise ProcessingError(errMsg + " : Product with id " + str(productId) + " does not exist.")
		except PriceBookEntry.DoesNotExist:
			raise ProcessingError(errMsg + " : PriceBookEntry does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removePriceBookEntries( self, productId, priceBookEntriesIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.PriceBookEntryDelegate import PriceBookEntryDelegate

		errMsg = "Failed to remove elements " + str(priceBookEntriesIds) + " for PriceBookEntries on Product"

		try:
			# get the Product
			product = self.get( productId ).first()
				
			# split on a comma with no spaces
			idList = priceBookEntriesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the PriceBookEntry		
				priceBookEntry = PriceBookEntryDelegate().get(id).first();	
				# add the PriceBookEntry
				product.priceBookEntries.remove(priceBookEntry)
				
			# save it		
			product.save()
			
			# reload and return the appropriate version
			return self.get( productId );
		except Product.DoesNotExist:
			raise ProcessingError(errMsg + " : Product with id " + str(productId) + " does not exist.")
		except PriceBookEntry.DoesNotExist:
			raise ProcessingError(errMsg + " : PriceBookEntry does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addOpportunityLineItems( self, productId, opportunityLineItemsIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.OpportunityLineItemDelegate import OpportunityLineItemDelegate

		errMsg = "Failed to add elements " + str(opportunityLineItemsIds) + " for OpportunityLineItems on Product"

		try:
			# get the Product
			product = self.get( productId ).first()
				
			# split on a comma with no spaces
			idList = opportunityLineItemsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the OpportunityLineItem		
				opportunityLineItem = OpportunityLineItemDelegate().get(id).first();	
				# add the OpportunityLineItem
				product.opportunityLineItems.add(opportunityLineItem)
				
			# save it		
			product.save()
			
			# reload and return the appropriate version
			return self.get( productId );
		except Product.DoesNotExist:
			raise ProcessingError(errMsg + " : Product with id " + str(productId) + " does not exist.")
		except OpportunityLineItem.DoesNotExist:
			raise ProcessingError(errMsg + " : OpportunityLineItem does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeOpportunityLineItems( self, productId, opportunityLineItemsIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.OpportunityLineItemDelegate import OpportunityLineItemDelegate

		errMsg = "Failed to remove elements " + str(opportunityLineItemsIds) + " for OpportunityLineItems on Product"

		try:
			# get the Product
			product = self.get( productId ).first()
				
			# split on a comma with no spaces
			idList = opportunityLineItemsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the OpportunityLineItem		
				opportunityLineItem = OpportunityLineItemDelegate().get(id).first();	
				# add the OpportunityLineItem
				product.opportunityLineItems.remove(opportunityLineItem)
				
			# save it		
			product.save()
			
			# reload and return the appropriate version
			return self.get( productId );
		except Product.DoesNotExist:
			raise ProcessingError(errMsg + " : Product with id " + str(productId) + " does not exist.")
		except OpportunityLineItem.DoesNotExist:
			raise ProcessingError(errMsg + " : OpportunityLineItem does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addQuoteLineItems( self, productId, quoteLineItemsIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.QuoteLineItemDelegate import QuoteLineItemDelegate

		errMsg = "Failed to add elements " + str(quoteLineItemsIds) + " for QuoteLineItems on Product"

		try:
			# get the Product
			product = self.get( productId ).first()
				
			# split on a comma with no spaces
			idList = quoteLineItemsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the QuoteLineItem		
				quoteLineItem = QuoteLineItemDelegate().get(id).first();	
				# add the QuoteLineItem
				product.quoteLineItems.add(quoteLineItem)
				
			# save it		
			product.save()
			
			# reload and return the appropriate version
			return self.get( productId );
		except Product.DoesNotExist:
			raise ProcessingError(errMsg + " : Product with id " + str(productId) + " does not exist.")
		except QuoteLineItem.DoesNotExist:
			raise ProcessingError(errMsg + " : QuoteLineItem does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeQuoteLineItems( self, productId, quoteLineItemsIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.QuoteLineItemDelegate import QuoteLineItemDelegate

		errMsg = "Failed to remove elements " + str(quoteLineItemsIds) + " for QuoteLineItems on Product"

		try:
			# get the Product
			product = self.get( productId ).first()
				
			# split on a comma with no spaces
			idList = quoteLineItemsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the QuoteLineItem		
				quoteLineItem = QuoteLineItemDelegate().get(id).first();	
				# add the QuoteLineItem
				product.quoteLineItems.remove(quoteLineItem)
				
			# save it		
			product.save()
			
			# reload and return the appropriate version
			return self.get( productId );
		except Product.DoesNotExist:
			raise ProcessingError(errMsg + " : Product with id " + str(productId) + " does not exist.")
		except QuoteLineItem.DoesNotExist:
			raise ProcessingError(errMsg + " : QuoteLineItem does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addOrderItems( self, productId, orderItemsIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.OrderItemDelegate import OrderItemDelegate

		errMsg = "Failed to add elements " + str(orderItemsIds) + " for OrderItems on Product"

		try:
			# get the Product
			product = self.get( productId ).first()
				
			# split on a comma with no spaces
			idList = orderItemsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the OrderItem		
				orderItem = OrderItemDelegate().get(id).first();	
				# add the OrderItem
				product.orderItems.add(orderItem)
				
			# save it		
			product.save()
			
			# reload and return the appropriate version
			return self.get( productId );
		except Product.DoesNotExist:
			raise ProcessingError(errMsg + " : Product with id " + str(productId) + " does not exist.")
		except OrderItem.DoesNotExist:
			raise ProcessingError(errMsg + " : OrderItem does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeOrderItems( self, productId, orderItemsIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.OrderItemDelegate import OrderItemDelegate

		errMsg = "Failed to remove elements " + str(orderItemsIds) + " for OrderItems on Product"

		try:
			# get the Product
			product = self.get( productId ).first()
				
			# split on a comma with no spaces
			idList = orderItemsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the OrderItem		
				orderItem = OrderItemDelegate().get(id).first();	
				# add the OrderItem
				product.orderItems.remove(orderItem)
				
			# save it		
			product.save()
			
			# reload and return the appropriate version
			return self.get( productId );
		except Product.DoesNotExist:
			raise ProcessingError(errMsg + " : Product with id " + str(productId) + " does not exist.")
		except OrderItem.DoesNotExist:
			raise ProcessingError(errMsg + " : OrderItem does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
