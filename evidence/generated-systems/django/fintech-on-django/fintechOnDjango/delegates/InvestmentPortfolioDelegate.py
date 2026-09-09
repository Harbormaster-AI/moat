from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from fintechOnDjango.models.InvestmentPortfolio import InvestmentPortfolio
from fintechOnDjango.models.Customer import Customer
from fintechOnDjango.models.InvestmentAccount import InvestmentAccount
from fintechOnDjango.models.TradeOrder import TradeOrder
from fintechOnDjango.models.Position import Position
from fintechOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model InvestmentPortfolio
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class InvestmentPortfolioDelegate Declaration
#======================================================================
class InvestmentPortfolioDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, investmentPortfolioId ):
		try:	
			investmentPortfolio = InvestmentPortfolio.objects.filter(id=investmentPortfolioId)
			return investmentPortfolio.first();
		except InvestmentPortfolio.DoesNotExist:
			raise ProcessingError("InvestmentPortfolio with id " + str(investmentPortfolioId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, investmentPortfolio):
		for model in serializers.deserialize("json", investmentPortfolio):
			model.save()
			return model;

	def create(self, investmentPortfolio):
		investmentPortfolio.save()
		return investmentPortfolio;

	def saveFromJson(self, investmentPortfolio):
		for model in serializers.deserialize("json", investmentPortfolio):
			model.save()
			return investmentPortfolio;
	
	def save(self, investmentPortfolio):
		investmentPortfolio.save()
		return investmentPortfolio;
	
	def delete(self, investmentPortfolioId ):
		errMsg = "Failed to delete InvestmentPortfolio from db using id " + str(investmentPortfolioId)
		
		try:
			investmentPortfolio = InvestmentPortfolio.objects.get(id=investmentPortfolioId)
			investmentPortfolio.delete()
			return True
		except InvestmentPortfolio.DoesNotExist:
			raise ProcessingError("InvestmentPortfolio with id " + str(investmentPortfolioId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = InvestmentPortfolio.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all InvestmentPortfolio from db")
		except Exception:
			return None;
		
	def assignCustomer( self, investmentPortfolioId, customerId ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.CustomerDelegate import CustomerDelegate

		errMsg = "Failed to assign element " + str(customerId) + " for Customer on InvestmentPortfolio"

		try:
			# get the InvestmentPortfolio from db
			investmentPortfolio = self.get( investmentPortfolioId ).first()	
			
			# get the Customer from db
			customer = CustomerDelegate().get(customerId).first();
			
			# assign the Customer		
			investmentPortfolio.customer = customer
			
			#save it
			investmentPortfolio.save()

			# reload and return the appropriate version					
			return self.get( investmentPortfolioId );
		except InvestmentPortfolio.DoesNotExist:
			raise ProcessingError(errMsg + " : InvestmentPortfolio with id " + str(investmentPortfolioId) + " does not exist.")
		except Customer.DoesNotExist:
			raise ProcessingError(errMsg + " : Customer with id " + str(customerId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignCustomer( self, investmentPortfolioId ):
		errMsg = "Failed to unassign element " + str(customerId) + " for Customer on InvestmentPortfolio"

		try:
			# get the InvestmentPortfolio from db
			investmentPortfolio = self.get( investmentPortfolioId ).first()	
			
			# assign to None for unassignment
			investmentPortfolio.customer = None			

			#save it
			investmentPortfolio.save()

			# reload and return the appropriate version					
			return self.get( investmentPortfolioId );
		except InvestmentPortfolio.DoesNotExist:
			raise ProcessingError(errMsg + " : InvestmentPortfolio with id " + str(investmentPortfolioId) + " does not exist.")
		except Exception:
			return None;
		
	def addAccounts( self, investmentPortfolioId, accountsIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.InvestmentAccountDelegate import InvestmentAccountDelegate

		errMsg = "Failed to add elements " + str(accountsIds) + " for Accounts on InvestmentPortfolio"

		try:
			# get the InvestmentPortfolio
			investmentPortfolio = self.get( investmentPortfolioId ).first()
				
			# split on a comma with no spaces
			idList = accountsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the InvestmentAccount		
				investmentAccount = InvestmentAccountDelegate().get(id).first();	
				# add the InvestmentAccount
				investmentPortfolio.accounts.add(investmentAccount)
				
			# save it		
			investmentPortfolio.save()
			
			# reload and return the appropriate version
			return self.get( investmentPortfolioId );
		except InvestmentPortfolio.DoesNotExist:
			raise ProcessingError(errMsg + " : InvestmentPortfolio with id " + str(investmentPortfolioId) + " does not exist.")
		except InvestmentAccount.DoesNotExist:
			raise ProcessingError(errMsg + " : InvestmentAccount does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeAccounts( self, investmentPortfolioId, accountsIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.InvestmentAccountDelegate import InvestmentAccountDelegate

		errMsg = "Failed to remove elements " + str(accountsIds) + " for Accounts on InvestmentPortfolio"

		try:
			# get the InvestmentPortfolio
			investmentPortfolio = self.get( investmentPortfolioId ).first()
				
			# split on a comma with no spaces
			idList = accountsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the InvestmentAccount		
				investmentAccount = InvestmentAccountDelegate().get(id).first();	
				# add the InvestmentAccount
				investmentPortfolio.accounts.remove(investmentAccount)
				
			# save it		
			investmentPortfolio.save()
			
			# reload and return the appropriate version
			return self.get( investmentPortfolioId );
		except InvestmentPortfolio.DoesNotExist:
			raise ProcessingError(errMsg + " : InvestmentPortfolio with id " + str(investmentPortfolioId) + " does not exist.")
		except InvestmentAccount.DoesNotExist:
			raise ProcessingError(errMsg + " : InvestmentAccount does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addOrders( self, investmentPortfolioId, ordersIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.TradeOrderDelegate import TradeOrderDelegate

		errMsg = "Failed to add elements " + str(ordersIds) + " for Orders on InvestmentPortfolio"

		try:
			# get the InvestmentPortfolio
			investmentPortfolio = self.get( investmentPortfolioId ).first()
				
			# split on a comma with no spaces
			idList = ordersIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the TradeOrder		
				tradeOrder = TradeOrderDelegate().get(id).first();	
				# add the TradeOrder
				investmentPortfolio.orders.add(tradeOrder)
				
			# save it		
			investmentPortfolio.save()
			
			# reload and return the appropriate version
			return self.get( investmentPortfolioId );
		except InvestmentPortfolio.DoesNotExist:
			raise ProcessingError(errMsg + " : InvestmentPortfolio with id " + str(investmentPortfolioId) + " does not exist.")
		except TradeOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : TradeOrder does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeOrders( self, investmentPortfolioId, ordersIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.TradeOrderDelegate import TradeOrderDelegate

		errMsg = "Failed to remove elements " + str(ordersIds) + " for Orders on InvestmentPortfolio"

		try:
			# get the InvestmentPortfolio
			investmentPortfolio = self.get( investmentPortfolioId ).first()
				
			# split on a comma with no spaces
			idList = ordersIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the TradeOrder		
				tradeOrder = TradeOrderDelegate().get(id).first();	
				# add the TradeOrder
				investmentPortfolio.orders.remove(tradeOrder)
				
			# save it		
			investmentPortfolio.save()
			
			# reload and return the appropriate version
			return self.get( investmentPortfolioId );
		except InvestmentPortfolio.DoesNotExist:
			raise ProcessingError(errMsg + " : InvestmentPortfolio with id " + str(investmentPortfolioId) + " does not exist.")
		except TradeOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : TradeOrder does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addHoldings( self, investmentPortfolioId, holdingsIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.PositionDelegate import PositionDelegate

		errMsg = "Failed to add elements " + str(holdingsIds) + " for Holdings on InvestmentPortfolio"

		try:
			# get the InvestmentPortfolio
			investmentPortfolio = self.get( investmentPortfolioId ).first()
				
			# split on a comma with no spaces
			idList = holdingsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Position		
				position = PositionDelegate().get(id).first();	
				# add the Position
				investmentPortfolio.holdings.add(position)
				
			# save it		
			investmentPortfolio.save()
			
			# reload and return the appropriate version
			return self.get( investmentPortfolioId );
		except InvestmentPortfolio.DoesNotExist:
			raise ProcessingError(errMsg + " : InvestmentPortfolio with id " + str(investmentPortfolioId) + " does not exist.")
		except Position.DoesNotExist:
			raise ProcessingError(errMsg + " : Position does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeHoldings( self, investmentPortfolioId, holdingsIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.PositionDelegate import PositionDelegate

		errMsg = "Failed to remove elements " + str(holdingsIds) + " for Holdings on InvestmentPortfolio"

		try:
			# get the InvestmentPortfolio
			investmentPortfolio = self.get( investmentPortfolioId ).first()
				
			# split on a comma with no spaces
			idList = holdingsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Position		
				position = PositionDelegate().get(id).first();	
				# add the Position
				investmentPortfolio.holdings.remove(position)
				
			# save it		
			investmentPortfolio.save()
			
			# reload and return the appropriate version
			return self.get( investmentPortfolioId );
		except InvestmentPortfolio.DoesNotExist:
			raise ProcessingError(errMsg + " : InvestmentPortfolio with id " + str(investmentPortfolioId) + " does not exist.")
		except Position.DoesNotExist:
			raise ProcessingError(errMsg + " : Position does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
