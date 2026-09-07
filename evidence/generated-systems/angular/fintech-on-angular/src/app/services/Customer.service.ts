import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Customer} from '../models/Customer';
import {FinancialInstitutionService} from '../services/FinancialInstitution.service';
import {AccountService} from '../services/Account.service';
import {WalletService} from '../services/Wallet.service';
import {PaymentCardService} from '../services/PaymentCard.service';
import {KYCProfileService} from '../services/KYCProfile.service';
import {ConsentService} from '../services/Consent.service';
import {AgreementService} from '../services/Agreement.service';
import {LoanApplicationService} from '../services/LoanApplication.service';
import {LoanService} from '../services/Loan.service';
import {InvestmentPortfolioService} from '../services/InvestmentPortfolio.service';
import {DisputeService} from '../services/Dispute.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class CustomerService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	customer : Customer;

	//********************************************************************
	// Catch all for the return value of a service call
	//********************************************************************
	result: any;

	//********************************************************************
	// sole constructor, injected with the HttpClient
	//********************************************************************
	constructor(private http: HttpClient) {
		super();
	}

		//********************************************************************
	// add a Customer
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addCustomer(firstName, lastName, dateOfBirth, email, phone, address, taxId, riskScore, Institution, Accounts, Wallets, Cards, KycProfiles, Consents, Agreements, LoanApplications, Loans, Portfolios, Disputes, CustomerType) : Observable<any> {
		const uri_ = this.apiUrl + '/Customer/create';
		const obj = {
			      		firstName: firstName,
      		lastName: lastName,
      		dateOfBirth: dateOfBirth,
      		email: email,
      		phone: phone,
      		address: address,
      		taxId: taxId,
      		riskScore: riskScore,
      		Institution: Institution != null && Institution.length > 0 ? Institution : null,
      		Accounts: Accounts != null && Accounts.length > 0 ? Accounts : null,
      		Wallets: Wallets != null && Wallets.length > 0 ? Wallets : null,
      		Cards: Cards != null && Cards.length > 0 ? Cards : null,
      		KycProfiles: KycProfiles != null && KycProfiles.length > 0 ? KycProfiles : null,
      		Consents: Consents != null && Consents.length > 0 ? Consents : null,
      		Agreements: Agreements != null && Agreements.length > 0 ? Agreements : null,
      		LoanApplications: LoanApplications != null && LoanApplications.length > 0 ? LoanApplications : null,
      		Loans: Loans != null && Loans.length > 0 ? Loans : null,
      		Portfolios: Portfolios != null && Portfolios.length > 0 ? Portfolios : null,
      		Disputes: Disputes != null && Disputes.length > 0 ? Disputes : null,
			CustomerType: CustomerType
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Customer
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateCustomer(firstName, lastName, dateOfBirth, email, phone, address, taxId, riskScore, Institution, Accounts, Wallets, Cards, KycProfiles, Consents, Agreements, LoanApplications, Loans, Portfolios, Disputes, CustomerType, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Customer/update/' + id;
		const obj = {
				      		firstName: firstName,
      		lastName: lastName,
      		dateOfBirth: dateOfBirth,
      		email: email,
      		phone: phone,
      		address: address,
      		taxId: taxId,
      		riskScore: riskScore,
      		Institution: Institution != null && Institution.length > 0 ? Institution : null,
      		Accounts: Accounts != null && Accounts.length > 0 ? Accounts : null,
      		Wallets: Wallets != null && Wallets.length > 0 ? Wallets : null,
      		Cards: Cards != null && Cards.length > 0 ? Cards : null,
      		KycProfiles: KycProfiles != null && KycProfiles.length > 0 ? KycProfiles : null,
      		Consents: Consents != null && Consents.length > 0 ? Consents : null,
      		Agreements: Agreements != null && Agreements.length > 0 ? Agreements : null,
      		LoanApplications: LoanApplications != null && LoanApplications.length > 0 ? LoanApplications : null,
      		Loans: Loans != null && Loans.length > 0 ? Loans : null,
      		Portfolios: Portfolios != null && Portfolios.length > 0 ? Portfolios : null,
      		Disputes: Disputes != null && Disputes.length > 0 ? Disputes : null,
			CustomerType: CustomerType
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Customer
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteCustomer(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Customer/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Customer
	// returns the results untouched as an Observable Customer
	// Customer model
	// delegates via URI
	//********************************************************************
	getCustomer(id) : Observable<Customer> {
		const uri_ = this.apiUrl + '/Customer/load/' + id;

		return this.http.get<Customer>(uri_);
	}
	
	//********************************************************************
	// gets all Customer
	// returns the results untouched as JSON representation of an
	// Observable array of Customer models
	// delegates via URI
	//********************************************************************
	getCustomers() : Observable<Customer[]> {
		const uri_ = this.apiUrl + '/Customer/';

		return this
			.http.get<Customer[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Institution on a Customer
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignInstitution( customerId, _institutionId ): Observable<any> {

		// get the Customer from storage
		this.loadHelper( customerId );

	// get the FinancialInstitution from storage
	var tmp 	= new FinancialInstitutionService(this.http).getFinancialInstitution(_institutionId);

	// assign the Institution
	this.customer.institution = tmp;

	// save the Customer
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Institution on a Customer
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignInstitution( customerId ): Observable<any> {

		// get the Customer from storage
		this.loadHelper( customerId );

	// assign Institution to null
	this.customer.institution = null;

	// save the Customer
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more accountsIds as a Accounts
	// to a Customer
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addAccounts( customerId, accountsIds ): Observable<any> {

		// get the Customer
		this.loadHelper( customerId );

	// split on a comma with no spaces
	var idList = accountsIds.split(',')

	// iterate over array of accounts ids
	idList.forEach(function (id) {
		// read the Account
		var account = new AccountService(this.http).getAccount(id);
		// add the Account if not already assigned
		if ( this.customer.accounts.indexOf(account) == -1 )
		this.customer.accounts.push(account);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more accountsIds as a Accounts
	// from a Customer
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeAccounts( customerId, accountsIds ): Observable<any> {

		// get the Customer
		this.loadHelper( customerId );


	// split on a comma with no spaces
	var idList 					= accountsIds.split(',');
	var accounts 	= this.customer.accounts;

	if ( accounts != null && accountsIds != null ) {

		// iterate over array of accounts ids
		accounts.forEach(function (obj) {
			if ( accountsIds.indexOf(obj._id) > -1 ) {
				// remove the Account
				this.customer.accounts.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more walletsIds as a Wallets
	// to a Customer
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addWallets( customerId, walletsIds ): Observable<any> {

		// get the Customer
		this.loadHelper( customerId );

	// split on a comma with no spaces
	var idList = walletsIds.split(',')

	// iterate over array of wallets ids
	idList.forEach(function (id) {
		// read the Wallet
		var wallet = new WalletService(this.http).getWallet(id);
		// add the Wallet if not already assigned
		if ( this.customer.wallets.indexOf(wallet) == -1 )
		this.customer.wallets.push(wallet);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more walletsIds as a Wallets
	// from a Customer
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeWallets( customerId, walletsIds ): Observable<any> {

		// get the Customer
		this.loadHelper( customerId );


	// split on a comma with no spaces
	var idList 					= walletsIds.split(',');
	var wallets 	= this.customer.wallets;

	if ( wallets != null && walletsIds != null ) {

		// iterate over array of wallets ids
		wallets.forEach(function (obj) {
			if ( walletsIds.indexOf(obj._id) > -1 ) {
				// remove the Wallet
				this.customer.wallets.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more cardsIds as a Cards
	// to a Customer
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addCards( customerId, cardsIds ): Observable<any> {

		// get the Customer
		this.loadHelper( customerId );

	// split on a comma with no spaces
	var idList = cardsIds.split(',')

	// iterate over array of cards ids
	idList.forEach(function (id) {
		// read the PaymentCard
		var paymentCard = new PaymentCardService(this.http).getPaymentCard(id);
		// add the PaymentCard if not already assigned
		if ( this.customer.cards.indexOf(paymentCard) == -1 )
		this.customer.cards.push(paymentCard);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more cardsIds as a Cards
	// from a Customer
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeCards( customerId, cardsIds ): Observable<any> {

		// get the Customer
		this.loadHelper( customerId );


	// split on a comma with no spaces
	var idList 					= cardsIds.split(',');
	var cards 	= this.customer.cards;

	if ( cards != null && cardsIds != null ) {

		// iterate over array of cards ids
		cards.forEach(function (obj) {
			if ( cardsIds.indexOf(obj._id) > -1 ) {
				// remove the PaymentCard
				this.customer.cards.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more kycProfilesIds as a KycProfiles
	// to a Customer
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addKycProfiles( customerId, kycProfilesIds ): Observable<any> {

		// get the Customer
		this.loadHelper( customerId );

	// split on a comma with no spaces
	var idList = kycProfilesIds.split(',')

	// iterate over array of kycProfiles ids
	idList.forEach(function (id) {
		// read the KYCProfile
		var kYCProfile = new KYCProfileService(this.http).getKYCProfile(id);
		// add the KYCProfile if not already assigned
		if ( this.customer.kycProfiles.indexOf(kYCProfile) == -1 )
		this.customer.kycProfiles.push(kYCProfile);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more kycProfilesIds as a KycProfiles
	// from a Customer
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeKycProfiles( customerId, kycProfilesIds ): Observable<any> {

		// get the Customer
		this.loadHelper( customerId );


	// split on a comma with no spaces
	var idList 					= kycProfilesIds.split(',');
	var kycProfiles 	= this.customer.kycProfiles;

	if ( kycProfiles != null && kycProfilesIds != null ) {

		// iterate over array of kycProfiles ids
		kycProfiles.forEach(function (obj) {
			if ( kycProfilesIds.indexOf(obj._id) > -1 ) {
				// remove the KYCProfile
				this.customer.kycProfiles.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more consentsIds as a Consents
	// to a Customer
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addConsents( customerId, consentsIds ): Observable<any> {

		// get the Customer
		this.loadHelper( customerId );

	// split on a comma with no spaces
	var idList = consentsIds.split(',')

	// iterate over array of consents ids
	idList.forEach(function (id) {
		// read the Consent
		var consent = new ConsentService(this.http).getConsent(id);
		// add the Consent if not already assigned
		if ( this.customer.consents.indexOf(consent) == -1 )
		this.customer.consents.push(consent);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more consentsIds as a Consents
	// from a Customer
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeConsents( customerId, consentsIds ): Observable<any> {

		// get the Customer
		this.loadHelper( customerId );


	// split on a comma with no spaces
	var idList 					= consentsIds.split(',');
	var consents 	= this.customer.consents;

	if ( consents != null && consentsIds != null ) {

		// iterate over array of consents ids
		consents.forEach(function (obj) {
			if ( consentsIds.indexOf(obj._id) > -1 ) {
				// remove the Consent
				this.customer.consents.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more agreementsIds as a Agreements
	// to a Customer
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addAgreements( customerId, agreementsIds ): Observable<any> {

		// get the Customer
		this.loadHelper( customerId );

	// split on a comma with no spaces
	var idList = agreementsIds.split(',')

	// iterate over array of agreements ids
	idList.forEach(function (id) {
		// read the Agreement
		var agreement = new AgreementService(this.http).getAgreement(id);
		// add the Agreement if not already assigned
		if ( this.customer.agreements.indexOf(agreement) == -1 )
		this.customer.agreements.push(agreement);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more agreementsIds as a Agreements
	// from a Customer
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeAgreements( customerId, agreementsIds ): Observable<any> {

		// get the Customer
		this.loadHelper( customerId );


	// split on a comma with no spaces
	var idList 					= agreementsIds.split(',');
	var agreements 	= this.customer.agreements;

	if ( agreements != null && agreementsIds != null ) {

		// iterate over array of agreements ids
		agreements.forEach(function (obj) {
			if ( agreementsIds.indexOf(obj._id) > -1 ) {
				// remove the Agreement
				this.customer.agreements.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more loanApplicationsIds as a LoanApplications
	// to a Customer
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addLoanApplications( customerId, loanApplicationsIds ): Observable<any> {

		// get the Customer
		this.loadHelper( customerId );

	// split on a comma with no spaces
	var idList = loanApplicationsIds.split(',')

	// iterate over array of loanApplications ids
	idList.forEach(function (id) {
		// read the LoanApplication
		var loanApplication = new LoanApplicationService(this.http).getLoanApplication(id);
		// add the LoanApplication if not already assigned
		if ( this.customer.loanApplications.indexOf(loanApplication) == -1 )
		this.customer.loanApplications.push(loanApplication);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more loanApplicationsIds as a LoanApplications
	// from a Customer
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeLoanApplications( customerId, loanApplicationsIds ): Observable<any> {

		// get the Customer
		this.loadHelper( customerId );


	// split on a comma with no spaces
	var idList 					= loanApplicationsIds.split(',');
	var loanApplications 	= this.customer.loanApplications;

	if ( loanApplications != null && loanApplicationsIds != null ) {

		// iterate over array of loanApplications ids
		loanApplications.forEach(function (obj) {
			if ( loanApplicationsIds.indexOf(obj._id) > -1 ) {
				// remove the LoanApplication
				this.customer.loanApplications.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more loansIds as a Loans
	// to a Customer
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addLoans( customerId, loansIds ): Observable<any> {

		// get the Customer
		this.loadHelper( customerId );

	// split on a comma with no spaces
	var idList = loansIds.split(',')

	// iterate over array of loans ids
	idList.forEach(function (id) {
		// read the Loan
		var loan = new LoanService(this.http).getLoan(id);
		// add the Loan if not already assigned
		if ( this.customer.loans.indexOf(loan) == -1 )
		this.customer.loans.push(loan);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more loansIds as a Loans
	// from a Customer
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeLoans( customerId, loansIds ): Observable<any> {

		// get the Customer
		this.loadHelper( customerId );


	// split on a comma with no spaces
	var idList 					= loansIds.split(',');
	var loans 	= this.customer.loans;

	if ( loans != null && loansIds != null ) {

		// iterate over array of loans ids
		loans.forEach(function (obj) {
			if ( loansIds.indexOf(obj._id) > -1 ) {
				// remove the Loan
				this.customer.loans.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more portfoliosIds as a Portfolios
	// to a Customer
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addPortfolios( customerId, portfoliosIds ): Observable<any> {

		// get the Customer
		this.loadHelper( customerId );

	// split on a comma with no spaces
	var idList = portfoliosIds.split(',')

	// iterate over array of portfolios ids
	idList.forEach(function (id) {
		// read the InvestmentPortfolio
		var investmentPortfolio = new InvestmentPortfolioService(this.http).getInvestmentPortfolio(id);
		// add the InvestmentPortfolio if not already assigned
		if ( this.customer.portfolios.indexOf(investmentPortfolio) == -1 )
		this.customer.portfolios.push(investmentPortfolio);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more portfoliosIds as a Portfolios
	// from a Customer
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removePortfolios( customerId, portfoliosIds ): Observable<any> {

		// get the Customer
		this.loadHelper( customerId );


	// split on a comma with no spaces
	var idList 					= portfoliosIds.split(',');
	var portfolios 	= this.customer.portfolios;

	if ( portfolios != null && portfoliosIds != null ) {

		// iterate over array of portfolios ids
		portfolios.forEach(function (obj) {
			if ( portfoliosIds.indexOf(obj._id) > -1 ) {
				// remove the InvestmentPortfolio
				this.customer.portfolios.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more disputesIds as a Disputes
	// to a Customer
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addDisputes( customerId, disputesIds ): Observable<any> {

		// get the Customer
		this.loadHelper( customerId );

	// split on a comma with no spaces
	var idList = disputesIds.split(',')

	// iterate over array of disputes ids
	idList.forEach(function (id) {
		// read the Dispute
		var dispute = new DisputeService(this.http).getDispute(id);
		// add the Dispute if not already assigned
		if ( this.customer.disputes.indexOf(dispute) == -1 )
		this.customer.disputes.push(dispute);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more disputesIds as a Disputes
	// from a Customer
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeDisputes( customerId, disputesIds ): Observable<any> {

		// get the Customer
		this.loadHelper( customerId );


	// split on a comma with no spaces
	var idList 					= disputesIds.split(',');
	var disputes 	= this.customer.disputes;

	if ( disputes != null && disputesIds != null ) {

		// iterate over array of disputes ids
		disputes.forEach(function (obj) {
			if ( disputesIds.indexOf(obj._id) > -1 ) {
				// remove the Dispute
				this.customer.disputes.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a Customer
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Customer/update/' + this.customer;

	return  this.http.post(uri_, this.customer );
}

	//********************************************************************
	// loadHelper - internal helper to load a Customer
	//********************************************************************	
	loadHelper( id ) {
		this.getCustomer(id)
			.subscribe((res : Customer) => {
				this.customer = res;
			});
	}
}