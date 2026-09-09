import React, { Component } from 'react'
import EndorsementService from '../services/EndorsementService';

class UpdateEndorsementComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                endorsementNumber: '',
                effectiveDate: '',
                description: ''
        }
        this.updateEndorsement = this.updateEndorsement.bind(this);

        this.changeendorsementNumberHandler = this.changeendorsementNumberHandler.bind(this);
        this.changeeffectiveDateHandler = this.changeeffectiveDateHandler.bind(this);
        this.changedescriptionHandler = this.changedescriptionHandler.bind(this);
    }

    componentDidMount(){
        EndorsementService.getEndorsementById(this.state.id).then( (res) =>{
            let endorsement = res.data;
            this.setState({
                endorsementNumber: endorsement.endorsementNumber,
                effectiveDate: endorsement.effectiveDate,
                description: endorsement.description
            });
        });
    }

    updateEndorsement = (e) => {
        e.preventDefault();
        let endorsement = {
            endorsementId: this.state.id,
            endorsementNumber: this.state.endorsementNumber,
            effectiveDate: this.state.effectiveDate,
            description: this.state.description
        };
        console.log('endorsement => ' + JSON.stringify(endorsement));
        console.log('id => ' + JSON.stringify(this.state.id));
        EndorsementService.updateEndorsement(endorsement).then( res => {
            this.props.history.push('/endorsements');
        });
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

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Endorsement</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> endorsementNumber: </label>
                                                <input placeholder="endorsementNumber" name="endorsementNumber" className="form-control" value={this.state.endorsementNumber} onChange={this.changeendorsementNumberHandler}/>

                                            <label> effectiveDate: </label>
                                                <input type="date" placeholder="effectiveDate" name="effectiveDate" className="form-control" value={this.state.effectiveDate} onChange={this.changeeffectiveDateHandler}/>

                                            <label> description: </label>
                                                <input placeholder="description" name="description" className="form-control" value={this.state.description} onChange={this.changedescriptionHandler}/>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateEndorsement}>Save</button>
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

export default UpdateEndorsementComponent
