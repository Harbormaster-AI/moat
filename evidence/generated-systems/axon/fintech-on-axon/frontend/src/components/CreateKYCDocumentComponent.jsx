import React, { Component } from 'react'
import KYCDocumentService from '../services/KYCDocumentService';

class CreateKYCDocumentComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                reference: '',
                issuedCountry: '',
                expirationDate: '',
                documentType: '',
                status: ''
        }
        this.changereferenceHandler = this.changereferenceHandler.bind(this);
        this.changeissuedCountryHandler = this.changeissuedCountryHandler.bind(this);
        this.changeexpirationDateHandler = this.changeexpirationDateHandler.bind(this);
        this.changeDocumentTypeHandler = this.changeDocumentTypeHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            KYCDocumentService.getKYCDocumentById(this.state.id).then( (res) =>{
                let kYCDocument = res.data;
                this.setState({
                    reference: kYCDocument.reference,
                    issuedCountry: kYCDocument.issuedCountry,
                    expirationDate: kYCDocument.expirationDate,
                    documentType: kYCDocument.documentType,
                    status: kYCDocument.status
                });
            });
        }        
    }
    saveOrUpdateKYCDocument = (e) => {
        e.preventDefault();
        let kYCDocument = {
                kYCDocumentId: this.state.id,
                reference: this.state.reference,
                issuedCountry: this.state.issuedCountry,
                expirationDate: this.state.expirationDate,
                documentType: this.state.documentType,
                status: this.state.status
            };
        console.log('kYCDocument => ' + JSON.stringify(kYCDocument));

        // step 5
        if(this.state.id === '_add'){
            kYCDocument.kYCDocumentId=''
            KYCDocumentService.createKYCDocument(kYCDocument).then(res =>{
                this.props.history.push('/kYCDocuments');
            });
        }else{
            KYCDocumentService.updateKYCDocument(kYCDocument).then( res => {
                this.props.history.push('/kYCDocuments');
            });
        }
    }
    
    changereferenceHandler= (event) => {
        this.setState({reference: event.target.value});
    }
    changeissuedCountryHandler= (event) => {
        this.setState({issuedCountry: event.target.value});
    }
    changeexpirationDateHandler= (event) => {
        this.setState({expirationDate: event.target.value});
    }
    changeDocumentTypeHandler= (event) => {
        this.setState({documentType: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/kYCDocuments');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add KYCDocument</h3>
        }else{
            return <h3 className="text-center">Update KYCDocument</h3>
        }
    }
    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                {
                                    this.getTitle()
                                }
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> reference:&emsp; </label>
                                                <input placeholder="reference" name="reference" className="form-control" value={this.state.reference} onChange={this.changereferenceHandler}/>

                                            <label> issuedCountry:&emsp; </label>
                                                <input placeholder="issuedCountry" name="issuedCountry" className="form-control" value={this.state.issuedCountry} onChange={this.changeissuedCountryHandler}/>

                                            <label> expirationDate:&emsp; </label>
                                                <input type="date" placeholder="expirationDate" name="expirationDate" className="form-control" value={this.state.expirationDate} onChange={this.changeexpirationDateHandler}/>

                                            <label> DocumentType:&emsp; </label>
                                                <select value={this.state.documentType} onChange={this.changeDocumentTypeHandler}>
                      <option name="DocumentType" className="form-control" >
                          Passport
                      </option>
                      <option name="DocumentType" className="form-control" >
                          NationalID
                      </option>
                      <option name="DocumentType" className="form-control" >
                          DriverLicense
                      </option>
                      <option name="DocumentType" className="form-control" >
                          BusinessRegistration
                      </option>
                      <option name="DocumentType" className="form-control" >
                          ProofOfAddress
                      </option>
                    </select>

                                            <label> Status:&emsp; </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Submitted
                      </option>
                      <option name="Status" className="form-control" >
                          Approved
                      </option>
                      <option name="Status" className="form-control" >
                          Rejected
                      </option>
                      <option name="Status" className="form-control" >
                          Expired
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateKYCDocument}>Save</button>
                                        <button className="btn btn-danger" onClick={this.cancel.bind(this)} style={{marginLeft: "10px"}}>Cancel</button>
                                    </form>
                                </div>
                            </div>
                        </div>
                   </div>
            </div>
        )
    }
}

export default CreateKYCDocumentComponent
