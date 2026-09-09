from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from hrOnDjango.models.PaymentMethod import PaymentMethod
from hrOnDjango.models.Employee import Employee
from hrOnDjango.models.BankAccount import BankAccount
from hrOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model PaymentMethod
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PaymentMethodDelegate Declaration
#======================================================================
class PaymentMethodDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, paymentMethodId ):
		try:	
			paymentMethod = PaymentMethod.objects.filter(id=paymentMethodId)
			return paymentMethod.first();
		except PaymentMethod.DoesNotExist:
			raise ProcessingError("PaymentMethod with id " + str(paymentMethodId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, paymentMethod):
		for model in serializers.deserialize("json", paymentMethod):
			model.save()
			return model;

	def create(self, paymentMethod):
		paymentMethod.save()
		return paymentMethod;

	def saveFromJson(self, paymentMethod):
		for model in serializers.deserialize("json", paymentMethod):
			model.save()
			return paymentMethod;
	
	def save(self, paymentMethod):
		paymentMethod.save()
		return paymentMethod;
	
	def delete(self, paymentMethodId ):
		errMsg = "Failed to delete PaymentMethod from db using id " + str(paymentMethodId)
		
		try:
			paymentMethod = PaymentMethod.objects.get(id=paymentMethodId)
			paymentMethod.delete()
			return True
		except PaymentMethod.DoesNotExist:
			raise ProcessingError("PaymentMethod with id " + str(paymentMethodId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = PaymentMethod.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all PaymentMethod from db")
		except Exception:
			return None;
		
	def assignEmployee( self, paymentMethodId, employeeId ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.EmployeeDelegate import EmployeeDelegate

		errMsg = "Failed to assign element " + str(employeeId) + " for Employee on PaymentMethod"

		try:
			# get the PaymentMethod from db
			paymentMethod = self.get( paymentMethodId ).first()	
			
			# get the Employee from db
			employee = EmployeeDelegate().get(employeeId).first();
			
			# assign the Employee		
			paymentMethod.employee = employee
			
			#save it
			paymentMethod.save()

			# reload and return the appropriate version					
			return self.get( paymentMethodId );
		except PaymentMethod.DoesNotExist:
			raise ProcessingError(errMsg + " : PaymentMethod with id " + str(paymentMethodId) + " does not exist.")
		except Employee.DoesNotExist:
			raise ProcessingError(errMsg + " : Employee with id " + str(employeeId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignEmployee( self, paymentMethodId ):
		errMsg = "Failed to unassign element " + str(employeeId) + " for Employee on PaymentMethod"

		try:
			# get the PaymentMethod from db
			paymentMethod = self.get( paymentMethodId ).first()	
			
			# assign to None for unassignment
			paymentMethod.employee = None			

			#save it
			paymentMethod.save()

			# reload and return the appropriate version					
			return self.get( paymentMethodId );
		except PaymentMethod.DoesNotExist:
			raise ProcessingError(errMsg + " : PaymentMethod with id " + str(paymentMethodId) + " does not exist.")
		except Exception:
			return None;
		
	def assignBankAccount( self, paymentMethodId, bankAccountId ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.BankAccountDelegate import BankAccountDelegate

		errMsg = "Failed to assign element " + str(bankAccountId) + " for BankAccount on PaymentMethod"

		try:
			# get the PaymentMethod from db
			paymentMethod = self.get( paymentMethodId ).first()	
			
			# get the BankAccount from db
			bankAccount = BankAccountDelegate().get(bankAccountId).first();
			
			# assign the BankAccount		
			paymentMethod.bankAccount = bankAccount
			
			#save it
			paymentMethod.save()

			# reload and return the appropriate version					
			return self.get( paymentMethodId );
		except PaymentMethod.DoesNotExist:
			raise ProcessingError(errMsg + " : PaymentMethod with id " + str(paymentMethodId) + " does not exist.")
		except BankAccount.DoesNotExist:
			raise ProcessingError(errMsg + " : BankAccount with id " + str(bankAccountId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignBankAccount( self, paymentMethodId ):
		errMsg = "Failed to unassign element " + str(bankAccountId) + " for BankAccount on PaymentMethod"

		try:
			# get the PaymentMethod from db
			paymentMethod = self.get( paymentMethodId ).first()	
			
			# assign to None for unassignment
			paymentMethod.bankAccount = None			

			#save it
			paymentMethod.save()

			# reload and return the appropriate version					
			return self.get( paymentMethodId );
		except PaymentMethod.DoesNotExist:
			raise ProcessingError(errMsg + " : PaymentMethod with id " + str(paymentMethodId) + " does not exist.")
		except Exception:
			return None;
		
