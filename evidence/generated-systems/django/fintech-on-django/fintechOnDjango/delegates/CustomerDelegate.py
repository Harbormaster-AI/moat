from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from fintechOnDjango.models.Customer import Customer
from fintechOnDjango.models.FinancialInstitution import FinancialInstitution
from fintechOnDjango.models.Account import Account
from fintechOnDjango.models.Wallet import Wallet
from fintechOnDjango.models.PaymentCard import PaymentCard
from fintechOnDjango.models.KYCProfile import KYCProfile
from fintechOnDjango.models.Consent import Consent
from fintechOnDjango.models.Agreement import Agreement
from fintechOnDjango.models.LoanApplication import LoanApplication
from fintechOnDjango.models.Loan import Loan
from fintechOnDjango.models.InvestmentPortfolio import InvestmentPortfolio
from fintechOnDjango.models.Dispute import Dispute
from fintechOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Customer
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CustomerDelegate Declaration
#======================================================================
class CustomerDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, customerId ):
		try:	
			customer = Customer.objects.filter(id=customerId)
			return customer.first();
		except Customer.DoesNotExist:
			raise ProcessingError("Customer with id " + str(customerId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, customer):
		for model in serializers.deserialize("json", customer):
			model.save()
			return model;

	def create(self, customer):
		customer.save()
		return customer;

	def saveFromJson(self, customer):
		for model in serializers.deserialize("json", customer):
			model.save()
			return customer;
	
	def save(self, customer):
		customer.save()
		return customer;
	
	def delete(self, customerId ):
		errMsg = "Failed to delete Customer from db using id " + str(customerId)
		
		try:
			customer = Customer.objects.get(id=customerId)
			customer.delete()
			return True
		except Customer.DoesNotExist:
			raise ProcessingError("Customer with id " + str(customerId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Customer.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Customer from db")
		except Exception:
			return None;
		
	def assignInstitution( self, customerId, institutionId ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.FinancialInstitutionDelegate import FinancialInstitutionDelegate

		errMsg = "Failed to assign element " + str(institutionId) + " for Institution on Customer"

		try:
			# get the Customer from db
			customer = self.get( customerId ).first()	
			
			# get the FinancialInstitution from db
			financialInstitution = FinancialInstitutionDelegate().get(institutionId).first();
			
			# assign the Institution		
			customer.institution = financialInstitution
			
			#save it
			customer.save()

			# reload and return the appropriate version					
			return self.get( customerId );
		except Customer.DoesNotExist:
			raise ProcessingError(errMsg + " : Customer with id " + str(customerId) + " does not exist.")
		except FinancialInstitution.DoesNotExist:
			raise ProcessingError(errMsg + " : FinancialInstitution with id " + str(institutionId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignInstitution( self, customerId ):
		errMsg = "Failed to unassign element " + str(institutionId) + " for Institution on Customer"

		try:
			# get the Customer from db
			customer = self.get( customerId ).first()	
			
			# assign to None for unassignment
			customer.financialInstitution = None			

			#save it
			customer.save()

			# reload and return the appropriate version					
			return self.get( customerId );
		except Customer.DoesNotExist:
			raise ProcessingError(errMsg + " : Customer with id " + str(customerId) + " does not exist.")
		except Exception:
			return None;
		
	def addAccounts( self, customerId, accountsIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.AccountDelegate import AccountDelegate

		errMsg = "Failed to add elements " + str(accountsIds) + " for Accounts on Customer"

		try:
			# get the Customer
			customer = self.get( customerId ).first()
				
			# split on a comma with no spaces
			idList = accountsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Account		
				account = AccountDelegate().get(id).first();	
				# add the Account
				customer.accounts.add(account)
				
			# save it		
			customer.save()
			
			# reload and return the appropriate version
			return self.get( customerId );
		except Customer.DoesNotExist:
			raise ProcessingError(errMsg + " : Customer with id " + str(customerId) + " does not exist.")
		except Account.DoesNotExist:
			raise ProcessingError(errMsg + " : Account does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeAccounts( self, customerId, accountsIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.AccountDelegate import AccountDelegate

		errMsg = "Failed to remove elements " + str(accountsIds) + " for Accounts on Customer"

		try:
			# get the Customer
			customer = self.get( customerId ).first()
				
			# split on a comma with no spaces
			idList = accountsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Account		
				account = AccountDelegate().get(id).first();	
				# add the Account
				customer.accounts.remove(account)
				
			# save it		
			customer.save()
			
			# reload and return the appropriate version
			return self.get( customerId );
		except Customer.DoesNotExist:
			raise ProcessingError(errMsg + " : Customer with id " + str(customerId) + " does not exist.")
		except Account.DoesNotExist:
			raise ProcessingError(errMsg + " : Account does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addWallets( self, customerId, walletsIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.WalletDelegate import WalletDelegate

		errMsg = "Failed to add elements " + str(walletsIds) + " for Wallets on Customer"

		try:
			# get the Customer
			customer = self.get( customerId ).first()
				
			# split on a comma with no spaces
			idList = walletsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Wallet		
				wallet = WalletDelegate().get(id).first();	
				# add the Wallet
				customer.wallets.add(wallet)
				
			# save it		
			customer.save()
			
			# reload and return the appropriate version
			return self.get( customerId );
		except Customer.DoesNotExist:
			raise ProcessingError(errMsg + " : Customer with id " + str(customerId) + " does not exist.")
		except Wallet.DoesNotExist:
			raise ProcessingError(errMsg + " : Wallet does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeWallets( self, customerId, walletsIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.WalletDelegate import WalletDelegate

		errMsg = "Failed to remove elements " + str(walletsIds) + " for Wallets on Customer"

		try:
			# get the Customer
			customer = self.get( customerId ).first()
				
			# split on a comma with no spaces
			idList = walletsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Wallet		
				wallet = WalletDelegate().get(id).first();	
				# add the Wallet
				customer.wallets.remove(wallet)
				
			# save it		
			customer.save()
			
			# reload and return the appropriate version
			return self.get( customerId );
		except Customer.DoesNotExist:
			raise ProcessingError(errMsg + " : Customer with id " + str(customerId) + " does not exist.")
		except Wallet.DoesNotExist:
			raise ProcessingError(errMsg + " : Wallet does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addCards( self, customerId, cardsIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.PaymentCardDelegate import PaymentCardDelegate

		errMsg = "Failed to add elements " + str(cardsIds) + " for Cards on Customer"

		try:
			# get the Customer
			customer = self.get( customerId ).first()
				
			# split on a comma with no spaces
			idList = cardsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the PaymentCard		
				paymentCard = PaymentCardDelegate().get(id).first();	
				# add the PaymentCard
				customer.cards.add(paymentCard)
				
			# save it		
			customer.save()
			
			# reload and return the appropriate version
			return self.get( customerId );
		except Customer.DoesNotExist:
			raise ProcessingError(errMsg + " : Customer with id " + str(customerId) + " does not exist.")
		except PaymentCard.DoesNotExist:
			raise ProcessingError(errMsg + " : PaymentCard does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeCards( self, customerId, cardsIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.PaymentCardDelegate import PaymentCardDelegate

		errMsg = "Failed to remove elements " + str(cardsIds) + " for Cards on Customer"

		try:
			# get the Customer
			customer = self.get( customerId ).first()
				
			# split on a comma with no spaces
			idList = cardsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the PaymentCard		
				paymentCard = PaymentCardDelegate().get(id).first();	
				# add the PaymentCard
				customer.cards.remove(paymentCard)
				
			# save it		
			customer.save()
			
			# reload and return the appropriate version
			return self.get( customerId );
		except Customer.DoesNotExist:
			raise ProcessingError(errMsg + " : Customer with id " + str(customerId) + " does not exist.")
		except PaymentCard.DoesNotExist:
			raise ProcessingError(errMsg + " : PaymentCard does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addKycProfiles( self, customerId, kycProfilesIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.KYCProfileDelegate import KYCProfileDelegate

		errMsg = "Failed to add elements " + str(kycProfilesIds) + " for KycProfiles on Customer"

		try:
			# get the Customer
			customer = self.get( customerId ).first()
				
			# split on a comma with no spaces
			idList = kycProfilesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the KYCProfile		
				kYCProfile = KYCProfileDelegate().get(id).first();	
				# add the KYCProfile
				customer.kycProfiles.add(kYCProfile)
				
			# save it		
			customer.save()
			
			# reload and return the appropriate version
			return self.get( customerId );
		except Customer.DoesNotExist:
			raise ProcessingError(errMsg + " : Customer with id " + str(customerId) + " does not exist.")
		except KYCProfile.DoesNotExist:
			raise ProcessingError(errMsg + " : KYCProfile does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeKycProfiles( self, customerId, kycProfilesIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.KYCProfileDelegate import KYCProfileDelegate

		errMsg = "Failed to remove elements " + str(kycProfilesIds) + " for KycProfiles on Customer"

		try:
			# get the Customer
			customer = self.get( customerId ).first()
				
			# split on a comma with no spaces
			idList = kycProfilesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the KYCProfile		
				kYCProfile = KYCProfileDelegate().get(id).first();	
				# add the KYCProfile
				customer.kycProfiles.remove(kYCProfile)
				
			# save it		
			customer.save()
			
			# reload and return the appropriate version
			return self.get( customerId );
		except Customer.DoesNotExist:
			raise ProcessingError(errMsg + " : Customer with id " + str(customerId) + " does not exist.")
		except KYCProfile.DoesNotExist:
			raise ProcessingError(errMsg + " : KYCProfile does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addConsents( self, customerId, consentsIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.ConsentDelegate import ConsentDelegate

		errMsg = "Failed to add elements " + str(consentsIds) + " for Consents on Customer"

		try:
			# get the Customer
			customer = self.get( customerId ).first()
				
			# split on a comma with no spaces
			idList = consentsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Consent		
				consent = ConsentDelegate().get(id).first();	
				# add the Consent
				customer.consents.add(consent)
				
			# save it		
			customer.save()
			
			# reload and return the appropriate version
			return self.get( customerId );
		except Customer.DoesNotExist:
			raise ProcessingError(errMsg + " : Customer with id " + str(customerId) + " does not exist.")
		except Consent.DoesNotExist:
			raise ProcessingError(errMsg + " : Consent does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeConsents( self, customerId, consentsIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.ConsentDelegate import ConsentDelegate

		errMsg = "Failed to remove elements " + str(consentsIds) + " for Consents on Customer"

		try:
			# get the Customer
			customer = self.get( customerId ).first()
				
			# split on a comma with no spaces
			idList = consentsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Consent		
				consent = ConsentDelegate().get(id).first();	
				# add the Consent
				customer.consents.remove(consent)
				
			# save it		
			customer.save()
			
			# reload and return the appropriate version
			return self.get( customerId );
		except Customer.DoesNotExist:
			raise ProcessingError(errMsg + " : Customer with id " + str(customerId) + " does not exist.")
		except Consent.DoesNotExist:
			raise ProcessingError(errMsg + " : Consent does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addAgreements( self, customerId, agreementsIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.AgreementDelegate import AgreementDelegate

		errMsg = "Failed to add elements " + str(agreementsIds) + " for Agreements on Customer"

		try:
			# get the Customer
			customer = self.get( customerId ).first()
				
			# split on a comma with no spaces
			idList = agreementsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Agreement		
				agreement = AgreementDelegate().get(id).first();	
				# add the Agreement
				customer.agreements.add(agreement)
				
			# save it		
			customer.save()
			
			# reload and return the appropriate version
			return self.get( customerId );
		except Customer.DoesNotExist:
			raise ProcessingError(errMsg + " : Customer with id " + str(customerId) + " does not exist.")
		except Agreement.DoesNotExist:
			raise ProcessingError(errMsg + " : Agreement does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeAgreements( self, customerId, agreementsIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.AgreementDelegate import AgreementDelegate

		errMsg = "Failed to remove elements " + str(agreementsIds) + " for Agreements on Customer"

		try:
			# get the Customer
			customer = self.get( customerId ).first()
				
			# split on a comma with no spaces
			idList = agreementsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Agreement		
				agreement = AgreementDelegate().get(id).first();	
				# add the Agreement
				customer.agreements.remove(agreement)
				
			# save it		
			customer.save()
			
			# reload and return the appropriate version
			return self.get( customerId );
		except Customer.DoesNotExist:
			raise ProcessingError(errMsg + " : Customer with id " + str(customerId) + " does not exist.")
		except Agreement.DoesNotExist:
			raise ProcessingError(errMsg + " : Agreement does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addLoanApplications( self, customerId, loanApplicationsIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.LoanApplicationDelegate import LoanApplicationDelegate

		errMsg = "Failed to add elements " + str(loanApplicationsIds) + " for LoanApplications on Customer"

		try:
			# get the Customer
			customer = self.get( customerId ).first()
				
			# split on a comma with no spaces
			idList = loanApplicationsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the LoanApplication		
				loanApplication = LoanApplicationDelegate().get(id).first();	
				# add the LoanApplication
				customer.loanApplications.add(loanApplication)
				
			# save it		
			customer.save()
			
			# reload and return the appropriate version
			return self.get( customerId );
		except Customer.DoesNotExist:
			raise ProcessingError(errMsg + " : Customer with id " + str(customerId) + " does not exist.")
		except LoanApplication.DoesNotExist:
			raise ProcessingError(errMsg + " : LoanApplication does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeLoanApplications( self, customerId, loanApplicationsIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.LoanApplicationDelegate import LoanApplicationDelegate

		errMsg = "Failed to remove elements " + str(loanApplicationsIds) + " for LoanApplications on Customer"

		try:
			# get the Customer
			customer = self.get( customerId ).first()
				
			# split on a comma with no spaces
			idList = loanApplicationsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the LoanApplication		
				loanApplication = LoanApplicationDelegate().get(id).first();	
				# add the LoanApplication
				customer.loanApplications.remove(loanApplication)
				
			# save it		
			customer.save()
			
			# reload and return the appropriate version
			return self.get( customerId );
		except Customer.DoesNotExist:
			raise ProcessingError(errMsg + " : Customer with id " + str(customerId) + " does not exist.")
		except LoanApplication.DoesNotExist:
			raise ProcessingError(errMsg + " : LoanApplication does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addLoans( self, customerId, loansIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.LoanDelegate import LoanDelegate

		errMsg = "Failed to add elements " + str(loansIds) + " for Loans on Customer"

		try:
			# get the Customer
			customer = self.get( customerId ).first()
				
			# split on a comma with no spaces
			idList = loansIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Loan		
				loan = LoanDelegate().get(id).first();	
				# add the Loan
				customer.loans.add(loan)
				
			# save it		
			customer.save()
			
			# reload and return the appropriate version
			return self.get( customerId );
		except Customer.DoesNotExist:
			raise ProcessingError(errMsg + " : Customer with id " + str(customerId) + " does not exist.")
		except Loan.DoesNotExist:
			raise ProcessingError(errMsg + " : Loan does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeLoans( self, customerId, loansIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.LoanDelegate import LoanDelegate

		errMsg = "Failed to remove elements " + str(loansIds) + " for Loans on Customer"

		try:
			# get the Customer
			customer = self.get( customerId ).first()
				
			# split on a comma with no spaces
			idList = loansIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Loan		
				loan = LoanDelegate().get(id).first();	
				# add the Loan
				customer.loans.remove(loan)
				
			# save it		
			customer.save()
			
			# reload and return the appropriate version
			return self.get( customerId );
		except Customer.DoesNotExist:
			raise ProcessingError(errMsg + " : Customer with id " + str(customerId) + " does not exist.")
		except Loan.DoesNotExist:
			raise ProcessingError(errMsg + " : Loan does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addPortfolios( self, customerId, portfoliosIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.InvestmentPortfolioDelegate import InvestmentPortfolioDelegate

		errMsg = "Failed to add elements " + str(portfoliosIds) + " for Portfolios on Customer"

		try:
			# get the Customer
			customer = self.get( customerId ).first()
				
			# split on a comma with no spaces
			idList = portfoliosIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the InvestmentPortfolio		
				investmentPortfolio = InvestmentPortfolioDelegate().get(id).first();	
				# add the InvestmentPortfolio
				customer.portfolios.add(investmentPortfolio)
				
			# save it		
			customer.save()
			
			# reload and return the appropriate version
			return self.get( customerId );
		except Customer.DoesNotExist:
			raise ProcessingError(errMsg + " : Customer with id " + str(customerId) + " does not exist.")
		except InvestmentPortfolio.DoesNotExist:
			raise ProcessingError(errMsg + " : InvestmentPortfolio does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removePortfolios( self, customerId, portfoliosIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.InvestmentPortfolioDelegate import InvestmentPortfolioDelegate

		errMsg = "Failed to remove elements " + str(portfoliosIds) + " for Portfolios on Customer"

		try:
			# get the Customer
			customer = self.get( customerId ).first()
				
			# split on a comma with no spaces
			idList = portfoliosIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the InvestmentPortfolio		
				investmentPortfolio = InvestmentPortfolioDelegate().get(id).first();	
				# add the InvestmentPortfolio
				customer.portfolios.remove(investmentPortfolio)
				
			# save it		
			customer.save()
			
			# reload and return the appropriate version
			return self.get( customerId );
		except Customer.DoesNotExist:
			raise ProcessingError(errMsg + " : Customer with id " + str(customerId) + " does not exist.")
		except InvestmentPortfolio.DoesNotExist:
			raise ProcessingError(errMsg + " : InvestmentPortfolio does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addDisputes( self, customerId, disputesIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.DisputeDelegate import DisputeDelegate

		errMsg = "Failed to add elements " + str(disputesIds) + " for Disputes on Customer"

		try:
			# get the Customer
			customer = self.get( customerId ).first()
				
			# split on a comma with no spaces
			idList = disputesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Dispute		
				dispute = DisputeDelegate().get(id).first();	
				# add the Dispute
				customer.disputes.add(dispute)
				
			# save it		
			customer.save()
			
			# reload and return the appropriate version
			return self.get( customerId );
		except Customer.DoesNotExist:
			raise ProcessingError(errMsg + " : Customer with id " + str(customerId) + " does not exist.")
		except Dispute.DoesNotExist:
			raise ProcessingError(errMsg + " : Dispute does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeDisputes( self, customerId, disputesIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.DisputeDelegate import DisputeDelegate

		errMsg = "Failed to remove elements " + str(disputesIds) + " for Disputes on Customer"

		try:
			# get the Customer
			customer = self.get( customerId ).first()
				
			# split on a comma with no spaces
			idList = disputesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Dispute		
				dispute = DisputeDelegate().get(id).first();	
				# add the Dispute
				customer.disputes.remove(dispute)
				
			# save it		
			customer.save()
			
			# reload and return the appropriate version
			return self.get( customerId );
		except Customer.DoesNotExist:
			raise ProcessingError(errMsg + " : Customer with id " + str(customerId) + " does not exist.")
		except Dispute.DoesNotExist:
			raise ProcessingError(errMsg + " : Dispute does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
