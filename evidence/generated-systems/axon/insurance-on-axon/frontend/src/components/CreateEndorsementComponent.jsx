import React, { Component } from 'react'
import EndorsementService from '../services/EndorsementService';

class CreateEndorsementComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                endorsementNumber: '',
                effectiveDate: '',
                description: ''
        }
        this.changeendorsementNumberHandler = this.changeendorsementNumberHandler.bind(this);
        this.changeeffectiveDateHandler = this.changeeffectiveDateHandler.bind(this);
        this.changedescriptionHandler = this.changedescriptionHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            EndorsementService.getEndorsementById(this.state.id).then( (res) =>{
                let endorsement = res.data;
                this.setState({
                    endorsementNumber: endorsement.endorsementNumber,
                    effectiveDate: endorsement.effectiveDate,
                    description: endorsement.description
                });
            });
        }        
    }
    saveOrUpdateEndorsement = (e) => {
        e.preventDefault();
        let endorsement = {
                endorsementId: this.state.id,
                endorsementNumber: this.state.endorsementNumber,
                effectiveDate: this.state.effectiveDate,
                description: this.state.description
            };
        console.log('endorsement => ' + JSON.stringify(endorsement));

        // step 5
        if(this.state.id === '_add'){
            endorsement.endorsementId=''
            EndorsementService.createEndorsement(endorsement).then(res =>{
                this.props.history.push('/endorsements');
            });
        }else{
            EndorsementService.updateEndorsement(endorsement).then( res => {
                this.props.history.push('/endorsements');
            });
        }
    }
    
    changeendorsementNumberHandler= (event) => {
        this.setState({endorsementNumber: event.target.value});
    }
    changeeffectiveDateHandler= (event) => {
        this.setState({effectiveDate: event.target.value});
    }
    changedescriptionHandler= (event) => {
        this.setState({description: event.target.value});
    }

    cancel(){
        this.props.history.push('/endorsements');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Endorsement</h3>
        }else{
            return <h3 className="text-center">Update Endorsement</h3>
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
                                            <label> endorsementNumber:&emsp; </label>
                                                <input placeholder="endorsementNumber" name="endorsementNumber" className="form-control" value={this.state.endorsementNumber} onChange={this.changeendorsementNumberHandler}/>

                                            <label> effectiveDate:&emsp; </label>
                                                <input type="date" placeholder="effectiveDate" name="effectiveDate" className="form-control" value={this.state.effectiveDate} onChange={this.changeeffectiveDateHandler}/>

                                            <label> description:&emsp; </label>
                                                <input placeholder="description" name="description" className="form-control" value={this.state.description} onChange={this.changedescriptionHandler}/>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateEndorsement}>Save</button>
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

export default CreateEndorsementComponent
