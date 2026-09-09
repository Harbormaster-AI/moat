import React, { Component } from 'react'
import ContractService from '../services/ContractService';

class CreateContractComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                title: '',
                effectiveDate: '',
                expiryDate: '',
                repositoryUrl: '',
                status: ''
        }
        this.changetitleHandler = this.changetitleHandler.bind(this);
        this.changeeffectiveDateHandler = this.changeeffectiveDateHandler.bind(this);
        this.changeexpiryDateHandler = this.changeexpiryDateHandler.bind(this);
        this.changerepositoryUrlHandler = this.changerepositoryUrlHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            ContractService.getContractById(this.state.id).then( (res) =>{
                let contract = res.data;
                this.setState({
                    title: contract.title,
                    effectiveDate: contract.effectiveDate,
                    expiryDate: contract.expiryDate,
                    repositoryUrl: contract.repositoryUrl,
                    status: contract.status
                });
            });
        }        
    }
    saveOrUpdateContract = (e) => {
        e.preventDefault();
        let contract = {
                contractId: this.state.id,
                title: this.state.title,
                effectiveDate: this.state.effectiveDate,
                expiryDate: this.state.expiryDate,
                repositoryUrl: this.state.repositoryUrl,
                status: this.state.status
            };
        console.log('contract => ' + JSON.stringify(contract));

        // step 5
        if(this.state.id === '_add'){
            contract.contractId=''
            ContractService.createContract(contract).then(res =>{
                this.props.history.push('/contracts');
            });
        }else{
            ContractService.updateContract(contract).then( res => {
                this.props.history.push('/contracts');
            });
        }
    }
    
    changetitleHandler= (event) => {
        this.setState({title: event.target.value});
    }
    changeeffectiveDateHandler= (event) => {
        this.setState({effectiveDate: event.target.value});
    }
    changeexpiryDateHandler= (event) => {
        this.setState({expiryDate: event.target.value});
    }
    changerepositoryUrlHandler= (event) => {
        this.setState({repositoryUrl: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/contracts');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Contract</h3>
        }else{
            return <h3 className="text-center">Update Contract</h3>
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
                                            <label> title:&emsp; </label>
                                                <input placeholder="title" name="title" className="form-control" value={this.state.title} onChange={this.changetitleHandler}/>

                                            <label> effectiveDate:&emsp; </label>
                                                <input type="date" placeholder="effectiveDate" name="effectiveDate" className="form-control" value={this.state.effectiveDate} onChange={this.changeeffectiveDateHandler}/>

                                            <label> expiryDate:&emsp; </label>
                                                <input type="date" placeholder="expiryDate" name="expiryDate" className="form-control" value={this.state.expiryDate} onChange={this.changeexpiryDateHandler}/>

                                            <label> repositoryUrl:&emsp; </label>
                                                <input placeholder="repositoryUrl" name="repositoryUrl" className="form-control" value={this.state.repositoryUrl} onChange={this.changerepositoryUrlHandler}/>

                                            <label> Status:&emsp; </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Draft
                      </option>
                      <option name="Status" className="form-control" >
                          Active
                      </option>
                      <option name="Status" className="form-control" >
                          Expiring
                      </option>
                      <option name="Status" className="form-control" >
                          Terminated
                      </option>
                      <option name="Status" className="form-control" >
                          Archived
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateContract}>Save</button>
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

export default CreateContractComponent
