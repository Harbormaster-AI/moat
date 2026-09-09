import React, { Component } from 'react'
import CreditorService from '../services/CreditorService';

class CreateCreditorComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                bic: '',
                address: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changebicHandler = this.changebicHandler.bind(this);
        this.changeaddressHandler = this.changeaddressHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            CreditorService.getCreditorById(this.state.id).then( (res) =>{
                let creditor = res.data;
                this.setState({
                    name: creditor.name,
                    bic: creditor.bic,
                    address: creditor.address
                });
            });
        }        
    }
    saveOrUpdateCreditor = (e) => {
        e.preventDefault();
        let creditor = {
                creditorId: this.state.id,
                name: this.state.name,
                bic: this.state.bic,
                address: this.state.address
            };
        console.log('creditor => ' + JSON.stringify(creditor));

        // step 5
        if(this.state.id === '_add'){
            creditor.creditorId=''
            CreditorService.createCreditor(creditor).then(res =>{
                this.props.history.push('/creditors');
            });
        }else{
            CreditorService.updateCreditor(creditor).then( res => {
                this.props.history.push('/creditors');
            });
        }
    }
    
    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changebicHandler= (event) => {
        this.setState({bic: event.target.value});
    }
    changeaddressHandler= (event) => {
        this.setState({address: event.target.value});
    }

    cancel(){
        this.props.history.push('/creditors');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Creditor</h3>
        }else{
            return <h3 className="text-center">Update Creditor</h3>
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
                                            <label> name:&emsp; </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> bic:&emsp; </label>
                                                <input placeholder="bic" name="bic" className="form-control" value={this.state.bic} onChange={this.changebicHandler}/>

                                            <label> address:&emsp; </label>
                                                <input placeholder="address" name="address" className="form-control" value={this.state.address} onChange={this.changeaddressHandler}/>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateCreditor}>Save</button>
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

export default CreateCreditorComponent
