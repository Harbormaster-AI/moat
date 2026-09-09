from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from hrOnDjango.models.BankAccount import BankAccount
from hrOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model BankAccount
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class BankAccountDelegate Declaration
#======================================================================
class BankAccountDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, bankAccountId ):
		try:	
			bankAccount = BankAccount.objects.filter(id=bankAccountId)
			return bankAccount.first();
		except BankAccount.DoesNotExist:
			raise ProcessingError("BankAccount with id " + str(bankAccountId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, bankAccount):
		for model in serializers.deserialize("json", bankAccount):
			model.save()
			return model;

	def create(self, bankAccount):
		bankAccount.save()
		return bankAccount;

	def saveFromJson(self, bankAccount):
		for model in serializers.deserialize("json", bankAccount):
			model.save()
			return bankAccount;
	
	def save(self, bankAccount):
		bankAccount.save()
		return bankAccount;
	
	def delete(self, bankAccountId ):
		errMsg = "Failed to delete BankAccount from db using id " + str(bankAccountId)
		
		try:
			bankAccount = BankAccount.objects.get(id=bankAccountId)
			bankAccount.delete()
			return True
		except BankAccount.DoesNotExist:
			raise ProcessingError("BankAccount with id " + str(bankAccountId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = BankAccount.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all BankAccount from db")
		except Exception:
			return None;
		
