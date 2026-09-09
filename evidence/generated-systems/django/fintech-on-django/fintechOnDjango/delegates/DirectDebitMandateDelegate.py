from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from fintechOnDjango.models.DirectDebitMandate import DirectDebitMandate
from fintechOnDjango.models.Account import Account
from fintechOnDjango.models.Creditor import Creditor
from fintechOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model DirectDebitMandate
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DirectDebitMandateDelegate Declaration
#======================================================================
class DirectDebitMandateDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, directDebitMandateId ):
		try:	
			directDebitMandate = DirectDebitMandate.objects.filter(id=directDebitMandateId)
			return directDebitMandate.first();
		except DirectDebitMandate.DoesNotExist:
			raise ProcessingError("DirectDebitMandate with id " + str(directDebitMandateId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, directDebitMandate):
		for model in serializers.deserialize("json", directDebitMandate):
			model.save()
			return model;

	def create(self, directDebitMandate):
		directDebitMandate.save()
		return directDebitMandate;

	def saveFromJson(self, directDebitMandate):
		for model in serializers.deserialize("json", directDebitMandate):
			model.save()
			return directDebitMandate;
	
	def save(self, directDebitMandate):
		directDebitMandate.save()
		return directDebitMandate;
	
	def delete(self, directDebitMandateId ):
		errMsg = "Failed to delete DirectDebitMandate from db using id " + str(directDebitMandateId)
		
		try:
			directDebitMandate = DirectDebitMandate.objects.get(id=directDebitMandateId)
			directDebitMandate.delete()
			return True
		except DirectDebitMandate.DoesNotExist:
			raise ProcessingError("DirectDebitMandate with id " + str(directDebitMandateId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = DirectDebitMandate.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all DirectDebitMandate from db")
		except Exception:
			return None;
		
	def assignAccount( self, directDebitMandateId, accountId ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.AccountDelegate import AccountDelegate

		errMsg = "Failed to assign element " + str(accountId) + " for Account on DirectDebitMandate"

		try:
			# get the DirectDebitMandate from db
			directDebitMandate = self.get( directDebitMandateId ).first()	
			
			# get the Account from db
			account = AccountDelegate().get(accountId).first();
			
			# assign the Account		
			directDebitMandate.account = account
			
			#save it
			directDebitMandate.save()

			# reload and return the appropriate version					
			return self.get( directDebitMandateId );
		except DirectDebitMandate.DoesNotExist:
			raise ProcessingError(errMsg + " : DirectDebitMandate with id " + str(directDebitMandateId) + " does not exist.")
		except Account.DoesNotExist:
			raise ProcessingError(errMsg + " : Account with id " + str(accountId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignAccount( self, directDebitMandateId ):
		errMsg = "Failed to unassign element " + str(accountId) + " for Account on DirectDebitMandate"

		try:
			# get the DirectDebitMandate from db
			directDebitMandate = self.get( directDebitMandateId ).first()	
			
			# assign to None for unassignment
			directDebitMandate.account = None			

			#save it
			directDebitMandate.save()

			# reload and return the appropriate version					
			return self.get( directDebitMandateId );
		except DirectDebitMandate.DoesNotExist:
			raise ProcessingError(errMsg + " : DirectDebitMandate with id " + str(directDebitMandateId) + " does not exist.")
		except Exception:
			return None;
		
	def assignCreditor( self, directDebitMandateId, creditorId ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.CreditorDelegate import CreditorDelegate

		errMsg = "Failed to assign element " + str(creditorId) + " for Creditor on DirectDebitMandate"

		try:
			# get the DirectDebitMandate from db
			directDebitMandate = self.get( directDebitMandateId ).first()	
			
			# get the Creditor from db
			creditor = CreditorDelegate().get(creditorId).first();
			
			# assign the Creditor		
			directDebitMandate.creditor = creditor
			
			#save it
			directDebitMandate.save()

			# reload and return the appropriate version					
			return self.get( directDebitMandateId );
		except DirectDebitMandate.DoesNotExist:
			raise ProcessingError(errMsg + " : DirectDebitMandate with id " + str(directDebitMandateId) + " does not exist.")
		except Creditor.DoesNotExist:
			raise ProcessingError(errMsg + " : Creditor with id " + str(creditorId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignCreditor( self, directDebitMandateId ):
		errMsg = "Failed to unassign element " + str(creditorId) + " for Creditor on DirectDebitMandate"

		try:
			# get the DirectDebitMandate from db
			directDebitMandate = self.get( directDebitMandateId ).first()	
			
			# assign to None for unassignment
			directDebitMandate.creditor = None			

			#save it
			directDebitMandate.save()

			# reload and return the appropriate version					
			return self.get( directDebitMandateId );
		except DirectDebitMandate.DoesNotExist:
			raise ProcessingError(errMsg + " : DirectDebitMandate with id " + str(directDebitMandateId) + " does not exist.")
		except Exception:
			return None;
		
