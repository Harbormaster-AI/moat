from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from fintechOnDjango.models.Terminal import Terminal
from fintechOnDjango.models.Merchant import Merchant
from fintechOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Terminal
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class TerminalDelegate Declaration
#======================================================================
class TerminalDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, terminalId ):
		try:	
			terminal = Terminal.objects.filter(id=terminalId)
			return terminal.first();
		except Terminal.DoesNotExist:
			raise ProcessingError("Terminal with id " + str(terminalId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, terminal):
		for model in serializers.deserialize("json", terminal):
			model.save()
			return model;

	def create(self, terminal):
		terminal.save()
		return terminal;

	def saveFromJson(self, terminal):
		for model in serializers.deserialize("json", terminal):
			model.save()
			return terminal;
	
	def save(self, terminal):
		terminal.save()
		return terminal;
	
	def delete(self, terminalId ):
		errMsg = "Failed to delete Terminal from db using id " + str(terminalId)
		
		try:
			terminal = Terminal.objects.get(id=terminalId)
			terminal.delete()
			return True
		except Terminal.DoesNotExist:
			raise ProcessingError("Terminal with id " + str(terminalId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Terminal.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Terminal from db")
		except Exception:
			return None;
		
	def assignMerchant( self, terminalId, merchantId ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.MerchantDelegate import MerchantDelegate

		errMsg = "Failed to assign element " + str(merchantId) + " for Merchant on Terminal"

		try:
			# get the Terminal from db
			terminal = self.get( terminalId ).first()	
			
			# get the Merchant from db
			merchant = MerchantDelegate().get(merchantId).first();
			
			# assign the Merchant		
			terminal.merchant = merchant
			
			#save it
			terminal.save()

			# reload and return the appropriate version					
			return self.get( terminalId );
		except Terminal.DoesNotExist:
			raise ProcessingError(errMsg + " : Terminal with id " + str(terminalId) + " does not exist.")
		except Merchant.DoesNotExist:
			raise ProcessingError(errMsg + " : Merchant with id " + str(merchantId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignMerchant( self, terminalId ):
		errMsg = "Failed to unassign element " + str(merchantId) + " for Merchant on Terminal"

		try:
			# get the Terminal from db
			terminal = self.get( terminalId ).first()	
			
			# assign to None for unassignment
			terminal.merchant = None			

			#save it
			terminal.save()

			# reload and return the appropriate version					
			return self.get( terminalId );
		except Terminal.DoesNotExist:
			raise ProcessingError(errMsg + " : Terminal with id " + str(terminalId) + " does not exist.")
		except Exception:
			return None;
		
