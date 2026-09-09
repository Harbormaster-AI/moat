from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from fintechOnDjango.models.Trade import Trade
from fintechOnDjango.models.TradeOrder import TradeOrder
from fintechOnDjango.models.Security import Security
from fintechOnDjango.models.InvestmentAccount import InvestmentAccount
from fintechOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Trade
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class TradeDelegate Declaration
#======================================================================
class TradeDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, tradeId ):
		try:	
			trade = Trade.objects.filter(id=tradeId)
			return trade.first();
		except Trade.DoesNotExist:
			raise ProcessingError("Trade with id " + str(tradeId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, trade):
		for model in serializers.deserialize("json", trade):
			model.save()
			return model;

	def create(self, trade):
		trade.save()
		return trade;

	def saveFromJson(self, trade):
		for model in serializers.deserialize("json", trade):
			model.save()
			return trade;
	
	def save(self, trade):
		trade.save()
		return trade;
	
	def delete(self, tradeId ):
		errMsg = "Failed to delete Trade from db using id " + str(tradeId)
		
		try:
			trade = Trade.objects.get(id=tradeId)
			trade.delete()
			return True
		except Trade.DoesNotExist:
			raise ProcessingError("Trade with id " + str(tradeId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Trade.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Trade from db")
		except Exception:
			return None;
		
	def assignOrder( self, tradeId, orderId ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.TradeOrderDelegate import TradeOrderDelegate

		errMsg = "Failed to assign element " + str(orderId) + " for Order on Trade"

		try:
			# get the Trade from db
			trade = self.get( tradeId ).first()	
			
			# get the TradeOrder from db
			tradeOrder = TradeOrderDelegate().get(orderId).first();
			
			# assign the Order		
			trade.order = tradeOrder
			
			#save it
			trade.save()

			# reload and return the appropriate version					
			return self.get( tradeId );
		except Trade.DoesNotExist:
			raise ProcessingError(errMsg + " : Trade with id " + str(tradeId) + " does not exist.")
		except TradeOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : TradeOrder with id " + str(orderId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignOrder( self, tradeId ):
		errMsg = "Failed to unassign element " + str(orderId) + " for Order on Trade"

		try:
			# get the Trade from db
			trade = self.get( tradeId ).first()	
			
			# assign to None for unassignment
			trade.tradeOrder = None			

			#save it
			trade.save()

			# reload and return the appropriate version					
			return self.get( tradeId );
		except Trade.DoesNotExist:
			raise ProcessingError(errMsg + " : Trade with id " + str(tradeId) + " does not exist.")
		except Exception:
			return None;
		
	def assignSecurity( self, tradeId, securityId ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.SecurityDelegate import SecurityDelegate

		errMsg = "Failed to assign element " + str(securityId) + " for Security on Trade"

		try:
			# get the Trade from db
			trade = self.get( tradeId ).first()	
			
			# get the Security from db
			security = SecurityDelegate().get(securityId).first();
			
			# assign the Security		
			trade.security = security
			
			#save it
			trade.save()

			# reload and return the appropriate version					
			return self.get( tradeId );
		except Trade.DoesNotExist:
			raise ProcessingError(errMsg + " : Trade with id " + str(tradeId) + " does not exist.")
		except Security.DoesNotExist:
			raise ProcessingError(errMsg + " : Security with id " + str(securityId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignSecurity( self, tradeId ):
		errMsg = "Failed to unassign element " + str(securityId) + " for Security on Trade"

		try:
			# get the Trade from db
			trade = self.get( tradeId ).first()	
			
			# assign to None for unassignment
			trade.security = None			

			#save it
			trade.save()

			# reload and return the appropriate version					
			return self.get( tradeId );
		except Trade.DoesNotExist:
			raise ProcessingError(errMsg + " : Trade with id " + str(tradeId) + " does not exist.")
		except Exception:
			return None;
		
	def assignInvestmentAccount( self, tradeId, investmentAccountId ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.InvestmentAccountDelegate import InvestmentAccountDelegate

		errMsg = "Failed to assign element " + str(investmentAccountId) + " for InvestmentAccount on Trade"

		try:
			# get the Trade from db
			trade = self.get( tradeId ).first()	
			
			# get the InvestmentAccount from db
			investmentAccount = InvestmentAccountDelegate().get(investmentAccountId).first();
			
			# assign the InvestmentAccount		
			trade.investmentAccount = investmentAccount
			
			#save it
			trade.save()

			# reload and return the appropriate version					
			return self.get( tradeId );
		except Trade.DoesNotExist:
			raise ProcessingError(errMsg + " : Trade with id " + str(tradeId) + " does not exist.")
		except InvestmentAccount.DoesNotExist:
			raise ProcessingError(errMsg + " : InvestmentAccount with id " + str(investmentAccountId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignInvestmentAccount( self, tradeId ):
		errMsg = "Failed to unassign element " + str(investmentAccountId) + " for InvestmentAccount on Trade"

		try:
			# get the Trade from db
			trade = self.get( tradeId ).first()	
			
			# assign to None for unassignment
			trade.investmentAccount = None			

			#save it
			trade.save()

			# reload and return the appropriate version					
			return self.get( tradeId );
		except Trade.DoesNotExist:
			raise ProcessingError(errMsg + " : Trade with id " + str(tradeId) + " does not exist.")
		except Exception:
			return None;
		
