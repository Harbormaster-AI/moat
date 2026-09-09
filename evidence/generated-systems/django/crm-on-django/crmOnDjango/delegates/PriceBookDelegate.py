from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from crmOnDjango.models.PriceBook import PriceBook
from crmOnDjango.models.Organization import Organization
from crmOnDjango.models.PriceBookEntry import PriceBookEntry
from crmOnDjango.models.Quote import Quote
from crmOnDjango.models.Order import Order
from crmOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model PriceBook
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PriceBookDelegate Declaration
#======================================================================
class PriceBookDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, priceBookId ):
		try:	
			priceBook = PriceBook.objects.filter(id=priceBookId)
			return priceBook.first();
		except PriceBook.DoesNotExist:
			raise ProcessingError("PriceBook with id " + str(priceBookId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, priceBook):
		for model in serializers.deserialize("json", priceBook):
			model.save()
			return model;

	def create(self, priceBook):
		priceBook.save()
		return priceBook;

	def saveFromJson(self, priceBook):
		for model in serializers.deserialize("json", priceBook):
			model.save()
			return priceBook;
	
	def save(self, priceBook):
		priceBook.save()
		return priceBook;
	
	def delete(self, priceBookId ):
		errMsg = "Failed to delete PriceBook from db using id " + str(priceBookId)
		
		try:
			priceBook = PriceBook.objects.get(id=priceBookId)
			priceBook.delete()
			return True
		except PriceBook.DoesNotExist:
			raise ProcessingError("PriceBook with id " + str(priceBookId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = PriceBook.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all PriceBook from db")
		except Exception:
			return None;
		
	def assignOrganization( self, priceBookId, organizationId ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.OrganizationDelegate import OrganizationDelegate

		errMsg = "Failed to assign element " + str(organizationId) + " for Organization on PriceBook"

		try:
			# get the PriceBook from db
			priceBook = self.get( priceBookId ).first()	
			
			# get the Organization from db
			organization = OrganizationDelegate().get(organizationId).first();
			
			# assign the Organization		
			priceBook.organization = organization
			
			#save it
			priceBook.save()

			# reload and return the appropriate version					
			return self.get( priceBookId );
		except PriceBook.DoesNotExist:
			raise ProcessingError(errMsg + " : PriceBook with id " + str(priceBookId) + " does not exist.")
		except Organization.DoesNotExist:
			raise ProcessingError(errMsg + " : Organization with id " + str(organizationId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignOrganization( self, priceBookId ):
		errMsg = "Failed to unassign element " + str(organizationId) + " for Organization on PriceBook"

		try:
			# get the PriceBook from db
			priceBook = self.get( priceBookId ).first()	
			
			# assign to None for unassignment
			priceBook.organization = None			

			#save it
			priceBook.save()

			# reload and return the appropriate version					
			return self.get( priceBookId );
		except PriceBook.DoesNotExist:
			raise ProcessingError(errMsg + " : PriceBook with id " + str(priceBookId) + " does not exist.")
		except Exception:
			return None;
		
	def addEntries( self, priceBookId, entriesIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.PriceBookEntryDelegate import PriceBookEntryDelegate

		errMsg = "Failed to add elements " + str(entriesIds) + " for Entries on PriceBook"

		try:
			# get the PriceBook
			priceBook = self.get( priceBookId ).first()
				
			# split on a comma with no spaces
			idList = entriesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the PriceBookEntry		
				priceBookEntry = PriceBookEntryDelegate().get(id).first();	
				# add the PriceBookEntry
				priceBook.entries.add(priceBookEntry)
				
			# save it		
			priceBook.save()
			
			# reload and return the appropriate version
			return self.get( priceBookId );
		except PriceBook.DoesNotExist:
			raise ProcessingError(errMsg + " : PriceBook with id " + str(priceBookId) + " does not exist.")
		except PriceBookEntry.DoesNotExist:
			raise ProcessingError(errMsg + " : PriceBookEntry does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeEntries( self, priceBookId, entriesIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.PriceBookEntryDelegate import PriceBookEntryDelegate

		errMsg = "Failed to remove elements " + str(entriesIds) + " for Entries on PriceBook"

		try:
			# get the PriceBook
			priceBook = self.get( priceBookId ).first()
				
			# split on a comma with no spaces
			idList = entriesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the PriceBookEntry		
				priceBookEntry = PriceBookEntryDelegate().get(id).first();	
				# add the PriceBookEntry
				priceBook.entries.remove(priceBookEntry)
				
			# save it		
			priceBook.save()
			
			# reload and return the appropriate version
			return self.get( priceBookId );
		except PriceBook.DoesNotExist:
			raise ProcessingError(errMsg + " : PriceBook with id " + str(priceBookId) + " does not exist.")
		except PriceBookEntry.DoesNotExist:
			raise ProcessingError(errMsg + " : PriceBookEntry does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addQuotes( self, priceBookId, quotesIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.QuoteDelegate import QuoteDelegate

		errMsg = "Failed to add elements " + str(quotesIds) + " for Quotes on PriceBook"

		try:
			# get the PriceBook
			priceBook = self.get( priceBookId ).first()
				
			# split on a comma with no spaces
			idList = quotesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Quote		
				quote = QuoteDelegate().get(id).first();	
				# add the Quote
				priceBook.quotes.add(quote)
				
			# save it		
			priceBook.save()
			
			# reload and return the appropriate version
			return self.get( priceBookId );
		except PriceBook.DoesNotExist:
			raise ProcessingError(errMsg + " : PriceBook with id " + str(priceBookId) + " does not exist.")
		except Quote.DoesNotExist:
			raise ProcessingError(errMsg + " : Quote does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeQuotes( self, priceBookId, quotesIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.QuoteDelegate import QuoteDelegate

		errMsg = "Failed to remove elements " + str(quotesIds) + " for Quotes on PriceBook"

		try:
			# get the PriceBook
			priceBook = self.get( priceBookId ).first()
				
			# split on a comma with no spaces
			idList = quotesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Quote		
				quote = QuoteDelegate().get(id).first();	
				# add the Quote
				priceBook.quotes.remove(quote)
				
			# save it		
			priceBook.save()
			
			# reload and return the appropriate version
			return self.get( priceBookId );
		except PriceBook.DoesNotExist:
			raise ProcessingError(errMsg + " : PriceBook with id " + str(priceBookId) + " does not exist.")
		except Quote.DoesNotExist:
			raise ProcessingError(errMsg + " : Quote does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addOrders( self, priceBookId, ordersIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.OrderDelegate import OrderDelegate

		errMsg = "Failed to add elements " + str(ordersIds) + " for Orders on PriceBook"

		try:
			# get the PriceBook
			priceBook = self.get( priceBookId ).first()
				
			# split on a comma with no spaces
			idList = ordersIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Order		
				order = OrderDelegate().get(id).first();	
				# add the Order
				priceBook.orders.add(order)
				
			# save it		
			priceBook.save()
			
			# reload and return the appropriate version
			return self.get( priceBookId );
		except PriceBook.DoesNotExist:
			raise ProcessingError(errMsg + " : PriceBook with id " + str(priceBookId) + " does not exist.")
		except Order.DoesNotExist:
			raise ProcessingError(errMsg + " : Order does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeOrders( self, priceBookId, ordersIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.OrderDelegate import OrderDelegate

		errMsg = "Failed to remove elements " + str(ordersIds) + " for Orders on PriceBook"

		try:
			# get the PriceBook
			priceBook = self.get( priceBookId ).first()
				
			# split on a comma with no spaces
			idList = ordersIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Order		
				order = OrderDelegate().get(id).first();	
				# add the Order
				priceBook.orders.remove(order)
				
			# save it		
			priceBook.save()
			
			# reload and return the appropriate version
			return self.get( priceBookId );
		except PriceBook.DoesNotExist:
			raise ProcessingError(errMsg + " : PriceBook with id " + str(priceBookId) + " does not exist.")
		except Order.DoesNotExist:
			raise ProcessingError(errMsg + " : Order does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
