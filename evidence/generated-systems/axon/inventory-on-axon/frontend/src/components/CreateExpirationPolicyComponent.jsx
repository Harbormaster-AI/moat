import React, { Component } from 'react'
import ExpirationPolicyService from '../services/ExpirationPolicyService';

class CreateExpirationPolicyComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                rejectIfDaysToExpireLessThan: '',
                autoQuarantineDaysToExpire: '',
                rotationMethod: ''
        }
        this.changerejectIfDaysToExpireLessThanHandler = this.changerejectIfDaysToExpireLessThanHandler.bind(this);
        this.changeautoQuarantineDaysToExpireHandler = this.changeautoQuarantineDaysToExpireHandler.bind(this);
        this.changeRotationMethodHandler = this.changeRotationMethodHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            ExpirationPolicyService.getExpirationPolicyById(this.state.id).then( (res) =>{
                let expirationPolicy = res.data;
                this.setState({
                    rejectIfDaysToExpireLessThan: expirationPolicy.rejectIfDaysToExpireLessThan,
                    autoQuarantineDaysToExpire: expirationPolicy.autoQuarantineDaysToExpire,
                    rotationMethod: expirationPolicy.rotationMethod
                });
            });
        }        
    }
    saveOrUpdateExpirationPolicy = (e) => {
        e.preventDefault();
        let expirationPolicy = {
                expirationPolicyId: this.state.id,
                rejectIfDaysToExpireLessThan: this.state.rejectIfDaysToExpireLessThan,
                autoQuarantineDaysToExpire: this.state.autoQuarantineDaysToExpire,
                rotationMethod: this.state.rotationMethod
            };
        console.log('expirationPolicy => ' + JSON.stringify(expirationPolicy));

        // step 5
        if(this.state.id === '_add'){
            expirationPolicy.expirationPolicyId=''
            ExpirationPolicyService.createExpirationPolicy(expirationPolicy).then(res =>{
                this.props.history.push('/expirationPolicys');
            });
        }else{
            ExpirationPolicyService.updateExpirationPolicy(expirationPolicy).then( res => {
                this.props.history.push('/expirationPolicys');
            });
        }
    }
    
    changerejectIfDaysToExpireLessThanHandler= (event) => {
        this.setState({rejectIfDaysToExpireLessThan: event.target.value});
    }
    changeautoQuarantineDaysToExpireHandler= (event) => {
        this.setState({autoQuarantineDaysToExpire: event.target.value});
    }
    changeRotationMethodHandler= (event) => {
        this.setState({rotationMethod: event.target.value});
    }

    cancel(){
        this.props.history.push('/expirationPolicys');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add ExpirationPolicy</h3>
        }else{
            return <h3 className="text-center">Update ExpirationPolicy</h3>
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
                                            <label> rejectIfDaysToExpireLessThan:&emsp; </label>
                                                <input type="number" placeholder="rejectIfDaysToExpireLessThan" name="rejectIfDaysToExpireLessThan" className="form-control" value={this.state.rejectIfDaysToExpireLessThan} onChange={this.changerejectIfDaysToExpireLessThanHandler}/>

                                            <label> autoQuarantineDaysToExpire:&emsp; </label>
                                                <input type="number" placeholder="autoQuarantineDaysToExpire" name="autoQuarantineDaysToExpire" className="form-control" value={this.state.autoQuarantineDaysToExpire} onChange={this.changeautoQuarantineDaysToExpireHandler}/>

                                            <label> RotationMethod:&emsp; </label>
                                                <select value={this.state.rotationMethod} onChange={this.changeRotationMethodHandler}>
                      <option name="RotationMethod" className="form-control" >
                          FIFO
                      </option>
                      <option name="RotationMethod" className="form-control" >
                          LIFO
                      </option>
                      <option name="RotationMethod" className="form-control" >
                          FEFO
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateExpirationPolicy}>Save</button>
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

export default CreateExpirationPolicyComponent
