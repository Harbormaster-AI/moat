import React from 'react';
import './App.css';
import {BrowserRouter as Router, Route, Switch} from 'react-router-dom'
import HomePageComponent from './components/HomePageComponent';
import HeaderComponent from './components/HeaderComponent';
import FooterComponent from './components/FooterComponent';
import ListInsurerComponent from './components/ListInsurerComponent';
import CreateInsurerComponent from './components/CreateInsurerComponent';
import ViewInsurerComponent from './components/ViewInsurerComponent';
import ListInsuranceProductComponent from './components/ListInsuranceProductComponent';
import CreateInsuranceProductComponent from './components/CreateInsuranceProductComponent';
import ViewInsuranceProductComponent from './components/ViewInsuranceProductComponent';
import ListCoverageDefinitionComponent from './components/ListCoverageDefinitionComponent';
import CreateCoverageDefinitionComponent from './components/CreateCoverageDefinitionComponent';
import ViewCoverageDefinitionComponent from './components/ViewCoverageDefinitionComponent';
import ListDistributorComponent from './components/ListDistributorComponent';
import CreateDistributorComponent from './components/CreateDistributorComponent';
import ViewDistributorComponent from './components/ViewDistributorComponent';
import ListAgentComponent from './components/ListAgentComponent';
import CreateAgentComponent from './components/CreateAgentComponent';
import ViewAgentComponent from './components/ViewAgentComponent';
import ListCustomerComponent from './components/ListCustomerComponent';
import CreateCustomerComponent from './components/CreateCustomerComponent';
import ViewCustomerComponent from './components/ViewCustomerComponent';
import ListApplicationComponent from './components/ListApplicationComponent';
import CreateApplicationComponent from './components/CreateApplicationComponent';
import ViewApplicationComponent from './components/ViewApplicationComponent';
import ListQuoteComponent from './components/ListQuoteComponent';
import CreateQuoteComponent from './components/CreateQuoteComponent';
import ViewQuoteComponent from './components/ViewQuoteComponent';
import ListUnderwritingDecisionComponent from './components/ListUnderwritingDecisionComponent';
import CreateUnderwritingDecisionComponent from './components/CreateUnderwritingDecisionComponent';
import ViewUnderwritingDecisionComponent from './components/ViewUnderwritingDecisionComponent';
import ListUnderwriterComponent from './components/ListUnderwriterComponent';
import CreateUnderwriterComponent from './components/CreateUnderwriterComponent';
import ViewUnderwriterComponent from './components/ViewUnderwriterComponent';
import ListPolicyComponent from './components/ListPolicyComponent';
import CreatePolicyComponent from './components/CreatePolicyComponent';
import ViewPolicyComponent from './components/ViewPolicyComponent';
import ListEndorsementComponent from './components/ListEndorsementComponent';
import CreateEndorsementComponent from './components/CreateEndorsementComponent';
import ViewEndorsementComponent from './components/ViewEndorsementComponent';
import ListPolicyCoverageComponent from './components/ListPolicyCoverageComponent';
import CreatePolicyCoverageComponent from './components/CreatePolicyCoverageComponent';
import ViewPolicyCoverageComponent from './components/ViewPolicyCoverageComponent';
import ListInsuredObjectComponent from './components/ListInsuredObjectComponent';
import CreateInsuredObjectComponent from './components/CreateInsuredObjectComponent';
import ViewInsuredObjectComponent from './components/ViewInsuredObjectComponent';
import ListBeneficiaryComponent from './components/ListBeneficiaryComponent';
import CreateBeneficiaryComponent from './components/CreateBeneficiaryComponent';
import ViewBeneficiaryComponent from './components/ViewBeneficiaryComponent';
import ListBillingAccountComponent from './components/ListBillingAccountComponent';
import CreateBillingAccountComponent from './components/CreateBillingAccountComponent';
import ViewBillingAccountComponent from './components/ViewBillingAccountComponent';
import ListInvoiceComponent from './components/ListInvoiceComponent';
import CreateInvoiceComponent from './components/CreateInvoiceComponent';
import ViewInvoiceComponent from './components/ViewInvoiceComponent';
import ListPaymentComponent from './components/ListPaymentComponent';
import CreatePaymentComponent from './components/CreatePaymentComponent';
import ViewPaymentComponent from './components/ViewPaymentComponent';
import ListClaimComponent from './components/ListClaimComponent';
import CreateClaimComponent from './components/CreateClaimComponent';
import ViewClaimComponent from './components/ViewClaimComponent';
import ListIncidentComponent from './components/ListIncidentComponent';
import CreateIncidentComponent from './components/CreateIncidentComponent';
import ViewIncidentComponent from './components/ViewIncidentComponent';
import ListExposureComponent from './components/ListExposureComponent';
import CreateExposureComponent from './components/CreateExposureComponent';
import ViewExposureComponent from './components/ViewExposureComponent';
import ListAdjusterComponent from './components/ListAdjusterComponent';
import CreateAdjusterComponent from './components/CreateAdjusterComponent';
import ViewAdjusterComponent from './components/ViewAdjusterComponent';
import ListClaimReserveComponent from './components/ListClaimReserveComponent';
import CreateClaimReserveComponent from './components/CreateClaimReserveComponent';
import ViewClaimReserveComponent from './components/ViewClaimReserveComponent';
import ListClaimPaymentComponent from './components/ListClaimPaymentComponent';
import CreateClaimPaymentComponent from './components/CreateClaimPaymentComponent';
import ViewClaimPaymentComponent from './components/ViewClaimPaymentComponent';
import ListServiceProviderComponent from './components/ListServiceProviderComponent';
import CreateServiceProviderComponent from './components/CreateServiceProviderComponent';
import ViewServiceProviderComponent from './components/ViewServiceProviderComponent';
import ListReinsuranceAgreementComponent from './components/ListReinsuranceAgreementComponent';
import CreateReinsuranceAgreementComponent from './components/CreateReinsuranceAgreementComponent';
import ViewReinsuranceAgreementComponent from './components/ViewReinsuranceAgreementComponent';
import ListSubrogationRecoveryComponent from './components/ListSubrogationRecoveryComponent';
import CreateSubrogationRecoveryComponent from './components/CreateSubrogationRecoveryComponent';
import ViewSubrogationRecoveryComponent from './components/ViewSubrogationRecoveryComponent';
import ListThirdPartyComponent from './components/ListThirdPartyComponent';
import CreateThirdPartyComponent from './components/CreateThirdPartyComponent';
import ViewThirdPartyComponent from './components/ViewThirdPartyComponent';
import ListDocumentComponent from './components/ListDocumentComponent';
import CreateDocumentComponent from './components/CreateDocumentComponent';
import ViewDocumentComponent from './components/ViewDocumentComponent';
function App() {
  return (
    <div>
        <Router>
                <HeaderComponent className="header"/>
                <div className="container">
                    <Switch>
                          <Route path = "/" exact component = {HomePageComponent}></Route>
                            <Route path = "/insurers" component = {ListInsurerComponent}></Route>
                            <Route path = "/add-insurer/:id" component = {CreateInsurerComponent}></Route>
                            <Route path = "/view-insurer/:id" component = {ViewInsurerComponent}></Route>
                          {/* <Route path = "/update-insurer/:id" component = {UpdateInsurerComponent}></Route> */}
                            <Route path = "/insuranceProducts" component = {ListInsuranceProductComponent}></Route>
                            <Route path = "/add-insuranceProduct/:id" component = {CreateInsuranceProductComponent}></Route>
                            <Route path = "/view-insuranceProduct/:id" component = {ViewInsuranceProductComponent}></Route>
                          {/* <Route path = "/update-insuranceProduct/:id" component = {UpdateInsuranceProductComponent}></Route> */}
                            <Route path = "/coverageDefinitions" component = {ListCoverageDefinitionComponent}></Route>
                            <Route path = "/add-coverageDefinition/:id" component = {CreateCoverageDefinitionComponent}></Route>
                            <Route path = "/view-coverageDefinition/:id" component = {ViewCoverageDefinitionComponent}></Route>
                          {/* <Route path = "/update-coverageDefinition/:id" component = {UpdateCoverageDefinitionComponent}></Route> */}
                            <Route path = "/distributors" component = {ListDistributorComponent}></Route>
                            <Route path = "/add-distributor/:id" component = {CreateDistributorComponent}></Route>
                            <Route path = "/view-distributor/:id" component = {ViewDistributorComponent}></Route>
                          {/* <Route path = "/update-distributor/:id" component = {UpdateDistributorComponent}></Route> */}
                            <Route path = "/agents" component = {ListAgentComponent}></Route>
                            <Route path = "/add-agent/:id" component = {CreateAgentComponent}></Route>
                            <Route path = "/view-agent/:id" component = {ViewAgentComponent}></Route>
                          {/* <Route path = "/update-agent/:id" component = {UpdateAgentComponent}></Route> */}
                            <Route path = "/customers" component = {ListCustomerComponent}></Route>
                            <Route path = "/add-customer/:id" component = {CreateCustomerComponent}></Route>
                            <Route path = "/view-customer/:id" component = {ViewCustomerComponent}></Route>
                          {/* <Route path = "/update-customer/:id" component = {UpdateCustomerComponent}></Route> */}
                            <Route path = "/applications" component = {ListApplicationComponent}></Route>
                            <Route path = "/add-application/:id" component = {CreateApplicationComponent}></Route>
                            <Route path = "/view-application/:id" component = {ViewApplicationComponent}></Route>
                          {/* <Route path = "/update-application/:id" component = {UpdateApplicationComponent}></Route> */}
                            <Route path = "/quotes" component = {ListQuoteComponent}></Route>
                            <Route path = "/add-quote/:id" component = {CreateQuoteComponent}></Route>
                            <Route path = "/view-quote/:id" component = {ViewQuoteComponent}></Route>
                          {/* <Route path = "/update-quote/:id" component = {UpdateQuoteComponent}></Route> */}
                            <Route path = "/underwritingDecisions" component = {ListUnderwritingDecisionComponent}></Route>
                            <Route path = "/add-underwritingDecision/:id" component = {CreateUnderwritingDecisionComponent}></Route>
                            <Route path = "/view-underwritingDecision/:id" component = {ViewUnderwritingDecisionComponent}></Route>
                          {/* <Route path = "/update-underwritingDecision/:id" component = {UpdateUnderwritingDecisionComponent}></Route> */}
                            <Route path = "/underwriters" component = {ListUnderwriterComponent}></Route>
                            <Route path = "/add-underwriter/:id" component = {CreateUnderwriterComponent}></Route>
                            <Route path = "/view-underwriter/:id" component = {ViewUnderwriterComponent}></Route>
                          {/* <Route path = "/update-underwriter/:id" component = {UpdateUnderwriterComponent}></Route> */}
                            <Route path = "/policys" component = {ListPolicyComponent}></Route>
                            <Route path = "/add-policy/:id" component = {CreatePolicyComponent}></Route>
                            <Route path = "/view-policy/:id" component = {ViewPolicyComponent}></Route>
                          {/* <Route path = "/update-policy/:id" component = {UpdatePolicyComponent}></Route> */}
                            <Route path = "/endorsements" component = {ListEndorsementComponent}></Route>
                            <Route path = "/add-endorsement/:id" component = {CreateEndorsementComponent}></Route>
                            <Route path = "/view-endorsement/:id" component = {ViewEndorsementComponent}></Route>
                          {/* <Route path = "/update-endorsement/:id" component = {UpdateEndorsementComponent}></Route> */}
                            <Route path = "/policyCoverages" component = {ListPolicyCoverageComponent}></Route>
                            <Route path = "/add-policyCoverage/:id" component = {CreatePolicyCoverageComponent}></Route>
                            <Route path = "/view-policyCoverage/:id" component = {ViewPolicyCoverageComponent}></Route>
                          {/* <Route path = "/update-policyCoverage/:id" component = {UpdatePolicyCoverageComponent}></Route> */}
                            <Route path = "/insuredObjects" component = {ListInsuredObjectComponent}></Route>
                            <Route path = "/add-insuredObject/:id" component = {CreateInsuredObjectComponent}></Route>
                            <Route path = "/view-insuredObject/:id" component = {ViewInsuredObjectComponent}></Route>
                          {/* <Route path = "/update-insuredObject/:id" component = {UpdateInsuredObjectComponent}></Route> */}
                            <Route path = "/beneficiarys" component = {ListBeneficiaryComponent}></Route>
                            <Route path = "/add-beneficiary/:id" component = {CreateBeneficiaryComponent}></Route>
                            <Route path = "/view-beneficiary/:id" component = {ViewBeneficiaryComponent}></Route>
                          {/* <Route path = "/update-beneficiary/:id" component = {UpdateBeneficiaryComponent}></Route> */}
                            <Route path = "/billingAccounts" component = {ListBillingAccountComponent}></Route>
                            <Route path = "/add-billingAccount/:id" component = {CreateBillingAccountComponent}></Route>
                            <Route path = "/view-billingAccount/:id" component = {ViewBillingAccountComponent}></Route>
                          {/* <Route path = "/update-billingAccount/:id" component = {UpdateBillingAccountComponent}></Route> */}
                            <Route path = "/invoices" component = {ListInvoiceComponent}></Route>
                            <Route path = "/add-invoice/:id" component = {CreateInvoiceComponent}></Route>
                            <Route path = "/view-invoice/:id" component = {ViewInvoiceComponent}></Route>
                          {/* <Route path = "/update-invoice/:id" component = {UpdateInvoiceComponent}></Route> */}
                            <Route path = "/payments" component = {ListPaymentComponent}></Route>
                            <Route path = "/add-payment/:id" component = {CreatePaymentComponent}></Route>
                            <Route path = "/view-payment/:id" component = {ViewPaymentComponent}></Route>
                          {/* <Route path = "/update-payment/:id" component = {UpdatePaymentComponent}></Route> */}
                            <Route path = "/claims" component = {ListClaimComponent}></Route>
                            <Route path = "/add-claim/:id" component = {CreateClaimComponent}></Route>
                            <Route path = "/view-claim/:id" component = {ViewClaimComponent}></Route>
                          {/* <Route path = "/update-claim/:id" component = {UpdateClaimComponent}></Route> */}
                            <Route path = "/incidents" component = {ListIncidentComponent}></Route>
                            <Route path = "/add-incident/:id" component = {CreateIncidentComponent}></Route>
                            <Route path = "/view-incident/:id" component = {ViewIncidentComponent}></Route>
                          {/* <Route path = "/update-incident/:id" component = {UpdateIncidentComponent}></Route> */}
                            <Route path = "/exposures" component = {ListExposureComponent}></Route>
                            <Route path = "/add-exposure/:id" component = {CreateExposureComponent}></Route>
                            <Route path = "/view-exposure/:id" component = {ViewExposureComponent}></Route>
                          {/* <Route path = "/update-exposure/:id" component = {UpdateExposureComponent}></Route> */}
                            <Route path = "/adjusters" component = {ListAdjusterComponent}></Route>
                            <Route path = "/add-adjuster/:id" component = {CreateAdjusterComponent}></Route>
                            <Route path = "/view-adjuster/:id" component = {ViewAdjusterComponent}></Route>
                          {/* <Route path = "/update-adjuster/:id" component = {UpdateAdjusterComponent}></Route> */}
                            <Route path = "/claimReserves" component = {ListClaimReserveComponent}></Route>
                            <Route path = "/add-claimReserve/:id" component = {CreateClaimReserveComponent}></Route>
                            <Route path = "/view-claimReserve/:id" component = {ViewClaimReserveComponent}></Route>
                          {/* <Route path = "/update-claimReserve/:id" component = {UpdateClaimReserveComponent}></Route> */}
                            <Route path = "/claimPayments" component = {ListClaimPaymentComponent}></Route>
                            <Route path = "/add-claimPayment/:id" component = {CreateClaimPaymentComponent}></Route>
                            <Route path = "/view-claimPayment/:id" component = {ViewClaimPaymentComponent}></Route>
                          {/* <Route path = "/update-claimPayment/:id" component = {UpdateClaimPaymentComponent}></Route> */}
                            <Route path = "/serviceProviders" component = {ListServiceProviderComponent}></Route>
                            <Route path = "/add-serviceProvider/:id" component = {CreateServiceProviderComponent}></Route>
                            <Route path = "/view-serviceProvider/:id" component = {ViewServiceProviderComponent}></Route>
                          {/* <Route path = "/update-serviceProvider/:id" component = {UpdateServiceProviderComponent}></Route> */}
                            <Route path = "/reinsuranceAgreements" component = {ListReinsuranceAgreementComponent}></Route>
                            <Route path = "/add-reinsuranceAgreement/:id" component = {CreateReinsuranceAgreementComponent}></Route>
                            <Route path = "/view-reinsuranceAgreement/:id" component = {ViewReinsuranceAgreementComponent}></Route>
                          {/* <Route path = "/update-reinsuranceAgreement/:id" component = {UpdateReinsuranceAgreementComponent}></Route> */}
                            <Route path = "/subrogationRecoverys" component = {ListSubrogationRecoveryComponent}></Route>
                            <Route path = "/add-subrogationRecovery/:id" component = {CreateSubrogationRecoveryComponent}></Route>
                            <Route path = "/view-subrogationRecovery/:id" component = {ViewSubrogationRecoveryComponent}></Route>
                          {/* <Route path = "/update-subrogationRecovery/:id" component = {UpdateSubrogationRecoveryComponent}></Route> */}
                            <Route path = "/thirdPartys" component = {ListThirdPartyComponent}></Route>
                            <Route path = "/add-thirdParty/:id" component = {CreateThirdPartyComponent}></Route>
                            <Route path = "/view-thirdParty/:id" component = {ViewThirdPartyComponent}></Route>
                          {/* <Route path = "/update-thirdParty/:id" component = {UpdateThirdPartyComponent}></Route> */}
                            <Route path = "/documents" component = {ListDocumentComponent}></Route>
                            <Route path = "/add-document/:id" component = {CreateDocumentComponent}></Route>
                            <Route path = "/view-document/:id" component = {ViewDocumentComponent}></Route>
                          {/* <Route path = "/update-document/:id" component = {UpdateDocumentComponent}></Route> */}
                    </Switch>
                </div>
              <FooterComponent />
        </Router>
    </div>
    
  );
}

export default App;
