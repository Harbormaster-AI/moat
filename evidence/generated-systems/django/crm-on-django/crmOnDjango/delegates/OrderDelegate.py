from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from crmOnDjango.models.Order import Order
from crmOnDjango.models.Organization import Organization
from crmOnDjango.models.Account import Account
from crmOnDjango.models.Opportunity import Opportunity
from crmOnDjango.models.Quote import Quote
from crmOnDjango.models.User import User
from crmOnDjango.models.OrderItem import OrderItem
from crmOnDjango.models.Contract import Contract
from crmOnDjango.models.PriceBook import PriceBook
from crmOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Order
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class OrderDelegate Declaration
#======================================================================
class OrderDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, orderId ):
		try:	
			order = Order.objects.filter(id=orderId)
			return order.first();
		except Order.DoesNotExist:
			raise ProcessingError("Order with id " + str(orderId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, order):
		for model in serializers.deserialize("json", order):
			model.save()
			return model;

	def create(self, order):
		order.save()
		return order;

	def saveFromJson(self, order):
		for model in serializers.deserialize("json", order):
			model.save()
			return order;
	
	def save(self, order):
		order.save()
		return order;
	
	def delete(self, orderId ):
		errMsg = "Failed to delete Order from db using id " + str(orderId)
		
		try:
			order = Order.objects.get(id=orderId)
			order.delete()
			return True
		except Order.DoesNotExist:
			raise ProcessingError("Order with id " + str(orderId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Order.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Order from db")
		except Exception:
			return None;
		
	def assignOrganization( self, orderId, organizationId ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.OrganizationDelegate import OrganizationDelegate

		errMsg = "Failed to assign element " + str(organizationId) + " for Organization on Order"

		try:
			# get the Order from db
			order = self.get( orderId ).first()	
			
			# get the Organization from db
			organization = OrganizationDelegate().get(organizationId).first();
			
			# assign the Organization		
			order.organization = organization
			
			#save it
			order.save()

			# reload and return the appropriate version					
			return self.get( orderId );
		except Order.DoesNotExist:
			raise ProcessingError(errMsg + " : Order with id " + str(orderId) + " does not exist.")
		except Organization.DoesNotExist:
			raise ProcessingError(errMsg + " : Organization with id " + str(organizationId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignOrganization( self, orderId ):
		errMsg = "Failed to unassign element " + str(organizationId) + " for Organization on Order"

		try:
			# get the Order from db
			order = self.get( orderId ).first()	
			
			# assign to None for unassignment
			order.organization = None			

			#save it
			order.save()

			# reload and return the appropriate version					
			return self.get( orderId );
		except Order.DoesNotExist:
			raise ProcessingError(errMsg + " : Order with id " + str(orderId) + " does not exist.")
		except Exception:
			return None;
		
	def assignAccount( self, orderId, accountId ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.AccountDelegate import AccountDelegate

		errMsg = "Failed to assign element " + str(accountId) + " for Account on Order"

		try:
			# get the Order from db
			order = self.get( orderId ).first()	
			
			# get the Account from db
			account = AccountDelegate().get(accountId).first();
			
			# assign the Account		
			order.account = account
			
			#save it
			order.save()

			# reload and return the appropriate version					
			return self.get( orderId );
		except Order.DoesNotExist:
			raise ProcessingError(errMsg + " : Order with id " + str(orderId) + " does not exist.")
		except Account.DoesNotExist:
			raise ProcessingError(errMsg + " : Account with id " + str(accountId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignAccount( self, orderId ):
		errMsg = "Failed to unassign element " + str(accountId) + " for Account on Order"

		try:
			# get the Order from db
			order = self.get( orderId ).first()	
			
			# assign to None for unassignment
			order.account = None			

			#save it
			order.save()

			# reload and return the appropriate version					
			return self.get( orderId );
		except Order.DoesNotExist:
			raise ProcessingError(errMsg + " : Order with id " + str(orderId) + " does not exist.")
		except Exception:
			return None;
		
	def assignOpportunity( self, orderId, opportunityId ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.OpportunityDelegate import OpportunityDelegate

		errMsg = "Failed to assign element " + str(opportunityId) + " for Opportunity on Order"

		try:
			# get the Order from db
			order = self.get( orderId ).first()	
			
			# get the Opportunity from db
			opportunity = OpportunityDelegate().get(opportunityId).first();
			
			# assign the Opportunity		
			order.opportunity = opportunity
			
			#save it
			order.save()

			# reload and return the appropriate version					
			return self.get( orderId );
		except Order.DoesNotExist:
			raise ProcessingError(errMsg + " : Order with id " + str(orderId) + " does not exist.")
		except Opportunity.DoesNotExist:
			raise ProcessingError(errMsg + " : Opportunity with id " + str(opportunityId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignOpportunity( self, orderId ):
		errMsg = "Failed to unassign element " + str(opportunityId) + " for Opportunity on Order"

		try:
			# get the Order from db
			order = self.get( orderId ).first()	
			
			# assign to None for unassignment
			order.opportunity = None			

			#save it
			order.save()

			# reload and return the appropriate version					
			return self.get( orderId );
		except Order.DoesNotExist:
			raise ProcessingError(errMsg + " : Order with id " + str(orderId) + " does not exist.")
		except Exception:
			return None;
		
	def assignQuote( self, orderId, quoteId ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.QuoteDelegate import QuoteDelegate

		errMsg = "Failed to assign element " + str(quoteId) + " for Quote on Order"

		try:
			# get the Order from db
			order = self.get( orderId ).first()	
			
			# get the Quote from db
			quote = QuoteDelegate().get(quoteId).first();
			
			# assign the Quote		
			order.quote = quote
			
			#save it
			order.save()

			# reload and return the appropriate version					
			return self.get( orderId );
		except Order.DoesNotExist:
			raise ProcessingError(errMsg + " : Order with id " + str(orderId) + " does not exist.")
		except Quote.DoesNotExist:
			raise ProcessingError(errMsg + " : Quote with id " + str(quoteId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignQuote( self, orderId ):
		errMsg = "Failed to unassign element " + str(quoteId) + " for Quote on Order"

		try:
			# get the Order from db
			order = self.get( orderId ).first()	
			
			# assign to None for unassignment
			order.quote = None			

			#save it
			order.save()

			# reload and return the appropriate version					
			return self.get( orderId );
		except Order.DoesNotExist:
			raise ProcessingError(errMsg + " : Order with id " + str(orderId) + " does not exist.")
		except Exception:
			return None;
		
	def assignOwner( self, orderId, ownerId ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.UserDelegate import UserDelegate

		errMsg = "Failed to assign element " + str(ownerId) + " for Owner on Order"

		try:
			# get the Order from db
			order = self.get( orderId ).first()	
			
			# get the User from db
			user = UserDelegate().get(ownerId).first();
			
			# assign the Owner		
			order.owner = user
			
			#save it
			order.save()

			# reload and return the appropriate version					
			return self.get( orderId );
		except Order.DoesNotExist:
			raise ProcessingError(errMsg + " : Order with id " + str(orderId) + " does not exist.")
		except User.DoesNotExist:
			raise ProcessingError(errMsg + " : User with id " + str(ownerId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignOwner( self, orderId ):
		errMsg = "Failed to unassign element " + str(ownerId) + " for Owner on Order"

		try:
			# get the Order from db
			order = self.get( orderId ).first()	
			
			# assign to None for unassignment
			order.user = None			

			#save it
			order.save()

			# reload and return the appropriate version					
			return self.get( orderId );
		except Order.DoesNotExist:
			raise ProcessingError(errMsg + " : Order with id " + str(orderId) + " does not exist.")
		except Exception:
			return None;
		
	def assignContract( self, orderId, contractId ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.ContractDelegate import ContractDelegate

		errMsg = "Failed to assign element " + str(contractId) + " for Contract on Order"

		try:
			# get the Order from db
			order = self.get( orderId ).first()	
			
			# get the Contract from db
			contract = ContractDelegate().get(contractId).first();
			
			# assign the Contract		
			order.contract = contract
			
			#save it
			order.save()

			# reload and return the appropriate version					
			return self.get( orderId );
		except Order.DoesNotExist:
			raise ProcessingError(errMsg + " : Order with id " + str(orderId) + " does not exist.")
		except Contract.DoesNotExist:
			raise ProcessingError(errMsg + " : Contract with id " + str(contractId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignContract( self, orderId ):
		errMsg = "Failed to unassign element " + str(contractId) + " for Contract on Order"

		try:
			# get the Order from db
			order = self.get( orderId ).first()	
			
			# assign to None for unassignment
			order.contract = None			

			#save it
			order.save()

			# reload and return the appropriate version					
			return self.get( orderId );
		except Order.DoesNotExist:
			raise ProcessingError(errMsg + " : Order with id " + str(orderId) + " does not exist.")
		except Exception:
			return None;
		
	def assignPriceBook( self, orderId, priceBookId ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.PriceBookDelegate import PriceBookDelegate

		errMsg = "Failed to assign element " + str(priceBookId) + " for PriceBook on Order"

		try:
			# get the Order from db
			order = self.get( orderId ).first()	
			
			# get the PriceBook from db
			priceBook = PriceBookDelegate().get(priceBookId).first();
			
			# assign the PriceBook		
			order.priceBook = priceBook
			
			#save it
			order.save()

			# reload and return the appropriate version					
			return self.get( orderId );
		except Order.DoesNotExist:
			raise ProcessingError(errMsg + " : Order with id " + str(orderId) + " does not exist.")
		except PriceBook.DoesNotExist:
			raise ProcessingError(errMsg + " : PriceBook with id " + str(priceBookId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignPriceBook( self, orderId ):
		errMsg = "Failed to unassign element " + str(priceBookId) + " for PriceBook on Order"

		try:
			# get the Order from db
			order = self.get( orderId ).first()	
			
			# assign to None for unassignment
			order.priceBook = None			

			#save it
			order.save()

			# reload and return the appropriate version					
			return self.get( orderId );
		except Order.DoesNotExist:
			raise ProcessingError(errMsg + " : Order with id " + str(orderId) + " does not exist.")
		except Exception:
			return None;
		
	def addItems( self, orderId, itemsIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.OrderItemDelegate import OrderItemDelegate

		errMsg = "Failed to add elements " + str(itemsIds) + " for Items on Order"

		try:
			# get the Order
			order = self.get( orderId ).first()
				
			# split on a comma with no spaces
			idList = itemsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the OrderItem		
				orderItem = OrderItemDelegate().get(id).first();	
				# add the OrderItem
				order.items.add(orderItem)
				
			# save it		
			order.save()
			
			# reload and return the appropriate version
			return self.get( orderId );
		except Order.DoesNotExist:
			raise ProcessingError(errMsg + " : Order with id " + str(orderId) + " does not exist.")
		except OrderItem.DoesNotExist:
			raise ProcessingError(errMsg + " : OrderItem does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeItems( self, orderId, itemsIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.OrderItemDelegate import OrderItemDelegate

		errMsg = "Failed to remove elements " + str(itemsIds) + " for Items on Order"

		try:
			# get the Order
			order = self.get( orderId ).first()
				
			# split on a comma with no spaces
			idList = itemsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the OrderItem		
				orderItem = OrderItemDelegate().get(id).first();	
				# add the OrderItem
				order.items.remove(orderItem)
				
			# save it		
			order.save()
			
			# reload and return the appropriate version
			return self.get( orderId );
		except Order.DoesNotExist:
			raise ProcessingError(errMsg + " : Order with id " + str(orderId) + " does not exist.")
		except OrderItem.DoesNotExist:
			raise ProcessingError(errMsg + " : OrderItem does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
