from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from fintechOnDjango.models.PaymentCard import PaymentCard
from fintechOnDjango.models.Customer import Customer
from fintechOnDjango.models.Account import Account
from fintechOnDjango.models.CardTokenization import CardTokenization
from fintechOnDjango.models.Dispute import Dispute
from fintechOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model PaymentCard
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PaymentCardDelegate Declaration
#======================================================================
class PaymentCardDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, paymentCardId ):
		try:	
			paymentCard = PaymentCard.objects.filter(id=paymentCardId)
			return paymentCard.first();
		except PaymentCard.DoesNotExist:
			raise ProcessingError("PaymentCard with id " + str(paymentCardId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, paymentCard):
		for model in serializers.deserialize("json", paymentCard):
			model.save()
			return model;

	def create(self, paymentCard):
		paymentCard.save()
		return paymentCard;

	def saveFromJson(self, paymentCard):
		for model in serializers.deserialize("json", paymentCard):
			model.save()
			return paymentCard;
	
	def save(self, paymentCard):
		paymentCard.save()
		return paymentCard;
	
	def delete(self, paymentCardId ):
		errMsg = "Failed to delete PaymentCard from db using id " + str(paymentCardId)
		
		try:
			paymentCard = PaymentCard.objects.get(id=paymentCardId)
			paymentCard.delete()
			return True
		except PaymentCard.DoesNotExist:
			raise ProcessingError("PaymentCard with id " + str(paymentCardId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = PaymentCard.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all PaymentCard from db")
		except Exception:
			return None;
		
	def assignCustomer( self, paymentCardId, customerId ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.CustomerDelegate import CustomerDelegate

		errMsg = "Failed to assign element " + str(customerId) + " for Customer on PaymentCard"

		try:
			# get the PaymentCard from db
			paymentCard = self.get( paymentCardId ).first()	
			
			# get the Customer from db
			customer = CustomerDelegate().get(customerId).first();
			
			# assign the Customer		
			paymentCard.customer = customer
			
			#save it
			paymentCard.save()

			# reload and return the appropriate version					
			return self.get( paymentCardId );
		except PaymentCard.DoesNotExist:
			raise ProcessingError(errMsg + " : PaymentCard with id " + str(paymentCardId) + " does not exist.")
		except Customer.DoesNotExist:
			raise ProcessingError(errMsg + " : Customer with id " + str(customerId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignCustomer( self, paymentCardId ):
		errMsg = "Failed to unassign element " + str(customerId) + " for Customer on PaymentCard"

		try:
			# get the PaymentCard from db
			paymentCard = self.get( paymentCardId ).first()	
			
			# assign to None for unassignment
			paymentCard.customer = None			

			#save it
			paymentCard.save()

			# reload and return the appropriate version					
			return self.get( paymentCardId );
		except PaymentCard.DoesNotExist:
			raise ProcessingError(errMsg + " : PaymentCard with id " + str(paymentCardId) + " does not exist.")
		except Exception:
			return None;
		
	def assignAccount( self, paymentCardId, accountId ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.AccountDelegate import AccountDelegate

		errMsg = "Failed to assign element " + str(accountId) + " for Account on PaymentCard"

		try:
			# get the PaymentCard from db
			paymentCard = self.get( paymentCardId ).first()	
			
			# get the Account from db
			account = AccountDelegate().get(accountId).first();
			
			# assign the Account		
			paymentCard.account = account
			
			#save it
			paymentCard.save()

			# reload and return the appropriate version					
			return self.get( paymentCardId );
		except PaymentCard.DoesNotExist:
			raise ProcessingError(errMsg + " : PaymentCard with id " + str(paymentCardId) + " does not exist.")
		except Account.DoesNotExist:
			raise ProcessingError(errMsg + " : Account with id " + str(accountId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignAccount( self, paymentCardId ):
		errMsg = "Failed to unassign element " + str(accountId) + " for Account on PaymentCard"

		try:
			# get the PaymentCard from db
			paymentCard = self.get( paymentCardId ).first()	
			
			# assign to None for unassignment
			paymentCard.account = None			

			#save it
			paymentCard.save()

			# reload and return the appropriate version					
			return self.get( paymentCardId );
		except PaymentCard.DoesNotExist:
			raise ProcessingError(errMsg + " : PaymentCard with id " + str(paymentCardId) + " does not exist.")
		except Exception:
			return None;
		
	def addTokenizations( self, paymentCardId, tokenizationsIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.CardTokenizationDelegate import CardTokenizationDelegate

		errMsg = "Failed to add elements " + str(tokenizationsIds) + " for Tokenizations on PaymentCard"

		try:
			# get the PaymentCard
			paymentCard = self.get( paymentCardId ).first()
				
			# split on a comma with no spaces
			idList = tokenizationsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the CardTokenization		
				cardTokenization = CardTokenizationDelegate().get(id).first();	
				# add the CardTokenization
				paymentCard.tokenizations.add(cardTokenization)
				
			# save it		
			paymentCard.save()
			
			# reload and return the appropriate version
			return self.get( paymentCardId );
		except PaymentCard.DoesNotExist:
			raise ProcessingError(errMsg + " : PaymentCard with id " + str(paymentCardId) + " does not exist.")
		except CardTokenization.DoesNotExist:
			raise ProcessingError(errMsg + " : CardTokenization does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeTokenizations( self, paymentCardId, tokenizationsIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.CardTokenizationDelegate import CardTokenizationDelegate

		errMsg = "Failed to remove elements " + str(tokenizationsIds) + " for Tokenizations on PaymentCard"

		try:
			# get the PaymentCard
			paymentCard = self.get( paymentCardId ).first()
				
			# split on a comma with no spaces
			idList = tokenizationsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the CardTokenization		
				cardTokenization = CardTokenizationDelegate().get(id).first();	
				# add the CardTokenization
				paymentCard.tokenizations.remove(cardTokenization)
				
			# save it		
			paymentCard.save()
			
			# reload and return the appropriate version
			return self.get( paymentCardId );
		except PaymentCard.DoesNotExist:
			raise ProcessingError(errMsg + " : PaymentCard with id " + str(paymentCardId) + " does not exist.")
		except CardTokenization.DoesNotExist:
			raise ProcessingError(errMsg + " : CardTokenization does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addDisputes( self, paymentCardId, disputesIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.DisputeDelegate import DisputeDelegate

		errMsg = "Failed to add elements " + str(disputesIds) + " for Disputes on PaymentCard"

		try:
			# get the PaymentCard
			paymentCard = self.get( paymentCardId ).first()
				
			# split on a comma with no spaces
			idList = disputesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Dispute		
				dispute = DisputeDelegate().get(id).first();	
				# add the Dispute
				paymentCard.disputes.add(dispute)
				
			# save it		
			paymentCard.save()
			
			# reload and return the appropriate version
			return self.get( paymentCardId );
		except PaymentCard.DoesNotExist:
			raise ProcessingError(errMsg + " : PaymentCard with id " + str(paymentCardId) + " does not exist.")
		except Dispute.DoesNotExist:
			raise ProcessingError(errMsg + " : Dispute does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeDisputes( self, paymentCardId, disputesIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.DisputeDelegate import DisputeDelegate

		errMsg = "Failed to remove elements " + str(disputesIds) + " for Disputes on PaymentCard"

		try:
			# get the PaymentCard
			paymentCard = self.get( paymentCardId ).first()
				
			# split on a comma with no spaces
			idList = disputesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Dispute		
				dispute = DisputeDelegate().get(id).first();	
				# add the Dispute
				paymentCard.disputes.remove(dispute)
				
			# save it		
			paymentCard.save()
			
			# reload and return the appropriate version
			return self.get( paymentCardId );
		except PaymentCard.DoesNotExist:
			raise ProcessingError(errMsg + " : PaymentCard with id " + str(paymentCardId) + " does not exist.")
		except Dispute.DoesNotExist:
			raise ProcessingError(errMsg + " : Dispute does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
