import { HttpClient } from '@angular/common/http';
import * as enumTypes from '../models/EnumTypes';

import {InsurerService} from '../services/Insurer.service';
import {InsuranceProductService} from '../services/InsuranceProduct.service';
import {CoverageDefinitionService} from '../services/CoverageDefinition.service';
import {DistributorService} from '../services/Distributor.service';
import {AgentService} from '../services/Agent.service';
import {CustomerService} from '../services/Customer.service';
import {ApplicationService} from '../services/Application.service';
import {QuoteService} from '../services/Quote.service';
import {UnderwritingDecisionService} from '../services/UnderwritingDecision.service';
import {UnderwriterService} from '../services/Underwriter.service';
import {PolicyService} from '../services/Policy.service';
import {EndorsementService} from '../services/Endorsement.service';
import {PolicyCoverageService} from '../services/PolicyCoverage.service';
import {InsuredObjectService} from '../services/InsuredObject.service';
import {BeneficiaryService} from '../services/Beneficiary.service';
import {BillingAccountService} from '../services/BillingAccount.service';
import {InvoiceService} from '../services/Invoice.service';
import {PaymentService} from '../services/Payment.service';
import {ClaimService} from '../services/Claim.service';
import {IncidentService} from '../services/Incident.service';
import {ExposureService} from '../services/Exposure.service';
import {AdjusterService} from '../services/Adjuster.service';
import {ClaimReserveService} from '../services/ClaimReserve.service';
import {ClaimPaymentService} from '../services/ClaimPayment.service';
import {ServiceProviderService} from '../services/ServiceProvider.service';
import {ReinsuranceAgreementService} from '../services/ReinsuranceAgreement.service';
import {SubrogationRecoveryService} from '../services/SubrogationRecovery.service';
import {ThirdPartyService} from '../services/ThirdParty.service';
import {DocumentService} from '../services/Document.service';

import { Directive } from '@angular/core';

/**
 Base class of all Components.
 For convenience, contains all enums and entity lists
 **/

@Directive()
export class BaseComponent {

    constructor (private http: HttpClient) {}

// enum instances
    LineOfBusinesss = Object.keys(enumTypes.LineOfBusiness);
    CoverageTypes = Object.keys(enumTypes.CoverageType);
    DistributionChannelTypes = Object.keys(enumTypes.DistributionChannelType);
    ProducerStatuss = Object.keys(enumTypes.ProducerStatus);
    CustomerTypes = Object.keys(enumTypes.CustomerType);
    ApplicationStatuss = Object.keys(enumTypes.ApplicationStatus);
    UnderwritingDecisionTypes = Object.keys(enumTypes.UnderwritingDecisionType);
    PolicyStatuss = Object.keys(enumTypes.PolicyStatus);
    PaymentPlanTypes = Object.keys(enumTypes.PaymentPlanType);
    BillingStatuss = Object.keys(enumTypes.BillingStatus);
    InvoiceStatuss = Object.keys(enumTypes.InvoiceStatus);
    PaymentMethods = Object.keys(enumTypes.PaymentMethod);
    PaymentStatuss = Object.keys(enumTypes.PaymentStatus);
    InsuredObjectTypes = Object.keys(enumTypes.InsuredObjectType);
    RelationshipTypes = Object.keys(enumTypes.RelationshipType);
    ClaimStatuss = Object.keys(enumTypes.ClaimStatus);
    PerilTypes = Object.keys(enumTypes.PerilType);
    CauseOfLosss = Object.keys(enumTypes.CauseOfLoss);
    ExposureTypes = Object.keys(enumTypes.ExposureType);
    ExposureStatuss = Object.keys(enumTypes.ExposureStatus);
    AdjusterTypes = Object.keys(enumTypes.AdjusterType);
    ReserveTypes = Object.keys(enumTypes.ReserveType);
    ReserveStatuss = Object.keys(enumTypes.ReserveStatus);
    PayeeTypes = Object.keys(enumTypes.PayeeType);
    ServiceProviderTypes = Object.keys(enumTypes.ServiceProviderType);
    NetworkStatuss = Object.keys(enumTypes.NetworkStatus);
    ReinsuranceTypes = Object.keys(enumTypes.ReinsuranceType);
    TreatyTypes = Object.keys(enumTypes.TreatyType);
    ThirdPartyTypes = Object.keys(enumTypes.ThirdPartyType);
    DocumentTypes = Object.keys(enumTypes.DocumentType);
    SubrogationStatuss = Object.keys(enumTypes.SubrogationStatus);

// all collection instances
    insurers : any;
    insuranceProducts : any;
    coverageDefinitions : any;
    distributors : any;
    agents : any;
    customers : any;
    applications : any;
    quotes : any;
    underwritingDecisions : any;
    underwriters : any;
    policys : any;
    endorsements : any;
    policyCoverages : any;
    insuredObjects : any;
    beneficiarys : any;
    billingAccounts : any;
    invoices : any;
    payments : any;
    claims : any;
    incidents : any;
    exposures : any;
    adjusters : any;
    claimReserves : any;
    claimPayments : any;
    serviceProviders : any;
    reinsuranceAgreements : any;
    subrogationRecoverys : any;
    thirdPartys : any;
    documents : any;
  
// initialization  
    ngOnInit() {
    }

    initInsurerList() {
        if ( this.insurers == null ) {
            new InsurerService(this.http).getInsurers().subscribe(res => {
                this.insurers = res;
            });
        }
    }
    
    initInsuranceProductList() {
        if ( this.insuranceProducts == null ) {
            new InsuranceProductService(this.http).getInsuranceProducts().subscribe(res => {
                this.insuranceProducts = res;
            });
        }
    }
    
    initCoverageDefinitionList() {
        if ( this.coverageDefinitions == null ) {
            new CoverageDefinitionService(this.http).getCoverageDefinitions().subscribe(res => {
                this.coverageDefinitions = res;
            });
        }
    }
    
    initDistributorList() {
        if ( this.distributors == null ) {
            new DistributorService(this.http).getDistributors().subscribe(res => {
                this.distributors = res;
            });
        }
    }
    
    initAgentList() {
        if ( this.agents == null ) {
            new AgentService(this.http).getAgents().subscribe(res => {
                this.agents = res;
            });
        }
    }
    
    initCustomerList() {
        if ( this.customers == null ) {
            new CustomerService(this.http).getCustomers().subscribe(res => {
                this.customers = res;
            });
        }
    }
    
    initApplicationList() {
        if ( this.applications == null ) {
            new ApplicationService(this.http).getApplications().subscribe(res => {
                this.applications = res;
            });
        }
    }
    
    initQuoteList() {
        if ( this.quotes == null ) {
            new QuoteService(this.http).getQuotes().subscribe(res => {
                this.quotes = res;
            });
        }
    }
    
    initUnderwritingDecisionList() {
        if ( this.underwritingDecisions == null ) {
            new UnderwritingDecisionService(this.http).getUnderwritingDecisions().subscribe(res => {
                this.underwritingDecisions = res;
            });
        }
    }
    
    initUnderwriterList() {
        if ( this.underwriters == null ) {
            new UnderwriterService(this.http).getUnderwriters().subscribe(res => {
                this.underwriters = res;
            });
        }
    }
    
    initPolicyList() {
        if ( this.policys == null ) {
            new PolicyService(this.http).getPolicys().subscribe(res => {
                this.policys = res;
            });
        }
    }
    
    initEndorsementList() {
        if ( this.endorsements == null ) {
            new EndorsementService(this.http).getEndorsements().subscribe(res => {
                this.endorsements = res;
            });
        }
    }
    
    initPolicyCoverageList() {
        if ( this.policyCoverages == null ) {
            new PolicyCoverageService(this.http).getPolicyCoverages().subscribe(res => {
                this.policyCoverages = res;
            });
        }
    }
    
    initInsuredObjectList() {
        if ( this.insuredObjects == null ) {
            new InsuredObjectService(this.http).getInsuredObjects().subscribe(res => {
                this.insuredObjects = res;
            });
        }
    }
    
    initBeneficiaryList() {
        if ( this.beneficiarys == null ) {
            new BeneficiaryService(this.http).getBeneficiarys().subscribe(res => {
                this.beneficiarys = res;
            });
        }
    }
    
    initBillingAccountList() {
        if ( this.billingAccounts == null ) {
            new BillingAccountService(this.http).getBillingAccounts().subscribe(res => {
                this.billingAccounts = res;
            });
        }
    }
    
    initInvoiceList() {
        if ( this.invoices == null ) {
            new InvoiceService(this.http).getInvoices().subscribe(res => {
                this.invoices = res;
            });
        }
    }
    
    initPaymentList() {
        if ( this.payments == null ) {
            new PaymentService(this.http).getPayments().subscribe(res => {
                this.payments = res;
            });
        }
    }
    
    initClaimList() {
        if ( this.claims == null ) {
            new ClaimService(this.http).getClaims().subscribe(res => {
                this.claims = res;
            });
        }
    }
    
    initIncidentList() {
        if ( this.incidents == null ) {
            new IncidentService(this.http).getIncidents().subscribe(res => {
                this.incidents = res;
            });
        }
    }
    
    initExposureList() {
        if ( this.exposures == null ) {
            new ExposureService(this.http).getExposures().subscribe(res => {
                this.exposures = res;
            });
        }
    }
    
    initAdjusterList() {
        if ( this.adjusters == null ) {
            new AdjusterService(this.http).getAdjusters().subscribe(res => {
                this.adjusters = res;
            });
        }
    }
    
    initClaimReserveList() {
        if ( this.claimReserves == null ) {
            new ClaimReserveService(this.http).getClaimReserves().subscribe(res => {
                this.claimReserves = res;
            });
        }
    }
    
    initClaimPaymentList() {
        if ( this.claimPayments == null ) {
            new ClaimPaymentService(this.http).getClaimPayments().subscribe(res => {
                this.claimPayments = res;
            });
        }
    }
    
    initServiceProviderList() {
        if ( this.serviceProviders == null ) {
            new ServiceProviderService(this.http).getServiceProviders().subscribe(res => {
                this.serviceProviders = res;
            });
        }
    }
    
    initReinsuranceAgreementList() {
        if ( this.reinsuranceAgreements == null ) {
            new ReinsuranceAgreementService(this.http).getReinsuranceAgreements().subscribe(res => {
                this.reinsuranceAgreements = res;
            });
        }
    }
    
    initSubrogationRecoveryList() {
        if ( this.subrogationRecoverys == null ) {
            new SubrogationRecoveryService(this.http).getSubrogationRecoverys().subscribe(res => {
                this.subrogationRecoverys = res;
            });
        }
    }
    
    initThirdPartyList() {
        if ( this.thirdPartys == null ) {
            new ThirdPartyService(this.http).getThirdPartys().subscribe(res => {
                this.thirdPartys = res;
            });
        }
    }
    
    initDocumentList() {
        if ( this.documents == null ) {
            new DocumentService(this.http).getDocuments().subscribe(res => {
                this.documents = res;
            });
        }
    }
    
    
// comparison function for select controls  
    compareFn(user1: any, user2: any) {
        return user1 == user2
    }    
}
