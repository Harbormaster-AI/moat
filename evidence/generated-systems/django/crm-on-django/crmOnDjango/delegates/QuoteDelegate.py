from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from crmOnDjango.models.Quote import Quote
from crmOnDjango.models.Organization import Organization
from crmOnDjango.models.Account import Account
from crmOnDjango.models.Opportunity import Opportunity
from crmOnDjango.models.User import User
from crmOnDjango.models.QuoteLineItem import QuoteLineItem
from crmOnDjango.models.PriceBook import PriceBook
from crmOnDjango.models.Order import Order
from crmOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Quote
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class QuoteDelegate Declaration
#======================================================================
class QuoteDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, quoteId ):
		try:	
			quote = Quote.objects.filter(id=quoteId)
			return quote.first();
		except Quote.DoesNotExist:
			raise ProcessingError("Quote with id " + str(quoteId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, quote):
		for model in serializers.deserialize("json", quote):
			model.save()
			return model;

	def create(self, quote):
		quote.save()
		return quote;

	def saveFromJson(self, quote):
		for model in serializers.deserialize("json", quote):
			model.save()
			return quote;
	
	def save(self, quote):
		quote.save()
		return quote;
	
	def delete(self, quoteId ):
		errMsg = "Failed to delete Quote from db using id " + str(quoteId)
		
		try:
			quote = Quote.objects.get(id=quoteId)
			quote.delete()
			return True
		except Quote.DoesNotExist:
			raise ProcessingError("Quote with id " + str(quoteId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Quote.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Quote from db")
		except Exception:
			return None;
		
	def assignOrganization( self, quoteId, organizationId ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.OrganizationDelegate import OrganizationDelegate

		errMsg = "Failed to assign element " + str(organizationId) + " for Organization on Quote"

		try:
			# get the Quote from db
			quote = self.get( quoteId ).first()	
			
			# get the Organization from db
			organization = OrganizationDelegate().get(organizationId).first();
			
			# assign the Organization		
			quote.organization = organization
			
			#save it
			quote.save()

			# reload and return the appropriate version					
			return self.get( quoteId );
		except Quote.DoesNotExist:
			raise ProcessingError(errMsg + " : Quote with id " + str(quoteId) + " does not exist.")
		except Organization.DoesNotExist:
			raise ProcessingError(errMsg + " : Organization with id " + str(organizationId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignOrganization( self, quoteId ):
		errMsg = "Failed to unassign element " + str(organizationId) + " for Organization on Quote"

		try:
			# get the Quote from db
			quote = self.get( quoteId ).first()	
			
			# assign to None for unassignment
			quote.organization = None			

			#save it
			quote.save()

			# reload and return the appropriate version					
			return self.get( quoteId );
		except Quote.DoesNotExist:
			raise ProcessingError(errMsg + " : Quote with id " + str(quoteId) + " does not exist.")
		except Exception:
			return None;
		
	def assignAccount( self, quoteId, accountId ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.AccountDelegate import AccountDelegate

		errMsg = "Failed to assign element " + str(accountId) + " for Account on Quote"

		try:
			# get the Quote from db
			quote = self.get( quoteId ).first()	
			
			# get the Account from db
			account = AccountDelegate().get(accountId).first();
			
			# assign the Account		
			quote.account = account
			
			#save it
			quote.save()

			# reload and return the appropriate version					
			return self.get( quoteId );
		except Quote.DoesNotExist:
			raise ProcessingError(errMsg + " : Quote with id " + str(quoteId) + " does not exist.")
		except Account.DoesNotExist:
			raise ProcessingError(errMsg + " : Account with id " + str(accountId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignAccount( self, quoteId ):
		errMsg = "Failed to unassign element " + str(accountId) + " for Account on Quote"

		try:
			# get the Quote from db
			quote = self.get( quoteId ).first()	
			
			# assign to None for unassignment
			quote.account = None			

			#save it
			quote.save()

			# reload and return the appropriate version					
			return self.get( quoteId );
		except Quote.DoesNotExist:
			raise ProcessingError(errMsg + " : Quote with id " + str(quoteId) + " does not exist.")
		except Exception:
			return None;
		
	def assignOpportunity( self, quoteId, opportunityId ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.OpportunityDelegate import OpportunityDelegate

		errMsg = "Failed to assign element " + str(opportunityId) + " for Opportunity on Quote"

		try:
			# get the Quote from db
			quote = self.get( quoteId ).first()	
			
			# get the Opportunity from db
			opportunity = OpportunityDelegate().get(opportunityId).first();
			
			# assign the Opportunity		
			quote.opportunity = opportunity
			
			#save it
			quote.save()

			# reload and return the appropriate version					
			return self.get( quoteId );
		except Quote.DoesNotExist:
			raise ProcessingError(errMsg + " : Quote with id " + str(quoteId) + " does not exist.")
		except Opportunity.DoesNotExist:
			raise ProcessingError(errMsg + " : Opportunity with id " + str(opportunityId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignOpportunity( self, quoteId ):
		errMsg = "Failed to unassign element " + str(opportunityId) + " for Opportunity on Quote"

		try:
			# get the Quote from db
			quote = self.get( quoteId ).first()	
			
			# assign to None for unassignment
			quote.opportunity = None			

			#save it
			quote.save()

			# reload and return the appropriate version					
			return self.get( quoteId );
		except Quote.DoesNotExist:
			raise ProcessingError(errMsg + " : Quote with id " + str(quoteId) + " does not exist.")
		except Exception:
			return None;
		
	def assignOwner( self, quoteId, ownerId ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.UserDelegate import UserDelegate

		errMsg = "Failed to assign element " + str(ownerId) + " for Owner on Quote"

		try:
			# get the Quote from db
			quote = self.get( quoteId ).first()	
			
			# get the User from db
			user = UserDelegate().get(ownerId).first();
			
			# assign the Owner		
			quote.owner = user
			
			#save it
			quote.save()

			# reload and return the appropriate version					
			return self.get( quoteId );
		except Quote.DoesNotExist:
			raise ProcessingError(errMsg + " : Quote with id " + str(quoteId) + " does not exist.")
		except User.DoesNotExist:
			raise ProcessingError(errMsg + " : User with id " + str(ownerId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignOwner( self, quoteId ):
		errMsg = "Failed to unassign element " + str(ownerId) + " for Owner on Quote"

		try:
			# get the Quote from db
			quote = self.get( quoteId ).first()	
			
			# assign to None for unassignment
			quote.user = None			

			#save it
			quote.save()

			# reload and return the appropriate version					
			return self.get( quoteId );
		except Quote.DoesNotExist:
			raise ProcessingError(errMsg + " : Quote with id " + str(quoteId) + " does not exist.")
		except Exception:
			return None;
		
	def assignPriceBook( self, quoteId, priceBookId ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.PriceBookDelegate import PriceBookDelegate

		errMsg = "Failed to assign element " + str(priceBookId) + " for PriceBook on Quote"

		try:
			# get the Quote from db
			quote = self.get( quoteId ).first()	
			
			# get the PriceBook from db
			priceBook = PriceBookDelegate().get(priceBookId).first();
			
			# assign the PriceBook		
			quote.priceBook = priceBook
			
			#save it
			quote.save()

			# reload and return the appropriate version					
			return self.get( quoteId );
		except Quote.DoesNotExist:
			raise ProcessingError(errMsg + " : Quote with id " + str(quoteId) + " does not exist.")
		except PriceBook.DoesNotExist:
			raise ProcessingError(errMsg + " : PriceBook with id " + str(priceBookId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignPriceBook( self, quoteId ):
		errMsg = "Failed to unassign element " + str(priceBookId) + " for PriceBook on Quote"

		try:
			# get the Quote from db
			quote = self.get( quoteId ).first()	
			
			# assign to None for unassignment
			quote.priceBook = None			

			#save it
			quote.save()

			# reload and return the appropriate version					
			return self.get( quoteId );
		except Quote.DoesNotExist:
			raise ProcessingError(errMsg + " : Quote with id " + str(quoteId) + " does not exist.")
		except Exception:
			return None;
		
	def assignOrder( self, quoteId, orderId ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.OrderDelegate import OrderDelegate

		errMsg = "Failed to assign element " + str(orderId) + " for Order on Quote"

		try:
			# get the Quote from db
			quote = self.get( quoteId ).first()	
			
			# get the Order from db
			order = OrderDelegate().get(orderId).first();
			
			# assign the Order		
			quote.order = order
			
			#save it
			quote.save()

			# reload and return the appropriate version					
			return self.get( quoteId );
		except Quote.DoesNotExist:
			raise ProcessingError(errMsg + " : Quote with id " + str(quoteId) + " does not exist.")
		except Order.DoesNotExist:
			raise ProcessingError(errMsg + " : Order with id " + str(orderId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignOrder( self, quoteId ):
		errMsg = "Failed to unassign element " + str(orderId) + " for Order on Quote"

		try:
			# get the Quote from db
			quote = self.get( quoteId ).first()	
			
			# assign to None for unassignment
			quote.order = None			

			#save it
			quote.save()

			# reload and return the appropriate version					
			return self.get( quoteId );
		except Quote.DoesNotExist:
			raise ProcessingError(errMsg + " : Quote with id " + str(quoteId) + " does not exist.")
		except Exception:
			return None;
		
	def addLineItems( self, quoteId, lineItemsIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.QuoteLineItemDelegate import QuoteLineItemDelegate

		errMsg = "Failed to add elements " + str(lineItemsIds) + " for LineItems on Quote"

		try:
			# get the Quote
			quote = self.get( quoteId ).first()
				
			# split on a comma with no spaces
			idList = lineItemsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the QuoteLineItem		
				quoteLineItem = QuoteLineItemDelegate().get(id).first();	
				# add the QuoteLineItem
				quote.lineItems.add(quoteLineItem)
				
			# save it		
			quote.save()
			
			# reload and return the appropriate version
			return self.get( quoteId );
		except Quote.DoesNotExist:
			raise ProcessingError(errMsg + " : Quote with id " + str(quoteId) + " does not exist.")
		except QuoteLineItem.DoesNotExist:
			raise ProcessingError(errMsg + " : QuoteLineItem does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeLineItems( self, quoteId, lineItemsIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.QuoteLineItemDelegate import QuoteLineItemDelegate

		errMsg = "Failed to remove elements " + str(lineItemsIds) + " for LineItems on Quote"

		try:
			# get the Quote
			quote = self.get( quoteId ).first()
				
			# split on a comma with no spaces
			idList = lineItemsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the QuoteLineItem		
				quoteLineItem = QuoteLineItemDelegate().get(id).first();	
				# add the QuoteLineItem
				quote.lineItems.remove(quoteLineItem)
				
			# save it		
			quote.save()
			
			# reload and return the appropriate version
			return self.get( quoteId );
		except Quote.DoesNotExist:
			raise ProcessingError(errMsg + " : Quote with id " + str(quoteId) + " does not exist.")
		except QuoteLineItem.DoesNotExist:
			raise ProcessingError(errMsg + " : QuoteLineItem does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
