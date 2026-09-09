from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from crmOnDjango.models.Contract import Contract
from crmOnDjango.models.Organization import Organization
from crmOnDjango.models.Account import Account
from crmOnDjango.models.User import User
from crmOnDjango.models.Order import Order
from crmOnDjango.models.Case_ import Case_
from crmOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Contract
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ContractDelegate Declaration
#======================================================================
class ContractDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, contractId ):
		try:	
			contract = Contract.objects.filter(id=contractId)
			return contract.first();
		except Contract.DoesNotExist:
			raise ProcessingError("Contract with id " + str(contractId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, contract):
		for model in serializers.deserialize("json", contract):
			model.save()
			return model;

	def create(self, contract):
		contract.save()
		return contract;

	def saveFromJson(self, contract):
		for model in serializers.deserialize("json", contract):
			model.save()
			return contract;
	
	def save(self, contract):
		contract.save()
		return contract;
	
	def delete(self, contractId ):
		errMsg = "Failed to delete Contract from db using id " + str(contractId)
		
		try:
			contract = Contract.objects.get(id=contractId)
			contract.delete()
			return True
		except Contract.DoesNotExist:
			raise ProcessingError("Contract with id " + str(contractId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Contract.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Contract from db")
		except Exception:
			return None;
		
	def assignOrganization( self, contractId, organizationId ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.OrganizationDelegate import OrganizationDelegate

		errMsg = "Failed to assign element " + str(organizationId) + " for Organization on Contract"

		try:
			# get the Contract from db
			contract = self.get( contractId ).first()	
			
			# get the Organization from db
			organization = OrganizationDelegate().get(organizationId).first();
			
			# assign the Organization		
			contract.organization = organization
			
			#save it
			contract.save()

			# reload and return the appropriate version					
			return self.get( contractId );
		except Contract.DoesNotExist:
			raise ProcessingError(errMsg + " : Contract with id " + str(contractId) + " does not exist.")
		except Organization.DoesNotExist:
			raise ProcessingError(errMsg + " : Organization with id " + str(organizationId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignOrganization( self, contractId ):
		errMsg = "Failed to unassign element " + str(organizationId) + " for Organization on Contract"

		try:
			# get the Contract from db
			contract = self.get( contractId ).first()	
			
			# assign to None for unassignment
			contract.organization = None			

			#save it
			contract.save()

			# reload and return the appropriate version					
			return self.get( contractId );
		except Contract.DoesNotExist:
			raise ProcessingError(errMsg + " : Contract with id " + str(contractId) + " does not exist.")
		except Exception:
			return None;
		
	def assignAccount( self, contractId, accountId ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.AccountDelegate import AccountDelegate

		errMsg = "Failed to assign element " + str(accountId) + " for Account on Contract"

		try:
			# get the Contract from db
			contract = self.get( contractId ).first()	
			
			# get the Account from db
			account = AccountDelegate().get(accountId).first();
			
			# assign the Account		
			contract.account = account
			
			#save it
			contract.save()

			# reload and return the appropriate version					
			return self.get( contractId );
		except Contract.DoesNotExist:
			raise ProcessingError(errMsg + " : Contract with id " + str(contractId) + " does not exist.")
		except Account.DoesNotExist:
			raise ProcessingError(errMsg + " : Account with id " + str(accountId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignAccount( self, contractId ):
		errMsg = "Failed to unassign element " + str(accountId) + " for Account on Contract"

		try:
			# get the Contract from db
			contract = self.get( contractId ).first()	
			
			# assign to None for unassignment
			contract.account = None			

			#save it
			contract.save()

			# reload and return the appropriate version					
			return self.get( contractId );
		except Contract.DoesNotExist:
			raise ProcessingError(errMsg + " : Contract with id " + str(contractId) + " does not exist.")
		except Exception:
			return None;
		
	def assignOwner( self, contractId, ownerId ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.UserDelegate import UserDelegate

		errMsg = "Failed to assign element " + str(ownerId) + " for Owner on Contract"

		try:
			# get the Contract from db
			contract = self.get( contractId ).first()	
			
			# get the User from db
			user = UserDelegate().get(ownerId).first();
			
			# assign the Owner		
			contract.owner = user
			
			#save it
			contract.save()

			# reload and return the appropriate version					
			return self.get( contractId );
		except Contract.DoesNotExist:
			raise ProcessingError(errMsg + " : Contract with id " + str(contractId) + " does not exist.")
		except User.DoesNotExist:
			raise ProcessingError(errMsg + " : User with id " + str(ownerId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignOwner( self, contractId ):
		errMsg = "Failed to unassign element " + str(ownerId) + " for Owner on Contract"

		try:
			# get the Contract from db
			contract = self.get( contractId ).first()	
			
			# assign to None for unassignment
			contract.user = None			

			#save it
			contract.save()

			# reload and return the appropriate version					
			return self.get( contractId );
		except Contract.DoesNotExist:
			raise ProcessingError(errMsg + " : Contract with id " + str(contractId) + " does not exist.")
		except Exception:
			return None;
		
	def addOrders( self, contractId, ordersIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.OrderDelegate import OrderDelegate

		errMsg = "Failed to add elements " + str(ordersIds) + " for Orders on Contract"

		try:
			# get the Contract
			contract = self.get( contractId ).first()
				
			# split on a comma with no spaces
			idList = ordersIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Order		
				order = OrderDelegate().get(id).first();	
				# add the Order
				contract.orders.add(order)
				
			# save it		
			contract.save()
			
			# reload and return the appropriate version
			return self.get( contractId );
		except Contract.DoesNotExist:
			raise ProcessingError(errMsg + " : Contract with id " + str(contractId) + " does not exist.")
		except Order.DoesNotExist:
			raise ProcessingError(errMsg + " : Order does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeOrders( self, contractId, ordersIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.OrderDelegate import OrderDelegate

		errMsg = "Failed to remove elements " + str(ordersIds) + " for Orders on Contract"

		try:
			# get the Contract
			contract = self.get( contractId ).first()
				
			# split on a comma with no spaces
			idList = ordersIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Order		
				order = OrderDelegate().get(id).first();	
				# add the Order
				contract.orders.remove(order)
				
			# save it		
			contract.save()
			
			# reload and return the appropriate version
			return self.get( contractId );
		except Contract.DoesNotExist:
			raise ProcessingError(errMsg + " : Contract with id " + str(contractId) + " does not exist.")
		except Order.DoesNotExist:
			raise ProcessingError(errMsg + " : Order does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addCases( self, contractId, casesIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.Case_Delegate import Case_Delegate

		errMsg = "Failed to add elements " + str(casesIds) + " for Cases on Contract"

		try:
			# get the Contract
			contract = self.get( contractId ).first()
				
			# split on a comma with no spaces
			idList = casesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Case_		
				case_ = Case_Delegate().get(id).first();	
				# add the Case_
				contract.cases.add(case_)
				
			# save it		
			contract.save()
			
			# reload and return the appropriate version
			return self.get( contractId );
		except Contract.DoesNotExist:
			raise ProcessingError(errMsg + " : Contract with id " + str(contractId) + " does not exist.")
		except Case_.DoesNotExist:
			raise ProcessingError(errMsg + " : Case_ does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeCases( self, contractId, casesIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.Case_Delegate import Case_Delegate

		errMsg = "Failed to remove elements " + str(casesIds) + " for Cases on Contract"

		try:
			# get the Contract
			contract = self.get( contractId ).first()
				
			# split on a comma with no spaces
			idList = casesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Case_		
				case_ = Case_Delegate().get(id).first();	
				# add the Case_
				contract.cases.remove(case_)
				
			# save it		
			contract.save()
			
			# reload and return the appropriate version
			return self.get( contractId );
		except Contract.DoesNotExist:
			raise ProcessingError(errMsg + " : Contract with id " + str(contractId) + " does not exist.")
		except Case_.DoesNotExist:
			raise ProcessingError(errMsg + " : Case_ does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
